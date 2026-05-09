# Standings Stub

## https://site.api.espn.com/apis/site/v2/sports/basketball/{league}/standings

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.
- Stub route; use the canonical standings endpoint instead.

## Example Response

```json
{
  "fullViewLink": {
    "text": "Full Standings",
    "href": "https://www.espn.com/nba/standings"
  }
}
```
