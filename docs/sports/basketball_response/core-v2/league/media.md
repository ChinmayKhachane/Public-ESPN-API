# Media

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/media

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "count": 1291,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 646,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/media/1?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/media/2?lang=en&region=us"
    }
  ]
}
```
