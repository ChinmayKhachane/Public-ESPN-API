# MLB Competition Status

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/status

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/stat...",
  "clock": 0.0,
  "displayClock": "0:00",
  "period": 9,
  "type": {
    "id": "3",
    "name": "STATUS_FINAL",
    "state": "post",
    "completed": true,
    "detail": "Final",
    "shortDetail": "Final"
  },
  "featuredAthletes": [
    {
      "name": "winningPitcher",
      "displayName": "Winning Pitcher",
      "abbreviation": "WP",
      "team": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/18?lang=en&region=us"
      },
      "statistics": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp..."
      },
      "shortDisplayName": "Win",
      "playerId": 4918155,
      "athlete": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/athletes/4918155?lang=en&region=us"
      }
    },
    {
      "name": "losingPitcher",
      "displayName": "Losing Pitcher",
      "abbreviation": "LP",
      "team": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/17?lang=en&region=us"
      },
      "statistics": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp..."
      },
      "shortDisplayName": "Loss",
      "playerId": 42433,
      "athlete": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/athletes/42433?lang=en&region=us"
      }
    }
  ],
  "halfInning": 17,
  "periodPrefix": "End"
}
```
