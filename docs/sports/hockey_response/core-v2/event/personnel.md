# Play Personnel

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/events/{event}/competitions/{competition}/plays/{play}/personnel

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/plays/401871412000000510/personnel`

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/plays/401871412000000510/personnel?lang=en&region=us",
  "playPersonnel": [
    {
      "competitor": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15?lang=en&region=us"
      },
      "entries": [
        {
          "athlete": {},
          "position": {},
          "statistics": {},
          "whereabouts": {},
          "projections": {}
        },
        {
          "athlete": {},
          "position": {},
          "statistics": {},
          "whereabouts": {},
          "projections": {}
        }
      ]
    },
    {
      "competitor": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/7?lang=en&region=us"
      },
      "entries": [
        {
          "athlete": {},
          "position": {},
          "statistics": {},
          "whereabouts": {},
          "projections": {}
        },
        {
          "athlete": {},
          "position": {},
          "statistics": {},
          "whereabouts": {},
          "projections": {}
        }
      ]
    }
  ]
}
```
