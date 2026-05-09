# Circuits

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/circuits

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.
- Live request returned an empty collection.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/circuits?limit=2`

## Example Response

```json
{
  "count": 0,
  "pageIndex": 0,
  "pageSize": 25,
  "pageCount": 0,
  "items": []
}
```
