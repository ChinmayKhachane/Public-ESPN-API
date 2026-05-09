# Site Rankings

## https://site.api.espn.com/apis/site/v2/sports/basketball/{league}/rankings

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `404`. The error payload is documented as observed.
- NBA currently returns `404` for this route.

## Example Response

```json
{
  "code": 404,
  "message": "Unable to get sport data for sport: basketball and league: nba"
}
```
