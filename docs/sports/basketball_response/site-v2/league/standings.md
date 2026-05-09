# Standings

## https://site.api.espn.com/apis/v2/sports/basketball/{league}/standings

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.
- Use `/apis/v2/`; the `/apis/site/v2/` standings path returned a stub payload in live testing.

## Example Response

```json
{
  "id": "7",
  "uid": "s:40~l:46~g:7",
  "name": "National Basketball Association",
  "abbreviation": "NBA",
  "season": {
    "displayName": "2025-26",
    "year": 2026,
    "startDate": "2025-10-01T07:00Z",
    "endDate": "2026-06-27T06:59Z"
  },
  "shortName": "NBA",
  "children": [
    {
      "id": "5",
      "uid": "s:40~l:46~g:5",
      "name": "Eastern Conference",
      "abbreviation": "East",
      "isConference": true,
      "standings": {
        "id": "0",
        "name": "overall",
        "displayName": "Standings",
        "season": 2026,
        "entries": [],
        "links": [],
        "seasonType": 2,
        "seasonDisplayName": "2025-26"
      }
    },
    {
      "id": "6",
      "uid": "s:40~l:46~g:6",
      "name": "Western Conference",
      "abbreviation": "West",
      "isConference": true,
      "standings": {
        "id": "0",
        "name": "overall",
        "displayName": "Standings",
        "season": 2026,
        "entries": [],
        "links": [],
        "seasonType": 2,
        "seasonDisplayName": "2025-26"
      }
    }
  ],
  "isConference": false,
  "links": [
    {
      "text": "Table",
      "shortText": "Standings",
      "language": "en-US",
      "rel": [
        "standings",
        "desktop"
      ],
      "href": "https://www.espn.com/nba/standings/_/group/7",
      "isExternal": false,
      "isPremium": false
    }
  ],
  "seasons": [
    {
      "displayName": "2025-26",
      "year": 2026,
      "startDate": "2025-10-01T07:00Z",
      "endDate": "2026-06-27T06:59Z",
      "types": [
        {},
        {}
      ],
      "seasonYears": "2025-26"
    },
    {
      "displayName": "2024-25",
      "year": 2025,
      "startDate": "2024-09-24T07:00Z",
      "endDate": "2025-06-27T06:59Z",
      "types": [
        {},
        {}
      ],
      "seasonYears": "2024-25"
    }
  ]
}
```
