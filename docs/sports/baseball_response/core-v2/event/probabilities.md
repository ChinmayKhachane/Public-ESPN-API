# MLB Win Probabilities

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/probabilities?limit=5

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "count": 73,
  "pageIndex": 1,
  "pageSize": 5,
  "pageCount": 15,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/prob...",
      "competition": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256?lang..."
      },
      "play": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/play..."
      },
      "homeTeam": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/17?lang=en&region=us"
      },
      "awayTeam": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/18?lang=en&region=us"
      },
      "tiePercentage": 0.0,
      "homeWinPercentage": 0.594,
      "awayWinPercentage": 0.406,
      "lastModified": "2026-05-08T22:11Z",
      "sequenceNumber": "159480244"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/prob...",
      "competition": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256?lang..."
      },
      "play": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/play..."
      },
      "homeTeam": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/17?lang=en&region=us"
      },
      "awayTeam": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/18?lang=en&region=us"
      },
      "tiePercentage": 0.0,
      "homeWinPercentage": 0.564,
      "awayWinPercentage": 0.43600000000000005,
      "lastModified": "2026-05-08T22:13Z",
      "sequenceNumber": "159584329"
    }
  ]
}
```
