# Competition Competitor Roster

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}/roster

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20/roster?lang=en&region=us",
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/teams/20?lang=en&region=us"
  },
  "entries": [
    {
      "displayName": "George",
      "athlete": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/athletes/4251?lang=en&region=us"
      },
      "statistics": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20/roster/4251/statistics/0?lang=en&region=us"
      },
      "period": 0,
      "playerId": 4251,
      "active": false,
      "starter": true,
      "forPlayerId": 0,
      "jersey": "8",
      "valid": false,
      "position": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/positions/7?lang=en&region=us"
      },
      "didNotPlay": false
    },
    {
      "displayName": "Embiid",
      "athlete": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/athletes/3059318?lang=en&region=us"
      },
      "statistics": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20/roster/3059318/statistics/0?lang=en&region=us"
      },
      "period": 0,
      "playerId": 3059318,
      "active": false,
      "starter": true,
      "forPlayerId": 0,
      "jersey": "21",
      "valid": false,
      "position": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/positions/9?lang=en&region=us"
      },
      "didNotPlay": false
    }
  ],
  "competition": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161?lang=en&region=us"
  }
}
```
