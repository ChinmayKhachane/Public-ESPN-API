# Calendar

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/calendar

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `dates` | `string` | Date or date range |
| `seasontype` | `int` | Season type |
| `limit` | `int` | Page size |

## Example Response

```json
{
  "count": 4,
  "pageIndex": 1,
  "pageSize": 25,
  "pageCount": 1,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/calendar/ondays?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/calendar/offdays?lang=en&region=us"
    }
  ]
}
```
