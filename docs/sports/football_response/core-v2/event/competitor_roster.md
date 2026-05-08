# Competition Competitor Roster

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}/roster

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988`, `competitor=17` on 2026-05-08.
- This is an object with `entries`, not a paginated collection.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17/roster?lang=en&region=us",
  "competition": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988?lang=en&region=us"
  },
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/teams/17?lang=en&region=us"
  },
  "entries": [
    {
      "playerId": 16771,
      "starter": true,
      "jersey": "76",
      "athlete": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/athletes/16771?lang=en&region=us"
      },
      "position": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/positions/46?lang=en&region=us"
      },
      "statistics": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17/roster/16771/statistics/0?lang=en&region=us"
      },
      "displayName": "Moses"
    }
  ]
}
```
