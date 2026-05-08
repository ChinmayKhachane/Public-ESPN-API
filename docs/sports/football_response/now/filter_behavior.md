# Now API Filter Behavior

## https://now.core.api.espn.com/v1/sports/news?sport=football&limit=2

Notes:
- Verified on 2026-05-07.
- `sport=football` returned football-scoped results.
- Tested `leagues=nfl`, `leagues=28`, `team=dal`, and `team=6` did not independently narrow the result set in the same way.
- Combining `sport=football` with `leagues=nfl` or `leagues=28` still returned a general football feed, including college football and NFL headlines.
- Combining `sport=football` with `team=6` also returned a general football feed in the tested response.

## Tested Filter Results

| Query | HTTP | Observed Behavior |
| --- | --- | --- |
| `?sport=football&limit=2` | `200` | Football-scoped feed |
| `?leagues=nfl&limit=2` | `200` | Looked like global feed |
| `?leagues=28&limit=2` | `200` | Looked like global feed |
| `?sport=football&leagues=nfl&limit=2` | `200` | Football feed, not NFL-only |
| `?team=dal&limit=2` | `200` | Looked like global feed |
| `?team=6&limit=2` | `200` | Looked like global feed |
| `?sport=football&team=6&limit=2` | `200` | Football feed, not Cowboys-only |

## Example Football-Scoped Response

```json
{
  "resultsCount": 2,
  "status": "success",
  "headlines": [
    {
      "headline": "Ex-Ohio coach Brian Smith suing school following December firing",
      "categories": [
        {
          "type": "league",
          "description": "NCAA Football"
        },
        {
          "type": "team",
          "description": "Ohio Bobcats"
        }
      ]
    },
    {
      "headline": "Browns' Owusu-Koramoah (neck) to miss second straight season",
      "categories": [
        {
          "type": "league",
          "description": "NFL"
        },
        {
          "type": "team",
          "description": "Cleveland Browns"
        }
      ]
    }
  ]
}
```
