# Competition Competitor Statistics

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}/statistics

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988`, `competitor=17` on 2026-05-08.
- This is team-level game stats with athlete refs nested under each stat category.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17/statistics?lang=en&region=us",
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/teams/17?lang=en&region=us"
  },
  "competition": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988?lang=en&region=us"
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
