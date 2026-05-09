# MLB Venue By ID

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/venues/17

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/venues/17?lang=en&region=us",
  "id": "17",
  "guid": "2b3d674f-5706-347a-901a-f034fb6eade6",
  "fullName": "Cinergy Field",
  "shortName": "Cinergy Field",
  "address": {
    "city": "Cincinnati",
    "state": "Ohio",
    "zipCode": "45999"
  },
  "grass": true,
  "indoor": false,
  "images": []
}
```
