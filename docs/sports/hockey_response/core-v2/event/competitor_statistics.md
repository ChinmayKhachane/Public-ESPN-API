# Competitor Statistics

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}/statistics

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15/statistics`

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15/statistics/0?lang=en&region=us",
  "competition": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412?lang=en&region=us"
  },
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/teams/15?lang=en&region=us"
  },
  "splits": {
    "id": "0",
    "name": "Season",
    "abbreviation": "Season",
    "type": "season",
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
        ],
        "athletes": [
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
        ],
        "athletes": [
          {},
          {}
        ]
      }
    ]
  }
}
```
