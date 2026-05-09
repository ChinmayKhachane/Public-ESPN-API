# League

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.
- NHL league root currently points at the 2025-26 postseason context.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl`

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl?lang=en&region=us",
  "id": "90",
  "guid": "1a5f0227-a13e-396c-8cea-8961bc288666",
  "uid": "s:70~l:90",
  "name": "National Hockey League",
  "displayName": "NHL",
  "abbreviation": "NHL",
  "shortName": "NHL",
  "slug": "nhl",
  "isTournament": false,
  "season": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026?lang=en&region=us",
    "year": 2026,
    "startDate": "2025-09-20T07:00Z",
    "endDate": "2026-07-01T06:59Z",
    "displayName": "2025-26",
    "type": {
      "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/3?lang=en&region=us",
      "id": "3",
      "type": 3,
      "name": "Postseason",
      "abbreviation": "post",
      "year": 2026,
      "startDate": "2026-04-18T07:00Z",
      "endDate": "2026-07-01T06:59Z",
      "hasGroups": false,
      "hasStandings": false,
      "hasLegs": false,
      "groups": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/3/groups?lang=en&region=us"
      },
      "week": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/3/weeks/4?lang=en&region=us",
        "number": 4,
        "startDate": "2026-05-09T07:00Z",
        "endDate": "2026-05-16T06:59Z",
        "text": "Week 4",
        "rankings": {
          "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/3/weeks/4/rankings?lang=en&region=us"
        }
      },
      "weeks": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/3/weeks?lang=en&region=us"
      },
      "corrections": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/3/corrections?lang=en&region=us"
      },
      "leaders": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/3/leaders?lang=en&region=us"
      }
    },
    "types": {
      "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types?lang=en&region=us",
      "count": 4,
      "pageIndex": 1,
      "pageSize": 4,
      "pageCount": 1,
      "items": [
        {
          "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/1?lang=en&region=us",
          "id": "1",
          "type": 1,
          "name": "Preseason",
          "abbreviation": "pre",
          "year": 2026,
          "startDate": "2025-09-20T07:00Z",
          "endDate": "2025-10-07T06:59Z",
          "hasGroups": false,
          "hasStandings": true,
          "hasLegs": false,
          "groups": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/1/groups?lang=en&region=us"
          },
          "weeks": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/1/weeks?lang=en&region=us"
          },
          "corrections": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/1/corrections?lang=en&region=us"
          },
          "leaders": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/1/leaders?lang=en&region=us"
          },
          "slug": "preseason"
        },
        {
          "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/2?lang=en&region=us",
          "id": "2",
          "type": 2,
          "name": "Regular Season",
          "abbreviation": "reg",
          "year": 2026,
          "startDate": "2025-10-07T07:00Z",
          "endDate": "2026-04-18T06:59Z",
          "hasGroups": false,
          "hasStandings": true,
          "hasLegs": false,
          "groups": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/2/groups?lang=en&region=us"
          },
          "weeks": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/2/weeks?lang=en&region=us"
          },
          "corrections": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/2/corrections?lang=en&region=us"
          },
          "leaders": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/2/leaders?lang=en&region=us"
          },
          "slug": "regular-season"
        }
      ]
    },
    "rankings": {
      "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/rankings?lang=en&region=us"
    },
    "coaches": {
      "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/coaches?lang=en&region=us"
    },
    "athletes": {
      "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/athletes?lang=en&region=us"
    },
    "awards": {
      "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/awards?lang=en&region=us"
    },
    "futures": {
      "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/futures?lang=en&region=us"
    },
    "leaders": {
      "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/3/leaders?lang=en&region=us"
    }
  },
  "seasons": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons?lang=en&region=us"
  },
  "franchises": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/franchises?lang=en&region=us"
  },
  "teams": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/teams?lang=en&region=us"
  },
  "group": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/3/groups/9?lang=en&region=us"
  },
  "groups": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/3/groups?lang=en&region=us"
  }
}
```
