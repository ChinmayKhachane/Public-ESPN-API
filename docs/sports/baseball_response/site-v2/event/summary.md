# MLB Game Summary

## https://site.api.espn.com/apis/site/v2/sports/baseball/mlb/summary?event=401815256

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "header": {
    "id": "401815256",
    "uid": "s:1~l:10~e:401815256",
    "competitions": [
      {
        "id": "401815256",
        "uid": "s:1~l:10~e:401815256~c:401815256",
        "competitors": [
          {
            "id": "17",
            "uid": "s:1~l:10~t:17",
            "team": {},
            "order": 0,
            "homeAway": "home",
            "winner": false,
            "score": "0",
            "linescores": [],
            "record": [],
            "probables": []
          },
          {
            "id": "18",
            "uid": "s:1~l:10~t:18",
            "team": {},
            "order": 1,
            "homeAway": "away",
            "winner": true,
            "score": "10",
            "linescores": [],
            "record": [],
            "probables": []
          }
        ],
        "status": {
          "type": {
            "id": "3",
            "name": "STATUS_FINAL",
            "state": "post",
            "completed": true,
            "detail": "Final",
            "shortDetail": "Final"
          },
          "featuredAthletes": [
            {},
            {}
          ],
          "periodPrefix": "End"
        },
        "date": "2026-05-08T22:10Z",
        "neutralSite": false,
        "conferenceCompetition": false,
        "boxscoreAvailable": true,
        "commentaryAvailable": false,
        "liveAvailable": false
      }
    ],
    "season": {
      "year": 2026,
      "current": true,
      "type": 2
    },
    "timeValid": true,
    "links": [
      {
        "rel": [
          "summary",
          "desktop"
        ],
        "href": "https://www.espn.com/mlb/game/_/gameId/401815256/astros-reds",
        "text": "Gamecast",
        "shortText": "Summary",
        "isExternal": false,
        "isPremium": false
      },
      {
        "rel": [
          "boxscore",
          "desktop"
        ],
        "href": "https://www.espn.com/mlb/boxscore/_/gameId/401815256",
        "text": "Box Score",
        "shortText": "Box Score",
        "isExternal": false,
        "isPremium": false
      }
    ],
    "week": 19,
    "league": {
      "id": "10",
      "uid": "s:1~l:10",
      "name": "Major League Baseball",
      "abbreviation": "MLB",
      "slug": "mlb",
      "midsizeName": "MLB",
      "isTournament": false,
      "links": [
        {
          "rel": [
            "index",
            "desktop"
          ],
          "href": "https://www.espn.com/mlb/",
          "text": "Index"
        },
        {
          "rel": [
            "index",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showClubhouse?uid=s:1~l:10",
          "text": "Index"
        }
      ],
      "logos": [
        {
          "href": "https://a.espncdn.com/i/teamlogos/leagues/500/mlb.png",
          "rel": [
            "full",
            "default"
          ]
        },
        {
          "href": "https://a.espncdn.com/combiner/i?img=/i/teamlogos/leagues/500-dark/mlb.png&w=500&h=500&transparent=true",
          "rel": [
            "full",
            "dark"
          ]
        }
      ]
    }
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
          "location": "Houston",
          "shortDisplayName": "Astros",
          "color": "002d62",
          "alternateColor": "eb6e1f"
        },
        "statistics": [
          {
            "name": "batting",
            "displayName": "Batting",
            "stats": []
          },
          {
            "name": "pitching",
            "displayName": "Pitching",
            "stats": []
          }
        ],
        "displayOrder": 1,
        "homeAway": "away",
        "details": [
          {
            "name": "battingDetails",
            "displayName": "Batting",
            "stats": []
          },
          {
            "name": "pitchingDetails",
            "displayName": "Pitching",
            "stats": []
          }
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
          "location": "Cincinnati",
          "shortDisplayName": "Reds",
          "color": "c6011f",
          "alternateColor": "ffffff"
        },
        "statistics": [
          {
            "name": "batting",
            "displayName": "Batting",
            "stats": []
          },
          {
            "name": "pitching",
            "displayName": "Pitching",
            "stats": []
          }
        ],
        "displayOrder": 2,
        "homeAway": "home",
        "details": [
          {
            "name": "battingDetails",
            "displayName": "Batting",
            "stats": []
          },
          {
            "name": "pitchingDetails",
            "displayName": "Pitching",
            "stats": []
          }
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
          "location": "Houston",
          "shortDisplayName": "Astros",
          "color": "002d62",
          "alternateColor": "eb6e1f"
        },
        "statistics": [
          {
            "athletes": [],
            "type": "batting",
            "names": [],
            "keys": [],
            "labels": [],
            "descriptions": [],
            "totals": []
          },
          {
            "athletes": [],
            "type": "pitching",
            "names": [],
            "keys": [],
            "labels": [],
            "descriptions": [],
            "totals": []
          }
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
          "location": "Cincinnati",
          "shortDisplayName": "Reds",
          "color": "c6011f",
          "alternateColor": "ffffff"
        },
        "statistics": [
          {
            "athletes": [],
            "type": "batting",
            "names": [],
            "keys": [],
            "labels": [],
            "descriptions": [],
            "totals": []
          },
          {
            "athletes": [],
            "type": "pitching",
            "names": [],
            "keys": [],
            "labels": [],
            "descriptions": [],
            "totals": []
          }
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
      "type": {
        "id": "59",
        "text": "Start Inning",
        "type": "start-inning"
      },
      "text": "Top of the 1st inning",
      "awayScore": 0,
      "homeScore": 0,
      "period": {
        "type": "Top",
        "number": 1,
        "displayValue": "1st Inning"
      },
      "scoringPlay": false,
      "scoreValue": 0
    },
    {
      "id": "4018152560001010001",
      "team": {
        "id": "18"
      },
      "sequenceNumber": "1",
      "type": {
        "id": "1",
        "text": "Start Batter/Pitcher",
        "alternativeText": "Now at bat",
        "type": "start-batterpitcher"
      },
      "text": "Nick Lodolo pitches to Jose Altuve",
      "awayScore": 0,
      "homeScore": 0,
      "period": {
        "type": "Top",
        "number": 1,
        "displayValue": "1st Inning"
      },
      "scoringPlay": false,
      "scoreValue": 0
    }
  ],
  "winprobability": [
    {
      "homeWinPercentage": 0.594,
      "tiePercentage": 0.0,
      "playId": "4018152560001990057"
    },
    {
      "homeWinPercentage": 0.564,
      "tiePercentage": 0.0,
      "playId": "4018152560002990057"
    }
  ],
  "notes": [],
  "format": {
    "regulation": {
      "displayName": "Inning",
      "slug": "inning",
      "periods": 9
    }
  },
  "gameInfo": {
    "venue": {
      "id": "83",
      "guid": "fc29fee9-3e88-343a-8f51-086d41354161",
      "fullName": "Great American Ball Park",
      "shortName": "Great American Ball Park",
      "address": {
        "city": "Cincinnati",
        "state": "Ohio",
        "zipCode": "45202"
      },
      "images": [
        {
          "href": "https://a.espncdn.com/i/venues/mlb/day/83.jpg",
          "width": 2000,
          "height": 1125,
          "alt": "",
          "rel": [
            "full",
            "day"
          ]
        }
      ]
    },
    "attendance": 24347,
    "officials": [
      {
        "displayName": "Jeremie Rehak",
        "position": {
          "id": "206",
          "name": "Home Plate Umpire",
          "displayName": "Home Plate Umpire"
        },
        "order": 1
      },
      {
        "displayName": "David Rackley",
        "position": {
          "id": "207",
          "name": "First Base Umpire",
          "displayName": "First Base Umpire"
        },
        "order": 2
      }
    ],
    "gameDuration": "2:32"
  },
  "seasonseries": [
    {
      "events": [
        {
          "id": "401815256",
          "uid": "s:1~l:10~e:401815256~c:401815256",
          "competitors": [
            {},
            {}
          ],
          "status": "post",
          "date": "2026-05-08T22:10:00Z",
          "timeValid": true,
          "statusType": {
            "id": "3",
            "name": "STATUS_FINAL",
            "state": "post",
            "completed": true,
            "detail": "Final",
            "shortDetail": "Final"
          },
          "neutralSite": false,
          "links": [
            {},
            {}
          ]
        },
        {
          "id": "401815271",
          "uid": "s:1~l:10~e:401815271~c:401815271",
          "competitors": [
            {},
            {}
          ],
          "status": "pre",
          "date": "2026-05-09T20:10:00Z",
          "timeValid": true,
          "statusType": {
            "id": "1",
            "name": "STATUS_SCHEDULED",
            "state": "pre",
            "completed": false,
            "detail": "Scheduled",
            "shortDetail": "5/9 - 4:10 PM EDT"
          },
          "neutralSite": false,
          "broadcasts": [
            {},
            {}
          ],
          "links": [
            {},
            {}
          ]
        }
      ],
      "type": "current",
      "title": "Current Series",
      "completed": false,
      "totalCompetitions": 3,
      "seriesScore": "1-0"
    },
    {
      "events": [
        {
          "id": "401815256",
          "uid": "s:1~l:10~e:401815256~c:401815256",
          "competitors": [
            {},
            {}
          ],
          "status": "post",
          "date": "2026-05-08T22:10:00Z",
          "timeValid": true,
          "statusType": {
            "id": "3",
            "name": "STATUS_FINAL",
            "state": "post",
            "completed": true,
            "detail": "Final",
            "shortDetail": "Final"
          },
          "neutralSite": false,
          "links": [
            {},
            {}
          ]
        },
        {
          "id": "401815271",
          "uid": "s:1~l:10~e:401815271~c:401815271",
          "competitors": [
            {},
            {}
          ],
          "status": "pre",
          "date": "2026-05-09T20:10:00Z",
          "timeValid": true,
          "statusType": {
            "id": "1",
            "name": "STATUS_SCHEDULED",
            "state": "pre",
            "completed": false,
            "detail": "Scheduled",
            "shortDetail": "5/9 - 4:10 PM EDT"
          },
          "neutralSite": false,
          "broadcasts": [
            {},
            {}
          ],
          "links": [
            {},
            {}
          ]
        }
      ],
      "type": "season",
      "title": "Regular Season Series",
      "completed": false,
      "totalCompetitions": 3,
      "seriesLabel": "Regular Season",
      "seriesScore": "1-0",
      "shortSummary": "HOU leads season"
    }
  ],
  "injuries": [
    {
      "team": {
        "id": "17",
        "uid": "s:1~l:10~t:17",
        "displayName": "Cincinnati Reds",
        "abbreviation": "CIN",
        "links": [
          {
            "href": "https://www.espn.com/mlb/team/_/name/cin/cincinnati-reds",
            "text": "Clubhouse"
          },
          {
            "href": "https://www.espn.com/mlb/team/schedule/_/name/cin",
            "text": "Schedule"
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
        "logos": [
          {
            "href": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-01T17:48Z"
          },
          {
            "href": "https://a.espncdn.com/i/teamlogos/mlb/500-dark/cin.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-01T17:49Z"
          }
        ]
      },
      "injuries": [
        {
          "status": "15-Day-IL",
          "date": "2026-05-08T11:02Z",
          "athlete": {
            "id": "41017",
            "uid": "s:1~l:10~a:41017",
            "guid": "e9899e53-ce21-c92e-a4ea-3d12347c0c90",
            "displayName": "Caleb Ferguson",
            "status": {},
            "lastName": "Ferguson",
            "fullName": "Caleb Ferguson",
            "shortName": "C. Ferguson",
            "links": [],
            "headshot": {}
          },
          "type": {
            "id": "10",
            "name": "INJURY_STATUS_15DAYIL",
            "abbreviation": "IL15"
          },
          "details": {
            "fantasyStatus": {},
            "type": "Oblique",
            "detail": "Strain",
            "side": "Right",
            "returnDate": "2026-05-08"
          }
        },
        {
          "status": "Day-To-Day",
          "date": "2026-05-08T11:02Z",
          "athlete": {
            "id": "36175",
            "uid": "s:1~l:10~a:36175",
            "guid": "d1c0c611-5713-413e-1bf7-5750d495cacf",
            "displayName": "Josh Staumont",
            "status": {},
            "lastName": "Staumont",
            "fullName": "Josh Staumont",
            "shortName": "J. Staumont",
            "links": [],
            "headshot": {}
          },
          "type": {
            "id": "6",
            "name": "INJURY_STATUS_DAYTODAY",
            "abbreviation": "DD"
          },
          "details": {
            "fantasyStatus": {},
            "type": "Undisclosed",
            "detail": "Not Specified",
            "side": "Not Specified",
            "returnDate": "2026-05-11"
          }
        }
      ]
    },
    {
      "team": {
        "id": "18",
        "uid": "s:1~l:10~t:18",
        "displayName": "Houston Astros",
        "abbreviation": "HOU",
        "links": [
          {
            "href": "https://www.espn.com/mlb/team/_/name/hou/houston-astros",
            "text": "Clubhouse"
          },
          {
            "href": "https://www.espn.com/mlb/team/schedule/_/name/hou",
            "text": "Schedule"
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/hou.png",
        "logos": [
          {
            "href": "https://a.espncdn.com/i/teamlogos/mlb/500/hou.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-01T17:48Z"
          },
          {
            "href": "https://a.espncdn.com/i/teamlogos/mlb/500-dark/hou.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-01T17:49Z"
          }
        ]
      },
      "injuries": [
        {
          "status": "Day-To-Day",
          "date": "2026-05-08T11:02Z",
          "athlete": {
            "id": "5273759",
            "uid": "s:1~l:10~a:5273759",
            "guid": "0a7a244a-a0f8-30bb-ac23-d9047e879ea0",
            "displayName": "Lucas Spence",
            "status": {},
            "lastName": "Spence",
            "fullName": "Lucas Spence",
            "shortName": "L. Spence",
            "links": [],
            "headshot": {}
          },
          "type": {
            "id": "6",
            "name": "INJURY_STATUS_DAYTODAY",
            "abbreviation": "DD"
          },
          "details": {
            "fantasyStatus": {},
            "type": "Undisclosed",
            "detail": "Not Specified",
            "side": "Not Specified",
            "returnDate": "2026-05-10"
          }
        },
        {
          "status": "15-Day-IL",
          "date": "2026-05-08T11:02Z",
          "athlete": {
            "id": "5330833",
            "uid": "s:1~l:10~a:5330833",
            "guid": "6f46a916-482d-3d3d-b32c-aee2790c71bb",
            "displayName": "Tatsuya Imai",
            "status": {},
            "lastName": "Imai",
            "fullName": "Tatsuya Imai",
            "shortName": "T. Imai",
            "links": [],
            "headshot": {}
          },
          "type": {
            "id": "10",
            "name": "INJURY_STATUS_15DAYIL",
            "abbreviation": "IL15"
          },
          "details": {
            "fantasyStatus": {},
            "type": "Arm",
            "detail": "Not Specified",
            "side": "Right",
            "returnDate": "2026-05-11"
          }
        }
      ]
    }
  ],
  "broadcasts": [
    {
      "type": {
        "id": "4",
        "slug": "streaming",
        "shortName": "Streaming",
        "longName": "Streaming"
      },
      "station": "MLB.TV",
      "market": {
        "id": "1",
        "type": "National"
      },
      "media": {
        "name": "MLB.TV",
        "callLetters": "MLB.TV",
        "shortName": "MLB.TV"
      },
      "lang": "en",
      "region": "us",
      "isNational": true
    },
    {
      "type": {
        "id": "4",
        "slug": "streaming",
        "shortName": "Streaming",
        "longName": "Streaming"
      },
      "station": "Reds.TV",
      "market": {
        "id": "2",
        "type": "Home"
      },
      "media": {
        "name": "Reds.TV",
        "callLetters": "Reds.TV",
        "shortName": "Reds.TV"
      },
      "lang": "en",
      "region": "us",
      "isNational": false
    }
  ],
  "pickcenter": [
    {
      "header": {
        "logo": {
          "dark": "https://a.espncdn.com/i/espnbet/dark/espn-bet-square-off.svg",
          "light": "https://a.espncdn.com/i/espnbet/espn-bet-square-off.svg",
          "exclusivesLogoDark": "https://a.espncdn.com/i/espnbet/espn-bet-square-mint.svg",
          "exclusivesLogoLight": "https://a.espncdn.com/i/espnbet/espn-bet-square-mint.svg"
        },
        "text": "Game Odds"
      },
      "provider": {
        "id": "100",
        "name": "DraftKings",
        "priority": 1,
        "logos": [
          {
            "href": "https://a.espncdn.com/i/betting/Draftkings_Light.svg",
            "rel": []
          },
          {
            "href": "https://a.espncdn.com/i/betting/Draftkings_Dark.svg",
            "rel": []
          }
        ]
      },
      "details": "CIN -132",
      "overUnder": 9.0,
      "spread": -1.5,
      "overOdds": -108.0,
      "underOdds": -112.0,
      "awayTeamOdds": {
        "team": {
          "$ref": "http://sports.core.api.espn.pvt/v2/sports/baseball/leagues/mlb/seasons/2026/teams/18?lang=en&region=us"
        },
        "favorite": false,
        "underdog": true,
        "moneyLine": 109,
        "teamId": "18",
        "favoriteAtOpen": false
      },
      "homeTeamOdds": {
        "team": {
          "$ref": "http://sports.core.api.espn.pvt/v2/sports/baseball/leagues/mlb/seasons/2026/teams/17?lang=en&region=us"
        },
        "favorite": true,
        "underdog": false,
        "moneyLine": -132,
        "teamId": "17",
        "favoriteAtOpen": true
      },
      "links": [
        {
          "language": "en-US",
          "rel": [
            "home",
            "desktop"
          ],
          "href": "https://sportsbook.draftkings.com/gateway?s=__s__&wpcid=__wpcid__&wpsrc=413&wpcn=ESPN&wpscn=Widget&wpcrn=Be...",
          "text": "Home Bet",
          "shortText": "Home Bet",
          "isExternal": true,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "away",
            "desktop"
          ],
          "href": "https://sportsbook.draftkings.com/gateway?s=__s__&wpcid=__wpcid__&wpsrc=413&wpcn=ESPN&wpscn=Widget&wpcrn=Be...",
          "text": "Away Bet",
          "shortText": "Away Bet",
          "isExternal": true,
          "isPremium": false
        }
      ]
    }
  ],
  "odds": [],
  "againstTheSpread": [
    {
      "team": {
        "id": "18",
        "uid": "s:1~l:10~t:18",
        "displayName": "Houston Astros",
        "abbreviation": "HOU",
        "links": [
          {
            "href": "https://www.espn.com/mlb/team/_/name/hou/houston-astros",
            "text": "Clubhouse"
          },
          {
            "href": "https://www.espn.com/mlb/team/schedule/_/name/hou",
            "text": "Schedule"
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/hou.png",
        "logos": [
          {
            "href": "https://a.espncdn.com/i/teamlogos/mlb/500/hou.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-01T17:48Z"
          },
          {
            "href": "https://a.espncdn.com/i/teamlogos/mlb/500-dark/hou.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-01T17:49Z"
          }
        ]
      },
      "records": []
    },
    {
      "team": {
        "id": "17",
        "uid": "s:1~l:10~t:17",
        "displayName": "Cincinnati Reds",
        "abbreviation": "CIN",
        "links": [
          {
            "href": "https://www.espn.com/mlb/team/_/name/cin/cincinnati-reds",
            "text": "Clubhouse"
          },
          {
            "href": "https://www.espn.com/mlb/team/schedule/_/name/cin",
            "text": "Schedule"
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
        "logos": [
          {
            "href": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-01T17:48Z"
          },
          {
            "href": "https://a.espncdn.com/i/teamlogos/mlb/500-dark/cin.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-01T17:49Z"
          }
        ]
      },
      "records": []
    }
  ],
  "news": {
    "header": "MLB News",
    "link": {
      "language": "en",
      "rel": [
        "index",
        "desktop"
      ],
      "href": "https://www.espn.com/mlb/",
      "text": "All MLB News",
      "shortText": "All News",
      "isExternal": false,
      "isPremium": false
    },
    "articles": [
      {
        "id": 48717767,
        "categories": [
          {
            "id": 80311,
            "uid": "s:1~l:10",
            "guid": "b38f959b-7865-31ac-8841-b88355519e10",
            "type": "league",
            "sportId": 10,
            "leagueId": 10,
            "league": {}
          },
          {
            "id": 724,
            "uid": "s:1~l:10~t:25",
            "guid": "4dec648c-3eb9-055c-aebc-2711f30975a0",
            "team": {},
            "type": "team",
            "sportId": 10,
            "teamId": 25
          }
        ],
        "nowId": "1-48717767",
        "contentKey": "48717767-1-5-1",
        "dataSourceIdentifier": "f60e80c46aee3",
        "type": "HeadlineNews",
        "headline": "Padres prospect pleads guilty to charge of transporting noncitizen immigrants",
        "lastModified": "2026-05-09T03:43:19Z",
        "published": "2026-05-09T03:43:19Z",
        "images": [
          {
            "id": 13368714,
            "name": "Baseball seams 150802 [600x400]",
            "dataSourceIdentifier": "886e0be6d23ad",
            "type": "header",
            "alt": "Baseball seams",
            "credit": "Pouya Dianat/Getty Images",
            "height": 400,
            "width": 600,
            "url": "https://a.espncdn.com/photo/2015/0802/mlb_baseball_b1_600x400.jpg"
          }
        ]
      },
      {
        "id": 48717747,
        "categories": [
          {
            "id": 525402,
            "uid": "s:1~l:10~a:4142424",
            "guid": "c8f56866-3f5a-304f-9961-10d6b8970628",
            "type": "athlete",
            "sportId": 10,
            "athleteId": 4142424,
            "athlete": {}
          },
          {
            "id": 468339,
            "uid": "s:1~l:10~a:4717833",
            "guid": "dcbeba5a-3fe9-3bf0-9f19-c077ef2a6f78",
            "type": "athlete",
            "sportId": 10,
            "athleteId": 4717833,
            "athlete": {}
          }
        ],
        "nowId": "1-48717747",
        "contentKey": "48717747-1-293-1",
        "dataSourceIdentifier": "c9a50ebd56d4c",
        "type": "Media",
        "headline": "Chicago Cubs vs. Texas Rangers: Game Highlights",
        "lastModified": "2026-05-09T03:30:14Z",
        "published": "2026-05-09T03:30:14Z",
        "images": [
          {
            "name": "Chicago Cubs vs. Texas Rangers: Game Highlights",
            "alt": "",
            "height": 324,
            "width": 576,
            "url": "https://a.espncdn.com/media/motion/wsc/2026/0509/e4a1a620-c370-4340-9b36-3c6550c5edc9/e4a1a620-c370-4340-9b..."
          }
        ]
      }
    ]
  },
  "rosters": [
    {
      "team": {
        "id": "17",
        "displayName": "Cincinnati Reds",
        "abbreviation": "CIN",
        "color": "c6011f",
        "alternateColor": "ffffff",
        "logos": [
          {
            "href": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-01T17:48Z"
          },
          {
            "href": "https://a.espncdn.com/i/teamlogos/mlb/500-dark/cin.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-01T17:49Z"
          }
        ]
      },
      "homeAway": "home",
      "winner": false,
      "roster": [
        {
          "active": true,
          "starter": true,
          "athlete": {
            "id": "36020",
            "uid": "s:1~l:10~a:36020",
            "guid": "9c422f07-8499-da29-dadb-539a4824ad17",
            "displayName": "TJ Friedl",
            "lastName": "Friedl",
            "fullName": "TJ Friedl",
            "shortName": "T. Friedl",
            "links": [],
            "headshot": {},
            "positions": []
          },
          "position": {
            "name": "Center Field",
            "displayName": "Center Fielder",
            "abbreviation": "CF"
          },
          "batOrder": 1,
          "subbedIn": false,
          "subbedOut": false,
          "media": {},
          "stats": [
            {},
            {}
          ],
          "jersey": "29"
        },
        {
          "active": true,
          "starter": true,
          "athlete": {
            "id": "42410",
            "uid": "s:1~l:10~a:42410",
            "guid": "04ad41a4-5394-355c-92cd-dcc55a79e720",
            "displayName": "JJ Bleday",
            "lastName": "Bleday",
            "fullName": "JJ Bleday",
            "shortName": "J. Bleday",
            "links": [],
            "headshot": {},
            "positions": []
          },
          "position": {
            "name": "Left Field",
            "displayName": "Left Fielder",
            "abbreviation": "LF"
          },
          "batOrder": 2,
          "subbedIn": false,
          "subbedOut": false,
          "media": {},
          "stats": [
            {},
            {}
          ],
          "jersey": "22"
        }
      ]
    },
    {
      "team": {
        "id": "18",
        "displayName": "Houston Astros",
        "abbreviation": "HOU",
        "color": "002d62",
        "alternateColor": "eb6e1f",
        "logos": [
          {
            "href": "https://a.espncdn.com/i/teamlogos/mlb/500/hou.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-01T17:48Z"
          },
          {
            "href": "https://a.espncdn.com/i/teamlogos/mlb/500-dark/hou.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-01T17:49Z"
          }
        ]
      },
      "homeAway": "away",
      "winner": true,
      "roster": [
        {
          "active": false,
          "starter": true,
          "athlete": {
            "id": "31662",
            "uid": "s:1~l:10~a:31662",
            "guid": "3e8ded2a-d813-324d-5925-28149ef72a30",
            "displayName": "Jose Altuve",
            "lastName": "Altuve",
            "fullName": "Jose Altuve",
            "shortName": "J. Altuve",
            "links": [],
            "headshot": {},
            "positions": []
          },
          "position": {
            "name": "Second Base",
            "displayName": "Second Baseman",
            "abbreviation": "2B"
          },
          "batOrder": 1,
          "subbedIn": false,
          "subbedOut": false,
          "media": {},
          "stats": [
            {},
            {}
          ],
          "jersey": "27"
        },
        {
          "active": true,
          "starter": false,
          "athlete": {
            "id": "42408",
            "uid": "s:1~l:10~a:42408",
            "guid": "7d4729e6-6a75-3635-8f9c-a5f76b3c72a1",
            "displayName": "Braden Shewmake",
            "lastName": "Shewmake",
            "fullName": "Braden Shewmake",
            "shortName": "B. Shewmake",
            "links": [],
            "headshot": {},
            "positions": []
          },
          "position": {
            "name": "Second Base",
            "displayName": "Second Baseman",
            "abbreviation": "2B"
          },
          "positions": [
            {},
            {}
          ],
          "notes": [
            {}
          ],
          "batOrder": 1,
          "subbedIn": false,
          "subbedOut": false,
          "media": {}
        }
      ]
    }
  ],
  "videos": [
    {
      "id": 48716451,
      "cerebroId": "69fe8549f795f455a4a6d052",
      "source": "espn",
      "headline": "Houston Astros vs. Cincinnati Reds: Game Highlights",
      "lastModified": "2026-05-09T00:58:22Z",
      "originalPublishDate": "2026-05-09T00:52:25Z",
      "duration": 88,
      "timeRestrictions": {
        "embargoDate": "2026-05-09T00:52:17Z",
        "expirationDate": "2026-05-11T00:52:17Z"
      },
      "deviceRestrictions": {
        "type": "whitelist",
        "devices": [
          "settop",
          "tablet"
        ]
      },
      "geoRestrictions": {
        "type": "whitelist",
        "countries": [
          "PR",
          "HN"
        ]
      }
    },
    {
      "id": 48715382,
      "cerebroId": "69fe6571c35ca554fbcb52d9",
      "source": "espn",
      "headline": "Zach Dezenzo hits a home run for the Astros",
      "lastModified": "2026-05-09T00:47:34Z",
      "originalPublishDate": "2026-05-08T22:36:34Z",
      "duration": 37,
      "timeRestrictions": {
        "embargoDate": "2026-05-08T22:36:26Z",
        "expirationDate": "2026-05-10T04:00:00Z"
      },
      "deviceRestrictions": {
        "type": "whitelist",
        "devices": [
          "settop",
          "tablet"
        ]
      },
      "geoRestrictions": {
        "type": "whitelist",
        "countries": [
          "PR",
          "HN"
        ]
      }
    }
  ],
  "playsMap": {
    "4018152560000000059": {
      "$ref": "#/plays/0"
    },
    "4018152560001010001": {
      "$ref": "#/plays/1"
    },
    "4018152560001020024": {
      "$ref": "#/plays/2"
    },
    "4018152560001990057": {
      "$ref": "#/plays/3"
    },
    "4018152560001990099": {
      "$ref": "#/plays/4"
    },
    "4018152560002010001": {
      "$ref": "#/plays/5"
    },
    "4018152560002020036": {
      "$ref": "#/plays/6"
    },
    "4018152560002030036": {
      "$ref": "#/plays/7"
    },
    "4018152560002040002": {
      "$ref": "#/plays/8"
    },
    "4018152560002990057": {
      "$ref": "#/plays/9"
    }
  },
  "atBats": {
    "4018152560001": [
      {
        "$ref": "#/plays/0"
      },
      {
        "$ref": "#/plays/1"
      }
    ],
    "4018152560002": [
      {
        "$ref": "#/plays/5"
      },
      {
        "$ref": "#/plays/6"
      }
    ],
    "4018152560003": [
      {
        "$ref": "#/plays/11"
      },
      {
        "$ref": "#/plays/12"
      }
    ],
    "4018152560101": [
      {
        "$ref": "#/plays/22"
      },
      {
        "$ref": "#/plays/23"
      }
    ],
    "4018152560102": [
      {
        "$ref": "#/plays/31"
      },
      {
        "$ref": "#/plays/32"
      }
    ],
    "4018152560103": [
      {
        "$ref": "#/plays/37"
      },
      {
        "$ref": "#/plays/38"
      }
    ],
    "4018152560104": [
      {
        "$ref": "#/plays/46"
      },
      {
        "$ref": "#/plays/47"
      }
    ],
    "4018152560201": [
      {
        "$ref": "#/plays/52"
      },
      {
        "$ref": "#/plays/53"
      }
    ],
    "4018152560202": [
      {
        "$ref": "#/plays/58"
      },
      {
        "$ref": "#/plays/59"
      }
    ],
    "4018152560203": [
      {
        "$ref": "#/plays/63"
      },
      {
        "$ref": "#/plays/64"
      }
    ]
  }
}
```
