# Football Core V3 Response Docs

Verified with live NFL requests on 2026-05-07.

- `league/` - league root and season detail
- `athlete/` - athlete collection, athlete detail, statistics log, and plays
- `team/` - team detail
- `unsupported_or_sparse.md` - tested unsupported or sparse paths

Notes:
- Core v3 uses `https://sports.core.api.espn.com/v3/sports/football/nfl/...` for NFL.
- `https://sports.core.api.espn.com/v3/sports/football/athletes` returned `404`; include the league slug.
- The `enable` query parameter is useful on detail endpoints. For example, `enable=team,position,links` expands athlete detail, and `enable=logos,links,groups,venue,record` expands team detail.
