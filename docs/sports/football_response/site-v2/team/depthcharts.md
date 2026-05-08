# Team Depth Charts

## https://site.api.espn.com/apis/site/v2/sports/football/{league}/teams/{id}/depthcharts

Notes:
- Verified with `league=nfl`, `id=6` on 2026-05-08.
- NFL returns `depthchart` as an array of schemes/packages, each with a `positions` object keyed by role codes.

## Example Response

```json
{
  "team": {
    "id": "6",
    "displayName": "Dallas Cowboys"
  },
  "depthchart": [
    {
      "id": "15",
      "name": "Base 3-4 D",
      "positions": {
        "lde": {
          "position": {
            "id": "11",
            "displayName": "Left Defensive End",
            "abbreviation": "LDE"
          },
          "athletes": [
            {
              "id": "4040982",
              "displayName": "Quinnen Williams"
            }
          ]
        }
      }
    }
  ]
}
```
