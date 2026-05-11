package football

import (
	"context"
	"fmt"

	"github.com/pseudo-r/Public-ESPN-API/espn-go/models"
)

// Source identifies an ESPN API family used by football workflows.
type Source string

const (
	SourceSiteV2   Source = "site-v2"
	SourceSiteAPI  Source = "site-api"
	SourceCoreV2   Source = "core-v2"
	SourceCoreV3   Source = "core-v3"
	SourceCommonV3 Source = "common-v3"
	SourceCDN      Source = "cdn"
	SourceSearchV2 Source = "search-v2"
	SourceNow      Source = "now"
)

// DataNeed names the user-facing data shape being resolved.
type DataNeed string

const (
	NeedTeamLookup        DataNeed = "team-lookup"
	NeedTeamProfile       DataNeed = "team-profile"
	NeedTeamSchedule      DataNeed = "team-schedule"
	NeedTeamRoster        DataNeed = "team-roster"
	NeedDepthChart        DataNeed = "depth-chart"
	NeedPlayerLookup      DataNeed = "player-lookup"
	NeedPlayerProfile     DataNeed = "player-profile"
	NeedPlayerSeasonStats DataNeed = "player-season-stats"
	NeedPlayerGameLog     DataNeed = "player-game-log"
	NeedPlayerSplits      DataNeed = "player-splits"
	NeedPlayerGameStats   DataNeed = "player-game-stats"
	NeedLeagueLeaders     DataNeed = "league-leaders"
	NeedTeamStatRankings  DataNeed = "team-stat-rankings"
	NeedScoreboard        DataNeed = "scoreboard"
	NeedGameSummary       DataNeed = "game-summary"
	NeedGameMetadata      DataNeed = "game-metadata"
	NeedTeamGameStats     DataNeed = "team-game-stats"
	NeedTeamGameLeaders   DataNeed = "team-game-leaders"
	NeedPlayByPlay        DataNeed = "play-by-play"
	NeedDrives            DataNeed = "drives"
	NeedWinProbability    DataNeed = "win-probability"
	NeedOdds              DataNeed = "odds"
	NeedVenue             DataNeed = "venue"
	NeedStandings         DataNeed = "standings"
	NeedInjuries          DataNeed = "injuries"
	NeedTransactions      DataNeed = "transactions"
	NeedNews              DataNeed = "news"
	NeedDraft             DataNeed = "draft"
)

// SourcePreference documents the preferred source and fallback/companion source
// order from sdk_chaining_matrix.md.
type SourcePreference struct {
	Need      DataNeed
	Primary   Source
	Fallbacks []Source
}

// SourcePreferences returns the football source policy table.
func SourcePreferences() map[DataNeed]SourcePreference {
	return map[DataNeed]SourcePreference{
		NeedTeamLookup:        {Need: NeedTeamLookup, Primary: SourceSiteV2, Fallbacks: []Source{SourceSearchV2, SourceCoreV2}},
		NeedTeamProfile:       {Need: NeedTeamProfile, Primary: SourceSiteV2, Fallbacks: []Source{SourceCoreV2, SourceCoreV3}},
		NeedTeamSchedule:      {Need: NeedTeamSchedule, Primary: SourceSiteV2, Fallbacks: []Source{SourceSiteV2}},
		NeedTeamRoster:        {Need: NeedTeamRoster, Primary: SourceSiteV2, Fallbacks: []Source{SourceCommonV3, SourceCoreV2}},
		NeedDepthChart:        {Need: NeedDepthChart, Primary: SourceSiteV2},
		NeedPlayerLookup:      {Need: NeedPlayerLookup, Primary: SourceSearchV2, Fallbacks: []Source{SourceSiteV2, SourceCommonV3}},
		NeedPlayerProfile:     {Need: NeedPlayerProfile, Primary: SourceCommonV3, Fallbacks: []Source{SourceCoreV2, SourceCoreV3}},
		NeedPlayerSeasonStats: {Need: NeedPlayerSeasonStats, Primary: SourceCommonV3, Fallbacks: []Source{SourceCoreV2}},
		NeedPlayerGameLog:     {Need: NeedPlayerGameLog, Primary: SourceCommonV3, Fallbacks: []Source{SourceCoreV2, SourceCoreV3}},
		NeedPlayerSplits:      {Need: NeedPlayerSplits, Primary: SourceCommonV3, Fallbacks: []Source{SourceCoreV2}},
		NeedPlayerGameStats:   {Need: NeedPlayerGameStats, Primary: SourceCoreV2, Fallbacks: []Source{SourceSiteV2}},
		NeedLeagueLeaders:     {Need: NeedLeagueLeaders, Primary: SourceCommonV3, Fallbacks: []Source{SourceSiteV2}},
		NeedTeamStatRankings:  {Need: NeedTeamStatRankings, Primary: SourceCommonV3, Fallbacks: []Source{SourceSiteV2}},
		NeedScoreboard:        {Need: NeedScoreboard, Primary: SourceSiteV2, Fallbacks: []Source{SourceCDN}},
		NeedGameSummary:       {Need: NeedGameSummary, Primary: SourceSiteV2, Fallbacks: []Source{SourceCDN, SourceCoreV2}},
		NeedGameMetadata:      {Need: NeedGameMetadata, Primary: SourceCoreV2, Fallbacks: []Source{SourceSiteV2}},
		NeedTeamGameStats:     {Need: NeedTeamGameStats, Primary: SourceCoreV2, Fallbacks: []Source{SourceSiteV2}},
		NeedTeamGameLeaders:   {Need: NeedTeamGameLeaders, Primary: SourceCoreV2, Fallbacks: []Source{SourceSiteV2}},
		NeedPlayByPlay:        {Need: NeedPlayByPlay, Primary: SourceCoreV2, Fallbacks: []Source{SourceSiteV2, SourceCDN}},
		NeedDrives:            {Need: NeedDrives, Primary: SourceCoreV2, Fallbacks: []Source{SourceSiteV2, SourceCDN}},
		NeedWinProbability:    {Need: NeedWinProbability, Primary: SourceCoreV2, Fallbacks: []Source{SourceSiteV2}},
		NeedOdds:              {Need: NeedOdds, Primary: SourceCoreV2, Fallbacks: []Source{SourceSiteV2}},
		NeedVenue:             {Need: NeedVenue, Primary: SourceCoreV2, Fallbacks: []Source{SourceCoreV3}},
		NeedStandings:         {Need: NeedStandings, Primary: SourceSiteAPI, Fallbacks: []Source{SourceSiteV2}},
		NeedInjuries:          {Need: NeedInjuries, Primary: SourceSiteV2},
		NeedTransactions:      {Need: NeedTransactions, Primary: SourceSiteV2},
		NeedNews:              {Need: NeedNews, Primary: SourceSiteV2, Fallbacks: []Source{SourceNow, SourceSearchV2}},
		NeedDraft:             {Need: NeedDraft, Primary: SourceSiteV2, Fallbacks: []Source{SourceCoreV2}},
	}
}

func (s *Service) SourcePreference(need DataNeed) (SourcePreference, bool) {
	preference, ok := SourcePreferences()[need]
	return preference, ok
}

type sourceAttempt struct {
	source Source
	call   func(context.Context) (models.JSON, error)
}

func (s *Service) trySources(ctx context.Context, need DataNeed, attempts ...sourceAttempt) (models.JSON, Source, error) {
	var lastErr error
	for _, attempt := range orderedAttempts(need, attempts) {
		payload, err := attempt.call(ctx)
		if err == nil && !emptyPayload(payload) {
			return payload, attempt.source, nil
		}
		if err != nil {
			lastErr = err
		} else {
			lastErr = fmt.Errorf("%s returned empty payload", attempt.source)
		}
	}
	if lastErr == nil {
		lastErr = fmt.Errorf("no source attempts configured for %s", need)
	}
	return nil, "", lastErr
}

func orderedAttempts(need DataNeed, attempts []sourceAttempt) []sourceAttempt {
	preference, ok := SourcePreferences()[need]
	if !ok {
		return attempts
	}
	order := append([]Source{preference.Primary}, preference.Fallbacks...)
	used := make([]bool, len(attempts))
	var ordered []sourceAttempt
	for _, source := range order {
		for i, attempt := range attempts {
			if used[i] || attempt.source != source {
				continue
			}
			ordered = append(ordered, attempt)
			used[i] = true
		}
	}
	for i, attempt := range attempts {
		if !used[i] {
			ordered = append(ordered, attempt)
		}
	}
	return ordered
}

func emptyPayload(payload models.JSON) bool {
	if len(payload) == 0 {
		return true
	}
	for _, value := range payload {
		if value != nil {
			return false
		}
	}
	return true
}
