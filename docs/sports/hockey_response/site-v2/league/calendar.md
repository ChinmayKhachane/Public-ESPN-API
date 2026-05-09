# Site Calendar

## https://site.api.espn.com/apis/site/v2/sports/hockey/{league}/calendar

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `404`.
- Endpoint returned an error payload; the response is documented as observed.

Tested URL:
`https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/calendar`

## Example Response

```json
{
  "code": 404
}
```
