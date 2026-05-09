# Competition Situation

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/events/{event}/competitions/{competition}/situation

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/situation?lang=en&region=us",
  "lastPlay": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/plays/401871161677?lang=en&region=us"
  },
  "homeTimeouts": {
    "timeoutsCurrent": 2,
    "timeoutsRemainingCurrent": 0
  },
  "awayTimeouts": {
    "timeoutsCurrent": 1,
    "timeoutsRemainingCurrent": 0
  },
  "homeFouls": {
    "teamFouls": 25,
    "teamFoulsCurrent": 5,
    "foulsToGive": 0,
    "bonusState": "DOUBLE"
  },
  "awayFouls": {
    "teamFouls": 21,
    "teamFoulsCurrent": 5,
    "foulsToGive": 0,
    "bonusState": "DOUBLE"
  }
}
```
