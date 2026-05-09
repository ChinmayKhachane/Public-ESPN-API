# Competition Competitor Linescores

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}/linescores

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988`, `competitor=17` on 2026-05-08.

## Example Response

```json
{
  "count": 4,
  "pageIndex": 1,
  "pageSize": 25,
  "pageCount": 1,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17/linescores/1/1?lang=en&region=us",
      "value": 0.0,
      "displayValue": "0",
      "source": {
        "id": "1",
        "description": "Basic/Manual"
      },
      "period": 1
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17/linescores/1/4?lang=en&region=us",
      "value": 13.0,
      "displayValue": "13",
      "source": {
        "id": "1",
        "description": "Basic/Manual"
      },
      "period": 4
    }
  ]
}
```
