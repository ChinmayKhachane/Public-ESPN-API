# Statistics By Athlete

## https://site.web.api.espn.com/apis/common/v3/sports/basketball/{league}/statistics/byathlete

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "athletes": [
    {
      "athlete": {
        "id": "3945274",
        "uid": "s:40~l:46~a:3945274",
        "guid": "583794eb-0f38-9bbd-3e25-9dd33b7f83b8",
        "displayName": "Luka Doncic",
        "slug": "luka-doncic",
        "type": "basketball",
        "status": {},
        "teams": [],
        "firstName": "Luka",
        "lastName": "Doncic",
        "shortName": "L. Doncic",
        "debutYear": 2018
      },
      "categories": [
        {},
        {}
      ]
    },
    {
      "athlete": {
        "id": "4278073",
        "uid": "s:40~l:46~a:4278073",
        "guid": "4dcec409-3ff9-2881-2bc3-b4289ce6c36d",
        "displayName": "Shai Gilgeous-Alexander",
        "slug": "shai-gilgeous-alexander",
        "type": "basketball",
        "status": {},
        "teams": [],
        "firstName": "Shai",
        "lastName": "Gilgeous-Alexander",
        "shortName": "S. Gilgeous-Alexander",
        "debutYear": 2018
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
      "displayName": "Own General",
      "labels": [
        "GP",
        "MIN"
      ],
      "names": [
        "gamesPlayed",
        "avgMinutes"
      ],
      "displayNames": [
        "Games Played",
        "Minutes Per Game"
      ],
      "descriptions": [
        "Games Played",
        "The average number of minutes per game."
      ]
    },
    {
      "name": "offensive",
      "displayName": "Own Offensive",
      "labels": [
        "PTS",
        "FGM"
      ],
      "names": [
        "avgPoints",
        "avgFieldGoalsMade"
      ],
      "displayNames": [
        "Points Per Game",
        "Average Field Goals Made"
      ],
      "descriptions": [
        "The average number of points scored per game.",
        "The average field goals made per game."
      ]
    }
  ],
  "glossary": [
    {
      "displayName": "3-Point Field Goal Percentage",
      "abbreviation": "3P%"
    },
    {
      "displayName": "Average 3-Point Field Goals Attempted",
      "abbreviation": "3PA"
    }
  ],
  "pagination": {
    "count": 578,
    "limit": 2,
    "page": 1,
    "pages": 289,
    "first": "http://site.api.espn.com:80/apis/common/v3/sports/basketball/nba/statistics/byathlete?season=2026&seasontype=2&limit=2",
    "next": "http://site.api.espn.com:80/apis/common/v3/sports/basketball/nba/statistics/byathlete?season=2026&seasontype=2&limit=2&page=2",
    "last": "http://site.api.espn.com:80/apis/common/v3/sports/basketball/nba/statistics/byathlete?season=2026&seasontype=2&limit=2&page=289"
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
