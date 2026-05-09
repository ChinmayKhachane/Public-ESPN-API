# Football API Hierarchy

Verified against the live NFL response docs and reference resolution map on
2026-05-09.

This file describes how football/NFL API calls flow from broad resources into
detail resources. It is intended for SDK design: collection endpoints often
return pagination plus `$ref` links, while detail endpoints can usually be
called directly once you know the relevant ID.

Primary fixtures used in the documented examples:
- `league=nfl`
- `season=2025`
- `event=401772988`
- `competition=401772988`
- `team id=22` for core-v2 season team examples
- `team id=6` for site-v2 and site-v3 team examples
- `competitor=17`
- `athlete id=4429202` or `4431452`, depending on endpoint family
- `drive=4017729881`
- `play=40177298840`

## Resolution Rules

Use these rules when designing SDK methods:

| Response shape | Field to follow | SDK behavior |
| --- | --- | --- |
| Detail object with useful fields | none | Return it directly. |
| Collection with `items[].$ref` only | `items[].$ref` | Collection methods can expose refs, or optionally resolve items. |
| Collection with `items[]` already expanded | none | Return collection directly; individual IDs can still be called directly. |
| Object with nested child refs | named child `$ref` fields | Do not resolve all children automatically; expose explicit child methods. |
| Empty `{}` or error payload | none | Return as observed or raise an SDK error based on client policy. |

Do not recursively expand every nested `$ref` by default. Football responses
can point back to league, season, team, event, competition, play, athlete, and
venue resources, which can create large graphs.

## Core V2 Root Flow

Base:

```text
https://sports.core.api.espn.com/v2/sports/football/leagues/nfl
```

The league root returns league metadata and embeds or references major child
resources.

```text
/leagues/nfl
  -> season
  -> teams
  -> athletes
  -> events
  -> calendar
  -> standings and season-type resources
```

The most useful SDK entry points are usually not the league root itself, but
the typed child methods below.

## Teams

Core-v2 team collections are season-scoped behind the scenes.

| SDK concept | Endpoint | Ref field | Detail endpoint |
| --- | --- | --- | --- |
| Team collection | `/leagues/{league}/teams` | `items[].$ref` | `/leagues/{league}/seasons/{season}/teams/{id}` |
| Season team detail | `/leagues/{league}/seasons/{season}/teams/{id}` | none for the team itself | direct by `season` and `id` |

Observed flow:

```text
GET /leagues/nfl/teams
  items[0].$ref
    -> /leagues/nfl/seasons/2025/teams/22?lang=en&region=us
```

SDK implication:

```text
football.core.teams()
  -> returns collection metadata and team refs

football.core.team("22", { season: 2025 })
  -> calls /leagues/nfl/seasons/2025/teams/22 directly
```

Yes, once the team ID is known, you can call the detail endpoint directly with
an arbitrary valid team ID. The collection is mainly useful for discovery,
pagination, and current-season ref discovery.

Related direct site endpoints:

| SDK concept | Endpoint | Notes |
| --- | --- | --- |
| Site teams | `/apis/site/v2/sports/football/{league}/teams` | User-facing team list. |
| Site team detail | `/apis/site/v2/sports/football/{league}/teams/{id}` | Direct by team ID. |
| Site roster | `/apis/site/v2/sports/football/{league}/teams/{id}/roster` | Direct by team ID. |
| Site schedule | `/apis/site/v2/sports/football/{league}/teams/{id}/schedule` | Direct by team ID. |
| Site depth chart | `/apis/site/v2/sports/football/{league}/teams/{id}/depthcharts` | Direct by team ID. |
| Site sparse team children | `/record`, `/news`, `/injuries`, `/leaders`, `/history`, `/transactions` | Returned `{}` for the tested NFL team. |

## Athletes

Athlete collections are ref-first. Athlete detail and athlete stat endpoints
are direct once the athlete ID is known.

| SDK concept | Endpoint | Ref field | Detail endpoint |
| --- | --- | --- | --- |
| Athlete collection | `/leagues/{league}/athletes` | `items[].$ref` | `/leagues/{league}/athletes/{id}` |
| Season athlete collection | `/leagues/{league}/seasons/{season}/athletes` | `items[].$ref` | `/leagues/{league}/seasons/{season}/athletes/{id}` |
| Athlete detail | `/leagues/{league}/athletes/{id}` | none for athlete itself | direct by athlete ID |
| Athlete statistics | `/leagues/{league}/athletes/{id}/statistics` | none for statistics object | direct by athlete ID |
| Athlete statistics log | `/leagues/{league}/athletes/{id}/statisticslog` | none for log object | direct by athlete ID |
| Athlete contracts | `/leagues/{league}/athletes/{id}/contracts` | no useful refs in tested collection | direct by athlete ID |
| Athlete records | `/leagues/{league}/athletes/{id}/records` | no useful refs in tested collection | direct by athlete ID |

Observed flow:

```text
GET /leagues/nfl/athletes
  items[0].$ref
    -> /leagues/nfl/athletes/4246273?lang=en&region=us
```

SDK implication:

```text
football.core.athletes()
  -> returns collection metadata and athlete refs

football.core.athlete("4429202")
  -> calls /leagues/nfl/athletes/4429202 directly

football.core.athleteStatistics("4429202")
  -> calls /leagues/nfl/athletes/4429202/statistics directly
```

Site/common-v3 athlete endpoints are separate user-facing views:

| SDK concept | Endpoint | Notes |
| --- | --- | --- |
| Athlete overview | `/apis/common/v3/sports/football/{league}/athletes/{id}/overview` | Direct by athlete ID. |
| Athlete stats | `/apis/common/v3/sports/football/{league}/athletes/{id}/stats` | Direct by athlete ID. |
| Athlete gamelog | `/apis/common/v3/sports/football/{league}/athletes/{id}/gamelog` | Direct by athlete ID. |
| Athlete splits | `/apis/common/v3/sports/football/{league}/athletes/{id}/splits` | Direct by athlete ID. |
| Athlete bio | `/apis/common/v3/sports/football/{league}/athletes/{id}/bio` | Direct by athlete ID. |
| Athlete news | `/apis/site/v2/sports/football/{league}/athletes/{id}/news` | Direct by athlete ID. |

## Events And Competitions

Events are the main hub for game-specific core data.

| SDK concept | Endpoint | Ref field | Detail endpoint |
| --- | --- | --- | --- |
| Event collection | `/leagues/{league}/events` | `items[].$ref` | `/leagues/{league}/events/{event}` |
| Event detail | `/leagues/{league}/events/{event}` | `competitions[].$ref` for competition detail | `/events/{event}/competitions/{competition}` |
| Competition detail | `/events/{event}/competitions/{competition}` | named child refs and child collections | child endpoints below |

Observed flow:

```text
GET /leagues/nfl/events
  items[0].$ref
    -> /leagues/nfl/events/401772988?lang=en&region=us
      competitions[0].$ref
        -> /leagues/nfl/events/401772988/competitions/401772988
```

SDK implication:

```text
football.core.events()
  -> returns event collection and refs

football.core.event("401772988")
  -> direct event detail

football.core.competition("401772988", "401772988")
  -> direct competition detail
```

The competition object is the richest core-v2 hub. It exposes competitors,
venue, broadcasts, odds, officials, drives, plays, probabilities, status,
situation, predictor, and power index data.

## Competition Children

These endpoints are direct once `event` and `competition` are known.

| SDK concept | Endpoint | Ref field | Notes |
| --- | --- | --- | --- |
| Broadcasts | `/events/{event}/competitions/{competition}/broadcasts` | none | Collection response. |
| Odds | `/events/{event}/competitions/{competition}/odds` | none | Collection response. |
| Officials | `/events/{event}/competitions/{competition}/officials` | none | Collection response. |
| Status | `/events/{event}/competitions/{competition}/status` | none | Direct status object. |
| Situation | `/events/{event}/competitions/{competition}/situation` | `lastPlay.$ref` available | SDK should expose `play(id)` rather than auto-resolve by default. |
| Details | `/events/{event}/competitions/{competition}/details` | none | Returned `404` for tested NFL event. |
| Play personnel | `/events/{event}/competitions/{competition}/plays/{play}/personnel` | none | Returned `500` for tested `play=4017729884651` in the hierarchy sweep. |

## Competitors

Competitors are event-specific teams.

| SDK concept | Endpoint | Ref field | Notes |
| --- | --- | --- | --- |
| Competitor detail | `/events/{event}/competitions/{competition}/competitors/{competitor}` | named child refs | Direct by event, competition, competitor ID. |
| Competitor score | `/competitors/{competitor}/score` | none | Direct score object. |
| Competitor linescores | `/competitors/{competitor}/linescores` | none | Collection response. |
| Competitor roster | `/competitors/{competitor}/roster` | none for top-level roster | Event roster entries. |
| Competitor statistics | `/competitors/{competitor}/statistics` | none for top-level stats | Team game stats. |
| Competitor leaders | `/competitors/{competitor}/leaders` | none for top-level leaders | Leader categories. |
| Roster athlete statistics | `/competitors/{competitor}/roster/{athlete}/statistics/{split}` | none for top-level stats | Direct by event roster athlete ID. |

Observed competition-to-competitor flow:

```text
GET /events/401772988/competitions/401772988
  competitors[0].$ref
    -> /events/401772988/competitions/401772988/competitors/17
```

SDK implication:

```text
football.core.competitor(eventId, competitionId, competitorId)
football.core.competitorScore(eventId, competitionId, competitorId)
football.core.competitorRoster(eventId, competitionId, competitorId)
```

Competitor IDs are usually team IDs within the event context, but SDK methods
should treat them as competitor IDs because ESPN URLs use the competition
competitor path.

## Drives And Plays

Drives and plays are game-specific resources under a competition.

| SDK concept | Endpoint | Ref field | Notes |
| --- | --- | --- | --- |
| Drives | `/events/{event}/competitions/{competition}/drives` | none for tested collection shape | Drive objects can include inline `plays`. |
| Drive detail | `/drives/{drive}` | none for drive itself | Direct by drive ID. |
| Drive plays | `/drives/{drive}/plays` | none for tested collection shape | Plays in one drive. |
| Plays | `/events/{event}/competitions/{competition}/plays` | none for tested collection shape | Full play collection. |
| Play detail | `/plays/{play}` | none for play itself | Direct by play ID. |
| Probabilities | `/probabilities` | none for tested collection shape | Win probability collection. |
| Probability detail | `/probabilities/{play}` | none for probability itself | Direct by play ID. |

Observed flow:

```text
GET /events/401772988/competitions/401772988/drives
  items[].id = 4017729881

GET /events/401772988/competitions/401772988/drives/4017729881
  plays.items[].id = 40177298840

GET /events/401772988/competitions/401772988/plays/40177298840
GET /events/401772988/competitions/401772988/probabilities/40177298840
```

SDK implication:

```text
football.core.drives(eventId, competitionId)
football.core.drive(eventId, competitionId, driveId)
football.core.drivePlays(eventId, competitionId, driveId)
football.core.plays(eventId, competitionId)
football.core.play(eventId, competitionId, playId)
football.core.probability(eventId, competitionId, playId)
```

## Calendar, Seasons, And Season Resources

| SDK concept | Endpoint | Ref field | Detail endpoint |
| --- | --- | --- | --- |
| Calendar | `/leagues/{league}/calendar` | `items[].$ref` | Calendar segment URLs such as `/calendar/ondays` |
| Seasons | `/leagues/{league}/seasons` | `items[].$ref` | `/leagues/{league}/seasons/{season}` |
| Current season | `/leagues/{league}/season` | none for season itself | Current season object |
| Season athletes | `/leagues/{league}/seasons/{season}/athletes` | `items[].$ref` | `/seasons/{season}/athletes/{id}` |
| Season draft | `/leagues/{league}/seasons/{season}/draft` | none for draft itself | Direct object |
| Season free agents | `/leagues/{league}/seasons/{season}/freeagents` | none for tested collection shape | Direct collection |
| Season manufacturers | `/leagues/{league}/seasons/{season}/manufacturers` | none | Returned `400`. |

Observed flows:

```text
GET /leagues/nfl/calendar
  items[0].$ref
    -> /leagues/nfl/calendar/ondays?lang=en&region=us

GET /leagues/nfl/seasons
  items[0].$ref
    -> /leagues/nfl/seasons/2026?lang=en&region=us

GET /leagues/nfl/seasons/2025/athletes
  items[0].$ref
    -> /leagues/nfl/seasons/2025/athletes/4246273?lang=en&region=us
```

## Misc Core Resources

These are mostly lookup collections. Some are collection-to-detail resources.

| SDK concept | Collection endpoint | Ref field | Detail endpoint |
| --- | --- | --- | --- |
| Casinos | `/leagues/{league}/casinos` | `items[].$ref` | `/leagues/{league}/casinos/{id}` |
| Countries | `/leagues/{league}/countries` | none in tested collection shape | `/leagues/{league}/countries/{id}` |
| Franchises | `/leagues/{league}/franchises` | none in tested collection shape | `/leagues/{league}/franchises/{id}` |
| Positions | `/leagues/{league}/positions` | none in tested collection shape | `/leagues/{league}/positions/{id}` |
| Providers | `/leagues/{league}/providers` | none in tested collection shape | `/leagues/{league}/providers/{id}` |
| Venues | `/leagues/{league}/venues` | none in tested collection shape | `/leagues/{league}/venues/{id}` |
| Circuits | `/leagues/{league}/circuits` | none | Empty collection in tested NFL docs. |
| Rankings | `/leagues/{league}/rankings` | none | Empty collection for NFL. |
| Recruiting | `/leagues/{league}/recruiting` | none | Empty collection for NFL. |
| Tournaments | `/leagues/{league}/tournaments` | none | Returned `400` for NFL. |

SDK implication:

```text
football.core.casinos()
football.core.casino(id)
football.core.venues()
football.core.venue(id)
football.core.franchises()
football.core.franchise(id)
```

## Site V2 User-Facing Flow

Site-v2 endpoints generally return user-facing JSON directly, not ref-first
core graph resources.

Base:

```text
https://site.api.espn.com/apis/site/v2/sports/football/nfl
```

| SDK concept | Endpoint | Notes |
| --- | --- | --- |
| Scoreboard | `/scoreboard` | Includes events and season calendar data. |
| Summary | `/summary?event={event}` | Useful game package: boxscore, drives, leaders, injuries, broadcasts, pickcenter. |
| Teams | `/teams` | User-facing team list. |
| Team detail | `/teams/{id}` | Direct by team ID. |
| Team roster | `/teams/{id}/roster` | Direct by team ID. |
| Team schedule | `/teams/{id}/schedule` | Direct by team ID. |
| Team depth charts | `/teams/{id}/depthcharts` | Direct by team ID. |
| News | `/news` | Direct league news. |
| Injuries | `/injuries` | Direct league injury report. |
| Transactions | `/transactions` | Direct transactions list. |
| Statistics | `/statistics` | Large stats object. |
| Groups | `/groups` | Conferences/divisions. |
| Draft | `/draft` | Draft data. |
| Rankings | `/rankings` | Returned `404` for NFL. |
| Calendar | `/calendar` | Returned `404` for NFL. |
| Standings stub | `/apis/site/v2/.../standings` | Stub with `fullViewLink`. |
| Standings data | `/apis/v2/sports/football/{league}/standings` | Use this for standings. |

SDK implication:

```text
football.site.scoreboard()
football.site.summary(eventId)
football.site.teams()
football.site.team(id)
football.site.standings()
```

## Site Web Common V3 Flow

Common-v3 is selective for football. The useful paths are mostly athlete,
team roster, and stat leaderboard views.

Base:

```text
https://site.web.api.espn.com/apis/common/v3/sports/football/nfl
```

| SDK concept | Endpoint | Status in docs |
| --- | --- | --- |
| Athlete overview | `/athletes/{id}/overview` | Works. |
| Athlete stats | `/athletes/{id}/stats` | Works. |
| Athlete gamelog | `/athletes/{id}/gamelog` | Works. |
| Athlete splits | `/athletes/{id}/splits` | Works. |
| Athlete bio | `/athletes/{id}/bio` | Works. |
| Team roster | `/teams/{id}/roster` | Works. |
| Stats by athlete | `/statistics/byathlete` | Works. |
| Stats by team | `/statistics/byteam` | Works. |
| Generic league/team/event roots | several paths | Mostly `404`; see `site-v3/unsupported_paths.md`. |

## CDN Game Package Flow

CDN endpoints return page-package data directly. They do not use `$ref`
resolution in the same way as core-v2.

Base:

```text
https://cdn.espn.com/core/nfl
```

| SDK concept | Endpoint | Useful payload location |
| --- | --- | --- |
| CDN scoreboard | `/scoreboard?xhr=1` | `content`, `nowFeed`, scoreboard page data |
| CDN game | `/game?xhr=1&gameId={event}` | `gamepackageJSON` |
| CDN boxscore | `/boxscore?xhr=1&gameId={event}` | `gamepackageJSON` |
| CDN matchup | `/matchup?xhr=1&gameId={event}` | `gamepackageJSON` |
| CDN play-by-play | `/playbyplay?xhr=1&gameId={event}` | `gamepackageJSON` |

SDK implication:

```text
football.cdn.game(eventId)
football.cdn.boxscore(eventId)
football.cdn.playByPlay(eventId)
```

## Search And News Flow

| SDK concept | Endpoint | Notes |
| --- | --- | --- |
| Scoreboard header | `https://site.web.api.espn.com/apis/v2/scoreboard/header?sport=football&league=nfl` | Use both `sport` and `league`; `sport=football` alone returned `400`. |
| Search | `https://site.web.api.espn.com/apis/search/v2?query={query}` | Results are grouped by type; hits are in `results[].contents[]`. |
| Now news | `https://now.core.api.espn.com/v1/sports/news?sport=football` | Real-time headline feed. |

## Suggested SDK Shape

Keep domain-specific clients, because ESPN's domains have different response
styles:

```text
football.core.*
  Ref-first graph resources and detailed event data.

football.site.*
  User-facing scoreboards, summaries, teams, news, standings, stats.

football.common.*
  Athlete profile views, team roster, and stat leaderboards.

football.cdn.*
  Game package/page-package endpoints.

football.search.*
  Search, scoreboard header, Now news.
```

Recommended resolution behavior:

| SDK method type | Default |
| --- | --- |
| Direct detail method, such as `team(id)` or `event(id)` | No ref traversal needed. |
| Collection method, such as `teams()` or `events()` | Return collection and refs by default; optionally resolve `items[].$ref`. |
| Convenience method, such as `allTeams({ resolve: true })` | Resolve each `items[].$ref` with pagination and limits. |
| Game hub method, such as `competition(event, competition)` | Return hub object; expose child methods instead of recursively expanding nested refs. |

## Known Direct-ID Patterns

These endpoint families can be called directly when the ID is known:

```text
/leagues/nfl/athletes/{id}
/leagues/nfl/seasons/{season}/athletes/{id}
/leagues/nfl/seasons/{season}/teams/{id}
/leagues/nfl/events/{event}
/leagues/nfl/events/{event}/competitions/{competition}
/leagues/nfl/events/{event}/competitions/{competition}/competitors/{competitor}
/leagues/nfl/events/{event}/competitions/{competition}/drives/{drive}
/leagues/nfl/events/{event}/competitions/{competition}/plays/{play}
/leagues/nfl/events/{event}/competitions/{competition}/probabilities/{play}
/leagues/nfl/venues/{id}
/leagues/nfl/casinos/{id}
/leagues/nfl/franchises/{id}
/leagues/nfl/positions/{id}
/leagues/nfl/providers/{id}
```

Use collection endpoints to discover IDs, current season scoping, and live ref
URLs; use direct detail endpoints for known IDs.
