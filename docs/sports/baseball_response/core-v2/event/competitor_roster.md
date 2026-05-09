# MLB Competitor Roster

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/competitors/17/roster?limit=5

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp...",
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/17?lang=en&region=us"
  },
  "entries": [
    {
      "displayName": "Friedl",
      "statistics": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp..."
      },
      "playerId": 36020,
      "period": 1,
      "active": true,
      "starter": true,
      "forPlayerId": 0,
      "valid": true,
      "athlete": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/athletes/36020?lang=en&region=us"
      },
      "position": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/positions/8?lang=en&region=us"
      }
    },
    {
      "displayName": "Bleday",
      "statistics": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp..."
      },
      "playerId": 42410,
      "period": 1,
      "active": true,
      "starter": true,
      "forPlayerId": 0,
      "valid": true,
      "athlete": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/athletes/42410?lang=en&region=us"
      },
      "position": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/positions/7?lang=en&region=us"
      }
    }
  ],
  "pitchers": [
    {
      "displayName": "Lodolo",
      "statistics": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp..."
      },
      "playerId": 42433,
      "period": 1,
      "active": false,
      "starter": true,
      "forPlayerId": 0,
      "valid": true,
      "athlete": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/athletes/42433?lang=en&region=us"
      },
      "position": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/positions/1?lang=en&region=us"
      }
    },
    {
      "displayName": "Johnson",
      "statistics": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp..."
      },
      "playerId": 32777,
      "period": 6,
      "active": false,
      "starter": false,
      "forPlayerId": 0,
      "valid": true,
      "athlete": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/athletes/32777?lang=en&region=us"
      },
      "position": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/positions/1?lang=en&region=us"
      }
    }
  ],
  "competition": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256?lang..."
  },
  "lastBatterOrder": "6"
}
```
