# Competition Status

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/events/{event}/competitions/{competition}/status

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/status?lang=en&region=us",
  "type": {
    "id": "3",
    "name": "STATUS_FINAL",
    "state": "post",
    "completed": true,
    "description": "Final",
    "detail": "Final",
    "shortDetail": "Final"
  },
  "period": 4,
  "clock": 720.0,
  "displayClock": "12:00"
}
```
