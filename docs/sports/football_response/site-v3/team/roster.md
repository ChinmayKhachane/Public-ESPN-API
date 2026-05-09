# Team Roster

## https://site.web.api.espn.com/apis/common/v3/sports/football/{league}/teams/{id}/roster

Notes:
- Verified with `league=nfl`, `id=6` and `id=17` on 2026-05-07.
- This is one of the few working NFL `common/v3` team endpoints.
- The response is organized by `positionGroups[]`, not a flat roster array.
- Top-level keys are `season`, `coach`, `positionGroups`, and `team`.
- `coach` is an array, even when there is only one head coach object.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `season` | `int` | Season year. Tested successfully with `season=2025`. |

## Example Response

```json
{
  "season": {
    "year": 2025,
    "type": 2,
    "name": "Regular Season"
  },
  "coach": [
    {
      "id": "17530",
      "firstName": "Brian",
      "lastName": "Schottenheimer",
      "experience": 1
    }
  ],
  "team": {
    "id": "6",
    "displayName": "Dallas Cowboys"
  },
  "positionGroups": [
    {
      "type": "offense",
      "displayName": "Offense",
      "athletes": [
        {
          "id": "4429202",
          "displayName": "Israel Abanikanda",
          "jersey": "30",
          "position": {
            "abbreviation": "RB"
          }
        }
      ]
    },
    {
      "type": "specialTeam",
      "displayName": "Special Teams",
      "athletes": [
        {
          "id": "14950",
          "displayName": "Bryan Anger",
          "jersey": "5",
          "position": {
            "abbreviation": "P"
          }
        }
      ]
    }
  ]
}
```
