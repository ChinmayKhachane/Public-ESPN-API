# MLB Team Power Index

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/powerindex/17

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `400`.
- HTTP `400` error payload.
- Root README lists this as a general specialized endpoint. MLB returned an unsupported error.

## Example Response

```json
{
  "error": {
    "code": 400,
    "message": "getSeasonPowerIndexByTeamSeason() not supported for baseball/mlb"
  }
}
```
