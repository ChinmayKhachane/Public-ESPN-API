# MLB Seasons

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons?limit=5

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "count": 151,
  "pageIndex": 1,
  "pageSize": 5,
  "pageCount": 31,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2025?lang=en&region=us"
    }
  ]
}
```
