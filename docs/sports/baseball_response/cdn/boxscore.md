# MLB CDN Boxscore

## https://cdn.espn.com/core/mlb/boxscore?xhr=1&gameId=401815256

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "gamepackageJSON": {
    "header": {
      "id": "401815256",
      "uid": "s:1~l:10~e:401815256",
      "competitions": [
        {
          "id": "401815256",
          "uid": "s:1~l:10~e:401815256~c:401815256",
          "competitors": [
            {},
            {}
          ],
          "status": {
            "periodPrefix": "End",
            "type": {},
            "featuredAthletes": []
          },
          "date": "2026-05-08T22:10Z",
          "commentaryAvailable": false,
          "conferenceCompetition": false,
          "liveAvailable": false,
          "broadcasts": [
            {},
            {}
          ],
          "playByPlaySource": "full"
        }
      ],
      "season": {
        "current": true,
        "year": 2026,
        "type": 2
      },
      "week": 19,
      "timeValid": true,
      "league": {
        "id": "10",
        "uid": "s:1~l:10",
        "name": "Major League Baseball",
        "abbreviation": "MLB",
        "slug": "mlb",
        "midsizeName": "MLB",
        "links": [
          {
            "rel": [],
            "href": "https://www.espn.com/mlb/",
            "text": "Index"
          },
          {
            "rel": [],
            "href": "sportscenter://x-callback-url/showClubhouse?uid=s:1~l:10",
            "text": "Index"
          }
        ],
        "logos": [
          {
            "rel": [],
            "href": "https://a.espncdn.com/i/teamlogos/leagues/500/mlb.png"
          },
          {
            "rel": [],
            "href": "https://a.espncdn.com/combiner/i?img=/i/teamlogos/leagues/500-dark/mlb.png&w=500&h=500&transparent=true"
          }
        ],
        "isTournament": false
      },
      "links": [
        {
          "isExternal": false,
          "shortText": "Summary",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/mlb/game/_/gameId/401815256/astros-reds",
          "text": "Gamecast",
          "isPremium": false
        },
        {
          "isExternal": false,
          "shortText": "Box Score",
          "rel": [
            "boxscore",
            "desktop"
          ],
          "href": "https://www.espn.com/mlb/boxscore/_/gameId/401815256",
          "text": "Box Score",
          "isPremium": false
        }
      ]
    },
    "boxscore": {
      "teams": [
        {
          "team": {
            "id": "18",
            "uid": "s:1~l:10~t:18",
            "name": "Astros",
            "displayName": "Houston Astros",
            "abbreviation": "HOU",
            "slug": "houston-astros",
            "shortDisplayName": "Astros",
            "alternateColor": "eb6e1f",
            "color": "002d62",
            "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/hou.png"
          },
          "statistics": [
            {},
            {}
          ],
          "homeAway": "away",
          "displayOrder": 1,
          "details": [
            {},
            {}
          ]
        },
        {
          "team": {
            "id": "17",
            "uid": "s:1~l:10~t:17",
            "name": "Reds",
            "displayName": "Cincinnati Reds",
            "abbreviation": "CIN",
            "slug": "cincinnati-reds",
            "shortDisplayName": "Reds",
            "alternateColor": "ffffff",
            "color": "c6011f",
            "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png"
          },
          "statistics": [
            {},
            {}
          ],
          "homeAway": "home",
          "displayOrder": 2,
          "details": [
            {},
            {}
          ]
        }
      ],
      "players": [
        {
          "team": {
            "id": "18",
            "uid": "s:1~l:10~t:18",
            "name": "Astros",
            "displayName": "Houston Astros",
            "abbreviation": "HOU",
            "slug": "houston-astros",
            "shortDisplayName": "Astros",
            "alternateColor": "eb6e1f",
            "color": "002d62",
            "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/hou.png"
          },
          "statistics": [
            {},
            {}
          ],
          "displayOrder": 1
        },
        {
          "team": {
            "id": "17",
            "uid": "s:1~l:10~t:17",
            "name": "Reds",
            "displayName": "Cincinnati Reds",
            "abbreviation": "CIN",
            "slug": "cincinnati-reds",
            "shortDisplayName": "Reds",
            "alternateColor": "ffffff",
            "color": "c6011f",
            "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png"
          },
          "statistics": [
            {},
            {}
          ],
          "displayOrder": 2
        }
      ]
    },
    "plays": [
      {
        "id": "4018152560000000059",
        "team": {
          "id": "18"
        },
        "sequenceNumber": "1",
        "period": {
          "displayValue": "1st Inning",
          "number": 1,
          "type": "Top"
        },
        "resultCount": {
          "balls": 0,
          "strikes": 0
        },
        "homeScore": 0,
        "atBatId": "4018152560001",
        "summaryType": "I",
        "scoringPlay": false,
        "type": {
          "id": "59",
          "text": "Start Inning",
          "type": "start-inning"
        }
      },
      {
        "id": "4018152560001010001",
        "team": {
          "id": "18"
        },
        "sequenceNumber": "1",
        "period": {
          "displayValue": "1st Inning",
          "number": 1,
          "type": "Top"
        },
        "resultCount": {
          "balls": 0,
          "strikes": 0
        },
        "homeScore": 0,
        "atBatId": "4018152560001",
        "summaryType": "A",
        "batOrder": 1,
        "scoringPlay": false
      }
    ],
    "winprobability": [],
    "news": {
      "header": "MLB News",
      "link": {
        "isExternal": false,
        "shortText": "All News",
        "rel": [
          "index",
          "desktop"
        ],
        "language": "en-US",
        "href": "https://www.espn.com/mlb/",
        "text": "All MLB News",
        "isPremium": false
      },
      "articles": [
        {
          "id": 48717767,
          "categories": [
            {},
            {}
          ],
          "contentKey": "48717767-1-5-1",
          "images": [
            {}
          ],
          "dataSourceIdentifier": "f60e80c46aee3",
          "published": "2026-05-09T03:43:19Z",
          "type": "HeadlineNews",
          "nowId": "1-48717767",
          "premium": false,
          "links": {
            "app": {},
            "web": {},
            "mobile": {},
            "api": {}
          }
        },
        {
          "id": 48717747,
          "categories": [
            {},
            {}
          ],
          "contentKey": "48717747-1-293-1",
          "images": [
            {}
          ],
          "dataSourceIdentifier": "c9a50ebd56d4c",
          "published": "2026-05-09T03:30:14Z",
          "type": "Media",
          "nowId": "1-48717747",
          "premium": false,
          "links": {
            "sportscenter": {},
            "web": {},
            "api": {}
          }
        }
      ]
    },
    "seasonseries": [
      {
        "events": [
          {
            "id": "401815256",
            "uid": "s:1~l:10~e:401815256~c:401815256",
            "competitors": [],
            "status": "post",
            "date": "2026-05-08T22:10:00Z",
            "statusType": {},
            "timeValid": true,
            "links": [],
            "neutralSite": false
          },
          {
            "id": "401815271",
            "uid": "s:1~l:10~e:401815271~c:401815271",
            "competitors": [],
            "status": "pre",
            "date": "2026-05-09T20:10:00Z",
            "statusType": {},
            "timeValid": true,
            "broadcasts": [],
            "links": [],
            "neutralSite": false
          }
        ],
        "totalCompetitions": 3,
        "seriesScore": "1-0",
        "completed": false,
        "type": "current",
        "title": "Current Series"
      },
      {
        "events": [
          {
            "id": "401815256",
            "uid": "s:1~l:10~e:401815256~c:401815256",
            "competitors": [],
            "status": "post",
            "date": "2026-05-08T22:10:00Z",
            "statusType": {},
            "timeValid": true,
            "links": [],
            "neutralSite": false
          },
          {
            "id": "401815271",
            "uid": "s:1~l:10~e:401815271~c:401815271",
            "competitors": [],
            "status": "pre",
            "date": "2026-05-09T20:10:00Z",
            "statusType": {},
            "timeValid": true,
            "broadcasts": [],
            "links": [],
            "neutralSite": false
          }
        ],
        "totalCompetitions": 3,
        "shortSummary": "HOU leads season",
        "seriesScore": "1-0",
        "seriesLabel": "Regular Season",
        "completed": false,
        "type": "season",
        "title": "Regular Season Series"
      }
    ],
    "broadcasts": [
      {
        "market": {
          "id": "1",
          "type": "National"
        },
        "station": "MLB.TV",
        "media": {
          "name": "MLB.TV",
          "callLetters": "MLB.TV",
          "shortName": "MLB.TV"
        },
        "type": {
          "id": "4",
          "slug": "streaming",
          "shortName": "Streaming",
          "longName": "Streaming"
        },
        "lang": "en",
        "region": "us",
        "isNational": true
      },
      {
        "market": {
          "id": "2",
          "type": "Home"
        },
        "station": "Reds.TV",
        "media": {
          "name": "Reds.TV",
          "callLetters": "Reds.TV",
          "shortName": "Reds.TV"
        },
        "type": {
          "id": "4",
          "slug": "streaming",
          "shortName": "Streaming",
          "longName": "Streaming"
        },
        "lang": "en",
        "region": "us",
        "isNational": false
      }
    ],
    "videos": [
      {
        "id": 48716451,
        "cerebroId": "69fe8549f795f455a4a6d052",
        "thumbnail": "https://a.espncdn.com/media/motion/wsc/2026/0509/1545d730-902f-4c9e-a0f2-14f383b9bc3e/1545d730-902f-4c9e-a0...",
        "ad": {
          "sport": "mlb",
          "bundle": "mlbhighlightsoneplay"
        },
        "timeRestrictions": {
          "embargoDate": "2026-05-09T00:52:17Z",
          "expirationDate": "2026-05-11T00:52:17Z"
        },
        "geoRestrictions": {
          "countries": [
            "PR",
            "HN"
          ],
          "type": "whitelist"
        },
        "source": "espn",
        "duration": 88,
        "deviceRestrictions": {
          "devices": [
            "settop",
            "tablet"
          ],
          "type": "whitelist"
        },
        "originalPublishDate": "2026-05-09T00:52:25Z"
      },
      {
        "id": 48715382,
        "cerebroId": "69fe6571c35ca554fbcb52d9",
        "thumbnail": "https://a.espncdn.com/media/motion/wsc/2026/0508/0b25d762-e51c-425e-b9d1-1a935f4b9be1/0b25d762-e51c-425e-b9...",
        "ad": {
          "sport": "mlb",
          "bundle": "mlbhighlightsoneplay"
        },
        "timeRestrictions": {
          "embargoDate": "2026-05-08T22:36:26Z",
          "expirationDate": "2026-05-10T04:00:00Z"
        },
        "geoRestrictions": {
          "countries": [
            "PR",
            "HN"
          ],
          "type": "whitelist"
        },
        "source": "espn",
        "duration": 37,
        "deviceRestrictions": {
          "devices": [
            "settop",
            "tablet"
          ],
          "type": "whitelist"
        },
        "originalPublishDate": "2026-05-08T22:36:34Z"
      }
    ],
    "standings": {
      "fullViewLink": {
        "text": "Full Standings",
        "href": "https://www.espn.com/mlb/standings"
      },
      "groups": [
        {
          "header": "2026 American League West Standings",
          "href": "https://www.espn.com/mlb/standings/_/group/3",
          "standings": {
            "entries": []
          }
        },
        {
          "header": "2026 National League Central Standings",
          "href": "https://www.espn.com/mlb/standings/_/group/5",
          "standings": {
            "entries": []
          }
        }
      ]
    }
  },
  "gameId": 401815256,
  "customStyleSheet": "game-package-baseball",
  "type": "boxscore",
  "content": {
    "title": "Astros vs. Reds - Box Score - May 8, 2026 - ESPN",
    "og_type": "website",
    "sport": "baseball",
    "league": "mlb",
    "tab": {
      "layout": "bc",
      "pageType": "Boxscore",
      "metaDescription": "Get box score updates on the {aDisplayName} vs. {hDisplayName} baseball game.",
      "columnsModuleTypes": {
        "default": [
          [
            "linescore",
            "boxscore"
          ],
          [
            "ad",
            "cliplist"
          ]
        ],
        "tablet": [
          [
            "linescore",
            "boxscore"
          ],
          []
        ],
        "mobile": [
          [
            "linescore",
            "boxscore"
          ],
          []
        ]
      },
      "metaTitle": "{aName} vs. {hName} - Box Score - {date}"
    },
    "statusState": "post",
    "canonical": "http://www.espn.com/mlb/boxscore/_/gameId/401815256",
    "tabType": "boxscore"
  },
  "__gamepackage__": {
    "playerHash": {
      "4422899": {
        "json": {
          "starter": true,
          "athlete": {
            "id": "4422899",
            "uid": "s:1~l:10~a:4422899",
            "guid": "413a1088-6b03-3030-ac99-7067979621bd",
            "displayName": "Matt McLain",
            "hotZones": [],
            "headshot": {},
            "links": [],
            "positions": [],
            "position": {},
            "shortName": "M. McLain"
          },
          "stats": [
            "1-3",
            "3"
          ],
          "active": true,
          "batOrder": 8,
          "atBats": [
            {},
            {}
          ],
          "position": {
            "id": "4",
            "name": "Second Base",
            "displayName": "Second Baseman",
            "abbreviation": "2B"
          }
        },
        "batting": {
          "BB": {
            "label": "BB",
            "displayValue": "0"
          },
          "AB": {
            "label": "AB",
            "displayValue": "3"
          },
          "R": {
            "label": "R",
            "displayValue": "0"
          },
          "AVG": {
            "label": "AVG",
            "displayValue": ".204"
          },
          "H-AB": {
            "label": "H-AB",
            "displayValue": "1-3"
          },
          "H": {
            "label": "H",
            "displayValue": "1"
          },
          "RBI": {
            "label": "RBI",
            "displayValue": "0"
          },
          "HR": {
            "label": "HR",
            "displayValue": "0"
          },
          "K": {
            "label": "K",
            "displayValue": "0"
          },
          "OBP": {
            "label": "OBP",
            "displayValue": ".310"
          }
        },
        "homeAway": "home",
        "teamColor": "c6011f"
      },
      "4722857": {
        "json": {
          "starter": true,
          "athlete": {
            "id": "4722857",
            "uid": "s:1~l:10~a:4722857",
            "guid": "9b8d6559-9c75-3d6b-bd28-b683001c0b83",
            "displayName": "Spencer Steer",
            "hotZones": [],
            "headshot": {},
            "links": [],
            "positions": [],
            "position": {},
            "shortName": "S. Steer"
          },
          "stats": [
            "0-4",
            "4"
          ],
          "active": true,
          "batOrder": 6,
          "atBats": [
            {},
            {}
          ],
          "positions": [
            {},
            {}
          ],
          "position": {
            "id": "3",
            "name": "First Base",
            "displayName": "First Baseman",
            "abbreviation": "1B"
          }
        },
        "batting": {
          "BB": {
            "label": "BB",
            "displayValue": "0"
          },
          "AB": {
            "label": "AB",
            "displayValue": "4"
          },
          "R": {
            "label": "R",
            "displayValue": "0"
          },
          "AVG": {
            "label": "AVG",
            "displayValue": ".246"
          },
          "H-AB": {
            "label": "H-AB",
            "displayValue": "0-4"
          },
          "H": {
            "label": "H",
            "displayValue": "0"
          },
          "RBI": {
            "label": "RBI",
            "displayValue": "0"
          },
          "HR": {
            "label": "HR",
            "displayValue": "0"
          },
          "K": {
            "label": "K",
            "displayValue": "2"
          },
          "OBP": {
            "label": "OBP",
            "displayValue": ".324"
          }
        },
        "homeAway": "home",
        "teamColor": "c6011f"
      },
      "5137883": {
        "json": {
          "starter": true,
          "athlete": {
            "id": "5137883",
            "uid": "s:1~l:10~a:5137883",
            "guid": "50137d73-0766-3fc4-bc70-e58098112705",
            "displayName": "Zach Dezenzo",
            "hotZones": [],
            "headshot": {},
            "links": [],
            "positions": [],
            "position": {},
            "shortName": "Z. Dezenzo"
          },
          "stats": [
            "1-3",
            "3"
          ],
          "active": false,
          "batOrder": 7,
          "atBats": [
            {},
            {}
          ],
          "position": {
            "id": "7",
            "name": "Left Field",
            "displayName": "Left Fielder",
            "abbreviation": "LF"
          }
        },
        "batting": {
          "BB": {
            "label": "BB",
            "displayValue": "0"
          },
          "AB": {
            "label": "AB",
            "displayValue": "3"
          },
          "R": {
            "label": "R",
            "displayValue": "1"
          },
          "AVG": {
            "label": "AVG",
            "displayValue": ".167"
          },
          "H-AB": {
            "label": "H-AB",
            "displayValue": "1-3"
          },
          "H": {
            "label": "H",
            "displayValue": "1"
          },
          "RBI": {
            "label": "RBI",
            "displayValue": "2"
          },
          "HR": {
            "label": "HR",
            "displayValue": "1"
          },
          "K": {
            "label": "K",
            "displayValue": "1"
          },
          "OBP": {
            "label": "OBP",
            "displayValue": ".167"
          }
        },
        "homeAway": "away",
        "teamColor": "002d62"
      },
      "42408": {
        "json": {
          "notes": [
            {}
          ],
          "starter": false,
          "athlete": {
            "id": "42408",
            "uid": "s:1~l:10~a:42408",
            "guid": "7d4729e6-6a75-3635-8f9c-a5f76b3c72a1",
            "displayName": "Braden Shewmake",
            "hotZones": [],
            "headshot": {},
            "links": [],
            "positions": [],
            "position": {},
            "shortName": "B. Shewmake"
          },
          "stats": [
            "1-1",
            "1"
          ],
          "active": true,
          "batOrder": 1,
          "atBats": [
            {}
          ],
          "positions": [
            {},
            {}
          ],
          "position": {
            "id": "4",
            "name": "Second Base",
            "displayName": "Second Baseman",
            "abbreviation": "2B"
          }
        },
        "batting": {
          "BB": {
            "label": "BB",
            "displayValue": "0"
          },
          "AB": {
            "label": "AB",
            "displayValue": "1"
          },
          "R": {
            "label": "R",
            "displayValue": "0"
          },
          "AVG": {
            "label": "AVG",
            "displayValue": ".353"
          },
          "H-AB": {
            "label": "H-AB",
            "displayValue": "1-1"
          },
          "H": {
            "label": "H",
            "displayValue": "1"
          },
          "RBI": {
            "label": "RBI",
            "displayValue": "0"
          },
          "HR": {
            "label": "HR",
            "displayValue": "0"
          },
          "K": {
            "label": "K",
            "displayValue": "0"
          },
          "OBP": {
            "label": "OBP",
            "displayValue": ".353"
          }
        },
        "homeAway": "away",
        "teamColor": "002d62"
      },
      "36020": {
        "json": {
          "starter": true,
          "athlete": {
            "id": "36020",
            "uid": "s:1~l:10~a:36020",
            "guid": "9c422f07-8499-da29-dadb-539a4824ad17",
            "displayName": "TJ Friedl",
            "hotZones": [],
            "headshot": {},
            "links": [],
            "positions": [],
            "position": {},
            "shortName": "T. Friedl"
          },
          "stats": [
            "0-4",
            "4"
          ],
          "active": true,
          "batOrder": 1,
          "atBats": [
            {},
            {}
          ],
          "position": {
            "id": "8",
            "name": "Center Field",
            "displayName": "Center Fielder",
            "abbreviation": "CF"
          }
        },
        "batting": {
          "BB": {
            "label": "BB",
            "displayValue": "0"
          },
          "AB": {
            "label": "AB",
            "displayValue": "4"
          },
          "R": {
            "label": "R",
            "displayValue": "0"
          },
          "AVG": {
            "label": "AVG",
            "displayValue": ".188"
          },
          "H-AB": {
            "label": "H-AB",
            "displayValue": "0-4"
          },
          "H": {
            "label": "H",
            "displayValue": "0"
          },
          "RBI": {
            "label": "RBI",
            "displayValue": "0"
          },
          "HR": {
            "label": "HR",
            "displayValue": "0"
          },
          "K": {
            "label": "K",
            "displayValue": "1"
          },
          "OBP": {
            "label": "OBP",
            "displayValue": ".265"
          }
        },
        "homeAway": "home",
        "teamColor": "c6011f"
      },
      "5136815": {
        "json": {
          "starter": false,
          "athlete": {
            "id": "5136815",
            "uid": "s:1~l:10~a:5136815",
            "guid": "5aa24b58-ff8d-39b2-8837-29764ebf6f84",
            "displayName": "Zach Cole",
            "hotZones": [],
            "headshot": {},
            "links": [],
            "positions": [],
            "position": {},
            "shortName": "Z. Cole"
          },
          "stats": [
            "1-1",
            "1"
          ],
          "active": true,
          "batOrder": 7,
          "atBats": [
            {}
          ],
          "position": {
            "id": "7",
            "name": "Left Field",
            "displayName": "Left Fielder",
            "abbreviation": "LF"
          }
        },
        "batting": {
          "BB": {
            "label": "BB",
            "displayValue": "0"
          },
          "AB": {
            "label": "AB",
            "displayValue": "1"
          },
          "R": {
            "label": "R",
            "displayValue": "1"
          },
          "AVG": {
            "label": "AVG",
            "displayValue": ".273"
          },
          "H-AB": {
            "label": "H-AB",
            "displayValue": "1-1"
          },
          "H": {
            "label": "H",
            "displayValue": "1"
          },
          "RBI": {
            "label": "RBI",
            "displayValue": "3"
          },
          "HR": {
            "label": "HR",
            "displayValue": "1"
          },
          "K": {
            "label": "K",
            "displayValue": "0"
          },
          "OBP": {
            "label": "OBP",
            "displayValue": ".273"
          }
        },
        "homeAway": "away",
        "teamColor": "002d62"
      },
      "4917694": {
        "json": {
          "starter": true,
          "athlete": {
            "id": "4917694",
            "uid": "s:1~l:10~a:4917694",
            "guid": "69d42781-0bf5-3c3a-94be-592c98ef717f",
            "displayName": "Elly De La Cruz",
            "hotZones": [],
            "headshot": {},
            "links": [],
            "positions": [],
            "position": {},
            "shortName": "E. De La Cruz"
          },
          "stats": [
            "2-4",
            "4"
          ],
          "active": true,
          "batOrder": 3,
          "atBats": [
            {},
            {}
          ],
          "position": {
            "id": "6",
            "name": "Shortstop",
            "displayName": "Shortstop",
            "abbreviation": "SS"
          }
        },
        "batting": {
          "BB": {
            "label": "BB",
            "displayValue": "0"
          },
          "AB": {
            "label": "AB",
            "displayValue": "4"
          },
          "R": {
            "label": "R",
            "displayValue": "0"
          },
          "AVG": {
            "label": "AVG",
            "displayValue": ".271"
          },
          "H-AB": {
            "label": "H-AB",
            "displayValue": "2-4"
          },
          "H": {
            "label": "H",
            "displayValue": "2"
          },
          "RBI": {
            "label": "RBI",
            "displayValue": "0"
          },
          "HR": {
            "label": "HR",
            "displayValue": "0"
          },
          "K": {
            "label": "K",
            "displayValue": "1"
          },
          "OBP": {
            "label": "OBP",
            "displayValue": ".341"
          }
        },
        "homeAway": "home",
        "teamColor": "c6011f"
      },
      "39216": {
        "json": {
          "notes": [
            {}
          ],
          "starter": false,
          "athlete": {
            "id": "39216",
            "uid": "s:1~l:10~a:39216",
            "guid": "ac261c9d-f174-3c5c-d5ab-24d7c986ed8f",
            "displayName": "Will Benson",
            "hotZones": [],
            "headshot": {},
            "links": [],
            "positions": [],
            "position": {},
            "shortName": "W. Benson"
          },
          "stats": [
            "0-1",
            "1"
          ],
          "active": true,
          "batOrder": 9,
          "atBats": [
            {}
          ],
          "positions": [
            {},
            {}
          ],
          "position": {
            "id": "9",
            "name": "Right Field",
            "displayName": "Right Fielder",
            "abbreviation": "RF"
          }
        },
        "batting": {
          "BB": {
            "label": "BB",
            "displayValue": "0"
          },
          "AB": {
            "label": "AB",
            "displayValue": "1"
          },
          "R": {
            "label": "R",
            "displayValue": "0"
          },
          "AVG": {
            "label": "AVG",
            "displayValue": ".211"
          },
          "H-AB": {
            "label": "H-AB",
            "displayValue": "0-1"
          },
          "H": {
            "label": "H",
            "displayValue": "0"
          },
          "RBI": {
            "label": "RBI",
            "displayValue": "0"
          },
          "HR": {
            "label": "HR",
            "displayValue": "0"
          },
          "K": {
            "label": "K",
            "displayValue": "1"
          },
          "OBP": {
            "label": "OBP",
            "displayValue": ".338"
          }
        },
        "homeAway": "home",
        "teamColor": "c6011f"
      },
      "5136495": {
        "json": {
          "starter": false,
          "athlete": {
            "id": "5136495",
            "uid": "s:1~l:10~a:5136495",
            "guid": "ce4ac8c9-c644-3a7d-a3ae-8135e9ada28d",
            "displayName": "Logan VanWey",
            "headshot": {},
            "links": [],
            "position": {},
            "shortName": "L. VanWey",
            "throws": "RHP"
          },
          "stats": [
            "1.0",
            "1"
          ],
          "active": true,
          "batOrder": 0,
          "position": {
            "id": "1",
            "name": "Pitcher",
            "displayName": "Pitcher",
            "abbreviation": "P"
          }
        },
        "pitching": {
          "BB": {
            "label": "BB",
            "displayValue": "0"
          },
          "R": {
            "label": "R",
            "displayValue": "0"
          },
          "PC": {
            "label": "PC",
            "displayValue": "20"
          },
          "ERA": {
            "label": "ERA",
            "displayValue": "0.00"
          },
          "IP": {
            "label": "IP",
            "displayValue": "1.0"
          },
          "H": {
            "label": "H",
            "displayValue": "1"
          },
          "HR": {
            "label": "HR",
            "displayValue": "0"
          },
          "K": {
            "label": "K",
            "displayValue": "2"
          },
          "ER": {
            "label": "ER",
            "displayValue": "0"
          },
          "PC-ST": {
            "label": "PC-ST",
            "displayValue": "20-14"
          }
        },
        "homeAway": "away",
        "teamColor": "002d62"
      },
      "5080771": {
        "json": {
          "starter": true,
          "athlete": {
            "id": "5080771",
            "uid": "s:1~l:10~a:5080771",
            "guid": "0784deac-211b-3475-896c-00bb70b18918",
            "displayName": "Sal Stewart",
            "hotZones": [],
            "headshot": {},
            "links": [],
            "positions": [],
            "position": {},
            "shortName": "S. Stewart"
          },
          "stats": [
            "1-4",
            "4"
          ],
          "active": true,
          "batOrder": 4,
          "atBats": [
            {},
            {}
          ],
          "positions": [
            {},
            {}
          ],
          "position": {
            "id": "5",
            "name": "Third Base",
            "displayName": "Third Baseman",
            "abbreviation": "3B"
          }
        },
        "batting": {
          "BB": {
            "label": "BB",
            "displayValue": "0"
          },
          "AB": {
            "label": "AB",
            "displayValue": "4"
          },
          "R": {
            "label": "R",
            "displayValue": "0"
          },
          "AVG": {
            "label": "AVG",
            "displayValue": ".245"
          },
          "H-AB": {
            "label": "H-AB",
            "displayValue": "1-4"
          },
          "H": {
            "label": "H",
            "displayValue": "1"
          },
          "RBI": {
            "label": "RBI",
            "displayValue": "0"
          },
          "HR": {
            "label": "HR",
            "displayValue": "0"
          },
          "K": {
            "label": "K",
            "displayValue": "0"
          },
          "OBP": {
            "label": "OBP",
            "displayValue": ".331"
          }
        },
        "homeAway": "home",
        "teamColor": "c6011f"
      }
    },
    "awayTeam": {
      "id": "18",
      "uid": "s:1~l:10~t:18",
      "team": {
        "id": "18",
        "uid": "s:1~l:10~t:18",
        "guid": "00a3015f-09ec-1b03-52af-656f5e0a18d5",
        "name": "Astros",
        "displayName": "Houston Astros",
        "abbreviation": "HOU",
        "alternateColor": "eb6e1f",
        "color": "002d62",
        "location": "Houston",
        "links": [
          {
            "rel": [],
            "href": "https://www.espn.com/mlb/team/_/name/hou/houston-astros",
            "text": "Clubhouse"
          }
        ]
      },
      "possession": false,
      "probables": [
        {
          "name": "probableStartingPitcher",
          "displayName": "Probable Starting Pitcher",
          "abbreviation": "SP",
          "statistics": {
            "splits": {}
          },
          "shortDisplayName": "Starter",
          "athlete": {
            "id": "4918155",
            "uid": "s:1~l:10~a:4918155",
            "guid": "f3b25d20-2c07-396f-83f5-dfc859f24c58",
            "displayName": "Mike Burrows",
            "status": {},
            "lastName": "Burrows",
            "fullName": "Mike Burrows",
            "throws": {},
            "headshot": {},
            "jersey": "50"
          },
          "playerId": 4918155
        }
      ],
      "hits": 13,
      "homeAway": "away",
      "score": "10",
      "winner": true,
      "record": [
        {
          "displayValue": "16-23",
          "type": "total"
        },
        {
          "displayValue": "7-13",
          "type": "road"
        }
      ]
    },
    "homeTeam": {
      "id": "17",
      "uid": "s:1~l:10~t:17",
      "team": {
        "id": "17",
        "uid": "s:1~l:10~t:17",
        "guid": "04b65a0b-3cca-d795-0e21-23606470418a",
        "name": "Reds",
        "displayName": "Cincinnati Reds",
        "abbreviation": "CIN",
        "alternateColor": "ffffff",
        "color": "c6011f",
        "location": "Cincinnati",
        "links": [
          {
            "rel": [],
            "href": "https://www.espn.com/mlb/team/_/name/cin/cincinnati-reds",
            "text": "Clubhouse"
          }
        ]
      },
      "possession": false,
      "probables": [
        {
          "name": "probableStartingPitcher",
          "displayName": "Probable Starting Pitcher",
          "abbreviation": "SP",
          "statistics": {
            "splits": {}
          },
          "shortDisplayName": "Starter",
          "athlete": {
            "id": "42433",
            "uid": "s:1~l:10~a:42433",
            "guid": "4e3529e7-ab6a-3a42-a59b-e3c04345ba31",
            "displayName": "Nick Lodolo",
            "status": {},
            "lastName": "Lodolo",
            "fullName": "Nick Lodolo",
            "throws": {},
            "headshot": {},
            "jersey": "40"
          },
          "playerId": 42433
        }
      ],
      "hits": 5,
      "homeAway": "home",
      "score": "0",
      "winner": false,
      "record": [
        {
          "displayValue": "20-19",
          "type": "total"
        },
        {
          "displayValue": "10-9",
          "type": "home"
        }
      ]
    },
    "awayTeamLogo": "https://a.espncdn.com/combiner/i?img=/i/teamlogos/mlb/500/hou.png&h=100&w=100",
    "homeTeamLogo": "https://a.espncdn.com/combiner/i?img=/i/teamlogos/mlb/500/cin.png&h=100&w=100",
    "highlightPlayers": false,
    "airingsHash": {
      "onWatch": false,
      "onDTC": false,
      "gameOnEPlus": false,
      "isOOM": false,
      "airingsAll": [],
      "airingsTVE": [],
      "airingsDTC": [],
      "networkHashTVE": {},
      "networkHashDTC": {},
      "userIsEntitledTVE": false
    }
  },
  "analytics": {
    "metrics": {
      "page_url": "/mlb/boxscore/_/gameId/401815256",
      "site": "espn",
      "game_state": "post",
      "content_type": "gamecast",
      "page_infrastructure": "sCore",
      "page_type": "boxscore",
      "game_detail": "401815256 Houston Astros vs Cincinnati Reds",
      "league": "mlb",
      "page_name": "espn:mlb:game:boxscore",
      "section": "baseball"
    },
    "omniture": {
      "gameInfo": "401815256 Houston Astros vs Cincinnati Reds",
      "league": "mlb",
      "countryRegion": "en-us",
      "channel": "mlb",
      "hier1": "mlb:game:boxscore",
      "section": "baseball",
      "pageName": "mlb:game:boxscore",
      "sections": "mlb:game",
      "site": "espn",
      "premium": "premium-no"
    },
    "chartbeat": {
      "domain": "www.espn.com",
      "sections": "baseball",
      "authors": "gamecast",
      "path": "/mlb/boxscore/_/gameId/401815256",
      "title": "Astros vs. Reds - Box Score - May 8, 2026 - ESPN",
      "zone": "www.espn.com.us.baseball",
      "loadPubJS": false,
      "loadVidJS": true
    },
    "nielsen": {
      "espnuk": {
        "apid": "P07264C85-15CD-4A80-8E56-B5BFA6D93296",
        "vc": "b01"
      },
      "espnau": {
        "apid": "P07264C85-15CD-4A80-8E56-B5BFA6D93296",
        "vc": "b01"
      },
      "espn": {
        "apid": "P07264C85-15CD-4A80-8E56-B5BFA6D93296",
        "vc": "b01"
      },
      "fantasy": {
        "apid": "P302B69D5-F1DD-4E7A-BF8D-3E60F0EB5E5A",
        "vc": "c07"
      },
      "espndeportes": {
        "apid": "P890E2723-EDBC-4CCE-96BA-F35EA3E50650",
        "vc": "c02"
      },
      "espnfc": {
        "apid": "PE6995AAE-0C49-4372-B5E7-54C61BFE2AA5",
        "vc": "c03"
      },
      "espnww": {
        "apid": "P07264C85-15CD-4A80-8E56-B5BFA6D93296",
        "vc": "b01"
      },
      "general": {
        "ci": "us-600140",
        "assetid": "N/A",
        "segB": "N/A",
        "sfcode": "dcr",
        "segA": "N/A",
        "section": "N/A",
        "segC": "N/A",
        "apn": "espnCOM"
      },
      "espnza": {
        "apid": "P07264C85-15CD-4A80-8E56-B5BFA6D93296",
        "vc": "b01"
      },
      "espnin": {
        "apid": "P07264C85-15CD-4A80-8E56-B5BFA6D93296",
        "vc": "b01"
      }
    },
    "device": "desktop",
    "cto": true,
    "qualtrics": false
  },
  "ads": {
    "id": 12129264,
    "page_url": "https://www.espn.com/mlb/boxscore/_/gameId/401815256",
    "prebidAdConfig": {
      "usePrebidBids": true,
      "timeout": 1000
    },
    "level": "espn.com/mlb/boxscore",
    "sizesEspnPlus": {
      "banner-index": {
        "excludedSize": [
          "728,90"
        ],
        "mappings": [
          {
            "viewport": [],
            "slot": []
          },
          {
            "viewport": [],
            "slot": []
          }
        ],
        "defaultSize": [
          970,
          66
        ],
        "excludedProfile": [
          "xl"
        ],
        "includedCountries": [
          "us"
        ],
        "pbjs": {
          "s": [
            []
          ],
          "xl": [
            []
          ],
          "l": [
            []
          ],
          "m": [
            []
          ]
        }
      },
      "gamecast": {
        "mappings": [
          {
            "viewport": [],
            "slot": []
          }
        ],
        "defaultSize": [
          320,
          50
        ]
      },
      "banner-scoreboard": {
        "excludedSize": [
          "970,250"
        ],
        "mappings": [
          {
            "viewport": [],
            "slot": []
          },
          {
            "viewport": [],
            "slot": []
          }
        ],
        "defaultSize": [
          970,
          66
        ],
        "includedCountries": [
          "us"
        ],
        "pbjs": {
          "s": [
            []
          ],
          "xl": [
            []
          ],
          "l": [
            []
          ],
          "m": [
            []
          ]
        }
      },
      "banner": {
        "mappings": [
          {
            "viewport": [],
            "slot": []
          },
          {
            "viewport": [],
            "slot": []
          }
        ],
        "defaultSize": [
          970,
          66
        ],
        "pbjs": {
          "s": [
            []
          ],
          "xl": [
            [],
            []
          ],
          "l": [
            [],
            []
          ],
          "m": [
            []
          ]
        }
      },
      "incontent-betting": {
        "mappings": [
          {
            "viewport": [],
            "slot": []
          },
          {
            "viewport": [],
            "slot": []
          }
        ],
        "defaultSize": [
          300,
          251
        ]
      },
      "native-betting": {
        "mappings": [
          {
            "viewport": [],
            "slot": []
          }
        ],
        "defaultSize": "fluid"
      },
      "instream": {
        "mappings": [
          {
            "viewport": [],
            "slot": []
          }
        ],
        "defaultSize": [
          1,
          3
        ]
      },
      "incontent": {
        "mappings": [
          {
            "viewport": [],
            "slot": []
          }
        ],
        "defaultSize": [
          300,
          250
        ]
      }
    },
    "delayInPageAdSlots": true,
    "incontentPositions": {
      "defaults": {
        "favorites": -1,
        "news": 4,
        "now": 4
      },
      "index": {
        "top": {
          "favorites": -1
        },
        "nfl": {}
      }
    },
    "showEspnPlusAds": false,
    "kvpsEspnPlus": [
      {
        "name": "ed",
        "value": "us"
      },
      {
        "name": "eplus",
        "value": "true"
      }
    ],
    "network": "21783347309"
  },
  "targeting": {},
  "meta": {
    "imageWidth": 1200,
    "image": "https://s.espncdn.com/stitcher/sports/baseball/mlb/events/401815256.png?templateId=espn.com.share.1",
    "twitter_card": "summary",
    "og_site_name": "ESPN.com",
    "twitter_app_id_iphone": "317469184",
    "og_type": "website",
    "twitter_app_name_googleplay": "ESPN",
    "label": "MLB",
    "canonical": "https://www.espn.com/mlb/boxscore/_/gameId/401815256",
    "type": "gamepackage"
  },
  "nowFeedSupported": true,
  "customNav": "<div id=\"gamepackage-header-wrap\" class=\"\"><div id=\"gamepackage-matchup-wrap\"><header class=\"game-strip gam...",
  "sport": [
    "mlb"
  ],
  "tier2Nav": {
    "subNavMenu": {
      "navigation": {
        "$ref": "/v2/navigation/12001873",
        "id": 12001873,
        "items": [
          {
            "$ref": "/v2/navigation/12001915",
            "id": 12001915,
            "links": [],
            "title": "MLB Home"
          },
          {
            "$ref": "/v2/navigation/11586778",
            "id": 11586778,
            "links": [],
            "title": "MLB Scores"
          }
        ],
        "links": [
          {
            "isExternal": false,
            "shortText": "MLB",
            "rel": [],
            "text": "MLB",
            "href": "/mlb/",
            "isPremium": false
          }
        ],
        "attributes": {
          "sport_id": "10",
          "root": "mlb"
        },
        "text": "MLB",
        "title": "MLB Menu - LIVE"
      },
      "navId": 12001873,
      "fallback": false
    }
  }
}
```
