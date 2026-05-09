# Athletes

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/athletes

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.
- Collection items are `$ref` links; the athlete detail doc resolves a representative player.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/athletes?limit=2&active=true`

## Example Response

```json
{
  "count": 1186,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 593,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/athletes/2273?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/athletes/2300?lang=en&region=us"
    }
  ]
}
```
