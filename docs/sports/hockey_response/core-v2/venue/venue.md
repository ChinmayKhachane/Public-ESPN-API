# Venues

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/venues

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/venues?limit=2`

## Example Response

```json
{
  "count": 180,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 90,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/venues/1?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/venues/4?lang=en&region=us"
    }
  ]
}
```
