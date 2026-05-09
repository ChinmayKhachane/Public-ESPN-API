# MLB Site Standings Stub

## https://site.api.espn.com/apis/site/v2/sports/baseball/mlb/standings

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.
- The `/apis/site/v2/` standings path is a stub/redirect-style payload for MLB; use `/apis/v2/` for standings.

## Example Response

```json
{
  "fullViewLink": {
    "text": "Full Standings",
    "href": "https://www.espn.com/mlb/standings"
  }
}
```
