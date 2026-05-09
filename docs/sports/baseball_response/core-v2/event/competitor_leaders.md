# MLB Competitor Leaders

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/competitors/17/leaders?limit=5

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp...",
  "id": "0",
  "name": "Total",
  "abbreviation": "Total",
  "categories": [
    {
      "name": "MLBRating",
      "displayName": "MLB Rating",
      "abbreviation": "RAT",
      "shortDisplayName": "RAT",
      "leaders": [
        {
          "team": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/17?lang=en&region=us"
          },
          "statistics": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp..."
          },
          "displayValue": "2-4, 2B, K",
          "value": 63.25,
          "rel": [
            "athlete"
          ],
          "athlete": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/athletes/4917694?lang=en&region=us"
          }
        },
        {
          "team": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/17?lang=en&region=us"
          },
          "statistics": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp..."
          },
          "displayValue": "1-3",
          "value": 60.5,
          "rel": [
            "athlete"
          ],
          "athlete": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/athletes/4422899?lang=en&region=us"
          }
        }
      ]
    },
    {
      "name": "avg",
      "displayName": "Batting Average",
      "abbreviation": "AVG",
      "shortDisplayName": "BA",
      "leaders": [
        {
          "team": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/17?lang=en&region=us"
          },
          "statistics": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp..."
          },
          "displayValue": "2-4, 2B, K",
          "value": 0.5,
          "rel": [
            "athlete"
          ],
          "athlete": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/athletes/4917694?lang=en&region=us"
          }
        },
        {
          "team": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/17?lang=en&region=us"
          },
          "statistics": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp..."
          },
          "displayValue": "1-4, 2 K",
          "value": 0.25,
          "rel": [
            "athlete"
          ],
          "athlete": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/athletes/40538?lang=en&region=us"
          }
        }
      ]
    }
  ],
  "type": "total"
}
```
