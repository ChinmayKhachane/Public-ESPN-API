# MLB Competitor Linescores

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/competitors/17/linescores?limit=5

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "count": 9,
  "pageIndex": 1,
  "pageSize": 5,
  "pageCount": 2,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp...",
      "value": 0.0,
      "displayValue": "0",
      "source": {
        "id": "2",
        "state": "full"
      },
      "period": 1,
      "hits": 1,
      "errors": 0
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp...",
      "value": 0.0,
      "displayValue": "0",
      "source": {
        "id": "2",
        "state": "full"
      },
      "period": 2,
      "hits": 0,
      "errors": 0
    }
  ]
}
```
