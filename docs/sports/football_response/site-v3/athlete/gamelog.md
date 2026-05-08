# Athlete Game Log

## https://site.web.api.espn.com/apis/common/v3/sports/football/{league}/athletes/{id}/gamelog

Notes:
- Verified with `league=nfl`, `id=4431452` on 2026-05-08.
- NFL game logs store `events` as an object keyed by event ID, not an array.
- Some players with limited usage may return only filters and no populated `events`.

## Example Response

```json
{
  "filters": [
    {
      "displayName": "League",
      "name": "league",
      "value": "nfl"
    }
  ],
  "categories": [
    {
      "name": "passing",
      "displayName": "Passing",
      "count": 11
    },
    {
      "name": "rushing",
      "displayName": "Rushing",
      "count": 5
    }
  ],
  "events": {
    "401772988": {
      "id": "401772988",
      "week": 5,
      "score": "29-13",
      "gameResult": "L",
      "opponent": {
        "id": "26",
        "displayName": "Seattle Seahawks",
        "abbreviation": "SEA"
      }
    }
  },
  "seasonTypes": []
}
```
