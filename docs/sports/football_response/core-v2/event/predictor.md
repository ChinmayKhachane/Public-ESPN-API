# Competition Predictor

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/predictor

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988` on 2026-05-08.
- NFL returns `homeTeam` and `awayTeam` blocks with predictor statistics rather than a single `gameProjection` field.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/predictor?lang=en&region=us",
  "name": "Matchup Predictor",
  "homeTeam": {
    "team": {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/teams/17?lang=en&region=us"
    },
    "statistics": [
      {
        "name": "gameProjection",
        "displayName": "WIN PROB",
        "value": 40.572,
        "displayValue": "40.6"
      },
      {
        "name": "matchupQuality",
        "displayName": "Matchup Quality",
        "value": 83.451,
        "displayValue": "83.5"
      }
    ]
  },
  "awayTeam": {
    "team": {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/teams/26?lang=en&region=us"
    },
    "statistics": [
      {
        "name": "gameProjection",
        "displayName": "WIN PROB",
        "value": 59.428000000000004,
        "displayValue": "59.4"
      }
    ]
  }
}
```
