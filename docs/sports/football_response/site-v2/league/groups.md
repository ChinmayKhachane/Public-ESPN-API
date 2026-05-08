# Groups

## https://site.api.espn.com/apis/site/v2/sports/football/{league}/groups

Notes:
- Verified with `league=nfl` on 2026-05-08.
- Response is a conference/division tree with team cards at the leaf level.

## Example Response

```json
{
  "status": "success",
  "groups": [
    {
      "name": "American Football Conference",
      "abbreviation": "AFC",
      "children": [
        {
          "name": "AFC East",
          "abbreviation": "EAST",
          "teams": [
            {
              "id": "2",
              "displayName": "Buffalo Bills",
              "abbreviation": "BUF"
            }
          ]
        }
      ]
    }
  ]
}
```
