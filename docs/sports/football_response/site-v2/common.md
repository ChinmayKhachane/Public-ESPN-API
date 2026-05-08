# Football Site V2 Response Docs

Verified with live NFL requests on 2026-05-08.

- `event/` - scoreboard and summary
- `team/` - teams, team detail, roster, schedule, depth charts, and sparse team subresources
- `league/` - news, standings, injuries, transactions, statistics, groups, draft, and unsupported calendar/rankings
- `athlete/` - athlete news

Notes:
- Standings are the one NFL exception that use `https://site.api.espn.com/apis/v2/...` rather than `/apis/site/v2/...`.
- Several team leaf endpoints currently return `{}` for NFL during the offseason.
