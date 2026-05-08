# Competition Play ID

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/plays/{play}

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988`, `play=40177298840` on 2026-05-08.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/plays/40177298840?lang=en&region=us",
  "id": "40177298840",
  "sequenceNumber": "4000",
  "type": {
    "id": "53",
    "text": "Kickoff",
    "abbreviation": "K"
  },
  "text": "A.Borregales kicks 65 yards from NE 35 to end zone, Touchback to the SEA 35.",
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/teams/26?lang=en&region=us"
  },
  "probability": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/probabilities/40177298840?lang=en&region=us"
  },
  "drive": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/drives/4017729881?lang=en&region=us"
  },
  "participants": [
    {
      "athlete": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/athletes/4569923?lang=en&region=us"
      },
      "type": "kicker"
    }
  ]
}
```
