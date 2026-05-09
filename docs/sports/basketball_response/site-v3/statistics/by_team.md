# Statistics By Team

## https://site.web.api.espn.com/apis/common/v3/sports/basketball/{league}/statistics/byteam

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "teams": [
    {
      "team": {
        "id": "7",
        "uid": "s:40~l:46~t:7",
        "guid": "c4aceb39-0eb9-a30b-1120-9cb5b12b677a",
        "name": "Nuggets",
        "displayName": "Denver Nuggets",
        "shortDisplayName": "Nuggets",
        "abbreviation": "DEN",
        "slug": "denver-nuggets",
        "logos": [],
        "links": [],
        "group": {},
        "ranks": {}
      },
      "categories": [
        {},
        {}
      ]
    },
    {
      "team": {
        "id": "14",
        "uid": "s:40~l:46~t:14",
        "guid": "81e3212c-30ef-9b1b-5edb-453b13ff265a",
        "name": "Heat",
        "displayName": "Miami Heat",
        "shortDisplayName": "Heat",
        "abbreviation": "MIA",
        "slug": "miami-heat",
        "logos": [],
        "links": [],
        "group": {},
        "ranks": {}
      },
      "categories": [
        {},
        {}
      ]
    }
  ],
  "categories": [
    {
      "name": "general",
      "labels": [
        "REB",
        "REB"
      ],
      "names": [
        "totalRebounds",
        "avgRebounds"
      ],
      "displayNames": [
        "Rebounds",
        "Rebounds Per Game"
      ],
      "descriptions": [
        "The total number of rebounds for a team or player",
        "The average rebounds per game."
      ]
    },
    {
      "name": "general"
    }
  ],
  "glossary": [
    {
      "displayName": "3-Point Field Goal Percentage",
      "abbreviation": "3P%"
    },
    {
      "displayName": "3-Point Field Goal Percentage Differential",
      "abbreviation": "3P%Diff"
    }
  ],
  "pagination": {
    "count": 30,
    "limit": 2,
    "page": 1,
    "pages": 15,
    "first": "http://site.api.espn.com:80/apis/common/v3/sports/basketball/nba/statistics/byteam?season=2026&seasontype=2&limit=2",
    "next": "http://site.api.espn.com:80/apis/common/v3/sports/basketball/nba/statistics/byteam?season=2026&seasontype=2&limit=2&page=2",
    "last": "http://site.api.espn.com:80/apis/common/v3/sports/basketball/nba/statistics/byteam?season=2026&seasontype=2&limit=2&page=15"
  },
  "currentSeason": {
    "displayName": "2025-26",
    "year": 2026,
    "type": {
      "id": "3",
      "name": "Postseason",
      "type": 3,
      "startDate": "2026-04-18T07:00:00.000+00:00",
      "endDate": "2026-06-27T06:59:00.000+00:00",
      "week": {}
    },
    "startDate": "2025-10-01T07:00:00.000+00:00",
    "endDate": "2026-06-27T06:59:00.000+00:00"
  },
  "requestedSeason": {
    "displayName": "2025-26",
    "year": 2026,
    "type": {
      "id": "2",
      "name": "Regular Season",
      "type": 2,
      "startDate": "2025-10-21T07:00:00.000+00:00",
      "endDate": "2026-04-13T06:59:00.000+00:00",
      "week": {
        "text": "Week 25",
        "number": 25,
        "startDate": "2026-04-07T07:00:00.000+00:00",
        "endDate": "2026-04-13T06:59:00.000+00:00"
      }
    },
    "startDate": "2025-10-01T07:00:00.000+00:00",
    "endDate": "2026-06-27T06:59:00.000+00:00"
  },
  "league": {
    "id": "46",
    "uid": "s:40~l:46",
    "name": "National Basketball Association",
    "abbreviation": "NBA",
    "slug": "nba",
    "shortName": "NBA"
  }
}
```
