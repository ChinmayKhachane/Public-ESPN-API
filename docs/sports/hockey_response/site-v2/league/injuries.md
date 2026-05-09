# League Injuries

## https://site.api.espn.com/apis/site/v2/sports/hockey/{league}/injuries

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/injuries`

## Example Response

```json
{
  "timestamp": "2026-05-09T17:03:06Z",
  "status": "success",
  "season": {
    "year": 2026,
    "type": 3,
    "name": "Postseason",
    "displayName": "2025-26"
  },
  "injuries": [
    {
      "id": "25",
      "displayName": "Anaheim Ducks",
      "injuries": [
        {
          "id": "591670",
          "longComment": "day-to-day",
          "shortComment": "day-to-day",
          "status": "Day-To-Day",
          "date": "2026-05-09T13:19Z",
          "athlete": {},
          "source": {},
          "type": {},
          "details": {}
        }
      ]
    },
    {
      "id": "1",
      "displayName": "Boston Bruins",
      "injuries": [
        {
          "id": "591478",
          "longComment": "out",
          "shortComment": "out",
          "status": "Out",
          "date": "2026-05-03T17:21Z",
          "athlete": {},
          "source": {},
          "type": {},
          "details": {}
        },
        {
          "id": "591442",
          "longComment": "out",
          "shortComment": "out",
          "status": "Out",
          "date": "2026-05-02T11:45Z",
          "athlete": {},
          "source": {},
          "type": {},
          "details": {}
        }
      ]
    }
  ]
}
```
