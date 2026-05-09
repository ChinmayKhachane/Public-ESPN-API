# MLB Competition Play By ID

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/plays/4018152560000000059

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/play...",
  "id": "4018152560000000059",
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/18?lang=en&region=us"
  },
  "sequenceNumber": "1",
  "type": {
    "id": "59",
    "text": "Start Inning",
    "type": "start-inning"
  },
  "text": "Top of the 1st inning",
  "shortText": "Top of the 1st inning",
  "alternativeText": "Top of the 1st inning",
  "shortAlternativeText": "Top of the 1st inning",
  "awayScore": 0,
  "homeScore": 0,
  "period": {
    "type": "Top",
    "number": 1,
    "displayValue": "1st Inning"
  },
  "valid": true,
  "scoringPlay": false,
  "priority": false,
  "scoreValue": 0,
  "modified": "2026-05-08T22:16Z",
  "wallclock": "2026-05-08T22:11:43Z"
}
```
