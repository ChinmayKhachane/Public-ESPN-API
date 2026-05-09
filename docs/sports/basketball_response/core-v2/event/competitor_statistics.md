# Competition Competitor Statistics

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}/statistics

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20/statistics/0?lang=en&region=us",
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/teams/20?lang=en&region=us"
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
        "athletes": [],
        "summary": "",
        "stats": []
      },
      {
        "name": "general",
        "displayName": "General",
        "shortDisplayName": "General",
        "abbreviation": "gen",
        "athletes": [],
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
