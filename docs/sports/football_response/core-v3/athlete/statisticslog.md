# Athlete Statistics Log

## https://sports.core.api.espn.com/v3/sports/football/{league}/athletes/{id}/statisticslog

Notes:
- Verified with `league=nfl`, `id=4431452` on 2026-05-07.
- The tested NFL response was very sparse: it returned only season IDs in `items`.

## Example Response

```json
{
  "count": 2,
  "items": [
    {
      "id": "2024"
    },
    {
      "id": "2025"
    }
  ]
}
```
