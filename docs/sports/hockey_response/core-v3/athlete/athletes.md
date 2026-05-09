# Core V3 Athletes

## https://sports.core.api.espn.com/v3/sports/hockey/{league}/athletes

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v3/sports/hockey/nhl/athletes?limit=2`

## Example Response

```json
{
  "count": 2596,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 1298,
  "items": [
    {
      "id": "3042264",
      "uid": "s:70~l:90~a:3042264",
      "guid": "0c1a14b5-2154-5563-d33d-2b2958c30d60",
      "firstName": "Tyler",
      "lastName": "Lewington",
      "fullName": "Tyler Lewington",
      "displayName": "Tyler Lewington",
      "shortName": "T. Lewington",
      "weight": 200.0,
      "displayWeight": "200 lbs",
      "height": 74.0,
      "displayHeight": "6' 2\"",
      "age": 31,
      "dateOfBirth": "1994-12-05T08:00Z"
    },
    {
      "id": "3942430",
      "uid": "s:70~l:90~a:3942430",
      "guid": "163d22dc-a33b-3ba3-97f5-395bd886059d",
      "firstName": "Kodie",
      "lastName": "Curran",
      "fullName": "Kodie Curran",
      "displayName": "Kodie Curran",
      "shortName": "K. Curran",
      "weight": 200.0,
      "displayWeight": "200 lbs",
      "height": 74.0,
      "displayHeight": "6' 2\"",
      "age": 36,
      "dateOfBirth": "1989-12-18T08:00Z"
    }
  ]
}
```
