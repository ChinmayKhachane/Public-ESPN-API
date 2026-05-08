# Competition Power Index Team

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/powerindex/{teamId}

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988`, `teamId=17` on 2026-05-08.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/powerindex/17?lang=en&region=us",
  "season": 2025,
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/teams/17?lang=en&region=us"
  },
  "stats": [
    {
      "name": "teampredptdiff",
      "displayName": "PRED PT DIFF",
      "value": -3.643,
      "displayValue": "-3.6"
    },
    {
      "name": "gameprojection",
      "displayName": "WIN PROB",
      "value": 40.572,
      "displayValue": "40.6%"
    },
    {
      "name": "matchupquality",
      "displayName": "MATCHUP QUALITY",
      "value": 83.451,
      "displayValue": "83.5"
    }
  ]
}
```
