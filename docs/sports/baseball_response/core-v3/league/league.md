# Core v3 MLB League

## https://sports.core.api.espn.com/v3/sports/baseball/mlb

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "id": "10",
  "uid": "s:1~l:10",
  "guid": "b38f959b-7865-31ac-8841-b88355519e10",
  "name": "Major League Baseball",
  "displayName": "MLB",
  "abbreviation": "MLB",
  "slug": "mlb",
  "groupId": "9",
  "shortName": "MLB",
  "midsizeName": "MLB",
  "color": "041e42"
}
```
