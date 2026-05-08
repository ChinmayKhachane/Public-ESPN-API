# Competition Situation

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/situation

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988` on 2026-05-08.
- For a completed game, possession is omitted, but the last play ref is still useful.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/situation?lang=en&region=us",
  "down": 0,
  "distance": 0,
  "yardLine": 71,
  "homeTimeouts": 0,
  "awayTimeouts": 0,
  "isRedZone": false,
  "lastPlay": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/plays/4017729884651?lang=en&region=us"
  }
}
```
