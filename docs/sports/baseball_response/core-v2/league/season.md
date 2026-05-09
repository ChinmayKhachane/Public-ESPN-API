# MLB Current Season

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/season

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026?lang=en&region=us",
  "displayName": "2026",
  "athletes": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/athletes?lang=en&region=us"
  },
  "year": 2026,
  "startDate": "2026-02-19T08:00Z",
  "endDate": "2026-11-12T07:59Z",
  "type": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/types/2?lang=en&region=us",
    "id": "2",
    "name": "Regular Season",
    "abbreviation": "reg",
    "slug": "regular-season",
    "type": 2,
    "year": 2026,
    "startDate": "2026-03-25T07:00Z",
    "endDate": "2026-09-30T06:59Z",
    "hasGroups": false
  },
  "types": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/types?lang=en&region=us",
    "count": 4,
    "pageIndex": 1,
    "pageSize": 4,
    "pageCount": 1,
    "items": [
      {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/types/1?lang=en&region=us",
        "id": "1",
        "name": "Spring Training",
        "abbreviation": "pre",
        "slug": "preseason",
        "type": 1,
        "year": 2026,
        "startDate": "2026-02-19T08:00Z",
        "endDate": "2026-03-25T06:59Z",
        "hasGroups": false
      },
      {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/types/2?lang=en&region=us",
        "id": "2",
        "name": "Regular Season",
        "abbreviation": "reg",
        "slug": "regular-season",
        "type": 2,
        "year": 2026,
        "startDate": "2026-03-25T07:00Z",
        "endDate": "2026-09-30T06:59Z",
        "hasGroups": false
      }
    ]
  },
  "rankings": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/rankings?lang=en&region=us"
  },
  "coaches": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/coaches?lang=en&region=us"
  },
  "futures": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/futures?lang=en&region=us"
  },
  "leaders": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/types/2/leaders?lang=en&region=us"
  }
}
```
