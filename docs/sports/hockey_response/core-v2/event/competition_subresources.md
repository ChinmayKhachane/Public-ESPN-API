# Competition Broadcasts

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/events/{event}/competitions/{competition}/broadcasts

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.
- One representative competition child resource; other competition child paths are documented separately.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/broadcasts?limit=2`

## Example Response

```json
{
  "count": 3,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 2,
  "items": [
    {
      "type": {
        "id": "1",
        "shortName": "TV",
        "longName": "Television",
        "slug": "tv"
      },
      "channel": 498,
      "station": "TNT",
      "slug": "tnt",
      "priority": 1,
      "market": {
        "id": "1",
        "type": "National"
      },
      "media": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/media/498?lang=en&region=us",
        "id": "498",
        "callLetters": "TNT",
        "name": "TNT",
        "shortName": "TNT",
        "slug": "tnt",
        "logos": [
          {},
          {}
        ]
      },
      "lang": "en",
      "region": "us",
      "competition": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412?lang=en&region=us"
      },
      "partnered": false
    },
    {
      "type": {
        "id": "1",
        "shortName": "TV",
        "longName": "Television",
        "slug": "tv"
      },
      "channel": 495,
      "station": "truTV",
      "slug": "trutv",
      "priority": 2,
      "market": {
        "id": "1",
        "type": "National"
      },
      "media": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/media/495?lang=en&region=us",
        "id": "495",
        "callLetters": "truTV",
        "name": "truTV",
        "shortName": "truTV",
        "slug": "trutv",
        "logos": [
          {},
          {}
        ]
      },
      "lang": "en",
      "region": "us",
      "competition": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412?lang=en&region=us"
      },
      "partnered": false
    }
  ]
}
```
