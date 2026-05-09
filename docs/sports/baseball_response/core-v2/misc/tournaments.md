# MLB Tournaments

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/tournaments?limit=5

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `400`.
- HTTP `400` error payload.

## Example Response

```json
{
  "error": {
    "code": 400,
    "message": "getTournaments() not supported for baseball/mlb"
  }
}
```
