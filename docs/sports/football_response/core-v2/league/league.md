# League

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}

Notes:
- Verified with `league=nfl` on 2026-05-08.
- The league root currently points at the 2025 NFL season and its offseason state.
- `draft` points to the upcoming draft resource, while `teams` points to a season-scoped team collection.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
|  |  |  |

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl?lang=en&region=us",
  "id": "28",
  "uid": "s:20~l:28",
  "name": "National Football League",
  "abbreviation": "NFL",
  "slug": "nfl",
  "season": {
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
  "teams": {
    "count": 32,
    "pageIndex": 1,
    "pageSize": 25,
    "pageCount": 2,
    "items": [
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {}
    ]
  },
  "athletes": {
    "count": 20173,
    "pageIndex": 1,
    "pageSize": 25,
    "pageCount": 807,
    "items": [
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {}
    ]
  },
  "events": {
    "$meta": {
      "parameters": {
        "week": [
          "5"
        ],
        "season": [
          "2025"
        ],
        "seasontypes": [
          "3"
        ]
      }
    },
    "count": 1,
    "pageIndex": 1,
    "pageSize": 25,
    "pageCount": 1,
    "items": [
      {}
    ]
  },
  "franchises": {
    "count": 32,
    "pageIndex": 1,
    "pageSize": 25,
    "pageCount": 2,
    "items": [
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {}
    ]
  },
  "calendar": {
    "count": 4,
    "pageIndex": 1,
    "pageSize": 25,
    "pageCount": 1,
    "items": [
      {},
      {},
      {},
      {}
    ]
  },
  "rankings": {
    "count": 0,
    "pageIndex": 0,
    "pageSize": 25,
    "pageCount": 0,
    "items": []
  },
  "draft": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2026/draft?lang=en&region=us",
    "uid": "s:20~l:28~e:draft~y:2026",
    "year": 2026,
    "numberOfRounds": 7,
    "displayName": "2026 National Football League Draft",
    "shortDisplayName": "2026 NFL Draft",
    "status": {},
    "athletes": {},
    "rounds": {},
    "positions": [
      {
        "id": "8",
        "name": "Quarterback",
        "displayName": "Quarterback",
        "abbreviation": "QB",
        "leaf": true,
        "parent": {}
      },
      {
        "id": "9",
        "name": "Running Back",
        "displayName": "Running Back",
        "abbreviation": "RB",
        "leaf": true,
        "parent": {}
      },
      {
        "id": "1",
        "name": "Wide Receiver",
        "displayName": "Wide Receiver",
        "abbreviation": "WR",
        "leaf": true,
        "parent": {}
      },
      {
        "id": "7",
        "name": "Tight End",
        "displayName": "Tight End",
        "abbreviation": "TE",
        "leaf": true,
        "parent": {}
      },
      {
        "id": "10",
        "name": "Fullback",
        "displayName": "Fullback",
        "abbreviation": "FB",
        "leaf": true,
        "parent": {}
      },
      {
        "id": "46",
        "name": "Offensive Tackle",
        "displayName": "Offensive Tackle",
        "abbreviation": "OT",
        "leaf": true,
        "parent": {}
      },
      {
        "id": "47",
        "name": "Offensive Guard",
        "displayName": "Offensive Guard",
        "abbreviation": "OG",
        "leaf": true,
        "parent": {}
      },
      {
        "id": "91",
        "name": "Center",
        "displayName": "Center",
        "abbreviation": "C",
        "leaf": true,
        "parent": {}
      },
      {
        "id": "32",
        "name": "Defensive Tackle",
        "displayName": "Defensive Tackle",
        "abbreviation": "DT",
        "leaf": true,
        "parent": {}
      },
      {
        "id": "30",
        "name": "Linebacker",
        "displayName": "Linebacker",
        "abbreviation": "LB",
        "leaf": true,
        "parent": {}
      },
      {
        "id": "264",
        "name": "EDGE",
        "displayName": "EDGE",
        "abbreviation": "EDGE",
        "leaf": true,
        "parent": {}
      },
      {
        "id": "29",
        "name": "Cornerback",
        "displayName": "Cornerback",
        "abbreviation": "CB",
        "leaf": true,
        "parent": {}
      },
      {
        "id": "36",
        "name": "Safety",
        "displayName": "Safety",
        "abbreviation": "S",
        "leaf": true,
        "parent": {}
      },
      {
        "id": "96",
        "name": "Long Snapper",
        "displayName": "Long Snapper",
        "abbreviation": "LS",
        "leaf": true,
        "parent": {}
      },
      {
        "id": "80",
        "name": "Place Kicker",
        "displayName": "Place Kicker",
        "abbreviation": "PK",
        "leaf": true,
        "parent": {}
      },
      {
        "id": "94",
        "name": "Punter",
        "displayName": "Punter",
        "abbreviation": "P",
        "leaf": true,
        "parent": {}
      }
    ],
    "needs": [
      {
        "team": {},
        "positions": [
          {
            "id": "32",
            "name": "Defensive Tackle",
            "displayName": "Defensive Tackle",
            "abbreviation": "DT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "30",
            "name": "Linebacker",
            "displayName": "Linebacker",
            "abbreviation": "LB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "46",
            "name": "Offensive Tackle",
            "displayName": "Offensive Tackle",
            "abbreviation": "OT",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "30",
            "name": "Linebacker",
            "displayName": "Linebacker",
            "abbreviation": "LB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "47",
            "name": "Offensive Guard",
            "displayName": "Offensive Guard",
            "abbreviation": "OG",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "36",
            "name": "Safety",
            "displayName": "Safety",
            "abbreviation": "S",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "32",
            "name": "Defensive Tackle",
            "displayName": "Defensive Tackle",
            "abbreviation": "DT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "91",
            "name": "Center",
            "displayName": "Center",
            "abbreviation": "C",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "46",
            "name": "Offensive Tackle",
            "displayName": "Offensive Tackle",
            "abbreviation": "OT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "30",
            "name": "Linebacker",
            "displayName": "Linebacker",
            "abbreviation": "LB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "47",
            "name": "Offensive Guard",
            "displayName": "Offensive Guard",
            "abbreviation": "OG",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "8",
            "name": "Quarterback",
            "displayName": "Quarterback",
            "abbreviation": "QB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "46",
            "name": "Offensive Tackle",
            "displayName": "Offensive Tackle",
            "abbreviation": "OT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "47",
            "name": "Offensive Guard",
            "displayName": "Offensive Guard",
            "abbreviation": "OG",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "36",
            "name": "Safety",
            "displayName": "Safety",
            "abbreviation": "S",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "30",
            "name": "Linebacker",
            "displayName": "Linebacker",
            "abbreviation": "LB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "7",
            "name": "Tight End",
            "displayName": "Tight End",
            "abbreviation": "TE",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "32",
            "name": "Defensive Tackle",
            "displayName": "Defensive Tackle",
            "abbreviation": "DT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "30",
            "name": "Linebacker",
            "displayName": "Linebacker",
            "abbreviation": "LB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "7",
            "name": "Tight End",
            "displayName": "Tight End",
            "abbreviation": "TE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "46",
            "name": "Offensive Tackle",
            "displayName": "Offensive Tackle",
            "abbreviation": "OT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "47",
            "name": "Offensive Guard",
            "displayName": "Offensive Guard",
            "abbreviation": "OG",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "32",
            "name": "Defensive Tackle",
            "displayName": "Defensive Tackle",
            "abbreviation": "DT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "46",
            "name": "Offensive Tackle",
            "displayName": "Offensive Tackle",
            "abbreviation": "OT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "9",
            "name": "Running Back",
            "displayName": "Running Back",
            "abbreviation": "RB",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "46",
            "name": "Offensive Tackle",
            "displayName": "Offensive Tackle",
            "abbreviation": "OT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "30",
            "name": "Linebacker",
            "displayName": "Linebacker",
            "abbreviation": "LB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "47",
            "name": "Offensive Guard",
            "displayName": "Offensive Guard",
            "abbreviation": "OG",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "47",
            "name": "Offensive Guard",
            "displayName": "Offensive Guard",
            "abbreviation": "OG",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "9",
            "name": "Running Back",
            "displayName": "Running Back",
            "abbreviation": "RB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "30",
            "name": "Linebacker",
            "displayName": "Linebacker",
            "abbreviation": "LB",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "30",
            "name": "Linebacker",
            "displayName": "Linebacker",
            "abbreviation": "LB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "36",
            "name": "Safety",
            "displayName": "Safety",
            "abbreviation": "S",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "91",
            "name": "Center",
            "displayName": "Center",
            "abbreviation": "C",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "32",
            "name": "Defensive Tackle",
            "displayName": "Defensive Tackle",
            "abbreviation": "DT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "7",
            "name": "Tight End",
            "displayName": "Tight End",
            "abbreviation": "TE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "8",
            "name": "Quarterback",
            "displayName": "Quarterback",
            "abbreviation": "QB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "32",
            "name": "Defensive Tackle",
            "displayName": "Defensive Tackle",
            "abbreviation": "DT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "47",
            "name": "Offensive Guard",
            "displayName": "Offensive Guard",
            "abbreviation": "OG",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "46",
            "name": "Offensive Tackle",
            "displayName": "Offensive Tackle",
            "abbreviation": "OT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "8",
            "name": "Quarterback",
            "displayName": "Quarterback",
            "abbreviation": "QB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "7",
            "name": "Tight End",
            "displayName": "Tight End",
            "abbreviation": "TE",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "7",
            "name": "Tight End",
            "displayName": "Tight End",
            "abbreviation": "TE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "36",
            "name": "Safety",
            "displayName": "Safety",
            "abbreviation": "S",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "36",
            "name": "Safety",
            "displayName": "Safety",
            "abbreviation": "S",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "91",
            "name": "Center",
            "displayName": "Center",
            "abbreviation": "C",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "47",
            "name": "Offensive Guard",
            "displayName": "Offensive Guard",
            "abbreviation": "OG",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "30",
            "name": "Linebacker",
            "displayName": "Linebacker",
            "abbreviation": "LB",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "46",
            "name": "Offensive Tackle",
            "displayName": "Offensive Tackle",
            "abbreviation": "OT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "32",
            "name": "Defensive Tackle",
            "displayName": "Defensive Tackle",
            "abbreviation": "DT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "7",
            "name": "Tight End",
            "displayName": "Tight End",
            "abbreviation": "TE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "36",
            "name": "Safety",
            "displayName": "Safety",
            "abbreviation": "S",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "47",
            "name": "Offensive Guard",
            "displayName": "Offensive Guard",
            "abbreviation": "OG",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "30",
            "name": "Linebacker",
            "displayName": "Linebacker",
            "abbreviation": "LB",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "47",
            "name": "Offensive Guard",
            "displayName": "Offensive Guard",
            "abbreviation": "OG",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "32",
            "name": "Defensive Tackle",
            "displayName": "Defensive Tackle",
            "abbreviation": "DT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "46",
            "name": "Offensive Tackle",
            "displayName": "Offensive Tackle",
            "abbreviation": "OT",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "8",
            "name": "Quarterback",
            "displayName": "Quarterback",
            "abbreviation": "QB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "47",
            "name": "Offensive Guard",
            "displayName": "Offensive Guard",
            "abbreviation": "OG",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "7",
            "name": "Tight End",
            "displayName": "Tight End",
            "abbreviation": "TE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "46",
            "name": "Offensive Tackle",
            "displayName": "Offensive Tackle",
            "abbreviation": "OT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "30",
            "name": "Linebacker",
            "displayName": "Linebacker",
            "abbreviation": "LB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "36",
            "name": "Safety",
            "displayName": "Safety",
            "abbreviation": "S",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "8",
            "name": "Quarterback",
            "displayName": "Quarterback",
            "abbreviation": "QB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "46",
            "name": "Offensive Tackle",
            "displayName": "Offensive Tackle",
            "abbreviation": "OT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "36",
            "name": "Safety",
            "displayName": "Safety",
            "abbreviation": "S",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "32",
            "name": "Defensive Tackle",
            "displayName": "Defensive Tackle",
            "abbreviation": "DT",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "8",
            "name": "Quarterback",
            "displayName": "Quarterback",
            "abbreviation": "QB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "47",
            "name": "Offensive Guard",
            "displayName": "Offensive Guard",
            "abbreviation": "OG",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "46",
            "name": "Offensive Tackle",
            "displayName": "Offensive Tackle",
            "abbreviation": "OT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "36",
            "name": "Safety",
            "displayName": "Safety",
            "abbreviation": "S",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "32",
            "name": "Defensive Tackle",
            "displayName": "Defensive Tackle",
            "abbreviation": "DT",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "47",
            "name": "Offensive Guard",
            "displayName": "Offensive Guard",
            "abbreviation": "OG",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "32",
            "name": "Defensive Tackle",
            "displayName": "Defensive Tackle",
            "abbreviation": "DT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "7",
            "name": "Tight End",
            "displayName": "Tight End",
            "abbreviation": "TE",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "46",
            "name": "Offensive Tackle",
            "displayName": "Offensive Tackle",
            "abbreviation": "OT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "91",
            "name": "Center",
            "displayName": "Center",
            "abbreviation": "C",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "36",
            "name": "Safety",
            "displayName": "Safety",
            "abbreviation": "S",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "9",
            "name": "Running Back",
            "displayName": "Running Back",
            "abbreviation": "RB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "46",
            "name": "Offensive Tackle",
            "displayName": "Offensive Tackle",
            "abbreviation": "OT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "32",
            "name": "Defensive Tackle",
            "displayName": "Defensive Tackle",
            "abbreviation": "DT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "30",
            "name": "Linebacker",
            "displayName": "Linebacker",
            "abbreviation": "LB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "7",
            "name": "Tight End",
            "displayName": "Tight End",
            "abbreviation": "TE",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "91",
            "name": "Center",
            "displayName": "Center",
            "abbreviation": "C",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "9",
            "name": "Running Back",
            "displayName": "Running Back",
            "abbreviation": "RB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "36",
            "name": "Safety",
            "displayName": "Safety",
            "abbreviation": "S",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "36",
            "name": "Safety",
            "displayName": "Safety",
            "abbreviation": "S",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "9",
            "name": "Running Back",
            "displayName": "Running Back",
            "abbreviation": "RB",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "32",
            "name": "Defensive Tackle",
            "displayName": "Defensive Tackle",
            "abbreviation": "DT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "30",
            "name": "Linebacker",
            "displayName": "Linebacker",
            "abbreviation": "LB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "9",
            "name": "Running Back",
            "displayName": "Running Back",
            "abbreviation": "RB",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "36",
            "name": "Safety",
            "displayName": "Safety",
            "abbreviation": "S",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "32",
            "name": "Defensive Tackle",
            "displayName": "Defensive Tackle",
            "abbreviation": "DT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "47",
            "name": "Offensive Guard",
            "displayName": "Offensive Guard",
            "abbreviation": "OG",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "1",
            "name": "Wide Receiver",
            "displayName": "Wide Receiver",
            "abbreviation": "WR",
            "leaf": true,
            "parent": {}
          }
        ]
      },
      {
        "team": {},
        "positions": [
          {
            "id": "47",
            "name": "Offensive Guard",
            "displayName": "Offensive Guard",
            "abbreviation": "OG",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "91",
            "name": "Center",
            "displayName": "Center",
            "abbreviation": "C",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "32",
            "name": "Defensive Tackle",
            "displayName": "Defensive Tackle",
            "abbreviation": "DT",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "29",
            "name": "Cornerback",
            "displayName": "Cornerback",
            "abbreviation": "CB",
            "leaf": true,
            "parent": {}
          },
          {
            "id": "264",
            "name": "EDGE",
            "displayName": "EDGE",
            "abbreviation": "EDGE",
            "leaf": true,
            "parent": {}
          }
        ]
      }
    ],
    "broadcasts": [
      {
        "type": {
          "id": "1",
          "shortName": "TV",
          "longName": "TV",
          "slug": "tv"
        },
        "station": "ABC",
        "slug": "abc",
        "priority": 0,
        "market": {
          "id": "1",
          "type": "National"
        },
        "media": {
          "id": "2",
          "callLetters": "ABC",
          "name": "ABC",
          "shortName": "ABC",
          "slug": "abc"
        },
        "lang": "en",
        "region": "us"
      },
      {
        "type": {
          "id": "1",
          "shortName": "TV",
          "longName": "TV",
          "slug": "tv"
        },
        "station": "ESPN",
        "slug": "espn",
        "priority": 0,
        "market": {
          "id": "1",
          "type": "National"
        },
        "media": {
          "id": "126",
          "callLetters": "ESPN",
          "name": "ESPN",
          "shortName": "ESPN",
          "slug": "espn"
        },
        "lang": "en",
        "region": "us"
      },
      {
        "type": {
          "id": "1",
          "shortName": "TV",
          "longName": "TV",
          "slug": "tv"
        },
        "station": "NFL Network",
        "slug": "nfl-network",
        "priority": 0,
        "market": {
          "id": "1",
          "type": "National"
        },
        "media": {
          "id": "398",
          "callLetters": "NFL Net",
          "name": "NFL Network",
          "shortName": "NFL Net",
          "slug": "nfl-network"
        },
        "lang": "en",
        "region": "us"
      },
      {
        "type": {
          "id": "4",
          "shortName": "Streaming",
          "longName": "Streaming",
          "slug": "streaming"
        },
        "station": "ESPN+",
        "slug": "espnplus",
        "priority": 0,
        "market": {
          "id": "1",
          "type": "National"
        },
        "media": {
          "id": "755",
          "callLetters": "ESPN+",
          "name": "ESPN+",
          "shortName": "ESPN+",
          "slug": "espn"
        },
        "lang": "en",
        "region": "us"
      }
    ],
    "links": [
      {
        "language": "en-US",
        "rel": [
          "index",
          "desktop",
          "draft"
        ],
        "href": "https://www.espn.com/nfl/draft",
        "text": "NFL Draft",
        "shortText": "NFL Draft",
        "isExternal": false,
        "isPremium": false
      }
    ],
    "startDate": "2026-04-24T00:00Z",
    "endDate": "2026-04-26T03:59Z"
  },
  "notes": {
    "count": 0,
    "pageIndex": 0,
    "pageSize": 25,
    "pageCount": 0,
    "items": []
  }
}
```
