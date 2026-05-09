# Play Detail

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/events/{event}/competitions/{competition}/plays/{play}

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/plays/401871412000000510`

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/plays/401871412000000510?lang=en&region=us",
  "id": "401871412000000510",
  "sequenceNumber": "3",
  "type": {
    "id": "502",
    "text": "Face Off",
    "abbreviation": "faceoff"
  },
  "text": "Luke Glendening faceoff won against Jordan Staal",
  "shortText": "Luke Glendening faceoff won against Jordan Staal",
  "alternativeText": "Luke Glendening faceoff won against Jordan Staal",
  "shortAlternativeText": "Luke Glendening faceoff won against Jordan Staal",
  "awayScore": 0,
  "homeScore": 0,
  "period": {
    "number": 1,
    "displayValue": "1st"
  },
  "clock": {
    "value": 0.0,
    "displayValue": "0:00"
  },
  "scoringPlay": false,
  "priority": false,
  "scoreValue": 0,
  "modified": "2026-05-08T15:11Z",
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/teams/15?lang=en&region=us"
  },
  "participants": [
    {
      "playerId": 2335062,
      "athlete": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/athletes/2335062?lang=en&region=us"
      },
      "position": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/positions/1?lang=en&region=us"
      },
      "statistics": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15/roster/2335062/statistics/0?lang=en&region=us"
      },
      "projections": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15/roster/2335062/projections?lang=en&region=us"
      },
      "order": 1,
      "type": "face-off-winner",
      "ytdGoals": 1
    },
    {
      "playerId": 3541,
      "athlete": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/athletes/3541?lang=en&region=us"
      },
      "position": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/positions/1?lang=en&region=us"
      },
      "statistics": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/7/roster/3541/statistics/0?lang=en&region=us"
      },
      "projections": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/7/roster/3541/projections?lang=en&region=us"
      },
      "order": 2,
      "ytdAssists": 2
    }
  ]
}
```
