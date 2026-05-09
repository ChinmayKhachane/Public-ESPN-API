# MLB Officials

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/officials?limit=5

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "count": 4,
  "pageIndex": 1,
  "pageSize": 5,
  "pageCount": 1,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/offi...",
      "displayName": "Jeremie Rehak",
      "firstName": "Jeremie",
      "lastName": "Rehak ",
      "position": {
        "id": "206",
        "name": "Home Plate Umpire",
        "displayName": "Home Plate Umpire"
      },
      "order": 1
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/offi...",
      "displayName": "David Rackley",
      "firstName": "David",
      "lastName": "Rackley ",
      "position": {
        "id": "207",
        "name": "First Base Umpire",
        "displayName": "First Base Umpire"
      },
      "order": 2
    }
  ]
}
```
