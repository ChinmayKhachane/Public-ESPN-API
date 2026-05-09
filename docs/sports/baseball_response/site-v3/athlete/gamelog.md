# MLB Athlete Gamelog

## https://site.web.api.espn.com/apis/common/v3/sports/baseball/mlb/athletes/4414528/gamelog

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "events": {
    "401815220": {
      "id": "401815220",
      "team": {
        "id": "17",
        "uid": "s:1~l:10~t:17",
        "abbreviation": "CIN",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/_/name/cin/cincinnati-reds",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/roster/_/name/cin/cincinnati-reds",
            "text": "Roster",
            "shortText": "Roster",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
        "isAllStar": false
      },
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/mlb/game/_/gameId/401815220/reds-cubs",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=baseball&leagueAbbrev=mlb&gameId=401815220",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "week": 18,
      "atVs": "@",
      "gameDate": "2026-05-05T23:40:00.000+00:00",
      "score": "3-2 F/10",
      "homeTeamId": "16",
      "awayTeamId": "17",
      "homeTeamScore": "3"
    },
    "401815148": {
      "id": "401815148",
      "team": {
        "id": "17",
        "uid": "s:1~l:10~t:17",
        "abbreviation": "CIN",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/_/name/cin/cincinnati-reds",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/roster/_/name/cin/cincinnati-reds",
            "text": "Roster",
            "shortText": "Roster",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
        "isAllStar": false
      },
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/mlb/game/_/gameId/401815148/rockies-reds",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=baseball&leagueAbbrev=mlb&gameId=401815148",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "week": 18,
      "atVs": "vs",
      "gameDate": "2026-04-30T16:40:00.000+00:00",
      "score": "6-4",
      "homeTeamId": "17",
      "awayTeamId": "27",
      "homeTeamScore": "6"
    },
    "401815068": {
      "id": "401815068",
      "team": {
        "id": "17",
        "uid": "s:1~l:10~t:17",
        "abbreviation": "CIN",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/_/name/cin/cincinnati-reds",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/roster/_/name/cin/cincinnati-reds",
            "text": "Roster",
            "shortText": "Roster",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
        "isAllStar": false
      },
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/mlb/game/_/gameId/401815068/tigers-reds",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=baseball&leagueAbbrev=mlb&gameId=401815068",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "week": 17,
      "atVs": "vs",
      "gameDate": "2026-04-24T22:40:00.000+00:00",
      "score": "9-8",
      "homeTeamId": "17",
      "awayTeamId": "6",
      "homeTeamScore": "9"
    },
    "401814993": {
      "id": "401814993",
      "team": {
        "id": "17",
        "uid": "s:1~l:10~t:17",
        "abbreviation": "CIN",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/_/name/cin/cincinnati-reds",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/roster/_/name/cin/cincinnati-reds",
            "text": "Roster",
            "shortText": "Roster",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
        "isAllStar": false
      },
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/mlb/game/_/gameId/401814993/reds-twins",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=baseball&leagueAbbrev=mlb&gameId=401814993",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "week": 16,
      "atVs": "@",
      "gameDate": "2026-04-18T18:10:00.000+00:00",
      "score": "5-4",
      "homeTeamId": "9",
      "awayTeamId": "17",
      "homeTeamScore": "4"
    },
    "401814906": {
      "id": "401814906",
      "team": {
        "id": "17",
        "uid": "s:1~l:10~t:17",
        "abbreviation": "CIN",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/_/name/cin/cincinnati-reds",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/roster/_/name/cin/cincinnati-reds",
            "text": "Roster",
            "shortText": "Roster",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
        "isAllStar": false
      },
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/mlb/game/_/gameId/401814906/angels-reds",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=baseball&leagueAbbrev=mlb&gameId=401814906",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "week": 15,
      "atVs": "vs",
      "gameDate": "2026-04-12T17:40:00.000+00:00",
      "score": "9-6",
      "homeTeamId": "17",
      "awayTeamId": "3",
      "homeTeamScore": "6"
    },
    "401814843": {
      "id": "401814843",
      "team": {
        "id": "17",
        "uid": "s:1~l:10~t:17",
        "abbreviation": "CIN",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/_/name/cin/cincinnati-reds",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/roster/_/name/cin/cincinnati-reds",
            "text": "Roster",
            "shortText": "Roster",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
        "isAllStar": false
      },
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/mlb/game/_/gameId/401814843/reds-marlins",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=baseball&leagueAbbrev=mlb&gameId=401814843",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "week": 14,
      "atVs": "@",
      "gameDate": "2026-04-07T22:40:00.000+00:00",
      "score": "6-3 F/10",
      "homeTeamId": "28",
      "awayTeamId": "17",
      "homeTeamScore": "3"
    },
    "401814767": {
      "id": "401814767",
      "team": {
        "id": "17",
        "uid": "s:1~l:10~t:17",
        "abbreviation": "CIN",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/_/name/cin/cincinnati-reds",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/roster/_/name/cin/cincinnati-reds",
            "text": "Roster",
            "shortText": "Roster",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
        "isAllStar": false
      },
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/mlb/game/_/gameId/401814767/pirates-reds",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=baseball&leagueAbbrev=mlb&gameId=401814767",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "week": 13,
      "atVs": "vs",
      "gameDate": "2026-04-01T16:40:00.000+00:00",
      "score": "8-3",
      "homeTeamId": "17",
      "awayTeamId": "23",
      "homeTeamScore": "3"
    },
    "401814689": {
      "id": "401814689",
      "team": {
        "id": "17",
        "uid": "s:1~l:10~t:17",
        "abbreviation": "CIN",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/_/name/cin/cincinnati-reds",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/roster/_/name/cin/cincinnati-reds",
            "text": "Roster",
            "shortText": "Roster",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
        "isAllStar": false
      },
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/mlb/game/_/gameId/401814689/red-sox-reds",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=baseball&leagueAbbrev=mlb&gameId=401814689",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "week": 13,
      "atVs": "vs",
      "gameDate": "2026-03-26T20:10:00.000+00:00",
      "score": "3-0",
      "homeTeamId": "17",
      "awayTeamId": "2",
      "homeTeamScore": "0"
    }
  },
  "filters": [
    {
      "name": "season",
      "displayName": "Season",
      "value": "2026",
      "options": [
        {
          "value": "2026",
          "displayValue": "2026"
        },
        {
          "value": "2025",
          "displayValue": "2025"
        }
      ]
    },
    {
      "name": "category",
      "displayName": "Category",
      "value": "pitching",
      "options": [
        {
          "value": "batting",
          "displayValue": "Batting",
          "shortDisplayName": "batting"
        },
        {
          "value": "pitching",
          "displayValue": "Pitching",
          "shortDisplayName": "pitching"
        }
      ]
    }
  ],
  "labels": [
    "IP",
    "H"
  ],
  "names": [
    "innings",
    "hits"
  ],
  "displayNames": [
    "Innings pitched",
    "Hits"
  ],
  "seasonTypes": [
    {
      "displayName": "2026 Regular Season",
      "categories": [
        {
          "displayName": "may",
          "events": [
            {}
          ],
          "type": "event",
          "splitType": "may",
          "totals": [
            "5.2",
            "4"
          ]
        },
        {
          "displayName": "april",
          "events": [
            {},
            {}
          ],
          "type": "event",
          "splitType": "april",
          "totals": [
            "28.2",
            "36"
          ]
        }
      ],
      "displayTeam": "CIN"
    }
  ],
  "glossary": [
    {
      "displayName": "Walks",
      "abbreviation": "BB"
    },
    {
      "displayName": "Decision",
      "abbreviation": "Dec"
    }
  ]
}
```
