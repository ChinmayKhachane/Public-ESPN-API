# Football SDK Chaining Matrix

Verified against the live NFL response docs and spot-checked with live requests
on 2026-05-09.

This file describes the football/NFL endpoint chains that are useful for an SDK.
It intentionally focuses on data that satisfies both criteria:

- The endpoint works and returns populated data in the documented NFL examples.
- The endpoint exposes a useful API-specific response shape, not just duplicate
  IDs, names, or scores that are available everywhere.

Simple identifiers are duplicated across ESPN domains. In this file,
"API-specific" means "this is the documented working source for this specific
shape." For example, `site-v2` summary and CDN game packages both expose game
data, but their schemas are different enough that both can be useful SDK
surfaces.

Fixtures used by the live football docs:

- `league=nfl`
- `season=2025`
- `event=401772988`
- `competition=401772988`
- `competitor=17`
- `team id=6` for site examples
- `team id=22` for core season-team examples
- `athlete id=4431452` for game/player stat examples
- `athlete id=4429202` for some athlete profile examples
- `drive=4017729881`
- `play=40177298840`

## Resolver Fields

SDK code should know these resolver fields before the request. Avoid traversing
arbitrary JSON looking for URLs at runtime unless the caller explicitly asks for
deep graph expansion.

| Source response | Field to follow | Resolves to | Use in SDK |
| --- | --- | --- | --- |
| Core-v2 paginated collections | `items[].$ref` | Detail resource for one item | Optional `resolveItems` behavior. |
| Core-v2 league root | named child refs such as `season`, `teams`, `athletes`, `events`, `calendar` | League child collections or details | SDK entrypoint discovery. |
| Core-v2 event detail | `competitions[].$ref` | Competition detail | Game detail chain. |
| Core-v2 competition detail | `competitors[].$ref` | Event-specific competitor detail | Team-in-game chain. |
| Core-v2 competition detail | `venue.$ref` | Venue detail | Stadium enrichment. |
| Core-v2 situation detail | `lastPlay.$ref` | Play detail | Live-game last-play chain. |
| Core-v2 probability detail | `play.$ref` | Play detail | Win-probability-to-play chain. |
| Site-v2 scoreboard | `events[].id` | Event ID | Start of game chain. |
| Site-v2 summary | `header.competitions[].id` | Competition ID | Usually same as event ID for NFL. |
| Site-v2 summary | `boxscore.teams[].team.id` | Competitor/team ID | Use as `competitor` in core-v2 event paths. |
| Site-v2 summary | `boxscore.players[].team.id`, player/stat rows | Team and athlete IDs | Player game-stat chain. |
| Site-v2 team schedule | `events[].id` | Event ID | Team season schedule to game chain. |
| Site-v2/site-v3 rosters | athlete IDs in roster entries | Athlete ID | Player profile/stat chains. |
| Search v2 player hits | `results[].contents[].uid` | Numeric athlete ID after `~a:` | Name-to-athlete-ID discovery. |
| Search v2 non-player hits | `results[].contents[].id` and links | Story/video IDs or web links | Query-to-content discovery. |

## Player Search And ID Persistence

Live tests on 2026-05-09 showed that Search v2 can be used for player
name-to-ID discovery, but the field choice matters.

Search endpoint:

```text
https://site.web.api.espn.com/apis/search/v2?query=drake%20maye
```

Representative player hit:

```json
{
  "id": "2fbd2f9e-624d-3e63-9bb5-9bb965782b68",
  "uid": "s:20~l:28~a:4431452",
  "type": "player",
  "displayName": "Drake Maye",
  "description": "NFL"
}
```

For player hits, `contents[].id` is the athlete GUID. It is useful to store as
metadata, but it is not the path ID accepted by the documented athlete
endpoints. The numeric ESPN athlete ID is embedded in `contents[].uid` after
`~a:`.

```text
uid: s:20~l:28~a:4431452
athleteId: 4431452
```

The SDK should parse athlete IDs from Search v2 player hits with a rule like:

```text
/~a:(\d+)/
```

The parsed numeric athlete ID persisted across the tested ESPN API families:

| Tested player | Search `id` | Search `uid` | Parsed athlete ID | Verified endpoints |
| --- | --- | --- | --- | --- |
| Drake Maye | `2fbd2f9e-624d-3e63-9bb5-9bb965782b68` | `s:20~l:28~a:4431452` | `4431452` | core-v2 athlete, core-v3 athlete, common-v3 athlete, common-v3 overview/gamelog, site-v2 athlete news, site-v2 roster, core-v2 player game stats |
| Patrick Mahomes | `37d87523-280a-9d4a-0adb-22cfc6d3619c` | `s:20~l:28~a:3139477` | `3139477` | core-v2 athlete, common-v3 athlete |
| Justin Jefferson | `c7fb1996-2530-7f98-9f2f-572debf989f4` | `s:20~l:28~a:4262921` | `4262921` | Search v2 player hit parsed successfully |

The GUID did not work as a replacement path parameter in the tested athlete
URLs. For Drake Maye, using
`2fbd2f9e-624d-3e63-9bb5-9bb965782b68` instead of `4431452` returned `400` from:

```text
https://sports.core.api.espn.com/v2/sports/football/leagues/nfl/athletes/{id}
https://sports.core.api.espn.com/v3/sports/football/nfl/athletes/{id}
https://site.web.api.espn.com/apis/common/v3/sports/football/nfl/athletes/{id}
https://site.api.espn.com/apis/site/v2/sports/football/nfl/athletes/{id}/news
```

Cross-sport spot check: Search v2 returned `uid=s:40~l:46~a:1966` for LeBron
James, and athlete ID `1966` worked on both NBA core-v2 and NBA common-v3
athlete endpoints. This suggests the numeric athlete ID is the correct SDK
canonical key for ESPN athlete URLs, while `uid` carries sport/league context
and `guid` should be stored as secondary metadata.

## Primary SDK Chains

These are the chains the SDK should expose as convenience workflows.

| Workflow | Chain | Terminal endpoint |
| --- | --- | --- |
| Name to player profile | `search/v2?query={name}` -> athlete ID -> common-v3 athlete views | `/apis/common/v3/.../athletes/{id}/overview`, `/stats`, `/gamelog`, `/splits`, `/bio` |
| Team name to team schedule | site-v2 `teams` or search -> team ID -> site-v2 team schedule | `/apis/site/v2/.../teams/{id}/schedule` |
| Team season game to event IDs | site-v2 team schedule -> `events[].id` | event IDs for summary/core calls |
| Event to user-facing game package | site-v2 scoreboard or schedule -> event ID -> site-v2 summary | `/apis/site/v2/.../summary?event={event}` |
| Event to exact core graph | core-v2 events -> `items[].$ref` -> event -> `competitions[].$ref` -> competition | `/v2/.../events/{event}/competitions/{competition}` |
| Event to player game stats | event summary -> competition/team/player IDs -> core-v2 roster athlete statistics | `/v2/.../competitors/{competitor}/roster/{athlete}/statistics/{split}` |
| Event to play detail | competition -> plays/drives -> play ID | `/v2/.../plays/{play}` |
| Event to win probability | competition -> probabilities -> probability/play ID | `/v2/.../probabilities/{play}` |
| Team to roster/depth chart | team ID -> site-v2 roster/depthcharts, or common-v3 roster | `/apis/site/v2/.../teams/{id}/roster`, `/depthcharts`, `/apis/common/v3/.../teams/{id}/roster` |
| League leaderboards | no ID or optional season params -> site/common stats endpoints | `/apis/site/v2/.../statistics`, `/apis/common/v3/.../statistics/byathlete`, `/byteam` |

## Recommended Data Sources

The SDK should hide most domain selection from normal callers. A user should ask
for `getPlayerGameStats(...)`, `getTeamSchedule(...)`, or `getPlayByPlay(...)`;
the SDK should choose the best ESPN API family internally and only expose
domain-specific methods for advanced users.

| Data needed | Preferred API | Why | Fallback or companion API |
| --- | --- | --- | --- |
| Team lookup by display name | Site-v2 `teams` or Search v2 | Lower-reference, user-facing names and IDs | Core-v2 `teams` when core refs are needed. |
| Team detail/profile | Site-v2 `teams/{id}` | User-facing team object | Core-v2 season team for logos, venue refs, and core graph fields; core-v3 team for v3 expansion shape. |
| Team schedule | Site-v2 `teams/{id}/schedule` | Direct team season schedule with event IDs | Site-v2 scoreboard by date/week for league-wide discovery. |
| Team roster | Site-v2 `teams/{id}/roster` | Friendly roster grouped for app use | Common-v3 team roster for position-group shape; core-v2 event roster for game-specific active roster. |
| Depth chart | Site-v2 `teams/{id}/depthcharts` | Only documented populated depth-chart source | None documented as a better populated NFL source. |
| Player lookup by name | Search v2 | Best name-to-ID entrypoint | Site-v2/common-v3 rosters when the team is already known. |
| Player profile page data | Common-v3 athlete overview/bio | UI-friendly profile, news, next game, fantasy, team history | Core-v2 or core-v3 athlete detail for graph identity and refs. |
| Player season stats | Common-v3 athlete stats | Friendly stat tables with labels, filters, glossary | Core-v2 athlete statistics for core `splits` shape. |
| Player game log | Common-v3 athlete gamelog | Game-by-game rows, labels, event map | Core-v2 athlete statisticslog for core log shape. |
| Player split stats | Common-v3 athlete splits | Split categories and descriptions | Core-v2 athlete statistics if a core split shape is required. |
| Player game-specific stat line | Core-v2 competitor roster athlete statistics | Exact single-player, single-game stat terminal under event/competition/competitor | Site-v2 summary can help discover the needed IDs and may contain display boxscore data. |
| League/player leaderboards | Common-v3 `statistics/byathlete` | Leaderboard-oriented athlete stat shape | Site-v2 `statistics` for site-style league stats. |
| Team statistical leaderboards | Common-v3 `statistics/byteam` | Team leaderboard-oriented shape | Site-v2 `statistics` for site-style league stats. |
| League scoreboard | Site-v2 scoreboard | Friendly events, week, season calendar, event IDs | CDN scoreboard for page-package payloads. |
| Game summary/boxscore | Site-v2 summary | Best general app-facing game package | CDN game/boxscore for ESPN page-package shape; core-v2 competition children for exact graph data. |
| Game metadata and child refs | Core-v2 event/competition | Exact event graph hub and subresource refs | Site-v2 summary for easier IDs and display data. |
| Team game stats | Core-v2 competitor statistics | Exact team-in-game stat splits | Site-v2 summary boxscore for display-oriented team stats. |
| Team game leaders | Core-v2 competitor leaders | Exact event/team leader categories | Site-v2 summary leaders for display package. |
| Play-by-play | Core-v2 plays/drives for structured play resources | Stable event/competition/play IDs and individual play detail endpoints | Site-v2 summary or CDN playbyplay for page/display package payloads. |
| Drive data | Core-v2 drives | Structured drive IDs and start/end fields | Site-v2 summary/CDN game package for display package shape. |
| Win probability | Core-v2 probabilities | Per-play probability terminal with `play.$ref` | Site-v2 summary `winprobability` for display package. |
| Odds and betting providers | Core-v2 competition odds and providers | Provider-specific odds resources and provider taxonomy | Site-v2 summary/pickcenter for user-facing betting package. |
| Venue/stadium data | Core-v2 venues or competition `venue.$ref` | Address, indoor/grass, images, core venue IDs | Core-v3/team expansions when already using v3 team detail. |
| Standings | Site API `/apis/v2/sports/football/{league}/standings` | Actual standings tree | Site-v2 `/standings` is only a web-link stub. |
| Injuries | Site-v2 league injuries | Populated league-wide injury report | Site-v2 team injury endpoint returned `{}` for tested NFL team. |
| Transactions | Site-v2 league transactions | Populated transaction collection | Site-v2 team transaction endpoint returned `{}` for tested NFL team. |
| News | Site-v2 league/athlete news or Now news | Site-v2 gives article packages; Now gives headline feed | Search v2 for story discovery by query. |

## SDK Abstraction Guidance

Normal SDK methods should be data-oriented, not ESPN-domain-oriented:

```text
football.getTeamSchedule(teamId, options)
football.getPlayerProfile(athleteId)
football.getPlayerGameStats(eventId, athleteId)
football.getPlayByPlay(eventId)
football.getStandings()
```

Those methods should call the recommended API source internally and perform the
known resolver steps from this file. Advanced users can still access raw domain
clients when they need native ESPN payloads:

```text
football.site.summary(eventId)
football.core.competition(eventId, competitionId)
football.common.athleteGamelog(athleteId)
football.cdn.playByPlay(eventId)
football.search.query(text)
```

Do not force normal callers to know that player profile data is best from
common-v3 while player game stat lines are best from core-v2. That is the
SDK's job. The caller should only opt into a raw domain client when they care
about the exact ESPN response shape.

## Core V2

Use core-v2 for ref-first graph resources, exact event subresources, and
machine-oriented IDs. Most collection endpoints are resolvers; most detail
endpoints are terminal data endpoints.

### League And Season

| Endpoint | SDK role | API-specific data | Chain |
| --- | --- | --- | --- |
| `/v2/sports/football/leagues/{league}` | Resolver and metadata | League metadata plus child refs for season, teams, athletes, events, and calendar | Start here only for generic discovery. |
| `/leagues/{league}/calendar` | Resolver/terminal | Calendar sections with season/window refs | League -> calendar -> calendar segment refs. |
| `/leagues/{league}/season` | Terminal | Current season object, season types, rankings refs | Direct from league. |
| `/leagues/{league}/seasons` | Resolver | Season list through `items[].$ref` | League -> seasons -> season detail. |
| `/leagues/{league}/seasons/{season}` | Terminal | Season detail by year | Direct when season is known. |
| `/leagues/{league}/seasons/{season}/athletes` | Resolver | Season-scoped athlete refs | Season -> athlete refs -> season athlete detail. |
| `/leagues/{league}/seasons/{season}/draft` | Terminal | Draft rounds, status, and draft athlete refs | Season -> draft. |
| `/leagues/{league}/seasons/{season}/freeagents` | Terminal collection | Free-agent collection shape | Season -> free agents. |

### Teams

| Endpoint | SDK role | API-specific data | Chain |
| --- | --- | --- | --- |
| `/leagues/{league}/teams` | Resolver | Current-season team refs in `items[].$ref` | League -> teams -> season team detail. |
| `/leagues/{league}/seasons/{season}/teams/{id}` | Terminal | Core team identity, logos, venue, links, and season scoping | Team collection or known ID -> team detail. |

Use site-v2 `teams` for friendlier name lookup. Use core-v2 team detail when
the SDK needs ESPN core refs, logos, venue, or season-specific team identity.

### Athletes

| Endpoint | SDK role | API-specific data | Chain |
| --- | --- | --- | --- |
| `/leagues/{league}/athletes` | Resolver | Athlete refs in `items[].$ref` | League -> athletes -> athlete detail. |
| `/leagues/{league}/athletes/{id}` | Terminal | Core athlete identity, alternate IDs, bio refs, team refs | Search/site roster/core collection -> athlete detail. |
| `/leagues/{league}/athletes/{id}/statistics` | Terminal | Core season/career statistics under `splits` | Athlete ID -> core stats. |
| `/leagues/{league}/athletes/{id}/statisticslog` | Terminal | Statistics log entries | Athlete ID -> core stat log. |
| `/leagues/{league}/athletes/{id}/contracts` | Terminal collection | Contract collection shape | Athlete ID -> contracts. |
| `/leagues/{league}/athletes/{id}/records` | Terminal collection | Records/achievements collection shape | Athlete ID -> records. |

Use common-v3 athlete endpoints for UI-friendly player pages and game logs. Use
core-v2 athlete endpoints for the core graph and ref-compatible stats objects.

### Events And Competitions

| Endpoint | SDK role | API-specific data | Chain |
| --- | --- | --- | --- |
| `/leagues/{league}/events` | Resolver | Event refs in `items[].$ref` | League -> events -> event detail. |
| `/leagues/{league}/events/{event}` | Resolver/terminal | Event metadata and `competitions[].$ref` | Event ID -> competition detail. |
| `/leagues/{league}/events/{event}/competitions/{competition}` | Hub | Rich core game hub: venue, competitors, broadcasts, odds, status, situation, drives, plays, probabilities, predictor | Event -> competition -> child endpoints. |
| `/competitions/{competition}/broadcasts` | Terminal collection | Broadcast/media entries | Competition -> broadcasts. |
| `/competitions/{competition}/odds` | Terminal collection | Provider-specific odds | Competition -> odds. |
| `/competitions/{competition}/officials` | Terminal collection | Officials list | Competition -> officials. |
| `/competitions/{competition}/notes` | Terminal collection | Game notes | Competition -> notes. |
| `/competitions/{competition}/status` | Terminal | Clock, period, and status type | Competition -> status. |
| `/competitions/{competition}/situation` | Terminal plus resolver | Down, distance, yard line, red zone, timeouts, `lastPlay.$ref` | Competition -> situation -> last play if needed. |
| `/competitions/{competition}/predictor` | Terminal | Predictor model values by home/away team | Competition -> predictor. |
| `/competitions/{competition}/powerindex` | Resolver | Team power index refs in `items[].$ref` | Competition -> power index -> team power index. |
| `/competitions/{competition}/powerindex/{teamId}` | Terminal | Team power index stats for the event | Competition/team -> power index detail. |

### Event Competitors

| Endpoint | SDK role | API-specific data | Chain |
| --- | --- | --- | --- |
| `/competitions/{competition}/competitors/{competitor}` | Hub | Event-specific team competitor, home/away, winner, child refs/placeholders | Competition -> competitor. |
| `/competitors/{competitor}/score` | Terminal | Competitor score object with winner/source | Competitor -> score. |
| `/competitors/{competitor}/linescores` | Terminal collection | Period-by-period scoring | Competitor -> linescores. |
| `/competitors/{competitor}/roster` | Resolver/terminal | Event roster entries, active game roster, team ref | Competitor -> roster -> athlete IDs. |
| `/competitors/{competitor}/statistics` | Terminal | Team game statistics under `splits` | Competitor -> team stats. |
| `/competitors/{competitor}/leaders` | Terminal | Team game leader categories | Competitor -> leaders. |
| `/competitors/{competitor}/roster/{athlete}/statistics/{split}` | Terminal | Player game stat line under `splits.categories[].stats[]` | Summary/roster -> athlete ID -> player game stats. |

This is the key chain for questions like "how many yards did a player have in a
game." Site-v2 summary is usually the easiest way to find the event,
competition, competitor, and athlete IDs; this core-v2 endpoint is the exact
single-player game-stat terminal.

### Drives, Plays, And Probability

| Endpoint | SDK role | API-specific data | Chain |
| --- | --- | --- | --- |
| `/competitions/{competition}/drives` | Resolver/terminal | Drive collection and drive IDs | Competition -> drives. |
| `/competitions/{competition}/drives/{drive}` | Terminal | Drive start/end, team, description, embedded play refs/items | Drives -> drive detail. |
| `/competitions/{competition}/drives/{drive}/plays` | Resolver/terminal | Plays within one drive | Drive -> drive plays. |
| `/competitions/{competition}/plays` | Resolver/terminal | Full play collection | Competition -> plays -> play detail. |
| `/competitions/{competition}/plays/{play}` | Terminal | Play text, type, sequence, scoring fields, participants when present | Play collection/situation/probability -> play detail. |
| `/competitions/{competition}/probabilities` | Resolver | Win probability refs/items keyed by play | Competition -> probabilities. |
| `/competitions/{competition}/probabilities/{play}` | Terminal plus resolver | Home/away/tie win percentages and `play.$ref` | Probability -> play detail. |

### Lookup Resources

| Endpoint | SDK role | API-specific data | Chain |
| --- | --- | --- | --- |
| `/leagues/{league}/venues` | Resolver/terminal | Venue collection with images/address fields | League -> venues -> venue detail. |
| `/leagues/{league}/venues/{id}` | Terminal | Venue detail with grass/indoor/address/images | Known venue ID -> venue. |
| `/leagues/{league}/countries` | Resolver/terminal | Countries with flags and athlete refs | League -> countries -> country detail. |
| `/leagues/{league}/countries/{id}` | Terminal | Country detail and related athletes | Known country ID -> country. |
| `/leagues/{league}/franchises` | Resolver/terminal | Franchise list with linked team/venue data | League -> franchises -> franchise detail. |
| `/leagues/{league}/franchises/{id}` | Terminal | Franchise detail, team, venue, logos, links | Known franchise ID -> franchise. |
| `/leagues/{league}/positions` | Resolver/terminal | Position taxonomy | League -> positions -> position detail. |
| `/leagues/{league}/positions/{id}` | Terminal | Position detail and abbreviation | Known position ID -> position. |
| `/leagues/{league}/providers` | Resolver/terminal | Odds/content provider taxonomy | League -> providers -> provider detail. |
| `/leagues/{league}/providers/{id}` | Terminal | Provider detail | Known provider ID -> provider. |
| `/leagues/{league}/casinos` | Resolver | Casino refs in `items[].$ref` | League -> casinos -> casino detail. |
| `/leagues/{league}/casinos/{id}` | Terminal | Casino detail with active/priority fields | Known casino ID -> casino. |

## Core V3

Use core-v3 only when the SDK wants v3-shaped core resources or expansion-style
queries. For NFL, the useful documented paths include league, season, team, and
athlete resources. Several v3 roots are sparse or unsupported.

| Endpoint | SDK role | API-specific data | Chain |
| --- | --- | --- | --- |
| `/v3/sports/football/{league}` | Terminal | V3 league identity fields | Direct. |
| `/v3/sports/football/{league}/seasons/{season}` | Terminal | V3 season object | Direct by season. |
| `/v3/sports/football/{league}/teams/{id}` | Terminal | V3 team identity, expandable logos/links/groups/venue/record | Team ID -> v3 team. |
| `/v3/sports/football/{league}/athletes` | Resolver/terminal | V3 athlete collection | League -> v3 athlete IDs. |
| `/v3/sports/football/{league}/athletes/{id}` | Terminal | V3 athlete identity, expandable team/position/links | Athlete ID -> v3 athlete. |
| `/v3/sports/football/{league}/athletes/{id}/plays` | Terminal collection | Athlete-related play collection | Athlete ID -> athlete plays. |
| `/v3/sports/football/{league}/athletes/{id}/statisticslog` | Terminal collection | V3 athlete stat-log entries | Athlete ID -> v3 stat log. |

## Site V2

Use site-v2 for user-facing and low-reference data. These endpoints are usually
the best SDK entrypoints for app workflows because they expose IDs and display
objects directly.

| Endpoint | SDK role | API-specific data | Chain |
| --- | --- | --- | --- |
| `/apis/site/v2/sports/football/{league}/scoreboard` | Entry terminal/resolver | Scoreboard events, week, league season calendar | Scoreboard -> `events[].id` -> summary/core event. |
| `/apis/site/v2/sports/football/{league}/summary?event={event}` | Hub | User-facing game package: boxscore, drives, leaders, injuries, broadcasts, pickcenter, odds, win probability | Event ID -> summary -> IDs for core chains. |
| `/apis/site/v2/sports/football/{league}/teams` | Entry resolver | User-facing team list | Team name -> team ID. |
| `/apis/site/v2/sports/football/{league}/teams/{id}` | Terminal | User-facing team detail | Team ID -> team detail. |
| `/apis/site/v2/sports/football/{league}/teams/{id}/roster` | Terminal/resolver | Team roster grouped for display, coach, team | Team ID -> athlete IDs -> athlete chains. |
| `/apis/site/v2/sports/football/{league}/teams/{id}/schedule` | Terminal/resolver | Team season schedule, events, bye week | Team ID/season -> event IDs. |
| `/apis/site/v2/sports/football/{league}/teams/{id}/depthcharts` | Terminal | Team depth chart | Team ID -> depth chart. |
| `/apis/site/v2/sports/football/{league}/athletes/{id}/news` | Terminal | Athlete-specific news articles | Athlete ID -> news. |
| `/apis/site/v2/sports/football/{league}/draft` | Terminal | Draft positions, rounds, status, broadcasts | League/season -> draft. |
| `/apis/site/v2/sports/football/{league}/groups` | Terminal | Conference/division groups | League -> groups. |
| `/apis/site/v2/sports/football/{league}/injuries` | Terminal | League-wide injury report | League -> injuries. |
| `/apis/site/v2/sports/football/{league}/news` | Terminal | League news articles | League -> news. |
| `/apis/v2/sports/football/{league}/standings` | Terminal | Actual standings tree with groups/children | League -> standings. |
| `/apis/site/v2/sports/football/{league}/standings` | Terminal stub | `fullViewLink` only | Use only for web-link discovery, not data. |
| `/apis/site/v2/sports/football/{league}/statistics` | Terminal | League statistics object | League -> stat leaders. |
| `/apis/site/v2/sports/football/{league}/transactions` | Terminal | Transaction collection and pagination | League/year -> transactions. |

## Site Web Common V3

Use common-v3 for athlete pages, athlete and team leaderboards, and the one
working common-v3 team roster shape. Generic common-v3 football league/team/event
roots are not useful for NFL in the documented tests.

| Endpoint | SDK role | API-specific data | Chain |
| --- | --- | --- | --- |
| `/apis/common/v3/sports/football/{league}/athletes/{id}` | Terminal | Athlete page shell: player switcher, quicklinks, standings, videos | Athlete ID -> page shell. |
| `/apis/common/v3/sports/football/{league}/athletes/{id}/overview` | Terminal | Combined player page data: statistics, news, nextGame, gameLog, rotowire, fantasy | Athlete ID -> overview. |
| `/apis/common/v3/sports/football/{league}/athletes/{id}/stats` | Terminal | Player stat tables with filters, teams, categories, glossary | Athlete ID -> season stats. |
| `/apis/common/v3/sports/football/{league}/athletes/{id}/gamelog` | Terminal | Game-by-game stat rows, labels, event map, season types | Athlete ID -> game log. |
| `/apis/common/v3/sports/football/{league}/athletes/{id}/splits` | Terminal | Split categories, labels, descriptions | Athlete ID -> splits. |
| `/apis/common/v3/sports/football/{league}/athletes/{id}/bio` | Terminal | Team history | Athlete ID -> bio. |
| `/apis/common/v3/sports/football/{league}/statistics/byathlete` | Terminal | Athlete leaderboards with pagination, categories, glossary | League/season/category -> athlete leaderboard. |
| `/apis/common/v3/sports/football/{league}/statistics/byteam` | Terminal | Team leaderboards with categories and glossary | League/season/category -> team leaderboard. |
| `/apis/common/v3/sports/football/{league}/teams/{id}/roster` | Terminal/resolver | Position-group roster, coach, team | Team ID -> grouped roster -> athlete IDs. |

## CDN Game Package

Use CDN endpoints when the SDK wants ESPN page-package payloads. These are less
normalized than site-v2/core-v2, but they can be useful when mirroring page views
or when `gamepackageJSON` has page-specific content.

| Endpoint | SDK role | API-specific data | Chain |
| --- | --- | --- | --- |
| `https://cdn.espn.com/core/nfl/scoreboard?xhr=1` | Terminal | CDN scoreboard page package, `content`, `nowFeed` | Direct. |
| `https://cdn.espn.com/core/nfl/game?xhr=1&gameId={event}` | Terminal | Full CDN game package in `gamepackageJSON` | Event ID -> CDN game. |
| `https://cdn.espn.com/core/nfl/boxscore?xhr=1&gameId={event}` | Terminal | CDN boxscore page package in `gamepackageJSON` | Event ID -> CDN boxscore. |
| `https://cdn.espn.com/core/nfl/matchup?xhr=1&gameId={event}` | Terminal | CDN matchup page package in `gamepackageJSON` | Event ID -> CDN matchup. |
| `https://cdn.espn.com/core/nfl/playbyplay?xhr=1&gameId={event}` | Terminal | CDN play-by-play page package in `gamepackageJSON` | Event ID -> CDN play-by-play. |

## Search, Header, And News Feeds

| Endpoint | SDK role | API-specific data | Chain |
| --- | --- | --- | --- |
| `https://site.web.api.espn.com/apis/search/v2?query={query}` | Entry resolver | Search result groups and content hits | Name/query -> player/team/story IDs and links. |
| `https://site.web.api.espn.com/apis/v2/scoreboard/header?sport=football&league=nfl` | Terminal | Header scoreboard sports/leagues/events shape | Direct. |
| `https://now.core.api.espn.com/v1/sports/news?sport=football` | Terminal | Now feed headlines and breaking-news shape | Direct or league/sport-scoped feed. |

## Not Useful As SDK Data Sources

These documented endpoints either returned an error, an empty object, or a stub
that should not be treated as a real data source.

| Endpoint family | Observed result | SDK decision |
| --- | --- | --- |
| Core-v2 `/events/{event}/competitions/{competition}/details` | `404` | Do not expose as working data. |
| Core-v2 `/plays/{play}/personnel` | `500` for tested NFL play | Keep internal/experimental only. |
| Core-v2 `/seasons/{season}/manufacturers` | `400` | Do not expose for NFL. |
| Core-v2 `/leagues/{league}/tournaments` | `400` | Do not expose for NFL. |
| Core-v2 `/leagues/{league}/circuits` | Empty collection in NFL docs | Not useful for NFL SDK data. |
| Core-v2 `/leagues/{league}/rankings` | Empty collection in NFL docs | Not useful for NFL SDK data. |
| Core-v2 `/leagues/{league}/recruiting` | Empty collection in NFL docs | Not useful for NFL SDK data. |
| Site-v2 `/calendar` | `404` for NFL | Use site-v2 scoreboard calendar fields instead. |
| Site-v2 `/rankings` | `404` for NFL | Not useful for NFL; may be college-football-specific. |
| Site-v2 team `/record`, `/news`, `/injuries`, `/leaders`, `/history`, `/transactions` | `{}` for tested NFL team | Do not expose as populated team endpoints. Use league-level or other domains where available. |
| Site-v3 `site.api` `/scoreboard` and `/summary` | `404` | Use site-v2 for these. |
| Site.web common-v3 generic league/team/event roots | Mostly `404` or `400` | Expose only the working common-v3 athlete/stat/team-roster endpoints above. |
| Site-v2 standings at `/apis/site/v2/.../standings` | `fullViewLink` only | Use `/apis/v2/sports/football/{league}/standings` for standings data. |

## SDK Surface Recommendation

Keep the client split by ESPN domain because each domain has a different
response style and different resolver requirements.

```text
football.site.*
  Friendly entrypoints: scoreboard, summary, teams, rosters, schedules,
  standings, injuries, transactions, news.

football.core.*
  Exact graph resources: events, competitions, competitors, player game stats,
  drives, plays, probabilities, lookup resources.

football.common.*
  Athlete profile/stat views, game logs, splits, leaderboards, common-v3 roster.

football.cdn.*
  Page-package game payloads in gamepackageJSON.

football.search.*
  Name/query resolution and scoreboard header data.

football.news.*
  Now-feed headline data.
```

Default SDK methods should return the endpoint's native shape. Add convenience
methods for high-value chains, but make automatic reference expansion explicit
and bounded, for example `resolveItems: true`, `limit`, and `maxDepth`.
