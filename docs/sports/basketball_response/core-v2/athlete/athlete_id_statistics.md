# Athlete Statistics

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/athletes/{id}/statistics

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/athletes/3059318/statistics/0?lang=en&region=us",
  "athlete": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/athletes/3059318?lang=en&region=us"
  },
  "splits": {
    "id": "0",
    "name": "All Splits",
    "abbreviation": "Total",
    "type": "total",
    "categories": [
      {
        "name": "defensive",
        "displayName": "Defensive",
        "shortDisplayName": "Defensive",
        "abbreviation": "def",
        "summary": "",
        "stats": []
      },
      {
        "name": "general",
        "displayName": "General",
        "shortDisplayName": "General",
        "abbreviation": "gen",
        "summary": "",
        "stats": []
      }
    ]
  }
}
```
