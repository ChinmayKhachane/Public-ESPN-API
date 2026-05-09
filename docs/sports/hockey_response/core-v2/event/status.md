# Status

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/events/{event}/competitions/{competition}/status

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/status`

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/status?lang=en&region=us",
  "clock": 0.0,
  "displayClock": "0:00",
  "period": 3,
  "type": {
    "id": "3",
    "name": "STATUS_FINAL",
    "state": "post",
    "completed": true,
    "description": "Final",
    "detail": "Final",
    "shortDetail": "Final"
  },
  "featuredAthletes": [
    {
      "name": "winningGoalie",
      "displayName": "Winning Goalie",
      "shortDisplayName": "Winning Goalie",
      "abbreviation": "W",
      "playerId": 2517899,
      "athlete": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/athletes/2517899?lang=en&region=us"
      },
      "team": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/teams/7?lang=en&region=us"
      },
      "statistics": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/7/roster/2517899/statistics/0?lang=en&region=us"
      },
      "projections": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/7/roster/2517899/projections?lang=en&region=us"
      }
    },
    {
      "name": "losingGoalie",
      "displayName": "Losing Goalie",
      "shortDisplayName": "Losing Goalie",
      "abbreviation": "L",
      "playerId": 3942459,
      "athlete": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/athletes/3942459?lang=en&region=us"
      },
      "team": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/teams/15?lang=en&region=us"
      },
      "statistics": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15/roster/3942459/statistics/0?lang=en&region=us"
      },
      "projections": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15/roster/3942459/projections?lang=en&region=us"
      }
    }
  ]
}
```
