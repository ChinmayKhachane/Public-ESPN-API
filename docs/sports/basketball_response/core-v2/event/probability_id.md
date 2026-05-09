# Competition Probability ID

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/events/{event}/competitions/{competition}/probabilities/{play}

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/probabilities/4018711617?lang=en&region=us",
  "competition": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161?lang=en&region=us"
  },
  "play": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/plays/4018711617?lang=en&region=us"
  },
  "homeTeam": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/teams/20?lang=en&region=us"
  },
  "awayTeam": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/teams/18?lang=en&region=us"
  },
  "tiePercentage": 0.0,
  "homeWinPercentage": 0.388,
  "awayWinPercentage": 0.612,
  "lastModified": "2026-05-09T02:14Z",
  "sequenceNumber": "7",
  "source": {
    "id": "2",
    "description": "feed",
    "state": "full"
  },
  "spreadCoverProbHome": 0.5368
}
```
