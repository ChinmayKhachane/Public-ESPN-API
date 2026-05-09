# Officials

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/events/{event}/competitions/{competition}/officials

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/officials?limit=2`

## Example Response

```json
{
  "count": 5,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 3,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/officials/2553179?lang=en&region=us",
      "id": "2553179",
      "firstName": "Francis",
      "lastName": "Charron",
      "fullName": "Francis Charron",
      "displayName": "Francis Charron",
      "position": {
        "name": "Referee",
        "displayName": "Referee",
        "id": "24"
      },
      "order": 0
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/officials/2553181?lang=en&region=us",
      "id": "2553181",
      "firstName": "Kyle",
      "lastName": "Rehman",
      "fullName": "Kyle Rehman",
      "displayName": "Kyle Rehman",
      "position": {
        "name": "Referee",
        "displayName": "Referee",
        "id": "24"
      },
      "order": 0
    }
  ]
}
```
