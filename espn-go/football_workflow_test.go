package espn_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	espn "github.com/pseudo-r/Public-ESPN-API/espn-go"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/football"
)

func TestFootballWorkflowResolvesIDsAndStats(t *testing.T) {
	ctx := context.Background()
	seen := map[string]bool{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seen[r.URL.Path] = true
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/apis/search/v2":
			if r.URL.Query().Get("query") != "Drake Maye" || r.URL.Query().Get("sport") != "football" || r.URL.Query().Get("league") != "nfl" {
				t.Fatalf("unexpected search query: %s", r.URL.RawQuery)
			}
			writeJSON(t, w, map[string]any{
				"results": []any{
					map[string]any{
						"type": "player",
						"contents": []any{
							map[string]any{
								"id":          "2fbd2f9e-624d-3e63-9bb5-9bb965782b68",
								"uid":         "s:20~l:28~a:4431452",
								"type":        "player",
								"displayName": "Drake Maye",
								"description": "NFL",
							},
						},
					},
				},
			})
		case "/apis/site/v2/sports/football/nfl/teams":
			writeJSON(t, w, map[string]any{
				"sports": []any{
					map[string]any{
						"leagues": []any{
							map[string]any{
								"teams": []any{
									map[string]any{
										"team": map[string]any{
											"id":           "17",
											"abbreviation": "NE",
											"displayName":  "New England Patriots",
											"location":     "New England",
											"nickname":     "Patriots",
										},
									},
								},
							},
						},
					},
				},
			})
		case "/apis/site/v2/sports/football/nfl/summary":
			if r.URL.Query().Get("event") != "401772988" {
				t.Fatalf("unexpected summary query: %s", r.URL.RawQuery)
			}
			writeJSON(t, w, summaryFixture())
		case "/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17/roster/4431452/statistics/0":
			writeJSON(t, w, map[string]any{
				"splits": map[string]any{
					"categories": []any{
						map[string]any{
							"name": "passing",
							"stats": []any{
								map[string]any{
									"name":         "passingYards",
									"displayName":  "Passing Yards",
									"value":        287,
									"displayValue": "287",
								},
							},
						},
					},
				},
			})
		default:
			t.Fatalf("unexpected path: %s", r.URL.String())
		}
	}))
	defer server.Close()

	client := espn.NewClient(espn.WithBaseURLs(testBaseURLs(server.URL)))
	football := client.Football("nfl")

	player, err := football.ResolvePlayerID(ctx, "Drake Maye")
	if err != nil {
		t.Fatalf("ResolvePlayerID returned error: %v", err)
	}
	if player.AthleteID != "4431452" || player.GUID == "" {
		t.Fatalf("unexpected player hit: %+v", player)
	}

	team, err := football.ResolveTeamID(ctx, "Patriots")
	if err != nil {
		t.Fatalf("ResolveTeamID returned error: %v", err)
	}
	if team.TeamID != "17" {
		t.Fatalf("unexpected team hit: %+v", team)
	}

	stats, err := football.GetPlayerGameStats(ctx, espn.PlayerGameStatsInput{
		EventID:    "401772988",
		AthleteID:  "4431452",
		PlayerName: "Drake Maye",
	})
	if err != nil {
		t.Fatalf("GetPlayerGameStats returned error: %v", err)
	}
	if stats.CompetitionID != "401772988" || stats.CompetitorID != "17" || len(stats.Stats) != 1 {
		t.Fatalf("unexpected stats: %+v", stats)
	}
	if stats.Stats[0].Name != "passingYards" || stats.Stats[0].DisplayValue != "287" {
		t.Fatalf("unexpected stat line: %+v", stats.Stats[0])
	}

	for _, path := range []string{
		"/apis/search/v2",
		"/apis/site/v2/sports/football/nfl/teams",
		"/apis/site/v2/sports/football/nfl/summary",
		"/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17/roster/4431452/statistics/0",
	} {
		if !seen[path] {
			t.Fatalf("expected path was not called: %s", path)
		}
	}
}

func TestFootballStandingsUsesPopulatedSiteAPIPath(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/apis/v2/sports/football/nfl/standings" {
			t.Fatalf("unexpected path: %s", r.URL.String())
		}
		writeJSON(t, w, map[string]any{"children": []any{map[string]any{"name": "AFC"}}})
	}))
	defer server.Close()

	client := espn.NewClient(espn.WithBaseURLs(testBaseURLs(server.URL)))
	standings, err := client.Football("nfl").GetStandings(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetStandings returned error: %v", err)
	}
	if len(standings["children"].([]any)) != 1 {
		t.Fatalf("unexpected standings payload: %+v", standings)
	}
}

func TestRawDomainClientsRemainAvailable(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/core/nfl/boxscore":
			if r.URL.Query().Get("xhr") != "1" || r.URL.Query().Get("gameId") != "401772988" {
				t.Fatalf("unexpected cdn query: %s", r.URL.RawQuery)
			}
			writeJSON(t, w, map[string]any{"type": "cdn-boxscore"})
		case "/apis/v2/scoreboard/header":
			if r.URL.Query().Get("sport") != "football" || r.URL.Query().Get("league") != "nfl" {
				t.Fatalf("unexpected header query: %s", r.URL.RawQuery)
			}
			writeJSON(t, w, map[string]any{"header": "ok"})
		default:
			t.Fatalf("unexpected path: %s", r.URL.String())
		}
	}))
	defer server.Close()

	football := espn.NewClient(espn.WithBaseURLs(testBaseURLs(server.URL))).Football("nfl")
	if _, err := football.CDN.Boxscore(context.Background(), "401772988"); err != nil {
		t.Fatalf("CDN.Boxscore returned error: %v", err)
	}
	if _, err := football.Search.ScoreboardHeader(context.Background(), url.Values{}); err != nil {
		t.Fatalf("Search.ScoreboardHeader returned error: %v", err)
	}
}

func TestAdditionalHierarchyMethodsExposeDocumentedPaths(t *testing.T) {
	expected := map[string]bool{
		"/apis/site/v2/sports/football/nfl/calendar":                                                          false,
		"/apis/site/v2/sports/football/nfl/rankings":                                                          false,
		"/apis/site/v2/sports/football/nfl/teams/6/record":                                                    false,
		"/apis/site/v2/sports/football/nfl/teams/6/news":                                                      false,
		"/apis/site/v2/sports/football/nfl/teams/6/leaders":                                                   false,
		"/apis/site/v2/sports/football/nfl/teams/6/history":                                                   false,
		"/apis/site/v2/sports/football/nfl/teams/6/transactions":                                              false,
		"/v2/sports/football/leagues/nfl/seasons/2025/athletes/4431452":                                       false,
		"/v2/sports/football/leagues/nfl/seasons/2025/draft/status":                                           false,
		"/v2/sports/football/leagues/nfl/seasons/2025/manufacturers":                                          false,
		"/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/details":                     false,
		"/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/notes":                       false,
		"/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/powerindex":                  false,
		"/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/powerindex/17":               false,
		"/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/plays/40177298840/personnel": false,
		"/v2/sports/football/leagues/nfl/casinos":                                                             false,
		"/v2/sports/football/leagues/nfl/casinos/1":                                                           false,
		"/v2/sports/football/leagues/nfl/circuits":                                                            false,
		"/v2/sports/football/leagues/nfl/circuits/1":                                                          false,
		"/v2/sports/football/leagues/nfl/countries":                                                           false,
		"/v2/sports/football/leagues/nfl/countries/1":                                                         false,
		"/v2/sports/football/leagues/nfl/franchises":                                                          false,
		"/v2/sports/football/leagues/nfl/franchises/1":                                                        false,
		"/v2/sports/football/leagues/nfl/positions":                                                           false,
		"/v2/sports/football/leagues/nfl/positions/1":                                                         false,
		"/v2/sports/football/leagues/nfl/providers":                                                           false,
		"/v2/sports/football/leagues/nfl/providers/1":                                                         false,
		"/v2/sports/football/leagues/nfl/rankings":                                                            false,
		"/v2/sports/football/leagues/nfl/recruiting":                                                          false,
		"/v2/sports/football/leagues/nfl/tournaments":                                                         false,
		"/v3/sports/football/nfl":                                                                             false,
		"/v3/sports/football/nfl/seasons/2025":                                                                false,
		"/v3/sports/football/nfl/teams/6":                                                                     false,
		"/v3/sports/football/nfl/athletes":                                                                    false,
		"/v3/sports/football/nfl/athletes/4431452":                                                            false,
		"/v3/sports/football/nfl/athletes/4431452/plays":                                                      false,
		"/v3/sports/football/nfl/athletes/4431452/statisticslog":                                              false,
		"/apis/common/v3/sports/football/nfl/athletes":                                                        false,
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := expected[r.URL.Path]; !ok {
			t.Fatalf("unexpected path: %s", r.URL.String())
		}
		expected[r.URL.Path] = true
		writeJSON(t, w, map[string]any{"ok": true})
	}))
	defer server.Close()

	football := espn.NewClient(espn.WithBaseURLs(testBaseURLs(server.URL))).Football("nfl")
	ctx := context.Background()
	calls := []func() error{
		func() error { _, err := football.Site.Calendar(ctx, nil); return err },
		func() error { _, err := football.Site.Rankings(ctx, nil); return err },
		func() error { _, err := football.Site.TeamRecord(ctx, "6", nil); return err },
		func() error { _, err := football.Site.TeamNews(ctx, "6", nil); return err },
		func() error { _, err := football.Site.TeamLeaders(ctx, "6", nil); return err },
		func() error { _, err := football.Site.TeamHistory(ctx, "6", nil); return err },
		func() error { _, err := football.Site.TeamTransactions(ctx, "6", nil); return err },
		func() error { _, err := football.Core.SeasonAthlete(ctx, "2025", "4431452", nil); return err },
		func() error { _, err := football.Core.DraftStatus(ctx, "2025", nil); return err },
		func() error { _, err := football.Core.Manufacturers(ctx, "2025", nil); return err },
		func() error { _, err := football.Core.Details(ctx, "401772988", "401772988", nil); return err },
		func() error { _, err := football.Core.Notes(ctx, "401772988", "401772988", nil); return err },
		func() error { _, err := football.Core.PowerIndex(ctx, "401772988", "401772988", nil); return err },
		func() error {
			_, err := football.Core.PowerIndexTeam(ctx, "401772988", "401772988", "17", nil)
			return err
		},
		func() error {
			_, err := football.Core.PlayPersonnel(ctx, "401772988", "401772988", "40177298840", nil)
			return err
		},
		func() error { _, err := football.Core.Casinos(ctx, nil); return err },
		func() error { _, err := football.Core.Casino(ctx, "1", nil); return err },
		func() error { _, err := football.Core.Circuits(ctx, nil); return err },
		func() error { _, err := football.Core.Circuit(ctx, "1", nil); return err },
		func() error { _, err := football.Core.Countries(ctx, nil); return err },
		func() error { _, err := football.Core.Country(ctx, "1", nil); return err },
		func() error { _, err := football.Core.Franchises(ctx, nil); return err },
		func() error { _, err := football.Core.Franchise(ctx, "1", nil); return err },
		func() error { _, err := football.Core.Positions(ctx, nil); return err },
		func() error { _, err := football.Core.Position(ctx, "1", nil); return err },
		func() error { _, err := football.Core.Providers(ctx, nil); return err },
		func() error { _, err := football.Core.Provider(ctx, "1", nil); return err },
		func() error { _, err := football.Core.Rankings(ctx, nil); return err },
		func() error { _, err := football.Core.Recruiting(ctx, nil); return err },
		func() error { _, err := football.Core.Tournaments(ctx, nil); return err },
		func() error { _, err := football.Core.V3League(ctx, nil); return err },
		func() error { _, err := football.Core.V3Season(ctx, "2025", nil); return err },
		func() error { _, err := football.Core.V3Team(ctx, "6", nil); return err },
		func() error { _, err := football.Core.V3Athletes(ctx, nil); return err },
		func() error { _, err := football.Core.V3Athlete(ctx, "4431452", nil); return err },
		func() error { _, err := football.Core.V3AthletePlays(ctx, "4431452", nil); return err },
		func() error { _, err := football.Core.V3AthleteStatisticsLog(ctx, "4431452", nil); return err },
		func() error { _, err := football.Common.Athletes(ctx, nil); return err },
	}
	for _, call := range calls {
		if err := call(); err != nil {
			t.Fatalf("hierarchy call returned error: %v", err)
		}
	}
	for path, called := range expected {
		if !called {
			t.Fatalf("expected path was not called: %s", path)
		}
	}
}

func TestAdditionalCapabilityWorkflows(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/apis/site/v2/sports/football/nfl/teams/17/schedule":
			writeJSON(t, w, map[string]any{
				"events": []any{
					scheduleEvent("old", "2025-01-01T18:00:00Z", true, "17", "6"),
					scheduleEvent("401772988", "2025-10-05T18:00:00Z", false, "17", "6"),
					scheduleEvent("later", "2025-10-12T18:00:00Z", false, "17", "8"),
				},
			})
		case "/apis/site/v2/sports/football/nfl/summary":
			if r.URL.Query().Get("event") != "401772988" {
				t.Fatalf("unexpected summary query: %s", r.URL.RawQuery)
			}
			writeJSON(t, w, summaryFixture())
		case "/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988":
			writeJSON(t, w, map[string]any{
				"venue": map[string]any{
					"$ref": serverURL(r) + "/v2/sports/football/leagues/nfl/venues/4738?lang=en&region=us",
				},
			})
		case "/v2/sports/football/leagues/nfl/venues/4738":
			writeJSON(t, w, map[string]any{"id": "4738", "fullName": "Gillette Stadium"})
		default:
			t.Fatalf("unexpected path: %s", r.URL.String())
		}
	}))
	defer server.Close()

	football := espn.NewClient(espn.WithBaseURLs(testBaseURLs(server.URL))).Football("nfl")
	ctx := context.Background()
	asOf := time.Date(2025, 10, 2, 0, 0, 0, 0, time.UTC)

	next, err := football.GetTeamNextGame(ctx, "17", "", asOf, nil)
	if err != nil {
		t.Fatalf("GetTeamNextGame returned error: %v", err)
	}
	if next["id"] != "401772988" {
		t.Fatalf("unexpected next game: %+v", next)
	}

	prev, err := football.GetTeamPreviousGame(ctx, "17", "", asOf, nil)
	if err != nil {
		t.Fatalf("GetTeamPreviousGame returned error: %v", err)
	}
	if prev["id"] != "old" {
		t.Fatalf("unexpected previous game: %+v", prev)
	}

	game, err := football.ResolveGameForTeam(ctx, espn.GameSearch{
		TeamID:     "17",
		Date:       "2025-10-05",
		OpponentID: "6",
	})
	if err != nil {
		t.Fatalf("ResolveGameForTeam returned error: %v", err)
	}
	if game.EventID != "401772988" || game.CompetitionID != "401772988" {
		t.Fatalf("unexpected game resolution: %+v", game)
	}

	venue, err := football.GetVenue(ctx, "", "401772988", "", "", "", nil)
	if err != nil {
		t.Fatalf("GetVenue returned error: %v", err)
	}
	if venue["id"] != "4738" {
		t.Fatalf("unexpected venue: %+v", venue)
	}
}

func TestSourcePreferencesExposeDocumentedOrder(t *testing.T) {
	preferences := football.SourcePreferences()
	tests := map[football.DataNeed]struct {
		primary   football.Source
		fallbacks []football.Source
	}{
		football.NeedPlayerProfile:    {football.SourceCommonV3, []football.Source{football.SourceCoreV2, football.SourceCoreV3}},
		football.NeedTeamProfile:      {football.SourceSiteV2, []football.Source{football.SourceCoreV2, football.SourceCoreV3}},
		football.NeedPlayByPlay:       {football.SourceCoreV2, []football.Source{football.SourceSiteV2, football.SourceCDN}},
		football.NeedStandings:        {football.SourceSiteAPI, []football.Source{football.SourceSiteV2}},
		football.NeedLeagueLeaders:    {football.SourceCommonV3, []football.Source{football.SourceSiteV2}},
		football.NeedWinProbability:   {football.SourceCoreV2, []football.Source{football.SourceSiteV2}},
		football.NeedTeamRoster:       {football.SourceSiteV2, []football.Source{football.SourceCommonV3, football.SourceCoreV2}},
		football.NeedPlayerGameStats:  {football.SourceCoreV2, []football.Source{football.SourceSiteV2}},
		football.NeedTeamStatRankings: {football.SourceCommonV3, []football.Source{football.SourceSiteV2}},
		football.NeedNews:             {football.SourceSiteV2, []football.Source{football.SourceNow, football.SourceSearchV2}},
	}

	for need, want := range tests {
		got, ok := preferences[need]
		if !ok {
			t.Fatalf("missing preference for %s", need)
		}
		if got.Primary != want.primary {
			t.Fatalf("%s primary mismatch: want %s got %s", need, want.primary, got.Primary)
		}
		if len(got.Fallbacks) != len(want.fallbacks) {
			t.Fatalf("%s fallback length mismatch: want %v got %v", need, want.fallbacks, got.Fallbacks)
		}
		for i, fallback := range want.fallbacks {
			if got.Fallbacks[i] != fallback {
				t.Fatalf("%s fallback[%d] mismatch: want %s got %s", need, i, fallback, got.Fallbacks[i])
			}
		}
	}
}

func TestFallbacksUseDocumentedSourceOrder(t *testing.T) {
	var calls []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls = append(calls, r.URL.Path)
		switch r.URL.Path {
		case "/apis/common/v3/sports/football/nfl/athletes/4431452/stats":
			http.Error(w, "common unavailable", http.StatusNotFound)
		case "/v2/sports/football/leagues/nfl/athletes/4431452/statistics":
			writeJSON(t, w, map[string]any{"source": "core-v2"})
		case "/apis/site/v2/sports/football/nfl/scoreboard":
			writeJSON(t, w, map[string]any{})
		case "/core/nfl/scoreboard":
			writeJSON(t, w, map[string]any{"source": "cdn"})
		default:
			t.Fatalf("unexpected path: %s", r.URL.String())
		}
	}))
	defer server.Close()

	footballClient := espn.NewClient(espn.WithBaseURLs(testBaseURLs(server.URL))).Football("nfl")
	stats, err := footballClient.GetPlayerSeasonStats(context.Background(), "4431452", "", nil)
	if err != nil {
		t.Fatalf("GetPlayerSeasonStats returned error: %v", err)
	}
	if stats["source"] != "core-v2" {
		t.Fatalf("expected core fallback payload, got %+v", stats)
	}

	scoreboard, err := footballClient.GetScoreboard(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetScoreboard returned error: %v", err)
	}
	if scoreboard["source"] != "cdn" {
		t.Fatalf("expected cdn fallback payload, got %+v", scoreboard)
	}

	want := []string{
		"/apis/common/v3/sports/football/nfl/athletes/4431452/stats",
		"/v2/sports/football/leagues/nfl/athletes/4431452/statistics",
		"/apis/site/v2/sports/football/nfl/scoreboard",
		"/core/nfl/scoreboard",
	}
	if !sameStrings(calls, want) {
		t.Fatalf("unexpected call order\nwant: %v\n got: %v", want, calls)
	}
}

func TestAdditionalChainHelpers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/apis/site/v2/sports/football/nfl/summary":
			writeJSON(t, w, summaryFixture())
		case "/v2/sports/football/leagues/nfl/events/401772988":
			writeJSON(t, w, map[string]any{"id": "401772988", "source": "core-event"})
		case "/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/drives/4017729881":
			writeJSON(t, w, map[string]any{"id": "4017729881"})
		case "/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/drives/4017729881/plays":
			writeJSON(t, w, map[string]any{"items": []any{map[string]any{"id": "40177298840"}}})
		case "/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/plays/40177298840":
			writeJSON(t, w, map[string]any{"id": "40177298840", "text": "pass complete"})
		case "/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/probabilities/40177298840":
			writeJSON(t, w, map[string]any{"id": "40177298840", "homeWinPercentage": 0.51})
		case "/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/situation":
			writeJSON(t, w, map[string]any{
				"lastPlay": map[string]any{
					"$ref": serverURL(r) + "/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/plays/40177298840?lang=en&region=us",
				},
			})
		case "/v2/sports/football/leagues/nfl/providers/58":
			writeJSON(t, w, map[string]any{"id": "58", "name": "ESPN BET"})
		case "/apis/site/v2/sports/football/nfl/injuries":
			writeJSON(t, w, map[string]any{
				"injuries": []any{
					map[string]any{"team": map[string]any{"id": "17"}, "athlete": map[string]any{"id": "1"}},
					map[string]any{"team": map[string]any{"id": "6"}, "athlete": map[string]any{"id": "2"}},
				},
			})
		case "/apis/site/v2/sports/football/nfl/transactions":
			writeJSON(t, w, map[string]any{
				"transactions": []any{
					map[string]any{"team": map[string]any{"id": "17"}, "id": "a"},
					map[string]any{"team": map[string]any{"id": "6"}, "id": "b"},
				},
			})
		default:
			t.Fatalf("unexpected path: %s", r.URL.String())
		}
	}))
	defer server.Close()

	footballClient := espn.NewClient(espn.WithBaseURLs(testBaseURLs(server.URL))).Football("nfl")
	ctx := context.Background()

	metadata, err := footballClient.GetGameMetadata(ctx, "401772988", "", nil)
	if err != nil {
		t.Fatalf("GetGameMetadata returned error: %v", err)
	}
	if metadata["source"] != "core-event" {
		t.Fatalf("unexpected metadata: %+v", metadata)
	}

	for name, call := range map[string]func() (espn.JSON, error){
		"drive": func() (espn.JSON, error) {
			return footballClient.GetDrive(ctx, "401772988", "", "4017729881", nil)
		},
		"drive plays": func() (espn.JSON, error) {
			return footballClient.GetDrivePlays(ctx, "401772988", "", "4017729881", nil)
		},
		"play": func() (espn.JSON, error) {
			return footballClient.GetPlay(ctx, "401772988", "", "40177298840", nil)
		},
		"probability": func() (espn.JSON, error) {
			return footballClient.GetProbability(ctx, "401772988", "", "40177298840", nil)
		},
		"last play": func() (espn.JSON, error) {
			return footballClient.GetLiveLastPlay(ctx, "401772988", "", nil)
		},
		"provider": func() (espn.JSON, error) {
			return footballClient.GetGameOddsProvider(ctx, "58", nil)
		},
	} {
		if _, err := call(); err != nil {
			t.Fatalf("%s returned error: %v", name, err)
		}
	}

	injuries, err := footballClient.GetTeamInjuries(ctx, "17", "", nil)
	if err != nil {
		t.Fatalf("GetTeamInjuries returned error: %v", err)
	}
	if got := len(injuries["injuries"].([]any)); got != 1 {
		t.Fatalf("expected one injury after filtering, got %d", got)
	}
	transactions, err := footballClient.GetTeamTransactions(ctx, "17", "", nil)
	if err != nil {
		t.Fatalf("GetTeamTransactions returned error: %v", err)
	}
	if got := len(transactions["transactions"].([]any)); got != 1 {
		t.Fatalf("expected one transaction after filtering, got %d", got)
	}
}

func testBaseURLs(base string) espn.BaseURLs {
	return espn.BaseURLs{Site: base, Core: base, Web: base, CDN: base, Now: base}
}

func writeJSON(t *testing.T, w http.ResponseWriter, value any) {
	t.Helper()
	if err := json.NewEncoder(w).Encode(value); err != nil {
		t.Fatalf("encode fixture: %v", err)
	}
}

func summaryFixture() map[string]any {
	return map[string]any{
		"header": map[string]any{
			"competitions": []any{
				map[string]any{
					"id": "401772988",
					"competitors": []any{
						map[string]any{
							"id":   "17",
							"team": map[string]any{"id": "17"},
						},
						map[string]any{
							"id":   "6",
							"team": map[string]any{"id": "6"},
						},
					},
				},
			},
		},
		"boxscore": map[string]any{
			"teams": []any{
				map[string]any{"team": map[string]any{"id": "17"}},
				map[string]any{"team": map[string]any{"id": "6"}},
			},
			"players": []any{
				map[string]any{
					"team": map[string]any{"id": "17"},
					"statistics": []any{
						map[string]any{
							"athletes": []any{
								map[string]any{
									"athlete": map[string]any{"id": "4431452"},
								},
							},
						},
					},
				},
			},
		},
	}
}

func scheduleEvent(id, date string, completed bool, teamIDs ...string) map[string]any {
	competitors := make([]any, 0, len(teamIDs))
	for _, id := range teamIDs {
		competitors = append(competitors, map[string]any{
			"team": map[string]any{
				"id":           id,
				"displayName":  "Team " + id,
				"abbreviation": id,
			},
		})
	}
	return map[string]any{
		"id":   id,
		"date": date,
		"status": map[string]any{
			"type": map[string]any{"completed": completed},
		},
		"competitions": []any{
			map[string]any{"competitors": competitors},
		},
	}
}

func serverURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + r.Host
}

func sameStrings(got, want []string) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}
