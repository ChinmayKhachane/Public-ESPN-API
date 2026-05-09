# MLB Athlete By ID

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/athletes/4414528

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/athletes/4414528?lang=en&region=us",
  "id": "4414528",
  "uid": "s:1~l:10~a:4414528",
  "guid": "3a81ab59-a326-31e7-9c98-0f9c09e26e18",
  "displayName": "Andrew Abbott",
  "slug": "andrew-abbott",
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/17?lang=en&region=us"
  },
  "statistics": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/athletes/4414528/statistics?lang=en&region=us"
  },
  "status": {
    "id": "1",
    "name": "Active",
    "abbreviation": "Active",
    "type": "active"
  },
  "type": "baseball",
  "alternateIds": {
    "sdr": "4414528"
  },
  "firstName": "Andrew",
  "lastName": "Abbott",
  "fullName": "Andrew Abbott",
  "shortName": "A. Abbott",
  "weight": 192.0,
  "displayWeight": "192 lbs",
  "height": 72.0
}
```
