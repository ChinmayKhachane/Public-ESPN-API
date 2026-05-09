# MLB Team By ID

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/17

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/17?lang=en&region=us",
  "id": "17",
  "uid": "s:1~l:10~t:17",
  "guid": "04b65a0b-3cca-d795-0e21-23606470418a",
  "name": "Reds",
  "displayName": "Cincinnati Reds",
  "abbreviation": "CIN",
  "slug": "cincinnati-reds",
  "events": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/17/events?lang=en&region=us"
  },
  "athletes": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/17/athletes?lang=en&regio..."
  },
  "statistics": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/types/2/teams/17/statistics?lan..."
  },
  "alternateIds": {
    "sdr": "25"
  },
  "location": "Cincinnati",
  "shortDisplayName": "Reds",
  "color": "c6011f",
  "alternateColor": "ffffff",
  "isActive": true,
  "isAllStar": false
}
```
