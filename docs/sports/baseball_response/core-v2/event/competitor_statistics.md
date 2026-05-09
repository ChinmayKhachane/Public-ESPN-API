# MLB Competitor Statistics

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/competitors/17/statistics?limit=5

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp...",
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/17?lang=en&region=us"
  },
  "competition": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256?lang..."
  },
  "splits": {
    "id": "0",
    "name": "All Splits",
    "abbreviation": "Total",
    "categories": [
      {
        "name": "batting",
        "displayName": "Batting",
        "abbreviation": "b",
        "athletes": [
          {
            "statistics": {},
            "athlete": {}
          },
          {
            "statistics": {},
            "athlete": {}
          }
        ],
        "shortDisplayName": "Batting",
        "stats": [
          {
            "name": "gamesPlayed",
            "displayName": "Games Played",
            "abbreviation": "GP",
            "shortDisplayName": "GP",
            "value": 1.0,
            "displayValue": "1"
          },
          {
            "name": "teamGamesPlayed",
            "displayName": "Team Games Played",
            "abbreviation": "G",
            "shortDisplayName": "G",
            "value": 1.0,
            "displayValue": "1"
          }
        ]
      },
      {
        "name": "pitching",
        "displayName": "Pitching",
        "abbreviation": "p",
        "athletes": [
          {
            "statistics": {},
            "athlete": {}
          },
          {
            "statistics": {},
            "athlete": {}
          }
        ],
        "shortDisplayName": "Pitching",
        "stats": [
          {
            "name": "gamesPlayed",
            "displayName": "Games Played",
            "abbreviation": "GP",
            "shortDisplayName": "GP",
            "value": 1.0,
            "displayValue": "1"
          },
          {
            "name": "teamGamesPlayed",
            "displayName": "Team Games Played",
            "abbreviation": "G",
            "shortDisplayName": "G",
            "value": 1.0,
            "displayValue": "1"
          }
        ]
      }
    ]
  }
}
```
