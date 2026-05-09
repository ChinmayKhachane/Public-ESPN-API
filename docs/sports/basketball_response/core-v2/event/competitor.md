# Competition Competitor

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20?lang=en&region=us",
  "id": "20",
  "uid": "s:40~l:46~t:20",
  "type": "team",
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/teams/20?lang=en&region=us"
  },
  "statistics": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20/statistics?lang=en&region=us"
  },
  "leaders": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20/leaders?lang=en&region=us"
  },
  "score": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20/score?lang=en&region=us"
  },
  "linescores": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20/linescores?lang=en&region=us"
  },
  "roster": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20/roster?lang=en&region=us"
  },
  "record": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/types/2/teams/20/record?lang=en&region=us"
  },
  "winner": false
}
```
