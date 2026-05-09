# Current Season

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/season

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026?lang=en&region=us",
  "displayName": "2025-26",
  "year": 2026,
  "type": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/types/3?lang=en&region=us",
    "id": "3",
    "name": "Postseason",
    "abbreviation": "post",
    "slug": "post-season",
    "year": 2026,
    "type": 3,
    "leaders": {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/types/3/leaders?lang=en&region=us"
    },
    "groups": {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/types/3/groups?lang=en&region=us"
    },
    "startDate": "2026-04-18T07:00Z",
    "endDate": "2026-06-27T06:59Z",
    "hasGroups": false
  },
  "athletes": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/athletes?lang=en&region=us"
  },
  "leaders": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/types/3/leaders?lang=en&region=us"
  },
  "startDate": "2025-10-01T07:00Z",
  "endDate": "2026-06-27T06:59Z",
  "types": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/types?lang=en&region=us",
    "count": 5,
    "pageIndex": 1,
    "pageSize": 5,
    "pageCount": 1,
    "items": [
      {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/types/1?lang=en&region=us",
        "id": "1",
        "name": "Preseason",
        "abbreviation": "pre",
        "slug": "preseason",
        "year": 2026,
        "type": 1,
        "groups": {},
        "startDate": "2025-10-01T07:00Z",
        "endDate": "2025-10-21T06:59Z",
        "hasGroups": false,
        "hasStandings": true
      },
      {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/types/2?lang=en&region=us",
        "id": "2",
        "name": "Regular Season",
        "abbreviation": "reg",
        "slug": "regular-season",
        "year": 2026,
        "type": 2,
        "leaders": {},
        "groups": {},
        "startDate": "2025-10-21T07:00Z",
        "endDate": "2026-04-13T06:59Z",
        "hasGroups": false
      }
    ]
  },
  "rankings": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/rankings?lang=en&region=us"
  },
  "powerIndexes": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/powerindex?lang=en&region=us"
  },
  "powerIndexLeaders": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/powerindex/leaders?lang=en&region=us"
  }
}
```
