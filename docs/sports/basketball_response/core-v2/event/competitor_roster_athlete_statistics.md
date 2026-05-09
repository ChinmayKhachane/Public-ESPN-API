# Competition Roster Athlete Statistics

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}/roster/{athlete}/statistics/{split}

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20/roster/3133603/statistics/0?lang=en&region=us",
  "athlete": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/athletes/3133603?lang=en&region=us"
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
  },
  "competition": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161?lang=en&region=us"
  }
}
```
