# Athlete News

## https://site.api.espn.com/apis/site/v2/sports/hockey/{league}/athletes/{athlete}/news

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/athletes/4565230/news?limit=2`

## Example Response

```json
{
  "header": "{0} News",
  "articles": []
}
```
