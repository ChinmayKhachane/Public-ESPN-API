# Core V3 Season

## https://sports.core.api.espn.com/v3/sports/hockey/{league}/seasons/{season}

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v3/sports/hockey/nhl/seasons/2026`

## Example Response

```json
{
  "year": 2026,
  "startDate": "2025-09-20T07:00Z",
  "endDate": "2026-07-01T06:59Z"
}
```
