# Competition Plays

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/events/{event}/competitions/{competition}/plays

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.
- NBA play objects include basketball-specific shot fields such as `shootingPlay`, coordinates, and points attempted.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/plays?lang=en&region=us",
  "count": 459,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 230,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/plays/4018711614?lang=en&region=us",
      "id": "4018711614",
      "type": {
        "id": "615",
        "text": "Jumpball"
      },
      "team": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/teams/20?lang=en&region=us"
      },
      "text": "Karl-Anthony Towns vs. Joel Embiid (Kelly Oubre Jr. gains possession)",
      "period": {
        "displayValue": "1st Quarter",
        "number": 1
      },
      "clock": {
        "value": 720.0,
        "displayValue": "12:00"
      },
      "homeScore": 0,
      "awayScore": 0,
      "scoringPlay": false,
      "scoreValue": 0,
      "probability": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/probabilities/4018711614?lang=en&region=us"
      }
    },
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
  ]
}
```
