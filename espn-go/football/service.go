package football

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/pseudo-r/Public-ESPN-API/espn-go/cdn"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/common"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/core"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/endpoints"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/internal/httpclient"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/models"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/now"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/search"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/site"
)

const defaultLeague = "nfl"

var athleteIDPattern = regexp.MustCompile(`~a:(\d+)`)

// Service exposes football-first workflows while keeping raw ESPN domain
// clients available for callers that need native payloads.
type Service struct {
	Sport  string
	League string

	Site   *site.Client
	Core   *core.Client
	Common *common.Client
	CDN    *cdn.Client
	Search *search.Client
	News   *now.Client
}

func New(transport *httpclient.Client, resolver endpoints.Resolver, league string) *Service {
	if league == "" {
		league = defaultLeague
	}
	sport := "football"
	return &Service{
		Sport:  sport,
		League: league,
		Site:   site.New(transport, resolver, sport, league),
		Core:   core.New(transport, resolver, sport, league),
		Common: common.New(transport, resolver, sport, league),
		CDN:    cdn.New(transport, resolver, league),
		Search: search.New(transport, resolver, sport, league),
		News:   now.New(transport, resolver, sport),
	}
}

// ParseAthleteID returns the canonical numeric ESPN athlete ID from Search v2
// UID strings such as s:20~l:28~a:4431452.
func ParseAthleteID(uid string) string {
	match := athleteIDPattern.FindStringSubmatch(uid)
	if len(match) != 2 {
		return ""
	}
	return match[1]
}

func (s *Service) SearchPlayers(ctx context.Context, query string, q url.Values) ([]models.PlayerHit, error) {
	payload, err := s.Search.Query(ctx, query, q)
	if err != nil {
		return nil, err
	}
	return parsePlayerHits(payload, s.League), nil
}

func (s *Service) ResolvePlayerID(ctx context.Context, playerName string) (models.PlayerHit, error) {
	hits, err := s.SearchPlayers(ctx, playerName, url.Values{"limit": {"10"}})
	if err != nil {
		return models.PlayerHit{}, err
	}
	if len(hits) == 0 {
		return models.PlayerHit{}, fmt.Errorf("football player %q not found", playerName)
	}
	return hits[0], nil
}

func (s *Service) GetPlayerProfile(ctx context.Context, athleteID, playerName string) (models.JSON, error) {
	id, err := s.resolveAthleteID(ctx, athleteID, playerName)
	if err != nil {
		return nil, err
	}
	payload, _, err := s.trySources(ctx, NeedPlayerProfile,
		sourceAttempt{source: SourceCommonV3, call: func(ctx context.Context) (models.JSON, error) {
			return s.Common.Athlete(ctx, id, nil)
		}},
		sourceAttempt{source: SourceCoreV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Core.Athlete(ctx, id, nil)
		}},
		sourceAttempt{source: SourceCoreV3, call: func(ctx context.Context) (models.JSON, error) {
			return s.Core.V3Athlete(ctx, id, nil)
		}},
	)
	return payload, err
}

func (s *Service) GetPlayerOverview(ctx context.Context, athleteID, playerName string, q url.Values) (models.JSON, error) {
	id, err := s.resolveAthleteID(ctx, athleteID, playerName)
	if err != nil {
		return nil, err
	}
	return s.Common.AthleteOverview(ctx, id, q)
}

func (s *Service) GetPlayerSeasonStats(ctx context.Context, athleteID, playerName string, q url.Values) (models.JSON, error) {
	id, err := s.resolveAthleteID(ctx, athleteID, playerName)
	if err != nil {
		return nil, err
	}
	payload, _, err := s.trySources(ctx, NeedPlayerSeasonStats,
		sourceAttempt{source: SourceCommonV3, call: func(ctx context.Context) (models.JSON, error) {
			return s.Common.AthleteStats(ctx, id, q)
		}},
		sourceAttempt{source: SourceCoreV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Core.AthleteStatistics(ctx, id, q)
		}},
	)
	return payload, err
}

func (s *Service) GetPlayerGameLog(ctx context.Context, athleteID, playerName string, q url.Values) (models.JSON, error) {
	id, err := s.resolveAthleteID(ctx, athleteID, playerName)
	if err != nil {
		return nil, err
	}
	payload, _, err := s.trySources(ctx, NeedPlayerGameLog,
		sourceAttempt{source: SourceCommonV3, call: func(ctx context.Context) (models.JSON, error) {
			return s.Common.AthleteGameLog(ctx, id, q)
		}},
		sourceAttempt{source: SourceCoreV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Core.AthleteStatisticsLog(ctx, id, q)
		}},
		sourceAttempt{source: SourceCoreV3, call: func(ctx context.Context) (models.JSON, error) {
			return s.Core.V3AthleteStatisticsLog(ctx, id, q)
		}},
	)
	return payload, err
}

func (s *Service) GetPlayerSplits(ctx context.Context, athleteID, playerName string, q url.Values) (models.JSON, error) {
	id, err := s.resolveAthleteID(ctx, athleteID, playerName)
	if err != nil {
		return nil, err
	}
	payload, _, err := s.trySources(ctx, NeedPlayerSplits,
		sourceAttempt{source: SourceCommonV3, call: func(ctx context.Context) (models.JSON, error) {
			return s.Common.AthleteSplits(ctx, id, q)
		}},
		sourceAttempt{source: SourceCoreV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Core.AthleteStatistics(ctx, id, q)
		}},
	)
	return payload, err
}

func (s *Service) GetPlayerBio(ctx context.Context, athleteID, playerName string, q url.Values) (models.JSON, error) {
	id, err := s.resolveAthleteID(ctx, athleteID, playerName)
	if err != nil {
		return nil, err
	}
	return s.Common.AthleteBio(ctx, id, q)
}

func (s *Service) GetPlayerNews(ctx context.Context, athleteID, playerName string, q url.Values) (models.JSON, error) {
	id, err := s.resolveAthleteID(ctx, athleteID, playerName)
	if err != nil {
		return nil, err
	}
	return s.Site.AthleteNews(ctx, id, q)
}

func (s *Service) ResolveTeamID(ctx context.Context, teamName string) (models.TeamHit, error) {
	payload, err := s.Site.Teams(ctx, nil)
	if err == nil {
		teams := parseTeamHits(payload)
		if match := bestTeamMatch(teamName, teams); match.TeamID != "" {
			return match, nil
		}
	}
	corePayload, coreErr := s.Core.Teams(ctx, nil)
	if coreErr == nil {
		if match := bestTeamMatch(teamName, parseTeamHits(corePayload)); match.TeamID != "" {
			return match, nil
		}
	}
	if err != nil {
		return models.TeamHit{}, err
	}
	if coreErr != nil {
		return models.TeamHit{}, coreErr
	}
	return models.TeamHit{}, fmt.Errorf("football team %q not found", teamName)
}

func (s *Service) GetTeamProfile(ctx context.Context, teamID, teamName string, q url.Values) (models.JSON, error) {
	id, err := s.resolveTeamID(ctx, teamID, teamName)
	if err != nil {
		return nil, err
	}
	season := q.Get("season")
	payload, _, err := s.trySources(ctx, NeedTeamProfile,
		sourceAttempt{source: SourceSiteV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Site.Team(ctx, id, q)
		}},
		sourceAttempt{source: SourceCoreV2, call: func(ctx context.Context) (models.JSON, error) {
			if season == "" {
				return nil, errors.New("season is required for core-v2 team fallback")
			}
			return s.Core.Team(ctx, season, id, q)
		}},
		sourceAttempt{source: SourceCoreV3, call: func(ctx context.Context) (models.JSON, error) {
			return s.Core.V3Team(ctx, id, q)
		}},
	)
	return payload, err
}

func (s *Service) GetTeamRoster(ctx context.Context, teamID, teamName string, q url.Values) (models.JSON, error) {
	id, err := s.resolveTeamID(ctx, teamID, teamName)
	if err != nil {
		return nil, err
	}
	payload, _, err := s.trySources(ctx, NeedTeamRoster,
		sourceAttempt{source: SourceSiteV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Site.Roster(ctx, id, q)
		}},
		sourceAttempt{source: SourceCommonV3, call: func(ctx context.Context) (models.JSON, error) {
			return s.Common.TeamRoster(ctx, id, q)
		}},
	)
	return payload, err
}

func (s *Service) GetTeamDepthChart(ctx context.Context, teamID, teamName string, q url.Values) (models.JSON, error) {
	id, err := s.resolveTeamID(ctx, teamID, teamName)
	if err != nil {
		return nil, err
	}
	return s.Site.DepthCharts(ctx, id, q)
}

func (s *Service) GetTeamSchedule(ctx context.Context, teamID, teamName string, q url.Values) (models.JSON, error) {
	id, err := s.resolveTeamID(ctx, teamID, teamName)
	if err != nil {
		return nil, err
	}
	payload, _, err := s.trySources(ctx, NeedTeamSchedule,
		sourceAttempt{source: SourceSiteV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Site.Schedule(ctx, id, q)
		}},
		sourceAttempt{source: SourceSiteV2, call: func(ctx context.Context) (models.JSON, error) {
			scoreboard, err := s.Site.Scoreboard(ctx, q)
			if err != nil {
				return nil, err
			}
			return models.JSON{"events": filterEventsByTeam(scoreboard, id), "source": "scoreboard"}, nil
		}},
	)
	return payload, err
}

func (s *Service) GetTeamNextGame(ctx context.Context, teamID, teamName string, asOf time.Time, q url.Values) (models.JSON, error) {
	id, err := s.resolveTeamID(ctx, teamID, teamName)
	if err != nil {
		return nil, err
	}
	schedule, err := s.Site.Schedule(ctx, id, q)
	if err != nil {
		return nil, err
	}
	event := selectScheduleEvent(schedule, asOf, true)
	if len(event) == 0 {
		return nil, fmt.Errorf("next game not found for team %s", id)
	}
	return event, nil
}

func (s *Service) GetTeamPreviousGame(ctx context.Context, teamID, teamName string, asOf time.Time, q url.Values) (models.JSON, error) {
	id, err := s.resolveTeamID(ctx, teamID, teamName)
	if err != nil {
		return nil, err
	}
	schedule, err := s.Site.Schedule(ctx, id, q)
	if err != nil {
		return nil, err
	}
	event := selectScheduleEvent(schedule, asOf, false)
	if len(event) == 0 {
		return nil, fmt.Errorf("previous game not found for team %s", id)
	}
	return event, nil
}

func (s *Service) GetTeamGameStats(ctx context.Context, eventID, teamID, teamName, competitionID, competitorID string, q url.Values) (models.JSON, error) {
	id, err := s.resolveTeamID(ctx, teamID, teamName)
	if err != nil {
		return nil, err
	}
	game, err := s.ResolveGame(ctx, eventID)
	if err != nil {
		return nil, err
	}
	if competitionID == "" {
		competitionID = game.CompetitionID
	}
	if competitorID == "" {
		competitorID = resolveCompetitorID(game.Summary, id, "")
	}
	if competitorID == "" {
		return nil, fmt.Errorf("competitor ID not found for team %s in event %s", id, eventID)
	}
	return s.Core.CompetitorStatistics(ctx, eventID, competitionID, competitorID, q)
}

func (s *Service) GetTeamGameLeaders(ctx context.Context, eventID, teamID, teamName, competitionID, competitorID string, q url.Values) (models.JSON, error) {
	id, err := s.resolveTeamID(ctx, teamID, teamName)
	if err != nil {
		return nil, err
	}
	game, err := s.ResolveGame(ctx, eventID)
	if err != nil {
		return nil, err
	}
	if competitionID == "" {
		competitionID = game.CompetitionID
	}
	if competitorID == "" {
		competitorID = resolveCompetitorID(game.Summary, id, "")
	}
	if competitorID == "" {
		return nil, fmt.Errorf("competitor ID not found for team %s in event %s", id, eventID)
	}
	payload, _, err := s.trySources(ctx, NeedTeamGameLeaders,
		sourceAttempt{source: SourceCoreV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Core.CompetitorLeaders(ctx, eventID, competitionID, competitorID, q)
		}},
		sourceAttempt{source: SourceSiteV2, call: func(ctx context.Context) (models.JSON, error) {
			summary, err := s.Site.Summary(ctx, eventID)
			if err != nil {
				return nil, err
			}
			return models.JSON{"leaders": summary["leaders"]}, nil
		}},
	)
	return payload, err
}

func (s *Service) GetScoreboard(ctx context.Context, q url.Values) (models.JSON, error) {
	payload, _, err := s.trySources(ctx, NeedScoreboard,
		sourceAttempt{source: SourceSiteV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Site.Scoreboard(ctx, q)
		}},
		sourceAttempt{source: SourceCDN, call: func(ctx context.Context) (models.JSON, error) {
			return s.CDN.Scoreboard(ctx, q)
		}},
	)
	return payload, err
}

func (s *Service) GetCurrentWeek(ctx context.Context) (models.JSON, error) {
	scoreboard, err := s.Site.Scoreboard(ctx, nil)
	if err != nil {
		return nil, err
	}
	return models.JSON{
		"season": scoreboard["season"],
		"week":   scoreboard["week"],
	}, nil
}

func (s *Service) ResolveGame(ctx context.Context, eventID string) (models.ResolvedGame, error) {
	if eventID == "" {
		return models.ResolvedGame{}, errors.New("eventID is required")
	}
	summary, err := s.Site.Summary(ctx, eventID)
	if err != nil {
		return models.ResolvedGame{}, err
	}
	return resolvedGameFromSummary(eventID, summary), nil
}

func (s *Service) ResolveGameForTeam(ctx context.Context, query models.GameSearch) (models.ResolvedGame, error) {
	if query.EventID != "" {
		return s.ResolveGame(ctx, query.EventID)
	}
	teamID, err := s.resolveTeamID(ctx, query.TeamID, query.TeamName)
	if err != nil {
		return models.ResolvedGame{}, err
	}
	q := scheduleQuery(query)
	schedule, err := s.Site.Schedule(ctx, teamID, q)
	if err != nil {
		return models.ResolvedGame{}, err
	}
	event := findScheduleEvent(schedule, query)
	eventID := firstString(event, nil, "id")
	if eventID == "" {
		return models.ResolvedGame{}, fmt.Errorf("game not found for team %s", teamID)
	}
	return s.ResolveGame(ctx, eventID)
}

func (s *Service) GetGameSummary(ctx context.Context, eventID string) (models.JSON, error) {
	payload, _, err := s.trySources(ctx, NeedGameSummary,
		sourceAttempt{source: SourceSiteV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Site.Summary(ctx, eventID)
		}},
		sourceAttempt{source: SourceCDN, call: func(ctx context.Context) (models.JSON, error) {
			return s.CDN.Game(ctx, eventID)
		}},
		sourceAttempt{source: SourceCoreV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Core.Event(ctx, eventID, nil)
		}},
	)
	return payload, err
}

func (s *Service) GetBoxScore(ctx context.Context, eventID string) (models.JSON, error) {
	payload, _, err := s.trySources(ctx, NeedGameSummary,
		sourceAttempt{source: SourceSiteV2, call: func(ctx context.Context) (models.JSON, error) {
			summary, err := s.Site.Summary(ctx, eventID)
			if err != nil {
				return nil, err
			}
			return models.AsMap(summary["boxscore"]), nil
		}},
		sourceAttempt{source: SourceCDN, call: func(ctx context.Context) (models.JSON, error) {
			return s.CDN.Boxscore(ctx, eventID)
		}},
	)
	return payload, err
}

func (s *Service) GetPlayByPlay(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	if competitionID == "" {
		game, err := s.ResolveGame(ctx, eventID)
		if err != nil {
			return nil, err
		}
		competitionID = game.CompetitionID
	}
	payload, _, err := s.trySources(ctx, NeedPlayByPlay,
		sourceAttempt{source: SourceCoreV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Core.Plays(ctx, eventID, competitionID, q)
		}},
		sourceAttempt{source: SourceSiteV2, call: func(ctx context.Context) (models.JSON, error) {
			summary, err := s.Site.Summary(ctx, eventID)
			if err != nil {
				return nil, err
			}
			return models.JSON{"plays": summary["plays"], "drives": summary["drives"]}, nil
		}},
		sourceAttempt{source: SourceCDN, call: func(ctx context.Context) (models.JSON, error) {
			return s.CDN.PlayByPlay(ctx, eventID)
		}},
	)
	return payload, err
}

func (s *Service) GetDrives(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	if competitionID == "" {
		game, err := s.ResolveGame(ctx, eventID)
		if err != nil {
			return nil, err
		}
		competitionID = game.CompetitionID
	}
	payload, _, err := s.trySources(ctx, NeedDrives,
		sourceAttempt{source: SourceCoreV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Core.Drives(ctx, eventID, competitionID, q)
		}},
		sourceAttempt{source: SourceSiteV2, call: func(ctx context.Context) (models.JSON, error) {
			summary, err := s.Site.Summary(ctx, eventID)
			if err != nil {
				return nil, err
			}
			return models.JSON{"drives": summary["drives"]}, nil
		}},
		sourceAttempt{source: SourceCDN, call: func(ctx context.Context) (models.JSON, error) {
			return s.CDN.Game(ctx, eventID)
		}},
	)
	return payload, err
}

func (s *Service) GetScoringPlays(ctx context.Context, eventID string) (models.JSON, error) {
	summary, err := s.Site.Summary(ctx, eventID)
	if err != nil {
		return nil, err
	}
	return models.JSON{"scoringPlays": summary["scoringPlays"]}, nil
}

func (s *Service) GetGameMetadata(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	payload, _, err := s.trySources(ctx, NeedGameMetadata,
		sourceAttempt{source: SourceCoreV2, call: func(ctx context.Context) (models.JSON, error) {
			if competitionID != "" {
				return s.Core.Competition(ctx, eventID, competitionID, q)
			}
			return s.Core.Event(ctx, eventID, q)
		}},
		sourceAttempt{source: SourceSiteV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Site.Summary(ctx, eventID)
		}},
	)
	return payload, err
}

func (s *Service) GetDrive(ctx context.Context, eventID, competitionID, driveID string, q url.Values) (models.JSON, error) {
	if competitionID == "" {
		game, err := s.ResolveGame(ctx, eventID)
		if err != nil {
			return nil, err
		}
		competitionID = game.CompetitionID
	}
	return s.Core.Drive(ctx, eventID, competitionID, driveID, q)
}

func (s *Service) GetDrivePlays(ctx context.Context, eventID, competitionID, driveID string, q url.Values) (models.JSON, error) {
	if competitionID == "" {
		game, err := s.ResolveGame(ctx, eventID)
		if err != nil {
			return nil, err
		}
		competitionID = game.CompetitionID
	}
	return s.Core.DrivePlays(ctx, eventID, competitionID, driveID, q)
}

func (s *Service) GetPlay(ctx context.Context, eventID, competitionID, playID string, q url.Values) (models.JSON, error) {
	if competitionID == "" {
		game, err := s.ResolveGame(ctx, eventID)
		if err != nil {
			return nil, err
		}
		competitionID = game.CompetitionID
	}
	return s.Core.Play(ctx, eventID, competitionID, playID, q)
}

func (s *Service) GetProbability(ctx context.Context, eventID, competitionID, playID string, q url.Values) (models.JSON, error) {
	if competitionID == "" {
		game, err := s.ResolveGame(ctx, eventID)
		if err != nil {
			return nil, err
		}
		competitionID = game.CompetitionID
	}
	return s.Core.Probability(ctx, eventID, competitionID, playID, q)
}

func (s *Service) GetWinProbability(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	if competitionID == "" {
		game, err := s.ResolveGame(ctx, eventID)
		if err != nil {
			return nil, err
		}
		competitionID = game.CompetitionID
	}
	payload, _, err := s.trySources(ctx, NeedWinProbability,
		sourceAttempt{source: SourceCoreV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Core.Probabilities(ctx, eventID, competitionID, q)
		}},
		sourceAttempt{source: SourceSiteV2, call: func(ctx context.Context) (models.JSON, error) {
			summary, err := s.Site.Summary(ctx, eventID)
			if err != nil {
				return nil, err
			}
			return models.JSON{"winprobability": summary["winprobability"]}, nil
		}},
	)
	return payload, err
}

func (s *Service) GetLiveSituation(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	if competitionID == "" {
		game, err := s.ResolveGame(ctx, eventID)
		if err != nil {
			return nil, err
		}
		competitionID = game.CompetitionID
	}
	return s.Core.Situation(ctx, eventID, competitionID, q)
}

func (s *Service) GetLiveLastPlay(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	if competitionID == "" {
		game, err := s.ResolveGame(ctx, eventID)
		if err != nil {
			return nil, err
		}
		competitionID = game.CompetitionID
	}
	situation, err := s.Core.Situation(ctx, eventID, competitionID, q)
	if err != nil {
		return nil, err
	}
	lastPlay := models.AsMap(situation["lastPlay"])
	if id := firstString(lastPlay, nil, "id"); id != "" {
		return s.Core.Play(ctx, eventID, competitionID, id, q)
	}
	if id := idFromRef(firstString(lastPlay, nil, "$ref")); id != "" {
		return s.Core.Play(ctx, eventID, competitionID, id, q)
	}
	if len(lastPlay) > 0 {
		return lastPlay, nil
	}
	return nil, fmt.Errorf("last play not found for event %s", eventID)
}

func (s *Service) GetGameOdds(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	if competitionID == "" {
		game, err := s.ResolveGame(ctx, eventID)
		if err != nil {
			return nil, err
		}
		competitionID = game.CompetitionID
	}
	payload, _, err := s.trySources(ctx, NeedOdds,
		sourceAttempt{source: SourceCoreV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Core.Odds(ctx, eventID, competitionID, q)
		}},
		sourceAttempt{source: SourceSiteV2, call: func(ctx context.Context) (models.JSON, error) {
			summary, err := s.Site.Summary(ctx, eventID)
			if err != nil {
				return nil, err
			}
			return models.JSON{"odds": summary["odds"], "pickcenter": summary["pickcenter"]}, nil
		}},
	)
	return payload, err
}

func (s *Service) GetGameOddsProvider(ctx context.Context, providerID string, q url.Values) (models.JSON, error) {
	return s.Core.Provider(ctx, providerID, q)
}

func (s *Service) GetBroadcasts(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	if competitionID == "" {
		game, err := s.ResolveGame(ctx, eventID)
		if err != nil {
			return nil, err
		}
		competitionID = game.CompetitionID
	}
	return s.Core.Broadcasts(ctx, eventID, competitionID, q)
}

func (s *Service) GetVenue(ctx context.Context, venueID, eventID, teamID, teamName, season string, q url.Values) (models.JSON, error) {
	if venueID != "" {
		return s.Core.Venue(ctx, venueID, q)
	}
	if eventID != "" {
		game, err := s.ResolveGame(ctx, eventID)
		if err != nil {
			return nil, err
		}
		competition, err := s.Core.Competition(ctx, eventID, game.CompetitionID, q)
		if err != nil {
			return nil, err
		}
		if venue := models.AsMap(competition["venue"]); len(venue) > 0 {
			if id := firstString(venue, nil, "id"); id != "" {
				return s.Core.Venue(ctx, id, q)
			}
			if id := idFromRef(firstString(venue, nil, "$ref")); id != "" {
				return s.Core.Venue(ctx, id, q)
			}
			return venue, nil
		}
		return nil, fmt.Errorf("venue not found for event %s", eventID)
	}
	id, err := s.resolveTeamID(ctx, teamID, teamName)
	if err != nil {
		return nil, err
	}
	if season == "" {
		return nil, errors.New("season is required when resolving venue from team")
	}
	team, err := s.Core.Team(ctx, season, id, q)
	if err != nil {
		return nil, err
	}
	if venue := models.AsMap(team["venue"]); len(venue) > 0 {
		if venueID := firstString(venue, nil, "id"); venueID != "" {
			return s.Core.Venue(ctx, venueID, q)
		}
		if venueID := idFromRef(firstString(venue, nil, "$ref")); venueID != "" {
			return s.Core.Venue(ctx, venueID, q)
		}
		return venue, nil
	}
	return nil, fmt.Errorf("venue not found for team %s", id)
}

func (s *Service) GetStandings(ctx context.Context, q url.Values) (models.JSON, error) {
	payload, _, err := s.trySources(ctx, NeedStandings,
		sourceAttempt{source: SourceSiteAPI, call: func(ctx context.Context) (models.JSON, error) {
			return s.Site.Standings(ctx, q)
		}},
		sourceAttempt{source: SourceSiteV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Site.StandingsStub(ctx, q)
		}},
	)
	return payload, err
}

func (s *Service) GetLeagueLeaders(ctx context.Context, q url.Values) (models.JSON, error) {
	payload, _, err := s.trySources(ctx, NeedLeagueLeaders,
		sourceAttempt{source: SourceCommonV3, call: func(ctx context.Context) (models.JSON, error) {
			return s.Common.StatisticsByAthlete(ctx, q)
		}},
		sourceAttempt{source: SourceSiteV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Site.Statistics(ctx, q)
		}},
	)
	return payload, err
}

func (s *Service) GetTeamStatRankings(ctx context.Context, q url.Values) (models.JSON, error) {
	payload, _, err := s.trySources(ctx, NeedTeamStatRankings,
		sourceAttempt{source: SourceCommonV3, call: func(ctx context.Context) (models.JSON, error) {
			return s.Common.StatisticsByTeam(ctx, q)
		}},
		sourceAttempt{source: SourceSiteV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Site.Statistics(ctx, q)
		}},
	)
	return payload, err
}

func (s *Service) GetInjuries(ctx context.Context, q url.Values) (models.JSON, error) {
	return s.Site.Injuries(ctx, q)
}

func (s *Service) GetTeamInjuries(ctx context.Context, teamID, teamName string, q url.Values) (models.JSON, error) {
	id, err := s.resolveTeamID(ctx, teamID, teamName)
	if err != nil {
		return nil, err
	}
	payload, err := s.Site.Injuries(ctx, q)
	if err == nil && !emptyPayload(payload) {
		return filterPayloadByTeam(payload, id), nil
	}
	teamPayload, teamErr := s.Site.TeamInjuries(ctx, id, q)
	if teamErr != nil {
		if err != nil {
			return nil, err
		}
		return nil, teamErr
	}
	return teamPayload, nil
}

func (s *Service) GetTransactions(ctx context.Context, q url.Values) (models.JSON, error) {
	return s.Site.Transactions(ctx, q)
}

func (s *Service) GetTeamTransactions(ctx context.Context, teamID, teamName string, q url.Values) (models.JSON, error) {
	id, err := s.resolveTeamID(ctx, teamID, teamName)
	if err != nil {
		return nil, err
	}
	payload, err := s.Site.Transactions(ctx, q)
	if err == nil && !emptyPayload(payload) {
		return filterPayloadByTeam(payload, id), nil
	}
	teamPayload, teamErr := s.Site.TeamTransactions(ctx, id, q)
	if teamErr != nil {
		if err != nil {
			return nil, err
		}
		return nil, teamErr
	}
	return teamPayload, nil
}

func (s *Service) GetLeagueNews(ctx context.Context, q url.Values) (models.JSON, error) {
	payload, _, err := s.trySources(ctx, NeedNews,
		sourceAttempt{source: SourceSiteV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Site.News(ctx, q)
		}},
		sourceAttempt{source: SourceNow, call: func(ctx context.Context) (models.JSON, error) {
			return s.News.Headlines(ctx, q)
		}},
	)
	return payload, err
}

func (s *Service) SearchContent(ctx context.Context, query string, q url.Values) (models.JSON, error) {
	return s.Search.Query(ctx, query, q)
}

func (s *Service) GetDraftInfo(ctx context.Context, q url.Values) (models.JSON, error) {
	season := q.Get("season")
	payload, _, err := s.trySources(ctx, NeedDraft,
		sourceAttempt{source: SourceSiteV2, call: func(ctx context.Context) (models.JSON, error) {
			return s.Site.Draft(ctx, q)
		}},
		sourceAttempt{source: SourceCoreV2, call: func(ctx context.Context) (models.JSON, error) {
			if season == "" {
				return nil, errors.New("season is required for core-v2 draft fallback")
			}
			return s.Core.Draft(ctx, season, q)
		}},
	)
	return payload, err
}

func (s *Service) GetPlayerGameStats(ctx context.Context, input models.PlayerGameStatsInput) (models.PlayerGameStats, error) {
	if input.EventID == "" {
		return models.PlayerGameStats{}, errors.New("eventID is required")
	}
	athleteID, err := s.resolveAthleteID(ctx, input.AthleteID, input.PlayerName)
	if err != nil {
		return models.PlayerGameStats{}, err
	}
	input.AthleteID = athleteID

	var summary models.JSON
	if input.CompetitionID == "" || input.CompetitorID == "" {
		game, err := s.ResolveGame(ctx, input.EventID)
		if err != nil {
			return models.PlayerGameStats{}, err
		}
		summary = game.Summary
		if input.CompetitionID == "" {
			input.CompetitionID = game.CompetitionID
		}
		if input.CompetitorID == "" {
			input.CompetitorID = resolveCompetitorID(summary, input.TeamID, athleteID)
		}
	}
	if input.CompetitionID == "" {
		input.CompetitionID = input.EventID
	}
	if input.CompetitorID == "" {
		return models.PlayerGameStats{}, fmt.Errorf("competitor ID not found for athlete %s in event %s", athleteID, input.EventID)
	}

	raw, err := s.Core.RosterAthleteStatistics(ctx, input.EventID, input.CompetitionID, input.CompetitorID, athleteID, input.Split, nil)
	if err != nil {
		return models.PlayerGameStats{}, err
	}
	return models.PlayerGameStats{
		EventID:       input.EventID,
		CompetitionID: input.CompetitionID,
		CompetitorID:  input.CompetitorID,
		AthleteID:     athleteID,
		Stats:         normalizeStatLines(raw),
		Raw:           raw,
		Summary:       summary,
	}, nil
}

func (s *Service) ResolveIDsBundle(ctx context.Context, playerName, teamName, eventID string) (models.IDBundle, error) {
	var bundle models.IDBundle
	if playerName != "" {
		player, err := s.ResolvePlayerID(ctx, playerName)
		if err != nil {
			return bundle, err
		}
		bundle.Player = &player
	}
	if teamName != "" {
		team, err := s.ResolveTeamID(ctx, teamName)
		if err != nil {
			return bundle, err
		}
		bundle.Team = &team
	}
	if eventID != "" {
		game, err := s.ResolveGame(ctx, eventID)
		if err != nil {
			return bundle, err
		}
		bundle.Game = &game
	}
	return bundle, nil
}

func (s *Service) ResolveIDsBundleForSearch(ctx context.Context, playerName string, gameSearch models.GameSearch) (models.IDBundle, error) {
	var bundle models.IDBundle
	if playerName != "" {
		player, err := s.ResolvePlayerID(ctx, playerName)
		if err != nil {
			return bundle, err
		}
		bundle.Player = &player
	}
	if gameSearch.TeamName != "" || gameSearch.TeamID != "" {
		if gameSearch.TeamName != "" {
			team, err := s.ResolveTeamID(ctx, gameSearch.TeamName)
			if err != nil {
				return bundle, err
			}
			bundle.Team = &team
		} else {
			bundle.Team = &models.TeamHit{TeamID: gameSearch.TeamID}
		}
		game, err := s.ResolveGameForTeam(ctx, gameSearch)
		if err != nil {
			return bundle, err
		}
		bundle.Game = &game
	}
	return bundle, nil
}

func (s *Service) resolveAthleteID(ctx context.Context, athleteID, playerName string) (string, error) {
	if athleteID != "" {
		return athleteID, nil
	}
	if playerName == "" {
		return "", errors.New("athleteID or playerName is required")
	}
	hit, err := s.ResolvePlayerID(ctx, playerName)
	if err != nil {
		return "", err
	}
	return hit.AthleteID, nil
}

func (s *Service) resolveTeamID(ctx context.Context, teamID, teamName string) (string, error) {
	if teamID != "" {
		return teamID, nil
	}
	if teamName == "" {
		return "", errors.New("teamID or teamName is required")
	}
	hit, err := s.ResolveTeamID(ctx, teamName)
	if err != nil {
		return "", err
	}
	return hit.TeamID, nil
}

func parsePlayerHits(payload models.JSON, league string) []models.PlayerHit {
	var hits []models.PlayerHit
	for _, result := range models.AsSlice(payload["results"]) {
		resultMap := models.AsMap(result)
		contents := models.AsSlice(resultMap["contents"])
		if contents == nil {
			contents = models.AsSlice(resultMap["items"])
		}
		if len(contents) == 0 {
			contents = []any{resultMap}
		}
		for _, content := range contents {
			raw := models.AsMap(content)
			kind := strings.ToLower(firstString(raw, resultMap, "type"))
			uid := firstString(raw, resultMap, "uid")
			athleteID := ParseAthleteID(uid)
			if kind != "player" && athleteID == "" {
				continue
			}
			description := firstString(raw, resultMap, "description", "league")
			if league != "" && description != "" && !strings.EqualFold(description, league) {
				// Search v2 often returns the league in description for player hits.
				// Keep non-league-described rows, but avoid obvious cross-league hits.
				if strings.Contains(strings.ToLower(description), "college") {
					continue
				}
			}
			hits = append(hits, models.PlayerHit{
				AthleteID:   athleteID,
				UID:         uid,
				GUID:        firstString(raw, resultMap, "id"),
				DisplayName: firstString(raw, resultMap, "displayName", "name"),
				Description: description,
				League:      description,
				Raw:         raw,
			})
		}
	}
	return hits
}

func parseTeamHits(payload models.JSON) []models.TeamHit {
	var hits []models.TeamHit
	add := func(value any) {
		raw := models.AsMap(value)
		if len(raw) == 0 {
			return
		}
		if nested := models.AsMap(raw["team"]); len(nested) > 0 {
			raw = nested
		}
		id := firstString(raw, nil, "id")
		if id == "" {
			return
		}
		hits = append(hits, models.TeamHit{
			TeamID:       id,
			Abbreviation: firstString(raw, nil, "abbreviation"),
			DisplayName:  firstString(raw, nil, "displayName", "name", "shortDisplayName"),
			Location:     firstString(raw, nil, "location"),
			Nickname:     firstString(raw, nil, "nickname"),
			Raw:          raw,
		})
	}

	for _, item := range models.AsSlice(payload["teams"]) {
		add(item)
	}
	for _, sport := range models.AsSlice(payload["sports"]) {
		for _, league := range models.AsSlice(models.AsMap(sport)["leagues"]) {
			for _, team := range models.AsSlice(models.AsMap(league)["teams"]) {
				add(team)
			}
		}
	}
	return hits
}

func bestTeamMatch(query string, teams []models.TeamHit) models.TeamHit {
	needle := normalizeMatch(query)
	for _, team := range teams {
		if normalizeMatch(team.Abbreviation) == needle {
			return team
		}
	}
	for _, team := range teams {
		values := []string{team.DisplayName, team.Location, team.Nickname, team.Abbreviation}
		for _, value := range values {
			if normalizeMatch(value) == needle {
				return team
			}
		}
	}
	for _, team := range teams {
		haystack := normalizeMatch(strings.Join([]string{team.DisplayName, team.Location, team.Nickname, team.Abbreviation}, " "))
		if strings.Contains(haystack, needle) {
			return team
		}
	}
	return models.TeamHit{}
}

func resolvedGameFromSummary(eventID string, summary models.JSON) models.ResolvedGame {
	game := models.ResolvedGame{EventID: eventID, Summary: summary}
	header := models.AsMap(summary["header"])
	for _, competition := range models.AsSlice(header["competitions"]) {
		comp := models.AsMap(competition)
		if game.CompetitionID == "" {
			game.CompetitionID = firstString(comp, nil, "id")
		}
		for _, competitor := range models.AsSlice(comp["competitors"]) {
			competitorMap := models.AsMap(competitor)
			if id := firstString(competitorMap, nil, "id"); id != "" {
				game.CompetitorIDs = appendIfMissing(game.CompetitorIDs, id)
			}
			team := models.AsMap(competitorMap["team"])
			if id := firstString(team, nil, "id"); id != "" {
				game.TeamIDs = appendIfMissing(game.TeamIDs, id)
			}
		}
	}
	boxscore := models.AsMap(summary["boxscore"])
	for _, teamRow := range models.AsSlice(boxscore["teams"]) {
		team := models.AsMap(models.AsMap(teamRow)["team"])
		if id := firstString(team, nil, "id"); id != "" {
			game.TeamIDs = appendIfMissing(game.TeamIDs, id)
			game.CompetitorIDs = appendIfMissing(game.CompetitorIDs, id)
		}
	}
	if game.CompetitionID == "" {
		game.CompetitionID = eventID
	}
	return game
}

func scheduleQuery(query models.GameSearch) url.Values {
	q := url.Values{}
	if query.Date != "" {
		q.Set("dates", strings.ReplaceAll(query.Date, "-", ""))
	}
	if query.Season != "" {
		q.Set("season", query.Season)
	}
	if query.SeasonType != "" {
		q.Set("seasontype", query.SeasonType)
	}
	if query.Week != "" {
		q.Set("week", query.Week)
	}
	if len(q) == 0 {
		return nil
	}
	return q
}

func findScheduleEvent(schedule models.JSON, query models.GameSearch) models.JSON {
	for _, item := range models.AsSlice(schedule["events"]) {
		event := models.AsMap(item)
		if query.EventID != "" && firstString(event, nil, "id") != query.EventID {
			continue
		}
		if query.Date != "" && !dateMatches(firstString(event, nil, "date"), query.Date) {
			continue
		}
		if query.Week != "" {
			week := models.AsMap(event["week"])
			if firstNonEmptyString(firstString(event, nil, "week"), firstString(week, nil, "number")) != query.Week {
				continue
			}
		}
		if query.OpponentID != "" && !eventHasTeam(event, query.OpponentID) {
			continue
		}
		if query.OpponentName != "" && !eventHasTeamName(event, query.OpponentName) {
			continue
		}
		return event
	}
	return models.JSON{}
}

func selectScheduleEvent(schedule models.JSON, asOf time.Time, next bool) models.JSON {
	if asOf.IsZero() {
		asOf = time.Now()
	}
	var selected models.JSON
	var selectedTime time.Time
	for _, item := range models.AsSlice(schedule["events"]) {
		event := models.AsMap(item)
		eventTime, ok := parseEventTime(firstString(event, nil, "date"))
		if !ok {
			continue
		}
		completed := eventCompleted(event)
		if next {
			if completed || eventTime.Before(asOf) {
				continue
			}
			if selected == nil || eventTime.Before(selectedTime) {
				selected = event
				selectedTime = eventTime
			}
			continue
		}
		if !completed && eventTime.After(asOf) {
			continue
		}
		if selected == nil || eventTime.After(selectedTime) {
			selected = event
			selectedTime = eventTime
		}
	}
	if selected == nil {
		return models.JSON{}
	}
	return selected
}

func parseEventTime(value string) (time.Time, bool) {
	if value == "" {
		return time.Time{}, false
	}
	if parsed, err := time.Parse(time.RFC3339, value); err == nil {
		return parsed, true
	}
	if parsed, err := time.Parse("2006-01-02", value); err == nil {
		return parsed, true
	}
	return time.Time{}, false
}

func dateMatches(eventDate, queryDate string) bool {
	normalizedQuery := strings.ReplaceAll(queryDate, "-", "")
	if eventTime, ok := parseEventTime(eventDate); ok {
		return eventTime.Format("20060102") == normalizedQuery
	}
	return strings.HasPrefix(strings.ReplaceAll(eventDate, "-", ""), normalizedQuery)
}

func eventCompleted(event models.JSON) bool {
	status := models.AsMap(event["status"])
	statusType := models.AsMap(status["type"])
	if completed, ok := statusType["completed"].(bool); ok {
		return completed
	}
	name := strings.ToLower(firstString(statusType, nil, "name", "state", "description"))
	return strings.Contains(name, "final") || strings.Contains(name, "post")
}

func eventHasTeam(event models.JSON, teamID string) bool {
	for _, competition := range models.AsSlice(event["competitions"]) {
		for _, competitor := range models.AsSlice(models.AsMap(competition)["competitors"]) {
			team := models.AsMap(models.AsMap(competitor)["team"])
			if firstString(team, nil, "id") == teamID {
				return true
			}
		}
	}
	return false
}

func filterEventsByTeam(scoreboard models.JSON, teamID string) []any {
	var events []any
	for _, item := range models.AsSlice(scoreboard["events"]) {
		event := models.AsMap(item)
		if eventHasTeam(event, teamID) {
			events = append(events, event)
		}
	}
	return events
}

func filterPayloadByTeam(payload models.JSON, teamID string) models.JSON {
	out := models.JSON{}
	for key, value := range payload {
		items := models.AsSlice(value)
		if items == nil {
			out[key] = value
			continue
		}
		var filtered []any
		for _, item := range items {
			itemMap := models.AsMap(item)
			if itemHasTeam(itemMap, teamID) {
				filtered = append(filtered, item)
			}
		}
		out[key] = filtered
	}
	return out
}

func itemHasTeam(item models.JSON, teamID string) bool {
	if firstString(item, nil, "teamId", "teamID") == teamID {
		return true
	}
	if firstString(models.AsMap(item["team"]), nil, "id") == teamID {
		return true
	}
	if firstString(models.AsMap(item["team"]), nil, "$ref") != "" && idFromRef(firstString(models.AsMap(item["team"]), nil, "$ref")) == teamID {
		return true
	}
	for _, key := range []string{"teams", "team"} {
		for _, child := range models.AsSlice(item[key]) {
			if firstString(models.AsMap(child), nil, "id") == teamID {
				return true
			}
		}
	}
	return false
}

func eventHasTeamName(event models.JSON, teamName string) bool {
	needle := normalizeMatch(teamName)
	for _, competition := range models.AsSlice(event["competitions"]) {
		for _, competitor := range models.AsSlice(models.AsMap(competition)["competitors"]) {
			team := models.AsMap(models.AsMap(competitor)["team"])
			haystack := normalizeMatch(strings.Join([]string{
				firstString(team, nil, "displayName"),
				firstString(team, nil, "shortDisplayName"),
				firstString(team, nil, "name"),
				firstString(team, nil, "location"),
				firstString(team, nil, "abbreviation"),
			}, " "))
			if strings.Contains(haystack, needle) {
				return true
			}
		}
	}
	return false
}

func resolveCompetitorID(summary models.JSON, teamID, athleteID string) string {
	boxscore := models.AsMap(summary["boxscore"])
	if teamID != "" {
		for _, teamRow := range models.AsSlice(boxscore["teams"]) {
			team := models.AsMap(models.AsMap(teamRow)["team"])
			if firstString(team, nil, "id") == teamID {
				return teamID
			}
		}
	}
	if athleteID == "" {
		return ""
	}
	for _, playerGroup := range models.AsSlice(boxscore["players"]) {
		group := models.AsMap(playerGroup)
		groupTeamID := firstString(models.AsMap(group["team"]), nil, "id")
		if groupTeamID == "" {
			groupTeamID = firstString(group, nil, "teamId")
		}
		if groupTeamID != "" && containsID(group, athleteID) {
			return groupTeamID
		}
	}
	return ""
}

func normalizeStatLines(raw models.JSON) []models.StatLine {
	var lines []models.StatLine
	splits := models.AsMap(raw["splits"])
	for _, category := range models.AsSlice(splits["categories"]) {
		categoryMap := models.AsMap(category)
		categoryName := firstString(categoryMap, nil, "name", "displayName")
		for _, stat := range models.AsSlice(categoryMap["stats"]) {
			statMap := models.AsMap(stat)
			lines = append(lines, models.StatLine{
				Category:     categoryName,
				Name:         firstString(statMap, nil, "name"),
				DisplayName:  firstString(statMap, nil, "displayName", "shortDisplayName"),
				Value:        statMap["value"],
				DisplayValue: firstString(statMap, nil, "displayValue"),
				Raw:          statMap,
			})
		}
	}
	return lines
}

func containsID(value any, id string) bool {
	switch typed := value.(type) {
	case map[string]any:
		for key, child := range typed {
			if (key == "id" || key == "athleteId" || key == "athleteID") && models.StringValue(child) == id {
				return true
			}
			if containsID(child, id) {
				return true
			}
		}
	case models.JSON:
		return containsID(map[string]any(typed), id)
	case []any:
		for _, child := range typed {
			if containsID(child, id) {
				return true
			}
		}
	}
	return false
}

func firstString(primary models.JSON, fallback models.JSON, keys ...string) string {
	for _, key := range keys {
		if primary != nil {
			if value := models.StringValue(primary[key]); value != "" {
				return value
			}
		}
		if fallback != nil {
			if value := models.StringValue(fallback[key]); value != "" {
				return value
			}
		}
	}
	return ""
}

func normalizeMatch(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	value = strings.ReplaceAll(value, ".", "")
	value = strings.ReplaceAll(value, "-", " ")
	return strings.Join(strings.Fields(value), " ")
}

func idFromRef(ref string) string {
	ref = strings.TrimSpace(ref)
	if ref == "" {
		return ""
	}
	if idx := strings.Index(ref, "?"); idx >= 0 {
		ref = ref[:idx]
	}
	ref = strings.TrimRight(ref, "/")
	if idx := strings.LastIndex(ref, "/"); idx >= 0 {
		return ref[idx+1:]
	}
	return ""
}

func firstNonEmptyString(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func appendIfMissing(values []string, value string) []string {
	for _, existing := range values {
		if existing == value {
			return values
		}
	}
	return append(values, value)
}
