# MLB League Statistics

## https://site.api.espn.com/apis/site/v2/sports/baseball/mlb/statistics

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "status": "success",
  "season": {
    "name": "Regular Season",
    "displayName": "2026",
    "year": 2026,
    "type": 2
  },
  "timestamp": "2026-05-09T03:45:30Z",
  "league": {
    "id": "10",
    "name": "Major League Baseball",
    "abbreviation": "MLB",
    "slug": "mlb",
    "shortName": "MLB",
    "midsizeName": "MLB",
    "isTournament": false,
    "links": [
      {
        "language": "en-US",
        "rel": [
          "index",
          "desktop"
        ],
        "href": "https://www.espn.com/mlb/",
        "text": "Index",
        "shortText": "Index",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "index",
          "sportscenter"
        ],
        "href": "sportscenter://x-callback-url/showClubhouse?uid=s:1~l:10",
        "text": "Index",
        "shortText": "Index",
        "isExternal": false,
        "isPremium": false
      }
    ],
    "logos": [
      {
        "href": "https://a.espncdn.com/i/teamlogos/leagues/500/mlb.png",
        "width": 500,
        "height": 500,
        "alt": "",
        "rel": [
          "full",
          "default"
        ],
        "lastUpdated": "2023-03-29T12:34Z"
      },
      {
        "href": "https://a.espncdn.com/combiner/i?img=/i/teamlogos/leagues/500-dark/mlb.png&w=500&h=500&transparent=true",
        "width": 500,
        "height": 500,
        "alt": "",
        "rel": [
          "full",
          "dark"
        ],
        "lastUpdated": "2026-05-08T20:30Z"
      }
    ]
  },
  "stats": {
    "id": "0",
    "name": "Total",
    "abbreviation": "Total",
    "categories": [
      {
        "name": "avg",
        "displayName": "Batting Average",
        "abbreviation": "AVG",
        "shortDisplayName": "BA",
        "leaders": [
          {
            "team": {},
            "statistics": {},
            "displayValue": "2-2, HR, 3 RBI, R",
            "value": 1.0,
            "athlete": {}
          },
          {
            "team": {},
            "statistics": {},
            "displayValue": "1-1, R",
            "value": 1.0,
            "athlete": {}
          }
        ]
      },
      {
        "name": "homeRuns",
        "displayName": "Home Runs",
        "abbreviation": "HR",
        "shortDisplayName": "HR",
        "leaders": [
          {
            "team": {},
            "statistics": {},
            "displayValue": "18-39, 7 HR, 6 2B, 9 RBI, 14 R, 8 BB, 3 SB, 7 K",
            "value": 7.0,
            "athlete": {}
          },
          {
            "team": {},
            "statistics": {},
            "displayValue": "16-46, 7 HR, 4 2B, 13 RBI, 15 R, 7 BB, 10 K",
            "value": 7.0,
            "athlete": {}
          }
        ]
      }
    ]
  }
}
```
