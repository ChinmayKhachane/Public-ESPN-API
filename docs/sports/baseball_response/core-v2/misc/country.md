# MLB Countries

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/countries?limit=5

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "count": 122,
  "pageIndex": 1,
  "pageSize": 5,
  "pageCount": 25,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/countries/1?lang=en&region=us",
      "id": "1",
      "abbreviation": "United States ",
      "athletes": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/countries/1/athletes?lang=en&region=us"
      }
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/countries/2?lang=en&region=us",
      "id": "2",
      "abbreviation": "Canada ",
      "athletes": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/countries/2/athletes?lang=en&region=us"
      }
    }
  ]
}
```
