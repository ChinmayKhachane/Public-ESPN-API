# Competition Power Index

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/powerindex

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988` on 2026-05-08.
- The collection returns one item per team.

## Example Response

```json
{
  "count": 2,
  "pageIndex": 1,
  "pageSize": 25,
  "pageCount": 1,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/powerindex/17?lang=en&region=us",
      "team": {},
      "season": 2025,
      "stats": [
        {
          "name": "teampredptdiff",
          "displayName": "PRED PT DIFF",
          "description": "Expected margin of victory for the FPI favorite.",
          "abbreviation": "PRED PT DIFF",
          "value": -3.643,
          "displayValue": "-3.6"
        },
        {
          "name": "gameprojection",
          "displayName": "WIN PROB",
          "description": "Team's predicted win percentage in this game at time of given BPI run",
          "abbreviation": "GAME PROJ",
          "value": 40.572,
          "displayValue": "40.6%"
        },
        {
          "name": "matchupquality",
          "displayName": "MATCHUP QUALITY",
          "description": "A measure of projected competitiveness and excitement in the game, using a 0 to 100 scale, with 100 as the most exciting",
          "abbreviation": "MATCHUP QUALITY",
          "value": 83.451,
          "displayValue": "83.5"
        },
        {
          "name": "teamadjgamescore",
          "displayName": "GAME SCORE",
          "description": "A measure of how well a team performed compared to their expected performance and the expected performance of a typical top 25 team.",
          "abbreviation": "TEAM ADJ GAMESCORE",
          "displayValue": ""
        }
      ]
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/powerindex/26?lang=en&region=us",
      "team": {},
      "season": 2025,
      "stats": [
        {
          "name": "teampredptdiff",
          "displayName": "PRED PT DIFF",
          "description": "Expected margin of victory for the FPI favorite.",
          "abbreviation": "PRED PT DIFF",
          "value": 3.643,
          "displayValue": "3.6"
        },
        {
          "name": "gameprojection",
          "displayName": "WIN PROB",
          "description": "Team's predicted win percentage in this game at time of given BPI run",
          "abbreviation": "GAME PROJ",
          "value": 59.428000000000004,
          "displayValue": "59.4%"
        },
        {
          "name": "matchupquality",
          "displayName": "MATCHUP QUALITY",
          "description": "A measure of projected competitiveness and excitement in the game, using a 0 to 100 scale, with 100 as the most exciting",
          "abbreviation": "MATCHUP QUALITY",
          "value": 83.451,
          "displayValue": "83.5"
        },
        {
          "name": "teamadjgamescore",
          "displayName": "GAME SCORE",
          "description": "A measure of how well a team performed compared to their expected performance and the expected performance of a typical top 25 team.",
          "abbreviation": "TEAM ADJ GAMESCORE",
          "displayValue": ""
        }
      ]
    }
  ]
}
```
