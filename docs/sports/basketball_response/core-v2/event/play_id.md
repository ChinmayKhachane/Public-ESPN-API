# Competition Play ID

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/events/{event}/competitions/{competition}/plays/{play}

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/plays/4018711617?lang=en&region=us",
  "id": "4018711617",
  "type": {
    "id": "110",
    "text": "Driving Layup Shot"
  },
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/teams/20?lang=en&region=us"
  },
  "text": "Kelly Oubre Jr. makes driving layup",
  "shortText": "K. Oubre Jr. makes driving layup",
  "period": {
    "displayValue": "1st Quarter",
    "number": 1
  },
  "clock": {
    "value": 715.0,
    "displayValue": "11:55"
  },
  "homeScore": 2,
  "awayScore": 0,
  "scoringPlay": true,
  "scoreValue": 2
}
```
