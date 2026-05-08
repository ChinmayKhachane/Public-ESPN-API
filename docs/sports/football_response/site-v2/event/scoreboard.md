# Scoreboard

## https://site.api.espn.com/apis/site/v2/sports/football/{league}/scoreboard

Notes:
- Verified with `league=nfl` on 2026-05-08.
- The scoreboard response includes the NFL season calendar directly, so the separate `/calendar` path is not needed for NFL.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `dates` | `string` | Specific date in `YYYYMMDD` format |
| `week` | `int` | Week number |
| `seasontype` | `int` | `1=pre`, `2=regular`, `3=post`, `4=off` |
| `limit` | `int` | Event count limit |

## Example Response

```json
{
  "leagues": [
    {
      "id": "28",
      "name": "National Football League",
      "abbreviation": "NFL",
      "slug": "nfl",
      "season": {
        "year": 2025,
        "displayName": "2025",
        "type": {
          "id": "2",
          "name": "Regular Season",
          "abbreviation": "reg"
        }
      },
      "calendar": [
        {
          "label": "Preseason",
          "value": "1"
        },
        {
          "label": "Regular Season",
          "value": "2"
        },
        {
          "label": "Postseason",
          "value": "3"
        },
        {
          "label": "Off Season",
          "value": "4"
        }
      ]
    }
  ],
  "season": {
    "type": 3,
    "year": 2025
  },
  "week": {
    "number": 5
  },
  "events": [
    {
      "id": "401772988"
    }
  ]
}
```
