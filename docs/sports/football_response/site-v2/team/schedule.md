# Team Schedule

## https://site.api.espn.com/apis/site/v2/sports/football/{league}/teams/{id}/schedule

Notes:
- Verified with `league=nfl`, `id=6` on 2026-05-08.

## Example Response

```json
{
  "team": {
    "id": "6",
    "displayName": "Dallas Cowboys",
    "recordSummary": "7-9-1",
    "standingSummary": "2nd in NFC East"
  },
  "byeWeek": 10,
  "season": {
    "year": 2025
  },
  "events": [
    {
      "id": "401772510"
    }
  ]
}
```
