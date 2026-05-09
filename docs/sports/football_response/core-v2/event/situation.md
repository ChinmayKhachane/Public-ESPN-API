# Competition Situation

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/situation

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988` on 2026-05-08.
- For a completed game, possession is omitted, but the last play ref is still useful.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/situation?lang=en&region=us",
  "down": 0,
  "distance": 0,
  "yardLine": 71,
  "homeTimeouts": 0,
  "awayTimeouts": 0,
  "isRedZone": false,
  "lastPlay": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/plays/4017729884651?lang=en&region=us",
    "id": "4017729884651",
    "sequenceNumber": "465100",
    "type": {
      "id": "66",
      "text": "End of Game",
      "abbreviation": "EG"
    },
    "text": "END GAME",
    "shortText": "End Game ",
    "alternativeText": "END GAME",
    "shortAlternativeText": "End Game ",
    "awayScore": 29,
    "homeScore": 13,
    "period": {
      "number": 4
    },
    "clock": {
      "value": 0.0,
      "displayValue": "0:00"
    },
    "scoringPlay": false,
    "priority": false,
    "scoreValue": 0,
    "modified": "2026-02-09T03:39Z",
    "probability": {},
    "teamParticipants": [
      {
        "team": {},
        "id": "17",
        "order": 1
      },
      {
        "team": {},
        "id": "26",
        "order": 2
      }
    ],
    "isPenalty": false,
    "statYardage": 0,
    "drive": {},
    "start": {
      "down": 0,
      "distance": 0,
      "yardLine": 0,
      "yardsToEndzone": 0
    },
    "end": {
      "down": 0,
      "distance": 0,
      "yardLine": 71,
      "yardsToEndzone": 29
    },
    "isTurnover": false
  }
}
```
