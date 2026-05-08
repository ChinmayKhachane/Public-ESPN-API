# Football Site V3 Response Docs

Verified with live NFL requests on 2026-05-07.

- `site-api/` - `site.api.espn.com/apis/site/v3/...` NFL endpoints
- `athlete/` - `site.web.api.espn.com/apis/common/v3/...` NFL athlete endpoints
- `team/` - working `site.web` team roster coverage
- `statistics/` - `byathlete` and `byteam` leaderboard-style endpoints
- `unsupported_paths.md` - tested `404`/unsupported NFL site-v3 paths

Notes:
- The `site.api` NFL v3 scoreboard and summary endpoints currently return `404`.
- The active NFL v3 surface is `site.web.api.espn.com/apis/common/v3/...`.
- NFL `common/v3` is selective: athlete pages, team rosters, and team stat leaderboards work, but most league, event, and generic team roots do not.
