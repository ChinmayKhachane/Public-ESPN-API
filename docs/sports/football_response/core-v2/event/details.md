# Competition Details

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/details

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988` on 2026-05-08.
- Even though the competition payload exposes a `details` ref, the direct NFL endpoint currently returns HTTP 404.
- Use `/plays` instead.

## Example Response

```json
{
  "error": {
    "code": 404
  }
}
```
