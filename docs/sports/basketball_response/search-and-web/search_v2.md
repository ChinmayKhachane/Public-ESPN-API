# Search V2

## https://site.web.api.espn.com/apis/search/v2?query={query}

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "totalFound": 4,
  "resultTypes": [
    {
      "displayName": "Players",
      "type": "player",
      "totalFound": 1
    },
    {
      "displayName": "Articles",
      "type": "article",
      "totalFound": 29641
    }
  ],
  "results": [
    {
      "displayName": "Players",
      "type": "player",
      "totalFound": 1,
      "page": 1,
      "limit": 2,
      "contents": [
        {}
      ]
    },
    {
      "displayName": "Articles",
      "type": "article",
      "totalFound": 2,
      "page": 1,
      "limit": 2,
      "contents": [
        {},
        {}
      ]
    }
  ]
}
```
