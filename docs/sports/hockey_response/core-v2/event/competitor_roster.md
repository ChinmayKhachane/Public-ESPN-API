# Competitor Roster

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}/roster

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15/roster?limit=2`

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15/roster?lang=en&region=us",
  "entries": [
    {
      "playerId": 5337,
      "jersey": "10",
      "athlete": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/athletes/5337?lang=en&region=us"
      },
      "position": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/positions/2?lang=en&region=us"
      },
      "displayName": "Wilson",
      "projections": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15/roster/5337/projections?lang=en&region=us"
      },
      "scratched": true,
      "scratchReason": "SCRATCHED"
    },
    {
      "playerId": 2335062,
      "jersey": "41",
      "athlete": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/athletes/2335062?lang=en&region=us"
      },
      "position": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/positions/1?lang=en&region=us"
      },
      "statistics": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15/roster/2335062/statistics/0?lang=en&region=us"
      },
      "displayName": "Glendening",
      "projections": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15/roster/2335062/projections?lang=en&region=us"
      },
      "scratched": false
    }
  ],
  "competition": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412?lang=en&region=us"
  },
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/teams/15?lang=en&region=us"
  }
}
```
