# Probabilities

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/events/{event}/competitions/{competition}/probabilities

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `400`.
- Content type: `application/json;charset=utf-8`.
- Endpoint returned an error payload; the response is documented as observed.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/probabilities?limit=2`

## Example Response

```json
{
  "error": {
    "message": "Probabilities are not supported for sport: hockey, league: nhl, competition: 401871412",
    "code": 400
  }
}
```
