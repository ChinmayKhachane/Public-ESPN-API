# MLB Standings

## https://site.api.espn.com/apis/v2/sports/baseball/mlb/standings

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "id": "9",
  "uid": "s:1~l:10~g:9",
  "name": "Major League Baseball",
  "abbreviation": "MLB",
  "season": {
    "displayName": "2026",
    "year": 2026,
    "startDate": "2026-02-19T08:00Z",
    "endDate": "2026-11-12T07:59Z"
  },
  "shortName": "MLB",
  "children": [
    {
      "id": "7",
      "uid": "s:1~l:10~g:7",
      "name": "American League",
      "abbreviation": "AL",
      "shortName": "AL",
      "standings": {
        "id": "0",
        "name": "overall",
        "displayName": "Standings",
        "season": 2026,
        "links": [
          {
            "language": "en-US",
            "rel": [],
            "href": "https://www.espn.com/mlb/standings/_/group/7",
            "text": "Table",
            "shortText": "Standings",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "seasonType": 2,
        "seasonDisplayName": "2026",
        "entries": [
          {
            "team": {},
            "stats": []
          },
          {
            "team": {},
            "stats": []
          }
        ]
      }
    },
    {
      "id": "8",
      "uid": "s:1~l:10~g:8",
      "name": "National League",
      "abbreviation": "NL",
      "shortName": "NL",
      "standings": {
        "id": "0",
        "name": "overall",
        "displayName": "Standings",
        "season": 2026,
        "links": [
          {
            "language": "en-US",
            "rel": [],
            "href": "https://www.espn.com/mlb/standings/_/group/8",
            "text": "Table",
            "shortText": "Standings",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "seasonType": 2,
        "seasonDisplayName": "2026",
        "entries": [
          {
            "team": {},
            "stats": []
          },
          {
            "team": {},
            "stats": []
          }
        ]
      }
    }
  ],
  "links": [
    {
      "language": "en-US",
      "rel": [
        "standings",
        "desktop"
      ],
      "href": "https://www.espn.com/mlb/standings/_/group/9",
      "text": "Table",
      "shortText": "Standings",
      "isExternal": false,
      "isPremium": false
    }
  ],
  "seasons": [
    {
      "displayName": "2026",
      "year": 2026,
      "startDate": "2026-02-19T08:00Z",
      "endDate": "2026-11-12T07:59Z",
      "types": [
        {
          "id": "1",
          "name": "Spring Training",
          "abbreviation": "pre",
          "startDate": "2026-02-19T08:00Z",
          "endDate": "2026-03-25T06:59Z",
          "hasStandings": true
        },
        {
          "id": "2",
          "name": "Regular Season",
          "abbreviation": "reg",
          "startDate": "2026-03-25T07:00Z",
          "endDate": "2026-09-30T06:59Z",
          "hasStandings": true
        }
      ]
    },
    {
      "displayName": "2025",
      "year": 2025,
      "startDate": "2025-02-15T08:00Z",
      "endDate": "2025-12-11T07:59Z",
      "types": [
        {
          "id": "1",
          "name": "Spring Training",
          "abbreviation": "pre",
          "startDate": "2025-02-15T08:00Z",
          "endDate": "2025-03-26T06:59Z",
          "hasStandings": true
        },
        {
          "id": "2",
          "name": "Regular Season",
          "abbreviation": "reg",
          "startDate": "2025-03-26T07:00Z",
          "endDate": "2025-09-30T06:59Z",
          "hasStandings": true
        }
      ]
    }
  ]
}
```
