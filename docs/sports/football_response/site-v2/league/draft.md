# Draft

## https://site.api.espn.com/apis/site/v2/sports/football/{league}/draft

Notes:
- Verified with `league=nfl` on 2026-05-08.
- This is a user-friendly site-facing draft board, different from the core-v2 draft object.

## Example Response

```json
{
  "uid": "s:20~l:28~e:draft~y:2026",
  "year": 2026,
  "displayName": "2026 National Football League Draft",
  "shortDisplayName": "2026 NFL Draft",
  "rounds": 7,
  "picks": [
    {
      "status": "SELECTION_MADE",
      "pick": 1,
      "overall": 1,
      "round": 1,
      "athlete": {
        "id": "110770",
        "displayName": "Fernando Mendoza"
      },
      "teamId": "13"
    }
  ]
}
```
