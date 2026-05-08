# Competition Probability ID

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/probabilities/{play}

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988`, `play=40177298840` on 2026-05-08.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/probabilities/40177298840?lang=en&region=us",
  "play": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/plays/40177298840?lang=en&region=us"
  },
  "homeWinPercentage": 0.3766,
  "awayWinPercentage": 0.6234,
  "tiePercentage": 0.0,
  "spreadCoverProbHome": 0.5071,
  "totalOverProb": 0.5049,
  "lastModified": "2026-02-09T02:41Z"
}
```
