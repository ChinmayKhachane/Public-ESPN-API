# Teams

## https://site.api.espn.com/apis/site/v2/sports/football/{league}/teams

Notes:
- Verified with `league=nfl` on 2026-05-08.
- NFL team collection is nested under `sports[].leagues[].teams[]` rather than using top-level pagination fields.

## Example Response

```json
{
  "sports": [
    {
      "leagues": [
        {
          "teams": [
            {
              "team": {
                "id": "22",
                "displayName": "Arizona Cardinals",
                "abbreviation": "ARI"
              }
            }
          ]
        }
      ]
    }
  ]
}
```
