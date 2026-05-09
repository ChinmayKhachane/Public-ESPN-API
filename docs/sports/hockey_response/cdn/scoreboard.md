# CDN Scoreboard

## https://cdn.espn.com/core/nhl/scoreboard?xhr=1

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `404`.
- Content type: `application/json;charset=UTF-8`.
- Endpoint returned an error payload; the response is documented as observed.

Tested URL:
`https://cdn.espn.com/core/nhl/scoreboard?xhr=1`

## Example Response

```json
{
  "timestamp": 1778346193729,
  "status": 404,
  "error": "Not Found",
  "message": "No message available",
  "path": "/core/nhl/scoreboard"
}
```
