# MLB Athlete Statistics Log

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/athletes/4414528/statisticslog

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/athletes/4414528/statisticslog?lang=en&regio...",
  "entries": [
    {
      "statistics": [
        {
          "statistics": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/types/2/athletes/4414528/statis..."
          },
          "type": "total"
        },
        {
          "team": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/17?lang=en&region=us"
          },
          "statistics": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/types/2/teams/17/athletes/44145..."
          },
          "type": "team"
        }
      ],
      "season": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026?lang=en&region=us"
      }
    },
    {
      "statistics": [
        {
          "statistics": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2025/types/2/athletes/4414528/statis..."
          },
          "type": "total"
        },
        {
          "team": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2025/teams/17?lang=en&region=us"
          },
          "statistics": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2025/types/2/teams/17/athletes/44145..."
          },
          "type": "team"
        }
      ],
      "season": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2025?lang=en&region=us"
      }
    }
  ]
}
```
