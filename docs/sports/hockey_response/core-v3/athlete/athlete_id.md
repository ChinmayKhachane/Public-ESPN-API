# Core V3 Athlete Detail

## https://sports.core.api.espn.com/v3/sports/hockey/{league}/athletes/{athlete}

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v3/sports/hockey/nhl/athletes/4565230`

## Example Response

```json
{
  "id": "4565230",
  "uid": "s:70~l:90~a:4565230",
  "guid": "bcca1a8b-c914-4ea9-2bde-abdf1d017ede",
  "firstName": "Trevor",
  "lastName": "Zegras",
  "fullName": "Trevor Zegras",
  "displayName": "Trevor Zegras",
  "shortName": "T. Zegras",
  "weight": 185.0,
  "displayWeight": "185 lbs",
  "height": 72.0,
  "displayHeight": "6' 0\"",
  "age": 25,
  "dateOfBirth": "2001-03-20T08:00Z"
}
```
