# Unsupported Or Sparse Core V3 Paths

Verified with live NFL requests on 2026-05-07.

## 404 Paths

```text
https://sports.core.api.espn.com/v3/sports/football/athletes?limit=1
https://sports.core.api.espn.com/v3/sports/football/nfl/leaders
https://sports.core.api.espn.com/v3/sports/football/nfl/leaders?season=2025
```

## Sparse Paths

```text
https://sports.core.api.espn.com/v3/sports/football/nfl/athletes/{id}/statisticslog
https://sports.core.api.espn.com/v3/sports/football/nfl/athletes/{id}/plays
```

Notes:
- `statisticslog` returned only season IDs for the tested NFL quarterback.
- `plays` returned `{ "count": 0, "items": [] }` for two tested NFL athletes.
- Detail endpoints become much more useful when paired with `enable=...` query parameters.
