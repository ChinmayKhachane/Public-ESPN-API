# Competition Competitor Score

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}/score

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988`, `competitor=17` on 2026-05-08.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17/scores/1?lang=en&region=us",
  "value": 13.0,
  "displayValue": "13",
  "winner": false,
  "source": {
    "id": "1",
    "description": "Basic/Manual"
  }
}
```
