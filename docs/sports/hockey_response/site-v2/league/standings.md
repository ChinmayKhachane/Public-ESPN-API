# Standings

## https://site.api.espn.com/apis/v2/sports/hockey/{league}/standings

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.api.espn.com/apis/v2/sports/hockey/nhl/standings`

## Example Response

```json
{
  "uid": "s:70~l:90~g:9",
  "id": "9",
  "name": "National Hockey League",
  "abbreviation": "NHL",
  "children": [
    {
      "uid": "s:70~l:90~g:7",
      "id": "7",
      "name": "Eastern Conference",
      "abbreviation": "East",
      "standings": {
        "id": "0",
        "name": "overall",
        "displayName": "Standings",
        "links": [
          {}
        ],
        "season": 2026,
        "seasonType": 2,
        "seasonDisplayName": "2025-26",
        "entries": [
          {},
          {}
        ]
      }
    },
    {
      "uid": "s:70~l:90~g:8",
      "id": "8",
      "name": "Western Conference",
      "abbreviation": "West",
      "standings": {
        "id": "0",
        "name": "overall",
        "displayName": "Standings",
        "links": [
          {}
        ],
        "season": 2026,
        "seasonType": 2,
        "seasonDisplayName": "2025-26",
        "entries": [
          {},
          {}
        ]
      }
    }
  ],
  "season": {
    "year": 2026,
    "startDate": "2025-09-20T07:00Z",
    "endDate": "2026-07-01T06:59Z",
    "displayName": "2025-26"
  },
  "links": [
    {
      "language": "en-US",
      "rel": [
        "standings",
        "desktop"
      ],
      "href": "https://www.espn.com/nhl/standings/_/group/9",
      "text": "Table",
      "shortText": "Standings",
      "isExternal": false,
      "isPremium": false
    }
  ],
  "seasons": [
    {
      "year": 2026,
      "startDate": "2025-09-20T07:00Z",
      "endDate": "2026-07-01T06:59Z",
      "displayName": "2025-26",
      "types": [
        {
          "id": "1",
          "name": "Preseason",
          "abbreviation": "pre",
          "startDate": "2025-09-20T07:00Z",
          "endDate": "2025-10-07T06:59Z",
          "hasStandings": true
        },
        {
          "id": "2",
          "name": "Regular Season",
          "abbreviation": "reg",
          "startDate": "2025-10-07T07:00Z",
          "endDate": "2026-04-18T06:59Z",
          "hasStandings": true
        }
      ],
      "seasonYears": "2025-26"
    },
    {
      "year": 2025,
      "startDate": "2024-09-21T07:00Z",
      "endDate": "2025-07-01T06:59Z",
      "displayName": "2024-25",
      "types": [
        {
          "id": "1",
          "name": "Preseason",
          "abbreviation": "pre",
          "startDate": "2024-09-21T07:00Z",
          "endDate": "2024-10-04T06:59Z",
          "hasStandings": true
        },
        {
          "id": "2",
          "name": "Regular Season",
          "abbreviation": "reg",
          "startDate": "2024-10-04T07:00Z",
          "endDate": "2025-04-18T06:59Z",
          "hasStandings": true
        }
      ],
      "seasonYears": "2024-25"
    }
  ]
}
```
