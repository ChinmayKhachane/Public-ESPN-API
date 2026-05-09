# Competitor Linescores

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}/linescores

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15/linescores?limit=2`

## Example Response

```json
{
  "count": 3,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 2,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15/linescores/2/1?lang=en&region=us",
      "value": 0.0,
      "displayValue": "0",
      "source": {
        "id": "2",
        "description": "feed",
        "state": "full"
      },
      "period": 1
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15/linescores/2/2?lang=en&region=us",
      "value": 1.0,
      "displayValue": "1",
      "source": {
        "id": "2",
        "description": "feed",
        "state": "full"
      },
      "period": 2
    }
  ]
}
```
