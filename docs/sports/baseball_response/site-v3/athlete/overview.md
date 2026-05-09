# MLB Athlete Overview

## https://site.web.api.espn.com/apis/common/v3/sports/baseball/mlb/athletes/4414528/overview

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "statistics": {
    "displayName": "2026 Pitching",
    "labels": [
      "GP",
      "GS"
    ],
    "names": [
      "gamesPlayed",
      "gamesStarted"
    ],
    "displayNames": [
      "Games Played",
      "Games Started"
    ],
    "splits": [
      {
        "displayName": "Regular Season",
        "stats": [
          "8",
          "8"
        ]
      },
      {
        "displayName": "Projected",
        "stats": [
          "34",
          "34"
        ]
      }
    ]
  },
  "news": [
    {
      "id": 48678093,
      "categories": [
        {
          "id": 12023,
          "type": "league",
          "sportId": 0,
          "leagueId": 3070,
          "league": {
            "id": 3070
          }
        },
        {
          "id": 80311,
          "uid": "s:1~l:10",
          "type": "league",
          "sportId": 10,
          "leagueId": 10,
          "league": {
            "id": 10,
            "abbreviation": "MLB",
            "links": {}
          }
        }
      ],
      "headline": "Fantasy baseball lineup advice for Tuesday: Jameson Taillon, Elmer Rodriguez in a good spot",
      "lastModified": "2026-05-05T16:53:20.000+00:00",
      "root": "flb/2026",
      "premium": false,
      "links": {
        "api": {
          "self": {
            "href": "https://content.core.api.espn.com/v1/sports/news/48678093"
          }
        },
        "mobile": {
          "href": "http://m.espn.go.com/wireless/story?storyId=48678093"
        },
        "web": {
          "href": "https://www.espn.com/fantasy/baseball/story/_/id/48678093/fantasy-baseball-pitcher-rankings-lineup-advice-t..."
        }
      },
      "type": "Story",
      "section": "Fantasy MLB",
      "linkText": "Fantasy baseball lineup advice for Tuesday: Taillon among top streaming options"
    },
    {
      "id": 48630465,
      "categories": [
        {
          "id": 12023,
          "type": "league",
          "sportId": 0,
          "leagueId": 3070,
          "league": {
            "id": 3070
          }
        },
        {
          "id": 80311,
          "uid": "s:1~l:10",
          "type": "league",
          "sportId": 10,
          "leagueId": 10,
          "league": {
            "id": 10,
            "abbreviation": "MLB",
            "links": {}
          }
        }
      ],
      "headline": "Fantasy baseball lineup advice for Thursday: Can Andrew Abbott get back on track?",
      "lastModified": "2026-04-30T12:00:28.000+00:00",
      "root": "flb/2026",
      "premium": false,
      "links": {
        "api": {
          "self": {
            "href": "https://content.core.api.espn.com/v1/sports/news/48630465"
          }
        },
        "mobile": {
          "href": "http://m.espn.go.com/wireless/story?storyId=48630465"
        },
        "web": {
          "href": "https://www.espn.com/fantasy/baseball/story/_/id/48630465/fantasy-baseball-pitcher-rankings-lineup-advice-t..."
        }
      },
      "type": "Story",
      "section": "Fantasy MLB",
      "linkText": "Fantasy baseball lineup advice for Thursday: Can Andrew Abbott get back on track?"
    }
  ],
  "nextGame": {
    "displayName": "Previous Game",
    "statistics": {
      "labels": [
        "GP",
        "GS"
      ],
      "names": [
        "gamesPlayed",
        "gamesStarted"
      ],
      "displayNames": [
        "Games Played",
        "Games Started"
      ],
      "splits": [
        {
          "displayName": "Last seven days",
          "stats": [
            "1",
            "1"
          ],
          "type": "lastSevenDays"
        },
        {
          "displayName": "Home",
          "stats": [
            "5",
            "5"
          ],
          "type": "splits"
        }
      ]
    },
    "league": {
      "id": "10",
      "uid": "s:1~l:10",
      "name": "Major League Baseball",
      "abbreviation": "MLB",
      "slug": "mlb",
      "events": [
        {
          "id": "401815256",
          "uid": "s:1~l:10~e:401815256",
          "name": "Houston Astros at Cincinnati Reds",
          "competitors": [
            {},
            {}
          ],
          "status": "post",
          "season": 2026,
          "competitionId": "401815256",
          "date": "2026-05-08T22:10:00.000+00:00",
          "timeValid": true,
          "shortName": "HOU @ CIN"
        }
      ],
      "shortName": "MLB"
    },
    "summaryStatistics": [
      {
        "name": "wins-losses",
        "displayName": "Wins-Losses",
        "abbreviation": "W-L",
        "splitId": "61",
        "splitAbbreviation": "Last seven days",
        "splitName": "Last 7 Days",
        "shortDisplayName": "W-L",
        "displayValue": "0-0",
        "splitBy": "lastSevenDays"
      },
      {
        "name": "ERA",
        "displayName": "Earned Run Average",
        "abbreviation": "ERA",
        "splitId": "61",
        "splitAbbreviation": "Last seven days",
        "splitName": "Last 7 Days",
        "shortDisplayName": "ERA",
        "displayValue": "0.00",
        "splitBy": "lastSevenDays"
      }
    ]
  },
  "gameLog": {
    "displayName": "Recent Games",
    "events": {
      "401815068": {
        "id": "401815068",
        "team": {
          "id": "17",
          "uid": "s:1~l:10~t:17",
          "abbreviation": "CIN",
          "links": [
            {},
            {}
          ],
          "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
          "isAllStar": false
        },
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/game/_/gameId/401815068/tigers-reds",
            "text": "Gamecast",
            "shortText": "Summary",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
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
      "401815220": {
        "id": "401815220",
        "team": {
          "id": "17",
          "uid": "s:1~l:10~t:17",
          "abbreviation": "CIN",
          "links": [
            {},
            {}
          ],
          "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
          "isAllStar": false
        },
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/game/_/gameId/401815220/reds-cubs",
            "text": "Gamecast",
            "shortText": "Summary",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
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
      "401814906": {
        "id": "401814906",
        "team": {
          "id": "17",
          "uid": "s:1~l:10~t:17",
          "abbreviation": "CIN",
          "links": [
            {},
            {}
          ],
          "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
          "isAllStar": false
        },
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/game/_/gameId/401814906/angels-reds",
            "text": "Gamecast",
            "shortText": "Summary",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
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
      "401815148": {
        "id": "401815148",
        "team": {
          "id": "17",
          "uid": "s:1~l:10~t:17",
          "abbreviation": "CIN",
          "links": [
            {},
            {}
          ],
          "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
          "isAllStar": false
        },
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/game/_/gameId/401815148/rockies-reds",
            "text": "Gamecast",
            "shortText": "Summary",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
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
      "401814993": {
        "id": "401814993",
        "team": {
          "id": "17",
          "uid": "s:1~l:10~t:17",
          "abbreviation": "CIN",
          "links": [
            {},
            {}
          ],
          "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
          "isAllStar": false
        },
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/game/_/gameId/401814993/reds-twins",
            "text": "Gamecast",
            "shortText": "Summary",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
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
      }
    },
    "statistics": [
      {
        "displayName": "Pitching",
        "events": [
          {
            "eventId": "401815220",
            "stats": []
          },
          {
            "eventId": "401815148",
            "stats": []
          }
        ],
        "labels": [
          "GS",
          "CG"
        ],
        "names": [
          "gamesStarted",
          "completeGames"
        ],
        "displayNames": [
          "Games Started",
          "Complete Games"
        ]
      }
    ]
  },
  "rotowire": {
    "headline": "Abbott didn't factor into the decision Tuesday against the Cubs, allowing four hits and four walks with fou...",
    "published": "Wed May 06 06:28:22 PDT 2026"
  }
}
```
