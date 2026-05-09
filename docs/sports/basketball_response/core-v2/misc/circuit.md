# Circuits

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/circuits

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.
- NBA returned an empty circuits collection.

## Example Response

```json
{
  "count": 0,
  "pageIndex": 0,
  "pageSize": 25,
  "pageCount": 0,
  "items": []
}
```
