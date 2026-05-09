# Agent Instructions

This repository documents ESPN's unofficial public API and includes Django
services that consume parts of it. Treat the documentation as the main product.

## Project Focus

- The API is undocumented, unstable, and sometimes inconsistent across sports,
  leagues, domains, and seasons.
- Prefer observed behavior over assumed URL patterns.
- Clearly document endpoints that are working, empty, sparse, stubbed,
  unsupported, or returning errors.
- Most response examples are samples or truncated examples unless a file
  explicitly says the full response is shown.

## Repository Map

- `docs/` is the main product. It contains the public API reference, sport
  overview pages, live-tested response examples, and SDK planning notes.
- `docs/README.md` is the documentation index. Update it when adding a new
  sport response directory or a major SDK/reference document.
- `docs/sports/{sport}.md` files are sport overview pages. They list league
  slugs, major endpoint families, domain variants, and links to live-tested
  response docs.
- `docs/sports/{sport}_response/` directories contain live-tested response
  documentation for one sport/league family. Treat these directories as the
  source of truth for observed endpoint behavior.
- `espn_service/` is a Django service that consumes parts of the ESPN API.
  Use it for service code changes, not as the primary source for endpoint
  documentation.
- `nhl_service/` is a Django service focused on NHL ingestion/consumption.
  Use it for NHL service changes, not as the primary source for cross-sport
  API behavior.

## Response Directory Layout

Most `docs/sports/{sport}_response/` directories use the same family layout:

- `core-v2/` documents `sports.core.api.espn.com/v2` endpoints. These are
  often ref-heavy graph resources. Collections commonly return `items[].$ref`
  and detail resources commonly expose additional child refs.
- `core-v3/` documents `sports.core.api.espn.com/v3` endpoints. These are
  selective and sometimes sparse or unsupported depending on sport/league.
- `site-v2/` documents `site.api.espn.com/apis/site/v2` endpoints. These are
  usually friendlier app-facing payloads for scoreboards, summaries, teams,
  schedules, rosters, injuries, transactions, news, and similar data.
- `site-v3/` documents `site.web.api.espn.com/apis/common/v3` and related
  site-v3 paths. These are often useful for athlete pages, statistics,
  game logs, splits, and selected team roster/leaderboard views.
- `cdn/` documents `cdn.espn.com/core/...` page-package endpoints. These
  typically return `gamepackageJSON` or page-view payloads for scoreboard,
  game, boxscore, matchup, and play-by-play views.
- `search-and-web/` documents search and web helper endpoints such as Search v2
  and scoreboard header APIs.
- `now/` documents Now feed/news endpoints.
- `specialized/`, when present, documents sport-specific or product-specific
  endpoints such as power index or futures.
- `common.md` files describe shared fixtures, conventions, or endpoint-family
  notes for that response directory.
- `unsupported_*`, `sparse_*`, or `*_stub` files are intentional. They document
  observed empty responses, stubs, `400`, `404`, `500`, or unsupported routes.
  Do not delete these just because they do not return useful data.

## Football SDK Planning Docs

For future football SDK work, read these files in this order:

1. `docs/sports/football_response/sdk_capabilities.md`
   - Defines what the SDK should retrieve from a football fan's perspective.
   - Uses the columns `Input`, `Output`, and `How we get there`.
   - Treat this as the product surface for normalized SDK methods such as
     player search, player game stats, team schedule, game summary,
     play-by-play, standings, injuries, transactions, and news.
2. `docs/sports/football_response/sdk_chaining_matrix.md`
   - Defines which ESPN API family should own each data type.
   - Documents resolver fields the SDK should know before requests, such as
     `items[].$ref`, `competitions[].$ref`, `competitors[].$ref`,
     `events[].id`, and Search v2 player `uid`.
   - Documents player search behavior: Search v2 player `contents[].id` is a
     GUID, while the numeric athlete path ID is parsed from `contents[].uid`
     after `~a:`.
   - Use this file to decide whether an SDK method should call `site-v2`,
     `core-v2`, `common-v3`, CDN, Search v2, or Now.
3. `docs/sports/football_response/api_hierarchy.md`
   - Describes the football API call graph and parent/child endpoint hierarchy.
   - Use this for understanding how league, team, athlete, event, competition,
     competitor, drive, play, probability, and lookup resources relate.
4. `docs/sports/football_response/reference_resolution.md`
   - Documents live-tested URL resolution paths from initial endpoint to final
     data URL for the documented football endpoints.
   - Use this when an endpoint returns `$ref` values or a collection of refs and
     the SDK needs to know which field to follow.
5. Individual endpoint docs under `docs/sports/football_response/**`
   - Use these for exact endpoint URLs, fixtures, query parameters, response
     keys, quirks, and examples.

SDK design principle for football: expose data-oriented methods first, such as
`getPlayerProfile`, `getPlayerGameStats`, `getTeamSchedule`, `getGameSummary`,
and `getPlayByPlay`. Those methods should hide ESPN domain selection and known
reference resolution from normal callers. Keep raw domain clients, such as
`football.site.*`, `football.core.*`, `football.common.*`, `football.cdn.*`,
and `football.search.*`, as escape hatches for callers who need native ESPN
payload shapes.

## Documentation Workflow

- Read existing docs before editing. Match the local file style.
- Use `docs/sports/{sport}_response/` for live-tested response docs for each
  sport.
- Keep endpoint docs concrete:
  - endpoint URL
  - tested league/team/event IDs
  - verification date
  - optional parameters
  - notes about quirks or unsupported behavior
  - example JSON response
- If an endpoint only works with a specific domain variant, document that
  explicitly. Example: NFL standings use `site.api.espn.com/apis/v2/...`, not
  `site.api.espn.com/apis/site/v2/...`.
- If a route returns `{}`, `404`, `400`, or another error, document that result
  instead of deleting the endpoint.

## Adding Response Docs For A Sport

When adding or expanding live-tested response documentation for any sport:

1. Identify the target sport slug and league slug first. If the user asks for a
   specific league, only document that league unless they expand the scope.
2. Review previous response docs such as `docs/sports/football_response/` and
   `docs/sports/basketball_response/` for reusable endpoint families, file
   organization, and response-doc style.
3. Translate the prior sport's endpoint families to the target sport and league
   using real ESPN URL patterns. Do not assume translated URLs work.
4. Choose realistic live fixtures for the target league: season, event,
   competition, team, competitor, athlete, play, and venue IDs as needed.
5. Make real HTTP requests with `curl` for every candidate endpoint. Record the
   HTTP status, content type when relevant, and whether the response is populated,
   sparse, empty, stubbed, or an error payload.
6. Resolve representative non-self `$ref` values with additional `curl` requests
   so examples prove that references are reachable. Do not recursively expand the
   full graph.
7. Document observed results in `docs/sports/{sport}_response/`, including
   working endpoints and endpoints that return `400`, `404`, `500`, `{}`, or
   other notable behavior.
8. After the prior-sport endpoint pass, inspect the root `README.md`,
   `docs/README.md`, and `docs/sports/{sport}.md` for additional general,
   cross-domain, and sport-specific endpoints that may not exist in the previous
   response docs.
9. Test those additional endpoints with `curl` as well, then document the
   results in the same response-doc structure.
10. Add or update links from `docs/sports/{sport}.md` and `docs/README.md` to
    the new response-doc directory when a live-tested response set exists.
11. Validate every JSON code block that was added or edited, and sweep the
    documented URLs to make sure statuses match the docs before reporting the
    final counts.

## Live Verification

- ESPN endpoints generally do not require an API key.
- When asked to verify docs, make real HTTP requests and inspect responses.
- Do not rely only on scripts or prior generated output. Scripts are useful for
  coverage, but manually inspect mismatches before editing docs.
- Use realistic fixtures already present in the docs for the requested sport and
  league. For football/NFL work, current commonly documented fixtures include:
  - `league=nfl`
  - `event=401772988`
  - `competition=401772988`
  - `competitor=17`
  - `team id=6` for site-v2 team examples
  - `team id=22` for core-v2 season team examples
  - `athlete id=4429202` for core-v2 athlete examples
  - `athlete id=4431452` for game/player stat examples

## `$ref` Handling

- ESPN core responses often contain `$ref` links.
- Keep self-referencing `$ref` values when they identify the current resource.
- Resolve non-self references only when doing so improves the example.
- Avoid recursively expanding the entire ESPN graph. It can become cyclic and
  enormous.
- If a fetched reference returns an error, include the error payload and note the
  observed HTTP status.

## Editing Rules

- Do not reformat unrelated documentation.
- Do not change generated-looking endpoint tables unless needed for the task.
- Preserve user or prior-agent edits in the working tree.
- Prefer small, reviewable documentation changes over broad rewrites.
- Use `https://` in new examples unless documenting an exact legacy URL already
  present in the repo.

## Code And Tests

- The Django services live under `espn_service/` and `nhl_service/`.
- Add or update tests when changing service behavior.
- Documentation-only changes usually do not require running the service test
  suites, but validate JSON code blocks when editing response examples.

## Useful Commands

```bash
rg -n 'pattern' docs/sports/{sport}_response
rg --files docs/sports/{sport}_response
python -m json.tool example.json
```
