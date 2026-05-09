# MLB Groups

## https://site.api.espn.com/apis/site/v2/sports/baseball/mlb/groups

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "status": "success",
  "groups": [
    {
      "name": "American League",
      "abbreviation": "AL",
      "children": [
        {
          "name": "American League East",
          "abbreviation": "ALE",
          "teams": [
            {},
            {}
          ]
        },
        {
          "name": "American League Central",
          "abbreviation": "ALC",
          "teams": [
            {},
            {}
          ]
        }
      ]
    },
    {
      "name": "National League",
      "abbreviation": "NL",
      "children": [
        {
          "name": "National League East",
          "abbreviation": "NLE",
          "teams": [
            {},
            {}
          ]
        },
        {
          "name": "National League Central",
          "abbreviation": "NLC",
          "teams": [
            {},
            {}
          ]
        }
      ]
    }
  ]
}
```
