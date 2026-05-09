# Casinos

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/casinos

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "count": 62,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 31,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/casinos/50?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/casinos/37?lang=en&region=us"
    }
  ]
}
```
