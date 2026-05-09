# Core V3 League

## https://sports.core.api.espn.com/v3/sports/hockey/{league}

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v3/sports/hockey/nhl`

## Example Response

```json
{
  "id": "90",
  "uid": "s:70~l:90",
  "guid": "1a5f0227-a13e-396c-8cea-8961bc288666",
  "groupId": "9",
  "name": "National Hockey League",
  "displayName": "NHL",
  "abbreviation": "NHL",
  "shortName": "NHL",
  "midsizeName": "NHL",
  "color": "000000",
  "slug": "nhl"
}
```
