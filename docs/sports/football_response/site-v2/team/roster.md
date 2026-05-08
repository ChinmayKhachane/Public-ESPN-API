# Team Roster

## https://site.api.espn.com/apis/site/v2/sports/football/{league}/teams/{id}/roster

Notes:
- Verified with `league=nfl`, `id=6` on 2026-05-08.
- NFL groups roster entries into offense/defense/special teams buckets.

## Example Response

```json
{
  "team": {
    "id": "6",
    "displayName": "Dallas Cowboys",
    "recordSummary": "7-9-1",
    "seasonSummary": "2025"
  },
  "coach": [
    {
      "id": "17530",
      "firstName": "Brian",
      "lastName": "Schottenheimer"
    }
  ],
  "athletes": [
    {
      "position": "offense",
      "items": [
        {
          "id": "4429202",
          "displayName": "Israel Abanikanda"
        }
      ]
    }
  ],
  "season": {
    "year": 2025
  }
}
```
