# Play Personnel

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/events/{event}/competitions/{competition}/plays/{play}/personnel

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/plays/4018711617/personnel?lang=en&region=us",
  "playPersonnel": [
    {
      "entries": [],
      "competitor": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20?lang=en&region=us"
      }
    },
    {
      "entries": [],
      "competitor": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/18?lang=en&region=us"
      }
    }
  ]
}
```
