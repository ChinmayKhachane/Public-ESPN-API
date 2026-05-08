# Teams

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/teams

Notes:
- Verified with `league=nfl` on 2026-05-08.
- The collection currently resolves to season-scoped team refs for the active league season.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `page` | `int` | Page number |
| `limit` | `int` | Page size |

## Example Response

```json
{
  "count": 32,
  "pageIndex": 1,
  "pageSize": 1,
  "pageCount": 32,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/teams/22?lang=en&region=us"
    }
  ]
}
```
