# Team Record

## https://site.api.espn.com/apis/site/v2/sports/hockey/{league}/teams/{team}/record

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.
- Several site-v2 team child paths can be sparse for NHL; this records the observed team record payload.
- Live request returned an empty `{}` payload.

Tested URL:
`https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/teams/15/record`

## Example Response

```json
{}
```
