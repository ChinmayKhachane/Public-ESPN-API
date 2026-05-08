# League Injuries

## https://site.api.espn.com/apis/site/v2/sports/football/{league}/injuries

Notes:
- Verified with `league=nfl` on 2026-05-08.
- Response is grouped by team and often contains offseason transaction-style notes, not just active in-game injury reports.

## Example Response

```json
{
  "season": {
    "year": 2025,
    "type": 4,
    "name": "Off Season",
    "displayName": "2025"
  },
  "status": "success",
  "injuries": [
    {
      "id": "22",
      "displayName": "Arizona Cardinals",
      "injuries": [
        {
          "id": "630675",
          "status": "Active",
          "athlete": {
            "displayName": "Jameson Geers",
            "position": {
              "abbreviation": "TE"
            },
            "team": {
              "id": "22",
              "displayName": "Arizona Cardinals"
            }
          }
        }
      ]
    }
  ]
}
```
