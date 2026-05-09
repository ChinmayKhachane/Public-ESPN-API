# MLB Competitor Roster Athlete Statistics

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/competitors/17/roster/4414528/statistics/0

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `404`.
- HTTP `404` error payload.

## Example Response

```json
{
  "error": {
    "code": 404,
    "message": "no instance found"
  }
}
```
