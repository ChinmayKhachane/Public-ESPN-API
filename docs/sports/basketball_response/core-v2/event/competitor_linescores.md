# Competition Competitor Linescores

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}/linescores

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "count": 4,
  "pageIndex": 1,
  "pageSize": 25,
  "pageCount": 1,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20/linescores/1/1?lang=en&region=us",
      "value": 31.0,
      "displayValue": "31",
      "period": 1,
      "source": {
        "id": "1",
        "description": "Basic/Manual"
      }
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20/linescores/1/2?lang=en&region=us",
      "value": 21.0,
      "displayValue": "21",
      "period": 2,
      "source": {
        "id": "1",
        "description": "Basic/Manual"
      }
    }
  ]
}
```
