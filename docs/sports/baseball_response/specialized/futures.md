# MLB Futures

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/futures?limit=5

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.
- Root README lists season futures. MLB returned a populated futures collection.

## Example Response

```json
{
  "count": 3,
  "pageIndex": 1,
  "pageSize": 5,
  "pageCount": 1,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/futures/2761?lang=en&region=us",
      "id": 2761,
      "name": "MLB - World Series - Winner",
      "displayName": "World Series Winner",
      "futures": [
        {
          "provider": {
            "id": "58",
            "name": "ESPN BET",
            "active": 1,
            "priority": 0
          },
          "books": [
            {},
            {}
          ]
        }
      ],
      "type": "winLeague"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/futures/3294?lang=en&region=us",
      "id": 3294,
      "name": "MLB - National League - Winner",
      "displayName": "National League Winner",
      "futures": [
        {
          "provider": {
            "id": "58",
            "name": "ESPN BET",
            "active": 1,
            "priority": 0
          },
          "books": [
            {},
            {}
          ]
        }
      ]
    }
  ]
}
```
