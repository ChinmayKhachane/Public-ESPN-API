# MLB Athlete Hot Zones

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/athletes/4414528/hotzones

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with an empty collection.
- Baseball-specific route from the root README. The tested pitcher returned an empty collection, but the route is valid.

## Example Response

```json
{
  "count": 0,
  "pageIndex": 0,
  "pageSize": 25,
  "pageCount": 0,
  "items": []
}
```
