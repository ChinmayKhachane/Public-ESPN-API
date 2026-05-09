# Athlete Statistics

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/athletes/{athlete}/statistics

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/athletes/4565230/statistics?seasontype=3`

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/athletes/4565230/statistics/0?lang=en&region=us",
  "athlete": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/athletes/4565230?lang=en&region=us"
  },
  "splits": {
    "id": "0",
    "name": "All Splits",
    "abbreviation": "TOTAL",
    "type": "total",
    "categories": [
      {
        "name": "defensive",
        "displayName": "Defensive",
        "shortDisplayName": "Defensive",
        "abbreviation": "def",
        "summary": "",
        "stats": [
          {},
          {}
        ]
      },
      {
        "name": "general",
        "displayName": "General",
        "shortDisplayName": "General",
        "abbreviation": "gen",
        "summary": "",
        "stats": [
          {},
          {}
        ]
      }
    ]
  }
}
```
