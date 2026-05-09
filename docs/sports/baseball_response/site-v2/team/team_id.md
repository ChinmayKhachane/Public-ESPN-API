# MLB Site Team By ID

## https://site.api.espn.com/apis/site/v2/sports/baseball/mlb/teams/17

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "team": {
    "id": "17",
    "uid": "s:1~l:10~t:17",
    "name": "Reds",
    "displayName": "Cincinnati Reds",
    "abbreviation": "CIN",
    "slug": "cincinnati-reds",
    "location": "Cincinnati",
    "shortDisplayName": "Reds",
    "color": "c6011f",
    "alternateColor": "ffffff"
  }
}
```
