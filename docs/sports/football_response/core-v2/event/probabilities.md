# Competition Probabilities

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/probabilities

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988` on 2026-05-08.
- This is a play-indexed win probability timeline.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `page` | `int` | Page number |
| `limit` | `int` | Page size |

## Example Response

```json
{
  "count": 206,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 103,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/probabilities/4017729881?lang=en&region=us",
      "competition": {},
      "play": {},
      "homeTeam": {},
      "awayTeam": {},
      "tiePercentage": 0.0,
      "homeWinPercentage": 0.4057,
      "awayWinPercentage": 0.5943,
      "lastModified": "2026-02-09T03:29Z",
      "sequenceNumber": "100",
      "source": {
        "id": "2",
        "description": "feed",
        "state": "full"
      },
      "secondsLeft": 0,
      "spreadCoverProbHome": 0.5,
      "spreadPushProb": 0.0,
      "totalOverProb": 0.5,
      "totalPushProb": 0.0
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/probabilities/40177298840?lang=en&region=us",
      "competition": {},
      "play": {},
      "homeTeam": {},
      "awayTeam": {},
      "tiePercentage": 0.0,
      "homeWinPercentage": 0.3766,
      "awayWinPercentage": 0.6234,
      "lastModified": "2026-02-09T02:41Z",
      "sequenceNumber": "4000",
      "source": {
        "id": "2",
        "description": "feed",
        "state": "full"
      },
      "secondsLeft": 0,
      "spreadCoverProbHome": 0.5071,
      "spreadPushProb": 0.0,
      "totalOverProb": 0.5049,
      "totalPushProb": 0.0
    }
  ]
}
```
