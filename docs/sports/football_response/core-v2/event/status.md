# Competition Status

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/status

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988` on 2026-05-08.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/status?lang=en&region=us",
  "clock": 0.0,
  "displayClock": "0:00",
  "period": 4,
  "type": {
    "id": "3",
    "name": "STATUS_FINAL",
    "state": "post",
    "completed": true,
    "description": "Final",
    "detail": "Final",
    "shortDetail": "Final"
  }
}
```
