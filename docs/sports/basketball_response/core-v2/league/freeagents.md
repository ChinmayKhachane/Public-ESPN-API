# Season Free Agents

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/seasons/{season}/freeagents

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `500`. The error payload is documented as observed.
- NBA returned an application error during live testing.

## Example Response

```json
{
  "error": {
    "code": 500,
    "message": "application error"
  }
}
```
