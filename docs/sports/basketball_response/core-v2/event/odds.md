# Competition Odds

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/events/{event}/competitions/{competition}/odds

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "count": 1,
  "pageIndex": 1,
  "pageSize": 25,
  "pageCount": 1,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/odds/100?lang=en&region=us",
      "provider": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/providers/100?lang=en&region=us",
        "id": "100",
        "name": "Draft Kings",
        "priority": 1
      },
      "details": "PHI -3.5",
      "overUnder": 214.5,
      "spread": -3.5,
      "overOdds": -110.0,
      "underOdds": -110.0,
      "awayTeamOdds": {
        "team": {},
        "favorite": false,
        "underdog": true,
        "moneyLine": 130,
        "spreadOdds": -115.0,
        "open": {},
        "close": {},
        "current": {}
      },
      "homeTeamOdds": {
        "team": {},
        "favorite": true,
        "underdog": false,
        "moneyLine": -155,
        "spreadOdds": -105.0,
        "open": {},
        "close": {},
        "current": {}
      },
      "links": [
        {},
        {}
      ],
      "moneylineWinner": false,
      "spreadWinner": false
    }
  ]
}
```
