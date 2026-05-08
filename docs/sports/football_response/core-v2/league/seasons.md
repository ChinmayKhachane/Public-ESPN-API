# Seasons

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/seasons

Notes:
- Verified with `league=nfl` on 2026-05-08.
- The collection is historical and includes future season containers.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `page` | `int` | Page number |
| `limit` | `int` | Page size |

## Example Response

```json
{
  "count": 105,
  "pageIndex": 1,
  "pageSize": 3,
  "pageCount": 35,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2026?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2024?lang=en&region=us"
    }
  ]
}
```
