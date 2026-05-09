# League Injuries

## https://site.api.espn.com/apis/site/v2/sports/basketball/{league}/injuries

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "season": {
    "name": "Postseason",
    "displayName": "2025-26",
    "year": 2026,
    "type": 3
  },
  "status": "success",
  "timestamp": "2026-05-09T03:16:05Z",
  "injuries": [
    {
      "id": "1",
      "displayName": "Atlanta Hawks",
      "injuries": [
        {},
        {}
      ]
    },
    {
      "id": "2",
      "displayName": "Boston Celtics",
      "injuries": [
        {}
      ]
    }
  ]
}
```
