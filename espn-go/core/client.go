package core

import (
	"context"
	"net/url"
	"strings"

	"github.com/pseudo-r/Public-ESPN-API/espn-go/endpoints"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/internal/httpclient"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/models"
)

// Client wraps sports.core.api.espn.com core-v2 and core-v3 resources.
type Client struct {
	transport *httpclient.Client
	resolver  endpoints.Resolver
	sport     string
	league    string
}

func New(transport *httpclient.Client, resolver endpoints.Resolver, sport, league string) *Client {
	return &Client{transport: transport, resolver: resolver, sport: sport, league: league}
}

func (c *Client) get(ctx context.Context, suffix string, q url.Values) (models.JSON, error) {
	return c.transport.GetJSON(ctx, c.resolver.CoreV2(c.sport, c.league, suffix, cloneValues(q)))
}

func cloneValues(q url.Values) url.Values {
	if q == nil {
		return nil
	}
	out := make(url.Values, len(q))
	for key, values := range q {
		out[key] = append([]string(nil), values...)
	}
	return out
}

func defaultCompetitionID(eventID, competitionID string) string {
	if competitionID != "" {
		return competitionID
	}
	return eventID
}

func (c *Client) LeagueInfo(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "", q)
}

func (c *Client) Calendar(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "calendar", q)
}

func (c *Client) Season(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "season", q)
}

func (c *Client) Seasons(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "seasons", q)
}

func (c *Client) SeasonDetail(ctx context.Context, season string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "seasons/"+season, q)
}

func (c *Client) SeasonAthletes(ctx context.Context, season string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "seasons/"+season+"/athletes", q)
}

func (c *Client) SeasonAthlete(ctx context.Context, season, athleteID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "seasons/"+season+"/athletes/"+athleteID, q)
}

func (c *Client) Draft(ctx context.Context, season string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "seasons/"+season+"/draft", q)
}

func (c *Client) DraftStatus(ctx context.Context, season string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "seasons/"+season+"/draft/status", q)
}

func (c *Client) FreeAgents(ctx context.Context, season string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "seasons/"+season+"/freeagents", q)
}

func (c *Client) Manufacturers(ctx context.Context, season string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "seasons/"+season+"/manufacturers", q)
}

func (c *Client) Teams(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "teams", q)
}

func (c *Client) Team(ctx context.Context, season, teamID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "seasons/"+season+"/teams/"+teamID, q)
}

func (c *Client) Athletes(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "athletes", q)
}

func (c *Client) Athlete(ctx context.Context, athleteID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "athletes/"+athleteID, q)
}

func (c *Client) AthleteStatistics(ctx context.Context, athleteID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "athletes/"+athleteID+"/statistics", q)
}

func (c *Client) AthleteStatisticsLog(ctx context.Context, athleteID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "athletes/"+athleteID+"/statisticslog", q)
}

func (c *Client) AthleteContracts(ctx context.Context, athleteID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "athletes/"+athleteID+"/contracts", q)
}

func (c *Client) AthleteRecords(ctx context.Context, athleteID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "athletes/"+athleteID+"/records", q)
}

func (c *Client) Events(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "events", q)
}

func (c *Client) Event(ctx context.Context, eventID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "events/"+eventID, q)
}

func (c *Client) Competition(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "events/"+eventID+"/competitions/"+defaultCompetitionID(eventID, competitionID), q)
}

func (c *Client) Broadcasts(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "broadcasts", q)
}

func (c *Client) Odds(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "odds", q)
}

func (c *Client) Officials(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "officials", q)
}

func (c *Client) Status(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "status", q)
}

func (c *Client) Situation(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "situation", q)
}

func (c *Client) Details(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "details", q)
}

func (c *Client) Notes(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "notes", q)
}

func (c *Client) Predictor(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "predictor", q)
}

func (c *Client) PowerIndex(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "powerindex", q)
}

func (c *Client) PowerIndexTeam(ctx context.Context, eventID, competitionID, teamID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "powerindex/"+teamID, q)
}

func (c *Client) Drives(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "drives", q)
}

func (c *Client) Drive(ctx context.Context, eventID, competitionID, driveID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "drives/"+driveID, q)
}

func (c *Client) DrivePlays(ctx context.Context, eventID, competitionID, driveID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "drives/"+driveID+"/plays", q)
}

func (c *Client) Plays(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "plays", q)
}

func (c *Client) Play(ctx context.Context, eventID, competitionID, playID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "plays/"+playID, q)
}

func (c *Client) PlayPersonnel(ctx context.Context, eventID, competitionID, playID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "plays/"+playID+"/personnel", q)
}

func (c *Client) Probabilities(ctx context.Context, eventID, competitionID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "probabilities", q)
}

func (c *Client) Probability(ctx context.Context, eventID, competitionID, probabilityID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "probabilities/"+probabilityID, q)
}

func (c *Client) Competitor(ctx context.Context, eventID, competitionID, competitorID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "competitors/"+competitorID, q)
}

func (c *Client) CompetitorScore(ctx context.Context, eventID, competitionID, competitorID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "competitors/"+competitorID+"/score", q)
}

func (c *Client) CompetitorLinescores(ctx context.Context, eventID, competitionID, competitorID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "competitors/"+competitorID+"/linescores", q)
}

func (c *Client) CompetitorRoster(ctx context.Context, eventID, competitionID, competitorID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "competitors/"+competitorID+"/roster", q)
}

func (c *Client) CompetitorStatistics(ctx context.Context, eventID, competitionID, competitorID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "competitors/"+competitorID+"/statistics", q)
}

func (c *Client) CompetitorLeaders(ctx context.Context, eventID, competitionID, competitorID string, q url.Values) (models.JSON, error) {
	return c.competitionChild(ctx, eventID, competitionID, "competitors/"+competitorID+"/leaders", q)
}

func (c *Client) RosterAthleteStatistics(ctx context.Context, eventID, competitionID, competitorID, athleteID, split string, q url.Values) (models.JSON, error) {
	if split == "" {
		split = "0"
	}
	suffix := "competitors/" + competitorID + "/roster/" + athleteID + "/statistics/" + split
	return c.competitionChild(ctx, eventID, competitionID, suffix, q)
}

func (c *Client) Venues(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "venues", q)
}

func (c *Client) Venue(ctx context.Context, venueID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "venues/"+venueID, q)
}

func (c *Client) Casinos(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "casinos", q)
}

func (c *Client) Casino(ctx context.Context, casinoID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "casinos/"+casinoID, q)
}

func (c *Client) Circuits(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "circuits", q)
}

func (c *Client) Circuit(ctx context.Context, circuitID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "circuits/"+circuitID, q)
}

func (c *Client) Countries(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "countries", q)
}

func (c *Client) Country(ctx context.Context, countryID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "countries/"+countryID, q)
}

func (c *Client) Franchises(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "franchises", q)
}

func (c *Client) Franchise(ctx context.Context, franchiseID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "franchises/"+franchiseID, q)
}

func (c *Client) Positions(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "positions", q)
}

func (c *Client) Position(ctx context.Context, positionID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "positions/"+positionID, q)
}

func (c *Client) Providers(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "providers", q)
}

func (c *Client) Provider(ctx context.Context, providerID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "providers/"+providerID, q)
}

func (c *Client) Rankings(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "rankings", q)
}

func (c *Client) Recruiting(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "recruiting", q)
}

func (c *Client) Tournaments(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "tournaments", q)
}

func (c *Client) V3(ctx context.Context, suffix string, q url.Values) (models.JSON, error) {
	return c.transport.GetJSON(ctx, c.resolver.CoreV3(c.sport, c.league, strings.TrimLeft(suffix, "/"), cloneValues(q)))
}

func (c *Client) V3League(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.V3(ctx, "", q)
}

func (c *Client) V3Season(ctx context.Context, season string, q url.Values) (models.JSON, error) {
	return c.V3(ctx, "seasons/"+season, q)
}

func (c *Client) V3Team(ctx context.Context, teamID string, q url.Values) (models.JSON, error) {
	return c.V3(ctx, "teams/"+teamID, q)
}

func (c *Client) V3Athletes(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.V3(ctx, "athletes", q)
}

func (c *Client) V3Athlete(ctx context.Context, athleteID string, q url.Values) (models.JSON, error) {
	return c.V3(ctx, "athletes/"+athleteID, q)
}

func (c *Client) V3AthletePlays(ctx context.Context, athleteID string, q url.Values) (models.JSON, error) {
	return c.V3(ctx, "athletes/"+athleteID+"/plays", q)
}

func (c *Client) V3AthleteStatisticsLog(ctx context.Context, athleteID string, q url.Values) (models.JSON, error) {
	return c.V3(ctx, "athletes/"+athleteID+"/statisticslog", q)
}

func (c *Client) URL(suffix string, q url.Values) string {
	return c.resolver.CoreV2(c.sport, c.league, strings.TrimLeft(suffix, "/"), cloneValues(q))
}

func (c *Client) competitionChild(ctx context.Context, eventID, competitionID, suffix string, q url.Values) (models.JSON, error) {
	path := "events/" + eventID + "/competitions/" + defaultCompetitionID(eventID, competitionID) + "/" + strings.TrimLeft(suffix, "/")
	return c.get(ctx, path, q)
}
