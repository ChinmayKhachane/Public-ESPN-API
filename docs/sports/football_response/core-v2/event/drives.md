# Competition Drives

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/drives

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988` on 2026-05-08.
- Drive objects inline a nested `plays` collection.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `page` | `int` | Page number |
| `limit` | `int` | Page size |

## Example Response

```json
{
  "count": 28,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 14,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/drives/4017729881?lang=en&region=us",
      "id": "4017729881",
      "description": "8 plays, 51 yards, 3:02",
      "result": "FG",
      "yards": 51,
      "offensivePlays": 8,
      "isScore": true,
      "plays": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/drives/4017729881/plays?lang=en&region=us"
      }
    }
  ]
}
```
