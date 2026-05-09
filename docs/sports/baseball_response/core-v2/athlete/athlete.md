# MLB Athletes Collection

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/athletes?limit=5&active=true

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "count": 3716,
  "pageIndex": 1,
  "pageSize": 5,
  "pageCount": 744,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/athletes/4781?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/athletes/4826?lang=en&region=us"
    }
  ]
}
```
