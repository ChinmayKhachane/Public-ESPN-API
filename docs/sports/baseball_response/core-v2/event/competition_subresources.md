# MLB Competition Broadcasts

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/broadcasts?limit=5

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "count": 3,
  "pageIndex": 1,
  "pageSize": 5,
  "pageCount": 1,
  "items": [
    {
      "slug": "mlbtv",
      "type": {
        "id": "4",
        "slug": "streaming",
        "shortName": "Streaming",
        "longName": "Streaming"
      },
      "channel": 885,
      "station": "MLB.TV",
      "priority": 1,
      "market": {
        "id": "1",
        "type": "National"
      },
      "media": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/media/885?lang=en&region=us",
        "id": "885",
        "name": "MLB.TV",
        "slug": "mlbtv",
        "threeLetterAbbreviation": "MLB",
        "callLetters": "MLB.TV",
        "shortName": "MLB.TV",
        "logos": [
          {
            "href": "https://a.espncdn.com/guid/0db644c3-9f87-37e7-9884-858c2ed45218/logos/default.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2025-03-05T17:28Z"
          }
        ]
      },
      "lang": "en",
      "region": "us",
      "competition": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256?lang..."
      }
    },
    {
      "slug": "redstv",
      "type": {
        "id": "4",
        "slug": "streaming",
        "shortName": "Streaming",
        "longName": "Streaming"
      },
      "channel": 1389,
      "station": "Reds.TV",
      "priority": 2,
      "market": {
        "id": "2",
        "type": "Home"
      },
      "media": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/media/1389?lang=en&region=us",
        "id": "1389",
        "name": "Reds.TV",
        "slug": "redstv",
        "threeLetterAbbreviation": "Red",
        "callLetters": "Reds.TV",
        "shortName": "Reds.TV",
        "logos": []
      },
      "lang": "en",
      "region": "us",
      "competition": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256?lang..."
      }
    }
  ]
}
```
