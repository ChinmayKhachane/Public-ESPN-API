# MLB Sparse Team Routes

## https://site.api.espn.com/apis/site/v2/sports/baseball/mlb/teams/17/record

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with an empty object.
- The documented `record` route is sparse for MLB on the tested URL.

## Example Response

```json
{}
```
