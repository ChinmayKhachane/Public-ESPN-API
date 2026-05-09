# Competition Probability ID

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/probabilities/{play}

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988`, `play=40177298840` on 2026-05-08.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/probabilities/40177298840?lang=en&region=us",
  "play": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/plays/40177298840?lang=en&region=us",
    "id": "40177298840",
    "sequenceNumber": "4000",
    "type": {
      "id": "53",
      "text": "Kickoff",
      "abbreviation": "K"
    },
    "text": "A.Borregales kicks 65 yards from NE 35 to end zone, Touchback to the SEA 35.",
    "shortText": "Andy Borregales 65 Yd Kickoff, Touchback",
    "alternativeText": "A.Borregales kicks 65 yards from NE 35 to end zone, Touchback to the SEA 35.",
    "shortAlternativeText": "Andy Borregales 65 Yd Kickoff, Touchback",
    "awayScore": 0,
    "homeScore": 0,
    "period": {
      "number": 1
    },
    "clock": {
      "value": 900.0,
      "displayValue": "15:00"
    },
    "scoringPlay": false,
    "priority": false,
    "scoreValue": 0,
    "modified": "2026-02-09T03:39Z",
    "team": {},
    "participants": [
      {
        "athlete": {},
        "position": {},
        "statistics": {},
        "playStatistics": {},
        "order": 1,
        "type": "kicker"
      }
    ],
    "probability": {},
    "wallclock": "2026-02-08T23:40:23Z",
    "teamParticipants": [
      {
        "team": {},
        "id": "26",
        "order": 1,
        "type": "offense"
      },
      {
        "team": {},
        "id": "17",
        "statistics": {},
        "playStatistics": {},
        "order": 2,
        "type": "defense"
      }
    ],
    "isPenalty": false,
    "statYardage": 0,
    "drive": {},
    "start": {
      "down": 0,
      "distance": 0,
      "yardLine": 35,
      "yardsToEndzone": 65,
      "team": {}
    },
    "end": {
      "down": 1,
      "distance": 10,
      "yardLine": 65,
      "yardsToEndzone": 65,
      "downDistanceText": "1st & 10 at SEA 35",
      "shortDownDistanceText": "1st & 10",
      "possessionText": "SEA 35",
      "team": {}
    },
    "isTurnover": false
  },
  "homeWinPercentage": 0.3766,
  "awayWinPercentage": 0.6234,
  "tiePercentage": 0.0,
  "spreadCoverProbHome": 0.5071,
  "totalOverProb": 0.5049,
  "lastModified": "2026-02-09T02:41Z"
}
```
