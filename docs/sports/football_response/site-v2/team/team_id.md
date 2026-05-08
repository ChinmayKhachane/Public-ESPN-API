# Team ID

## https://site.api.espn.com/apis/site/v2/sports/football/{league}/teams/{id}

Notes:
- Verified with `league=nfl`, `id=6` on 2026-05-08.
- The response is wrapped under a top-level `team` object.

## Example Response

```json
{
  "team": {
    "id": "6",
    "uid": "s:20~l:28~t:6",
    "slug": "dallas-cowboys",
    "location": "Dallas",
    "name": "Cowboys",
    "abbreviation": "DAL",
    "displayName": "Dallas Cowboys",
    "color": "002a5c",
    "groups": {
      "id": "1",
      "parent": {
        "id": "7"
      },
      "isConference": false
    },
    "standingSummary": "2nd in NFC East",
    "franchise": {
      "id": "6",
      "displayName": "Dallas Cowboys"
    },
    "nextEvent": []
  }
}
```
