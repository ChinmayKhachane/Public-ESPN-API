# Seasons

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/seasons

Notes:
- Verified with `league=nfl` on 2026-05-08.
- The collection is historical and includes future season containers.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `page` | `int` | Page number |
| `limit` | `int` | Page size |

## Example Response

```json
{
  "count": 105,
  "pageIndex": 1,
  "pageSize": 3,
  "pageCount": 35,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2026?lang=en&region=us",
      "year": 2026,
      "startDate": "2026-08-06T07:00Z",
      "endDate": "2027-02-16T07:59Z",
      "displayName": "2026",
      "type": {
        "id": "2",
        "type": 2,
        "name": "Regular Season",
        "abbreviation": "reg",
        "year": 2026,
        "startDate": "2026-09-09T07:00Z",
        "endDate": "2027-01-13T07:59Z",
        "hasGroups": false,
        "hasStandings": false,
        "hasLegs": false,
        "groups": {},
        "weeks": {},
        "corrections": {},
        "slug": "regular-season"
      },
      "types": {
        "count": 4,
        "pageIndex": 1,
        "pageSize": 4,
        "pageCount": 1,
        "items": [
          {
            "id": "1",
            "type": 1,
            "name": "Preseason",
            "abbreviation": "pre",
            "year": 2026,
            "startDate": "2026-08-06T07:00Z",
            "endDate": "2026-09-09T06:59Z",
            "hasGroups": false,
            "hasStandings": false,
            "hasLegs": false,
            "groups": {},
            "weeks": {},
            "corrections": {},
            "slug": "preseason"
          },
          {
            "id": "2",
            "type": 2,
            "name": "Regular Season",
            "abbreviation": "reg",
            "year": 2026,
            "startDate": "2026-09-09T07:00Z",
            "endDate": "2027-01-13T07:59Z",
            "hasGroups": false,
            "hasStandings": false,
            "hasLegs": false,
            "groups": {},
            "weeks": {},
            "corrections": {},
            "slug": "regular-season"
          },
          {
            "id": "3",
            "type": 3,
            "name": "Postseason",
            "abbreviation": "post",
            "year": 2026,
            "startDate": "2027-01-13T08:00Z",
            "endDate": "2027-02-16T07:59Z",
            "hasGroups": false,
            "hasStandings": false,
            "hasLegs": false,
            "groups": {},
            "weeks": {},
            "corrections": {},
            "slug": "post-season"
          },
          {
            "id": "4",
            "type": 4,
            "name": "Off Season",
            "abbreviation": "off",
            "year": 2026,
            "startDate": "2027-02-16T08:00Z",
            "endDate": "2027-08-01T06:59Z",
            "hasGroups": false,
            "hasStandings": false,
            "hasLegs": false,
            "groups": {},
            "weeks": {},
            "slug": "off-season"
          }
        ]
      },
      "rankings": {},
      "futures": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025?lang=en&region=us",
      "year": 2025,
      "startDate": "2025-07-31T07:00Z",
      "endDate": "2026-02-12T07:59Z",
      "displayName": "2025",
      "type": {
        "id": "2",
        "type": 2,
        "name": "Regular Season",
        "abbreviation": "reg",
        "year": 2025,
        "startDate": "2025-09-04T07:00Z",
        "endDate": "2026-01-07T07:59Z",
        "hasGroups": false,
        "hasStandings": true,
        "hasLegs": false,
        "groups": {},
        "weeks": {},
        "corrections": {},
        "leaders": {},
        "slug": "regular-season"
      },
      "types": {
        "count": 4,
        "pageIndex": 1,
        "pageSize": 4,
        "pageCount": 1,
        "items": [
          {
            "id": "1",
            "type": 1,
            "name": "Preseason",
            "abbreviation": "pre",
            "year": 2025,
            "startDate": "2025-07-31T07:00Z",
            "endDate": "2025-09-04T06:59Z",
            "hasGroups": false,
            "hasStandings": true,
            "hasLegs": false,
            "groups": {},
            "weeks": {},
            "corrections": {},
            "leaders": {},
            "slug": "preseason"
          },
          {
            "id": "2",
            "type": 2,
            "name": "Regular Season",
            "abbreviation": "reg",
            "year": 2025,
            "startDate": "2025-09-04T07:00Z",
            "endDate": "2026-01-07T07:59Z",
            "hasGroups": false,
            "hasStandings": true,
            "hasLegs": false,
            "groups": {},
            "weeks": {},
            "corrections": {},
            "leaders": {},
            "slug": "regular-season"
          },
          {
            "id": "3",
            "type": 3,
            "name": "Postseason",
            "abbreviation": "post",
            "year": 2025,
            "startDate": "2026-01-07T08:00Z",
            "endDate": "2026-02-12T07:59Z",
            "hasGroups": false,
            "hasStandings": false,
            "hasLegs": false,
            "groups": {},
            "weeks": {},
            "corrections": {},
            "leaders": {},
            "slug": "post-season"
          },
          {
            "id": "4",
            "type": 4,
            "name": "Off Season",
            "abbreviation": "off",
            "year": 2025,
            "startDate": "2026-02-12T08:00Z",
            "endDate": "2026-08-06T06:59Z",
            "hasGroups": false,
            "hasStandings": false,
            "hasLegs": false,
            "groups": {},
            "week": {
              "number": 1,
              "startDate": "2026-02-12T08:00Z",
              "endDate": "2026-08-01T06:59Z",
              "text": "Week 1",
              "rankings": {},
              "events": {}
            },
            "weeks": {},
            "slug": "off-season"
          }
        ]
      },
      "rankings": {},
      "coaches": {},
      "athletes": {},
      "awards": {},
      "futures": {},
      "leaders": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2024?lang=en&region=us",
      "year": 2024,
      "startDate": "2024-08-01T07:00Z",
      "endDate": "2025-02-15T07:59Z",
      "displayName": "2024",
      "type": {
        "id": "2",
        "type": 2,
        "name": "Regular Season",
        "abbreviation": "reg",
        "year": 2024,
        "startDate": "2024-09-05T07:00Z",
        "endDate": "2025-01-08T07:59Z",
        "hasGroups": false,
        "hasStandings": true,
        "hasLegs": false,
        "groups": {},
        "weeks": {},
        "corrections": {},
        "leaders": {},
        "slug": "regular-season"
      },
      "types": {
        "count": 4,
        "pageIndex": 1,
        "pageSize": 4,
        "pageCount": 1,
        "items": [
          {
            "id": "1",
            "type": 1,
            "name": "Preseason",
            "abbreviation": "pre",
            "year": 2024,
            "startDate": "2024-08-01T07:00Z",
            "endDate": "2024-09-05T06:59Z",
            "hasGroups": false,
            "hasStandings": true,
            "hasLegs": false,
            "groups": {},
            "weeks": {},
            "corrections": {},
            "leaders": {},
            "slug": "preseason"
          },
          {
            "id": "2",
            "type": 2,
            "name": "Regular Season",
            "abbreviation": "reg",
            "year": 2024,
            "startDate": "2024-09-05T07:00Z",
            "endDate": "2025-01-08T07:59Z",
            "hasGroups": false,
            "hasStandings": true,
            "hasLegs": false,
            "groups": {},
            "weeks": {},
            "corrections": {},
            "leaders": {},
            "slug": "regular-season"
          },
          {
            "id": "3",
            "type": 3,
            "name": "Postseason",
            "abbreviation": "post",
            "year": 2024,
            "startDate": "2025-01-08T08:00Z",
            "endDate": "2025-02-15T07:59Z",
            "hasGroups": false,
            "hasStandings": false,
            "hasLegs": false,
            "groups": {},
            "weeks": {},
            "corrections": {},
            "leaders": {},
            "slug": "post-season"
          },
          {
            "id": "4",
            "type": 4,
            "name": "Off Season",
            "abbreviation": "off",
            "year": 2024,
            "startDate": "2025-02-15T08:00Z",
            "endDate": "2025-07-31T06:59Z",
            "hasGroups": false,
            "hasStandings": false,
            "hasLegs": false,
            "groups": {},
            "weeks": {},
            "slug": "off-season"
          }
        ]
      },
      "rankings": {},
      "awards": {},
      "futures": {},
      "leaders": {}
    }
  ]
}
```
