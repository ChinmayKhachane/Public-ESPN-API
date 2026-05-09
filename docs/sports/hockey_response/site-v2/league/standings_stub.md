# Site V2 Standings Stub

## https://site.api.espn.com/apis/site/v2/sports/hockey/{league}/standings

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.
- The site/v2 standings path is kept because hockey.md notes it as a stub. Use `/apis/v2/` for standings data.

Tested URL:
`https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/standings`

## Example Response

```json
{
  "fullViewLink": {
    "text": "Full Standings",
    "href": "https://www.espn.com/nhl/standings"
  }
}
```
