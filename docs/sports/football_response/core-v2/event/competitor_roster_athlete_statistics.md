# Competition Competitor Roster Athlete Statistics

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}/roster/{athlete}/statistics/{split}

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988`, `competitor=17`, `athlete=4431452`, `split=0` on 2026-05-08.
- This is a player’s game stat line inside the event roster tree.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17/roster/4431452/statistics/0?lang=en&region=us",
  "athlete": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/athletes/4431452?lang=en&region=us"
  },
  "splits": {
    "categories": [
      {
        "name": "general",
        "displayName": "General"
      },
      {
        "name": "passing",
        "displayName": "Passing"
      },
      {
        "name": "rushing",
        "displayName": "Rushing"
      }
    ]
  }
}
```
