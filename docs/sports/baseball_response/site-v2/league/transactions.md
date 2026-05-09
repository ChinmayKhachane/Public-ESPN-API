# MLB Transactions

## https://site.api.espn.com/apis/site/v2/sports/baseball/mlb/transactions

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "count": 1126,
  "pageIndex": 1,
  "pageSize": 25,
  "pageCount": 46,
  "status": "success",
  "season": {
    "name": "Regular Season",
    "displayName": "2026",
    "year": 2026,
    "type": 2
  },
  "timestamp": "2026-05-09T03:57:33Z",
  "requestedYear": {
    "displayName": "2026",
    "year": 2026
  },
  "transactions": [
    {
      "team": {
        "id": "6",
        "name": "Tigers",
        "displayName": "Detroit Tigers",
        "abbreviation": "DET",
        "location": "Detroit",
        "color": "0a2240",
        "alternateColor": "ff4713",
        "logos": [
          {
            "href": "https://a.espncdn.com/i/teamlogos/mlb/500/det.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2018-06-05T12:07Z"
          },
          {
            "href": "https://a.espncdn.com/i/teamlogos/mlb/500-dark/det.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2023-08-14T14:41Z"
          }
        ],
        "links": [
          {
            "language": "en-US",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/_/name/det/detroit-tigers",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ]
      },
      "date": "2026-05-08T07:00Z"
    },
    {
      "team": {
        "id": "19",
        "name": "Dodgers",
        "displayName": "Los Angeles Dodgers",
        "abbreviation": "LAD",
        "location": "Los Angeles",
        "color": "005a9c",
        "alternateColor": "ffffff",
        "logos": [
          {
            "href": "https://a.espncdn.com/i/teamlogos/mlb/500/lad.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-01T17:48Z"
          },
          {
            "href": "https://a.espncdn.com/i/teamlogos/mlb/500-dark/lad.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-01T17:49Z"
          }
        ],
        "links": [
          {
            "language": "en-US",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/_/name/lad/los-angeles-dodgers",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ]
      },
      "date": "2026-05-08T07:00Z"
    }
  ]
}
```
