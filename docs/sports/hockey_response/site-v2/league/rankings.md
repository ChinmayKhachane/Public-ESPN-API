# Rankings

## https://site.api.espn.com/apis/site/v2/sports/hockey/{league}/rankings

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `404`.
- Content type: `application/json;charset=UTF-8`.
- Endpoint returned an error payload; the response is documented as observed.

Tested URL:
`https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/rankings`

## Example Response

```json
{
  "code": 404,
  "message": "Unable to get sport data for sport: hockey and league: nhl"
}
```
