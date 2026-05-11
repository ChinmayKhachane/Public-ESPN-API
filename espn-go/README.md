# espn-go

Go SDK for ESPN's unofficial public API. This module is scaffolded separately
from the Django service and follows the endpoint families documented under
`docs/sports/football_response`.

## Install

```bash
go get github.com/pseudo-r/Public-ESPN-API/espn-go
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"

	espn "github.com/pseudo-r/Public-ESPN-API/espn-go"
)

func main() {
	ctx := context.Background()
	client := espn.NewClient()
	football := client.Football("nfl")

	player, err := football.ResolvePlayerID(ctx, "Drake Maye")
	if err != nil {
		panic(err)
	}

	profile, err := football.GetPlayerProfile(ctx, player.AthleteID, "")
	if err != nil {
		panic(err)
	}

	fmt.Println(profile["athlete"])
}
```

## Project Shape

The SDK is split by ESPN API family and by data model responsibility:

```text
espn-go/
  espn.go                 root facade and shared options
  endpoints/              URL resolver with ESPN base domains and path rules
  internal/httpclient/    shared HTTP transport and JSON/error handling
  models/                 exported structured SDK types and raw JSON helpers
  site/                   site.api.espn.com site-v2 and standings routes
  core/                   sports.core.api.espn.com core-v2/core-v3 graph routes
  common/                 site.web.api.espn.com common-v3 profile/stat routes
  cdn/                    cdn.espn.com page-package game routes
  search/                 search-v2 and scoreboard-header routes
  now/                    now.core.api.espn.com news routes
  football/               football/NFL workflows that chain raw clients
```

`endpoints.Resolver` owns URL construction. Domain clients call resolver
methods instead of assembling hostnames, which keeps ESPN path quirks in one
place. For example, football standings use
`/apis/v2/sports/football/nfl/standings`, while normal site-v2 resources use
`/apis/site/v2/sports/football/nfl/...`.

## Football SDK

The football service exposes raw clients:

- `football.Site`: scoreboard, summary, teams, rosters, standings, news
- `football.Core`: core-v2 graph resources, competitions, plays, drives, odds
- `football.Common`: common-v3 athlete and leaderboard views
- `football.CDN`: CDN page-package game payloads
- `football.Search`: Search v2 and scoreboard header
- `football.News`: Now API headlines

It also exposes workflow methods that hide documented resolver chains:

- `SearchPlayers`, `ResolvePlayerID`, `ResolveTeamID`
- `GetPlayerProfile`, `GetPlayerSeasonStats`, `GetPlayerGameLog`
- `GetPlayerGameStats`
- `GetTeamRoster`, `GetTeamSchedule`, `GetTeamGameStats`
- `GetTeamNextGame`, `GetTeamPreviousGame`
- `GetScoreboard`, `ResolveGame`, `GetGameSummary`
- `GetPlayByPlay`, `GetDrives`, `GetWinProbability`
- `GetVenue`
- `GetStandings`, `GetLeagueLeaders`, `ResolveIDsBundle`

The raw clients expose the endpoint hierarchy from
`docs/sports/football_response/api_hierarchy.md`, including sparse or
unsupported documented paths. Those methods return ESPN JSON directly and let
the caller decide how to handle empty or error payloads. The football service
implements the user-level workflows from `sdk_capabilities.md` where the docs
describe resolver chains, such as team schedules to event IDs, summary to
competition/competitor IDs, and event/team IDs to venue or stat resources.

## Source Order

Football workflows use an explicit source policy from
`football.SourcePreferences()`. The policy follows the recommended source table
from `docs/sports/football_response/sdk_chaining_matrix.md`; for example:

- Player profile: common-v3, then core-v2, then core-v3.
- Team profile: site-v2, then core-v2, then core-v3.
- Scoreboard: site-v2, then CDN scoreboard.
- Play-by-play: core-v2, then site-v2 summary, then CDN play-by-play.
- Standings: populated site API `/apis/v2/.../standings`, then the site-v2 stub.

The workflow methods try the preferred source first and fall back when the
response errors or returns an empty object.

## Notes

- Search v2 player GUIDs are kept as metadata. The canonical player path ID is
  parsed from `uid` using `~a:(\d+)`.
- The SDK does not recursively expand ESPN `$ref` graphs by default. Use raw
  domain clients to follow explicit child resources.
- Add future sports by creating a sport workflow package over the same raw API
  family packages, then add sport-specific model types only where normalization
  is needed.
