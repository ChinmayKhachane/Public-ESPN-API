# Competition Officials

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/events/{event}/competitions/{competition}/officials

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "count": 4,
  "pageIndex": 1,
  "pageSize": 25,
  "pageCount": 1,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/officials/1?lang=en&region=us",
      "id": "6876",
      "displayName": "Josh Tiven",
      "firstName": "Josh",
      "lastName": "Tiven",
      "fullName": "Josh Tiven",
      "position": {
        "id": "40",
        "name": "Referee",
        "displayName": "Referee"
      },
      "order": 1
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/officials/2?lang=en&region=us",
      "id": "2615764",
      "displayName": "Marc Davis",
      "firstName": "Marc",
      "lastName": "Davis",
      "fullName": "Marc Davis",
      "position": {
        "id": "40",
        "name": "Referee",
        "displayName": "Referee"
      },
      "order": 2
    }
  ]
}
```
