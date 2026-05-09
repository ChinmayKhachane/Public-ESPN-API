# MLB Competitor By ID

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/competitors/17

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp...",
  "id": "17",
  "uid": "s:1~l:10~t:17",
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/17?lang=en&region=us"
  },
  "statistics": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp..."
  },
  "type": "team",
  "order": 0,
  "homeAway": "home",
  "winner": false,
  "score": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp..."
  },
  "linescores": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp..."
  },
  "roster": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp..."
  },
  "leaders": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp..."
  },
  "record": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp..."
  },
  "probables": [
    {
      "name": "probableStartingPitcher",
      "displayName": "Probable Starting Pitcher",
      "abbreviation": "SP",
      "statistics": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/types/2/athletes/42433/statisti..."
      },
      "shortDisplayName": "Starter",
      "playerId": 42433,
      "athlete": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/athletes/42433?lang=en&region=us"
      }
    }
  ]
}
```
