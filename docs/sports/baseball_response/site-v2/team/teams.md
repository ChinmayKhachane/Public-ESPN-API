# MLB Site Teams

## https://site.api.espn.com/apis/site/v2/sports/baseball/mlb/teams

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "sports": [
    {
      "id": "1",
      "uid": "s:1",
      "name": "Baseball",
      "slug": "baseball",
      "leagues": [
        {
          "id": "10",
          "uid": "s:1~l:10",
          "name": "Major League Baseball",
          "abbreviation": "MLB",
          "slug": "mlb",
          "season": {
            "displayName": "2026",
            "year": 2026
          },
          "shortName": "MLB",
          "teams": [
            {},
            {}
          ],
          "year": 2026
        }
      ]
    }
  ]
}
```
