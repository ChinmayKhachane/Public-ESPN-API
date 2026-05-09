# MLB Athlete Statistics

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/athletes/4414528/statistics

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/athletes/4414528/statistics/0?lang=en&region=us",
  "athlete": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/athletes/4414528?lang=en&region=us"
  },
  "splits": {
    "id": "0",
    "name": "All Splits",
    "abbreviation": "Total",
    "categories": [
      {
        "name": "pitching",
        "displayName": "Pitching",
        "abbreviation": "p",
        "shortDisplayName": "Pitching",
        "stats": [
          {
            "name": "gamesPlayed",
            "displayName": "Games Played",
            "abbreviation": "GP",
            "shortDisplayName": "GP",
            "value": 83.0,
            "displayValue": "83"
          },
          {
            "name": "teamGamesPlayed",
            "displayName": "Team Games Played",
            "abbreviation": "G",
            "shortDisplayName": "G",
            "value": 524.0,
            "displayValue": "524"
          }
        ]
      },
      {
        "name": "fielding",
        "displayName": "Fielding",
        "abbreviation": "f",
        "shortDisplayName": "Fielding",
        "stats": [
          {
            "name": "gamesPlayed",
            "displayName": "Games Played",
            "abbreviation": "GP",
            "shortDisplayName": "GP",
            "value": 62.0,
            "displayValue": "62"
          },
          {
            "name": "teamGamesPlayed",
            "displayName": "Team Games Played",
            "abbreviation": "G",
            "shortDisplayName": "G",
            "value": 362.0,
            "displayValue": "362"
          }
        ]
      }
    ]
  }
}
```
