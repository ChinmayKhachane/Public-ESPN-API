# Game Summary

## https://site.api.espn.com/apis/site/v2/sports/hockey/{league}/summary?event={event}

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/summary?event=401871412`

## Example Response

```json
{
  "boxscore": {
    "teams": [
      {
        "team": {
          "id": "7",
          "uid": "s:70~l:90~t:7",
          "slug": "carolina-hurricanes",
          "location": "Carolina",
          "name": "Hurricanes",
          "abbreviation": "CAR",
          "displayName": "Carolina Hurricanes",
          "shortDisplayName": "Hurricanes",
          "color": "e30426",
          "alternateColor": "000000",
          "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/car.png"
        },
        "statistics": [
          {
            "name": "blockedShots",
            "abbreviation": "BS",
            "displayValue": "9",
            "label": "Blocked Shots"
          },
          {
            "name": "hits",
            "abbreviation": "HT",
            "displayValue": "31",
            "label": "Hits"
          }
        ],
        "displayOrder": 1,
        "homeAway": "away"
      },
      {
        "team": {
          "id": "15",
          "uid": "s:70~l:90~t:15",
          "slug": "philadelphia-flyers",
          "location": "Philadelphia",
          "name": "Flyers",
          "abbreviation": "PHI",
          "displayName": "Philadelphia Flyers",
          "shortDisplayName": "Flyers",
          "color": "fe5823",
          "alternateColor": "000000",
          "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png"
        },
        "statistics": [
          {
            "name": "blockedShots",
            "abbreviation": "BS",
            "displayValue": "17",
            "label": "Blocked Shots"
          },
          {
            "name": "hits",
            "abbreviation": "HT",
            "displayValue": "41",
            "label": "Hits"
          }
        ],
        "displayOrder": 2,
        "homeAway": "home"
      }
    ],
    "players": [
      {
        "team": {
          "id": "7",
          "uid": "s:70~l:90~t:7",
          "slug": "carolina-hurricanes",
          "location": "Carolina",
          "name": "Hurricanes",
          "abbreviation": "CAR",
          "displayName": "Carolina Hurricanes",
          "shortDisplayName": "Hurricanes",
          "color": "e30426",
          "alternateColor": "000000",
          "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/car.png"
        },
        "statistics": [
          {
            "name": "forwards",
            "keys": [],
            "labels": [],
            "descriptions": [],
            "athletes": []
          },
          {
            "name": "defenses",
            "keys": [],
            "labels": [],
            "descriptions": [],
            "athletes": []
          }
        ],
        "displayOrder": 1
      },
      {
        "team": {
          "id": "15",
          "uid": "s:70~l:90~t:15",
          "slug": "philadelphia-flyers",
          "location": "Philadelphia",
          "name": "Flyers",
          "abbreviation": "PHI",
          "displayName": "Philadelphia Flyers",
          "shortDisplayName": "Flyers",
          "color": "fe5823",
          "alternateColor": "000000",
          "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png"
        },
        "statistics": [
          {
            "name": "forwards",
            "keys": [],
            "labels": [],
            "descriptions": [],
            "athletes": []
          },
          {
            "name": "defenses",
            "keys": [],
            "labels": [],
            "descriptions": [],
            "athletes": []
          }
        ],
        "displayOrder": 2
      }
    ]
  },
  "format": {
    "regulation": {
      "periods": 3,
      "displayName": "Period",
      "slug": "period",
      "clock": 1200.0
    }
  },
  "gameInfo": {
    "venue": {
      "id": "1845",
      "fullName": "Xfinity Mobile Arena",
      "address": {
        "city": "Philadelphia",
        "state": "PA",
        "country": "USA"
      },
      "grass": false,
      "images": [
        {
          "href": "https://a.espncdn.com/i/venues/nhl/day/1845.jpg",
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
    "attendance": 19970,
    "officials": [
      {
        "fullName": "Francis Charron",
        "displayName": "Francis Charron",
        "position": {
          "name": "Referee",
          "displayName": "Referee",
          "id": "24"
        },
        "order": 0
      },
      {
        "fullName": "Kyle Rehman",
        "displayName": "Kyle Rehman",
        "position": {
          "name": "Referee",
          "displayName": "Referee",
          "id": "24"
        },
        "order": 0
      }
    ]
  },
  "leaders": [
    {
      "team": {
        "id": "15",
        "uid": "s:70~l:90~t:15",
        "displayName": "Philadelphia Flyers",
        "abbreviation": "PHI",
        "links": [
          {
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse"
          },
          {
            "href": "https://www.espn.com/nhl/team/schedule/_/name/phi",
            "text": "Schedule"
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
        "logos": [
          {
            "href": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-02T15:51Z"
          },
          {
            "href": "https://a.espncdn.com/i/teamlogos/nhl/500-dark/phi.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-02T15:46Z"
          }
        ]
      },
      "leaders": [
        {
          "name": "goals",
          "displayName": "Goals",
          "leaders": [
            {}
          ]
        },
        {
          "name": "assists",
          "displayName": "Assists",
          "leaders": [
            {}
          ]
        }
      ]
    },
    {
      "team": {
        "id": "7",
        "uid": "s:70~l:90~t:7",
        "displayName": "Carolina Hurricanes",
        "abbreviation": "CAR",
        "links": [
          {
            "href": "https://www.espn.com/nhl/team/_/name/car/carolina-hurricanes",
            "text": "Clubhouse"
          },
          {
            "href": "https://www.espn.com/nhl/team/schedule/_/name/car",
            "text": "Schedule"
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/car.png",
        "logos": [
          {
            "href": "https://a.espncdn.com/i/teamlogos/nhl/500/car.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-02T15:51Z"
          },
          {
            "href": "https://a.espncdn.com/i/teamlogos/nhl/500-dark/car.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-02T15:46Z"
          }
        ]
      },
      "leaders": [
        {
          "name": "goals",
          "displayName": "Goals",
          "leaders": [
            {}
          ]
        },
        {
          "name": "assists",
          "displayName": "Assists",
          "leaders": [
            {}
          ]
        }
      ]
    }
  ],
  "seasonseries": [
    {
      "type": "season",
      "title": "Regular Season Series",
      "description": "Regular Season Series",
      "summary": "CAR wins series 3-1",
      "completed": true,
      "totalCompetitions": 4,
      "seriesLabel": "Regular Season",
      "seriesScore": "3-1",
      "shortSummary": "CAR wins season",
      "events": [
        {
          "id": "401802381",
          "uid": "s:70~l:90~e:401802381~c:401802381",
          "date": "2025-10-11T23:00:00Z",
          "timeValid": true,
          "status": "post",
          "statusType": {
            "id": "3",
            "name": "STATUS_FINAL",
            "state": "post",
            "completed": true,
            "description": "Final",
            "detail": "Final/OT",
            "shortDetail": "Final/OT",
            "altDetail": "OT"
          },
          "neutralSite": false,
          "competitors": [
            {},
            {}
          ],
          "links": [
            {},
            {}
          ]
        },
        {
          "id": "401802852",
          "uid": "s:70~l:90~e:401802852~c:401802852",
          "date": "2025-12-14T00:00:00Z",
          "timeValid": true,
          "status": "post",
          "statusType": {
            "id": "3",
            "name": "STATUS_FINAL",
            "state": "post",
            "completed": true,
            "description": "Final",
            "detail": "Final/SO",
            "shortDetail": "Final/SO",
            "altDetail": "SO"
          },
          "neutralSite": false,
          "competitors": [
            {},
            {}
          ],
          "links": [
            {},
            {}
          ]
        }
      ]
    },
    {
      "type": "playoff",
      "title": "Playoff Series",
      "description": "Playoff Series",
      "summary": "CAR leads series 3-0",
      "completed": false,
      "totalCompetitions": 7,
      "seriesLabel": "Playoffs",
      "seriesScore": "3-0",
      "shortSummary": "CAR leads series",
      "round": "East 2nd Round",
      "events": [
        {
          "id": "401871166",
          "uid": "s:70~l:90~e:401871166~c:401871166",
          "date": "2026-05-03T00:00:00Z",
          "timeValid": true,
          "status": "post",
          "statusType": {
            "id": "3",
            "name": "STATUS_FINAL",
            "state": "post",
            "completed": true,
            "description": "Final",
            "detail": "Final",
            "shortDetail": "Final"
          },
          "neutralSite": false,
          "competitors": [
            {},
            {}
          ],
          "links": [
            {},
            {}
          ]
        },
        {
          "id": "401871411",
          "uid": "s:70~l:90~e:401871411~c:401871411",
          "date": "2026-05-04T23:00:00Z",
          "timeValid": true,
          "status": "post",
          "statusType": {
            "id": "3",
            "name": "STATUS_FINAL",
            "state": "post",
            "completed": true,
            "description": "Final",
            "detail": "Final/OT",
            "shortDetail": "Final/OT",
            "altDetail": "OT"
          },
          "neutralSite": false,
          "competitors": [
            {},
            {}
          ],
          "links": [
            {},
            {}
          ]
        }
      ]
    }
  ],
  "injuries": [
    {
      "team": {
        "id": "15",
        "uid": "s:70~l:90~t:15",
        "displayName": "Philadelphia Flyers",
        "abbreviation": "PHI",
        "links": [
          {
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse"
          },
          {
            "href": "https://www.espn.com/nhl/team/schedule/_/name/phi",
            "text": "Schedule"
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
        "logos": [
          {
            "href": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-02T15:51Z"
          },
          {
            "href": "https://a.espncdn.com/i/teamlogos/nhl/500-dark/phi.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-02T15:46Z"
          }
        ]
      },
      "injuries": [
        {
          "status": "Out",
          "date": "2026-05-09T14:09Z",
          "athlete": {
            "id": "4392072",
            "uid": "s:70~l:90~a:4392072",
            "guid": "9b4c71a7-b8e9-94f4-662b-89af200a6017",
            "lastName": "Tippett",
            "fullName": "Owen Tippett",
            "displayName": "Owen Tippett",
            "shortName": "O. Tippett",
            "links": [],
            "headshot": {},
            "jersey": "74",
            "position": {},
            "status": {}
          },
          "type": {
            "id": "4",
            "name": "INJURY_STATUS_OUT",
            "description": "out",
            "abbreviation": "O"
          },
          "details": {
            "fantasyStatus": {},
            "type": "Undisclosed",
            "returnDate": "2026-05-11"
          }
        },
        {
          "status": "Out",
          "date": "2026-05-06T16:51Z",
          "athlete": {
            "id": "4419682",
            "uid": "s:70~l:90~a:4419682",
            "guid": "4353bdca-1851-39a3-8193-675e25d4eda2",
            "lastName": "Cates",
            "fullName": "Noah Cates",
            "displayName": "Noah Cates",
            "shortName": "N. Cates",
            "links": [],
            "headshot": {},
            "jersey": "27",
            "position": {},
            "status": {}
          },
          "type": {
            "id": "4",
            "name": "INJURY_STATUS_OUT",
            "description": "out",
            "abbreviation": "O"
          },
          "details": {
            "fantasyStatus": {},
            "type": "Lower Body",
            "returnDate": "2026-05-18"
          }
        }
      ]
    },
    {
      "team": {
        "id": "7",
        "uid": "s:70~l:90~t:7",
        "displayName": "Carolina Hurricanes",
        "abbreviation": "CAR",
        "links": [
          {
            "href": "https://www.espn.com/nhl/team/_/name/car/carolina-hurricanes",
            "text": "Clubhouse"
          },
          {
            "href": "https://www.espn.com/nhl/team/schedule/_/name/car",
            "text": "Schedule"
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/car.png",
        "logos": [
          {
            "href": "https://a.espncdn.com/i/teamlogos/nhl/500/car.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-02T15:51Z"
          },
          {
            "href": "https://a.espncdn.com/i/teamlogos/nhl/500-dark/car.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-02T15:46Z"
          }
        ]
      },
      "injuries": []
    }
  ],
  "broadcasts": [
    {
      "type": {
        "id": "4",
        "shortName": "Streaming",
        "longName": "Streaming",
        "slug": "streaming"
      },
      "station": "HBO Max",
      "market": {
        "id": "1",
        "type": "National"
      },
      "media": {
        "callLetters": "HBO Max",
        "name": "HBO Max",
        "shortName": "HBO Max"
      },
      "lang": "en",
      "region": "us",
      "isNational": true
    }
  ],
  "pickcenter": [
    {
      "provider": {
        "id": "100",
        "name": "Draft Kings",
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
      "details": "CAR -166",
      "overUnder": 5.5,
      "spread": 1.5,
      "overOdds": 105.0,
      "underOdds": -125.0,
      "awayTeamOdds": {
        "favorite": true,
        "underdog": false,
        "moneyLine": -166,
        "spreadOdds": 154.0,
        "team": {
          "$ref": "http://sports.core.api.espn.pvt/v2/sports/hockey/leagues/nhl/seasons/2026/teams/7?lang=en&region=us"
        },
        "teamId": "7",
        "favoriteAtOpen": true
      },
      "homeTeamOdds": {
        "favorite": false,
        "underdog": true,
        "moneyLine": 140,
        "spreadOdds": -185.0,
        "team": {
          "$ref": "http://sports.core.api.espn.pvt/v2/sports/hockey/leagues/nhl/seasons/2026/teams/15?lang=en&region=us"
        },
        "teamId": "15",
        "favoriteAtOpen": false
      },
      "links": [
        {
          "language": "en-US",
          "rel": [
            "home",
            "desktop"
          ],
          "href": "https://sportsbook.draftkings.com/gateway?s=__s__&wpcid=__wpcid__&wpsrc=413&wpcn=ESPN&wpscn=Widget&wpcrn=BetSlipDeepLink&wpscid=__wpscid__&wpcrid=xx&preurl=https%3A%2F%2Fsportsbook.draftkings.com%2Fevent%2F34104080%3Foutcomes%3D0ML84653447_1",
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
          "href": "https://sportsbook.draftkings.com/gateway?s=__s__&wpcid=__wpcid__&wpsrc=413&wpcn=ESPN&wpscn=Widget&wpcrn=BetSlipDeepLink&wpscid=__wpscid__&wpcrid=xx&preurl=https%3A%2F%2Fsportsbook.draftkings.com%2Fevent%2F34104080%3Foutcomes%3D0ML84653447_3",
          "text": "Away Bet",
          "shortText": "Away Bet",
          "isExternal": true,
          "isPremium": false
        }
      ],
      "moneyline": {
        "displayName": "Moneyline",
        "shortDisplayName": "ML",
        "home": {
          "close": {
            "odds": "+140"
          },
          "open": {
            "odds": "+130"
          }
        },
        "away": {
          "close": {
            "odds": "-166"
          },
          "open": {
            "odds": "-155"
          }
        }
      },
      "pointSpread": {
        "displayName": "Spread",
        "shortDisplayName": "Spread",
        "home": {
          "close": {
            "line": "+1.5",
            "odds": "-185"
          },
          "open": {
            "line": "+1.5",
            "odds": "-192"
          }
        },
        "away": {
          "close": {
            "line": "-1.5",
            "odds": "+154"
          },
          "open": {
            "line": "-1.5",
            "odds": "+160"
          }
        }
      },
      "total": {
        "displayName": "Total",
        "shortDisplayName": "Total",
        "over": {
          "close": {
            "line": "o5.5",
            "odds": "+105"
          },
          "open": {
            "line": "o5.5",
            "odds": "+100"
          }
        },
        "under": {
          "close": {
            "line": "u5.5",
            "odds": "-125"
          },
          "open": {
            "line": "u5.5",
            "odds": "-120"
          }
        }
      },
      "link": {
        "text": "See More Odds",
        "language": "en-US",
        "rel": [
          "main",
          "desktop"
        ],
        "href": "https://www.draftkings.com",
        "isExternal": true,
        "isPremium": false
      },
      "header": {
        "logo": {
          "dark": "https://a.espncdn.com/i/espnbet/dark/espn-bet-square-off.svg",
          "light": "https://a.espncdn.com/i/espnbet/espn-bet-square-off.svg",
          "exclusivesLogoDark": "https://a.espncdn.com/i/espnbet/espn-bet-square-mint.svg",
          "exclusivesLogoLight": "https://a.espncdn.com/i/espnbet/espn-bet-square-mint.svg"
        },
        "text": "Game Odds"
      },
      "footer": {
        "disclaimer": "Odds by DraftKings\nGAMBLING PROBLEM? CALL 1-800-GAMBLER, (800) 327-5050 or visit gamblinghelplinema.org (MA). Call 877-8-HOPENY/text HOPENY (467369) (NY).\nPlease Gamble Responsibly. 888-789-7777/visit ccpg.org (CT), or visit www.mdgamblinghelp.org (MD).\n21+ and present in most states. (18+ DC/KY/NH/WY). Void in ONT/OR/NH. Eligibility restrictions apply. On behalf of Boot Hill Casino & Resort (KS). Terms: sportsbook.draftkings.com/promos."
      }
    }
  ],
  "odds": [],
  "againstTheSpread": [
    {
      "team": {
        "id": "7",
        "uid": "s:70~l:90~t:7",
        "displayName": "Carolina Hurricanes",
        "abbreviation": "CAR",
        "links": [
          {
            "href": "https://www.espn.com/nhl/team/_/name/car/carolina-hurricanes",
            "text": "Clubhouse"
          },
          {
            "href": "https://www.espn.com/nhl/team/schedule/_/name/car",
            "text": "Schedule"
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/car.png",
        "logos": [
          {
            "href": "https://a.espncdn.com/i/teamlogos/nhl/500/car.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-02T15:51Z"
          },
          {
            "href": "https://a.espncdn.com/i/teamlogos/nhl/500-dark/car.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-02T15:46Z"
          }
        ]
      },
      "records": []
    },
    {
      "team": {
        "id": "15",
        "uid": "s:70~l:90~t:15",
        "displayName": "Philadelphia Flyers",
        "abbreviation": "PHI",
        "links": [
          {
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse"
          },
          {
            "href": "https://www.espn.com/nhl/team/schedule/_/name/phi",
            "text": "Schedule"
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
        "logos": [
          {
            "href": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-02T15:51Z"
          },
          {
            "href": "https://a.espncdn.com/i/teamlogos/nhl/500-dark/phi.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-02T15:46Z"
          }
        ]
      },
      "records": []
    }
  ],
  "news": {
    "header": "NHL News",
    "link": {
      "language": "en",
      "rel": [
        "index",
        "desktop"
      ],
      "href": "https://www.espn.com/nhl/",
      "text": "All NHL News",
      "shortText": "All News",
      "isExternal": false,
      "isPremium": false
    },
    "articles": [
      {
        "id": 44686270,
        "nowId": "1-44686270",
        "contentKey": "44686270-1-6-1",
        "dataSourceIdentifier": "25ce825f3281f",
        "type": "Story",
        "headline": "Stanley Cup odds: Eastern Conference favorite Hurricanes look to sweep Flyers",
        "description": "The Avalanche and Hurricanes are the top two favorites to win the 2026 Stanley Cup.",
        "lastModified": "2026-05-09T16:18:40Z",
        "published": "2026-05-09T16:18:40Z",
        "images": [
          {
            "dataSourceIdentifier": "36f5d7465189e",
            "id": 48666250,
            "type": "header",
            "name": "Carolina Hurricanes [1296x729]",
            "caption": "The Carolina Hurricanes are the current favorites to win the Eastern Conference.",
            "alt": "Mark Jankowski #77 of the Carolina Hurricanes faces off against Luke Glendening #41 of the Philadelphia Flyers in Game One of the Second Round of the 2026 Stanley Cup Playoffs at Lenovo Center on May 02, 2026 in Raleigh, North Carolina.",
            "credit": "Josh Lavallee/NHLI via Getty Images",
            "height": 729,
            "width": 1296,
            "url": "https://a.espncdn.com/photo/2026/0503/r1652863_1296x729_16-9.jpg"
          },
          {
            "dataSourceIdentifier": "e09e98e94bb75",
            "id": 48704137,
            "type": "header",
            "name": "Josh Doan [1296x729]",
            "caption": "The Sabres' Game 1 victory has their odds on the move.",
            "alt": "Buffalo Sabres right wing Josh Doan puts the puck past Montreal Canadiens goaltender Jakub Dobes (75) during the first period in Game 1 of a second-round NHL hockey Stanley Cup playoff series, Wednesday, May 6, 2026, in Buffalo, N.Y.",
            "credit": "AP Photo/Jeffrey T. Barnes",
            "height": 729,
            "width": 1296,
            "url": "https://a.espncdn.com/photo/2026/0507/r1654913_1296x729_16-9.jpg"
          }
        ],
        "categories": [
          {
            "id": 611584,
            "type": "league",
            "guid": "550a1e9e-377e-3019-bf5e-19c027d50a49",
            "description": "Sports Betting",
            "sportId": 22000,
            "leagueId": 22000,
            "league": {}
          },
          {
            "id": 137817,
            "type": "topic",
            "guid": "14f579db-79ec-a898-9405-8fec63e16063",
            "description": "news \u2013 sports betting",
            "sportId": 0,
            "topicId": 179
          }
        ],
        "premium": false,
        "links": {
          "web": {
            "href": "https://www.espn.com/espn/betting/story/_/id/44686270/2026-nhl-stanley-cup-playoffs-championship-odds"
          },
          "mobile": {
            "href": "http://m.espn.go.com/wireless/story?storyId=44686270"
          },
          "api": {
            "self": {}
          },
          "app": {
            "sportscenter": {}
          }
        },
        "byline": "Doug Greenberg"
      },
      {
        "id": 48714642,
        "nowId": "1-48714642",
        "contentKey": "48714642-1-6-1",
        "dataSourceIdentifier": "ee9ac16994ea7",
        "type": "Story",
        "headline": "Stanley Cup playoffs: How a culture of loyalty built the Wild's success",
        "description": "While some teams see veterans come and go, Minnesota has become a place players don't want to leave.",
        "lastModified": "2026-05-09T10:57:23Z",
        "published": "2026-05-09T10:57:23Z",
        "images": [
          {
            "dataSourceIdentifier": "35b4d2b893d72",
            "id": 48714797,
            "type": "header",
            "name": "Wild celebrate [1296x729]",
            "credit": "Bruce Kluckhohn/NHLI via Getty Images",
            "height": 729,
            "width": 1296,
            "url": "https://a.espncdn.com/photo/2026/0508/r1655517_1296x729_16-9.jpg"
          },
          {
            "dataSourceIdentifier": "424f1ee9a257e",
            "id": 48714776,
            "type": "header",
            "name": "Marcus Foligno [1296x729]",
            "caption": "Marcus Foligno is one of several veterans who chose to stay in Minnesota instead of potentially landing a bigger contract elsewhere.",
            "credit": "Luke Schmidt/NHLI via Getty Images",
            "height": 729,
            "width": 1296,
            "url": "https://a.espncdn.com/photo/2026/0508/r1655515_1296x729_16-9.jpg"
          }
        ],
        "categories": [
          {
            "id": 9580,
            "type": "league",
            "uid": "s:70~l:90",
            "guid": "1a5f0227-a13e-396c-8cea-8961bc288666",
            "description": "NHL",
            "sportId": 90,
            "leagueId": 90,
            "league": {}
          },
          {
            "id": 6601,
            "type": "team",
            "uid": "s:70~l:90~t:30",
            "guid": "b024c635-e192-6545-f8c1-b5da0df5ff6c",
            "description": "Minnesota Wild",
            "sportId": 90,
            "teamId": 30,
            "team": {}
          }
        ],
        "premium": false,
        "links": {
          "web": {
            "href": "https://www.espn.com/nhl/story/_/id/48714642/2026-nhl-playoffs-stanley-cup-minnesota-wild-franchise-contracts-free-agents-trades"
          },
          "mobile": {
            "href": "http://m.espn.go.com/nhl/story?storyId=48714642"
          },
          "api": {
            "self": {}
          },
          "app": {
            "sportscenter": {}
          }
        },
        "byline": "Ryan Clark"
      }
    ]
  },
  "header": {
    "id": "401871412",
    "uid": "s:70~l:90~e:401871412",
    "season": {
      "year": 2026,
      "current": true,
      "type": 3
    },
    "timeValid": true,
    "competitions": [
      {
        "id": "401871412",
        "uid": "s:70~l:90~e:401871412~c:401871412",
        "date": "2026-05-08T00:00Z",
        "neutralSite": false,
        "boxscoreAvailable": true,
        "commentaryAvailable": false,
        "liveAvailable": false,
        "shotChartAvailable": true,
        "onWatchESPN": false,
        "recent": false,
        "wallclockAvailable": true,
        "boxscoreSource": "full",
        "playByPlaySource": "full",
        "competitors": [
          {
            "id": "15",
            "uid": "s:70~l:90~t:15",
            "order": 0,
            "homeAway": "home",
            "winner": false,
            "team": {},
            "score": "1",
            "linescores": [],
            "record": [],
            "possession": false
          },
          {
            "id": "7",
            "uid": "s:70~l:90~t:7",
            "order": 1,
            "homeAway": "away",
            "winner": true,
            "team": {},
            "score": "4",
            "linescores": [],
            "record": [],
            "possession": false
          }
        ],
        "status": {
          "type": {
            "id": "3",
            "name": "STATUS_FINAL",
            "state": "post",
            "completed": true,
            "description": "Final",
            "detail": "Final",
            "shortDetail": "Final"
          },
          "featuredAthletes": [
            {},
            {}
          ]
        },
        "broadcasts": [
          {
            "type": {},
            "market": {},
            "media": {},
            "lang": "en",
            "region": "us",
            "isNational": true
          },
          {
            "type": {},
            "market": {},
            "media": {},
            "lang": "en",
            "region": "us",
            "isNational": true
          }
        ],
        "series": [
          {
            "type": "season",
            "title": "Regular Season Series",
            "description": "Regular Season Series",
            "summary": "CAR wins series 3-1",
            "completed": true,
            "totalCompetitions": 4,
            "competitors": [],
            "events": [],
            "startDate": "2026-05-02T04:00Z"
          },
          {
            "type": "playoff",
            "title": "Playoff Series",
            "description": "Playoff Series",
            "summary": "CAR leads series 3-0",
            "completed": false,
            "totalCompetitions": 7,
            "competitors": [],
            "events": [],
            "startDate": "2026-05-02T04:00Z"
          }
        ],
        "boxscoreMinutes": true
      }
    ],
    "links": [
      {
        "rel": [
          "summary",
          "desktop"
        ],
        "href": "https://www.espn.com/nhl/game/_/gameId/401871412/hurricanes-flyers",
        "text": "Gamecast",
        "shortText": "Gamecast",
        "isExternal": false,
        "isPremium": false
      },
      {
        "rel": [
          "recap",
          "desktop"
        ],
        "href": "https://www.espn.com/nhl/recap?gameId=401871412",
        "text": "Recap",
        "shortText": "Recap",
        "isExternal": false,
        "isPremium": false
      }
    ],
    "league": {
      "id": "90",
      "uid": "s:70~l:90",
      "name": "National Hockey League",
      "abbreviation": "NHL",
      "slug": "nhl",
      "isTournament": false,
      "links": [
        {
          "rel": [
            "index",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/",
          "text": "Index"
        },
        {
          "rel": [
            "index",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showClubhouse?uid=s:70~l:90",
          "text": "Index"
        }
      ],
      "logos": [
        {
          "href": "https://a.espncdn.com/i/teamlogos/leagues/500/nhl.png",
          "rel": [
            "full",
            "default"
          ]
        },
        {
          "href": "https://a.espncdn.com/i/teamlogos/leagues/500-dark/nhl.png",
          "rel": [
            "full",
            "dark"
          ]
        }
      ]
    },
    "gameNote": "East 2nd Round - Game 3",
    "standings": [
      {
        "team": "Carolina",
        "link": "https://www.espn.com/nhl/team/_/name/car/carolina-hurricanes",
        "id": "7",
        "uid": "s:70~l:90~t:7",
        "stats": [
          {
            "name": "otLosses",
            "displayName": "Overtime Losses",
            "shortDisplayName": "OTL",
            "description": "Number of Overtime Losses",
            "abbreviation": "OTL",
            "type": "otlosses",
            "value": 7.0,
            "displayValue": "7"
          },
          {
            "name": "losses",
            "displayName": "Losses",
            "shortDisplayName": "L",
            "description": "Losses",
            "abbreviation": "L",
            "type": "losses",
            "value": 22.0,
            "displayValue": "22"
          }
        ],
        "logo": [
          {
            "href": "https://a.espncdn.com/i/teamlogos/nhl/500/car.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-02T15:51Z"
          },
          {
            "href": "https://a.espncdn.com/i/teamlogos/nhl/500-dark/scoreboard/car.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-02T15:46Z"
          }
        ]
      },
      {
        "team": "Philadelphia",
        "link": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
        "id": "15",
        "uid": "s:70~l:90~t:15",
        "stats": [
          {
            "name": "otLosses",
            "displayName": "Overtime Losses",
            "shortDisplayName": "OTL",
            "description": "Number of Overtime Losses",
            "abbreviation": "OTL",
            "type": "otlosses",
            "value": 12.0,
            "displayValue": "12"
          },
          {
            "name": "losses",
            "displayName": "Losses",
            "shortDisplayName": "L",
            "description": "Losses",
            "abbreviation": "L",
            "type": "losses",
            "value": 27.0,
            "displayValue": "27"
          }
        ],
        "logo": [
          {
            "href": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-02T15:51Z"
          },
          {
            "href": "https://a.espncdn.com/i/teamlogos/nhl/500-dark/scoreboard/phi.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [],
            "lastUpdated": "2024-07-02T15:46Z"
          }
        ]
      }
    ]
  },
  "article": {
    "id": 48708465,
    "nowId": "1-48708465",
    "contentKey": "48708465-1-21-1",
    "dataSourceIdentifier": "3398684ea4ebc",
    "publishedkey": "nhl401871412",
    "type": "Recap",
    "gameId": "401871412",
    "headline": "Hurricanes beat the Flyers 4-1 in Game 3, take a 3-0 series lead",
    "description": "\u2014 Jordan Staal and Andrei Svechnikov scored on the power play and Jalen Chatfield added a short-handed goal, keying a special teams effort that helped the Carolina Hurricanes win their seventh straight playoff game, 4-1 over the Philadelphia Flyers in...",
    "linkText": "Hurricanes beat the Flyers 4-1 in Game 3, take a 3-0 series lead",
    "categorized": "2026-05-08T05:53:54Z",
    "originallyPosted": "2026-05-08T03:02:59Z",
    "lastModified": "2026-05-08T04:13:01Z",
    "published": "2026-05-08T03:02:59Z",
    "section": "NHL",
    "source": "AP",
    "images": [
      {
        "type": "Media",
        "name": "Multiple fights ensue in final minute of Hurricanes' win",
        "caption": "Multiple fights ensue in final minute of Hurricanes' win",
        "height": 324,
        "width": 576,
        "url": "https://a.espncdn.com/media/motion/2026/0508/ae14b8ed4b424d8b968128e4c8f31946113/ae14b8ed4b424d8b968128e4c8f31946113.jpg"
      }
    ],
    "video": [
      {
        "id": 48708480,
        "dataSourceIdentifier": "a4c4769549b76",
        "cerebroId": "69fd537bc4ffca19576e7ec2",
        "pccId": "b9fd762d-3151-4885-b496-549cd64d8d0f",
        "source": "espn",
        "headline": "Multiple fights ensue in final minute of Hurricanes' win",
        "shortHeadline": "Multiple fights ensue in final minute of Hurricanes' win",
        "caption": "Multiple fights ensue in final minute of Hurricanes' win",
        "title": "Multiple fights ensue in final minute of Hurricanes' win",
        "description": "Multiple fights ensue in final minute of Hurricanes' win",
        "shortDescription": "Multiple fights ensue in final minute of Hurricanes' win",
        "lastModified": "2026-05-08T05:53:52Z",
        "originalPublishDate": "2026-05-08T03:07:39Z",
        "premium": false,
        "syndicatable": true,
        "duration": 68,
        "videoRatio": "16:9,9:16",
        "timeRestrictions": {
          "embargoDate": "2026-05-08T03:07:34Z",
          "expirationDate": "2026-07-31T04:00:00Z"
        }
      }
    ]
  },
  "videos": [
    {
      "id": 48708577,
      "cerebroId": "69fd5b78c35ca554fbca235f",
      "source": "espn",
      "headline": "Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights",
      "description": "Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights",
      "lastModified": "2026-05-08T03:42:06Z",
      "originalPublishDate": "2026-05-08T03:41:45Z",
      "duration": 85,
      "timeRestrictions": {
        "embargoDate": "2026-05-08T03:41:35Z",
        "expirationDate": "2026-07-31T04:00:00Z"
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
          "PW"
        ]
      },
      "thumbnail": "https://a.espncdn.com/media/motion/wsc/2026/0508/ed49137f-122f-4bad-a14b-3e5bc904fd7e/ed49137f-122f-4bad-a14b-3e5bc904fd7e.jpg",
      "links": {
        "web": {
          "href": "https://www.espn.com/video/clip/_/id/48708577/game-highlights",
          "self": {
            "href": "https://www.espn.com/video/clip/_/id/48708577/game-highlights",
            "dsi": {}
          }
        },
        "mobile": {
          "source": {
            "href": "https://media.video-cdn.espn.com/motion/wsc/2026/0508/ed49137f-122f-4bad-a14b-3e5bc904fd7e/ed49137f-122f-4bad-a14b-3e5bc904fd7e.mp4"
          },
          "alert": {
            "href": "https://m.espn.com/general/video/videoAlert?vid=48708577"
          },
          "streaming": {
            "href": "https://watch.auth.api.espn.com/video/auth/brightcove/064f7bd0-b287-4019-be2c-f4e07dd50322/asset?UMADPARAMreferer=https://www.espn.com/video/clip/_/id/48708577/game-highlights"
          },
          "progressiveDownload": {
            "href": "https://watch.auth.api.espn.com/video/auth/brightcove/064f7bd0-b287-4019-be2c-f4e07dd50322/asset?UMADPARAMreferer=https://www.espn.com/video/clip/_/id/48708577/game-highlights"
          }
        },
        "api": {
          "self": {
            "href": "https://content.core.api.espn.com/v1/video/clips/48708577"
          },
          "artwork": {
            "href": "https://artwork.api.espn.com/artwork/collections/media/064f7bd0-b287-4019-be2c-f4e07dd50322"
          }
        },
        "source": {
          "href": "https://media.video-cdn.espn.com/motion/wsc/2026/0508/ed49137f-122f-4bad-a14b-3e5bc904fd7e/ed49137f-122f-4bad-a14b-3e5bc904fd7e_360p30_1464k.mp4",
          "mezzanine": {
            "href": "https://media.video-origin.espn.com/espnvideo/wsc/2026/0508/ed49137f-122f-4bad-a14b-3e5bc904fd7e/ed49137f-122f-4bad-a14b-3e5bc904fd7e.mp4"
          },
          "flash": {
            "href": "https://media.video-cdn.espn.com/motion/wsc/2026/0508/ed49137f-122f-4bad-a14b-3e5bc904fd7e/ed49137f-122f-4bad-a14b-3e5bc904fd7e.smil"
          },
          "hds": {
            "href": "https://hds.video-cdn.espn.com/z/motion/wsc/2026/0508/ed49137f-122f-4bad-a14b-3e5bc904fd7e/ed49137f-122f-4bad-a14b-3e5bc904fd7e_rel.smil/manifest.f4m"
          },
          "HLS": {
            "href": "https://service-pkgespn.akamaized.net/opp/hls/espn/wsc/2026/0508/ed49137f-122f-4bad-a14b-3e5bc904fd7e/ed49137f-122f-4bad-a14b-3e5bc904fd7e/playlist.m3u8",
            "HD": {},
            "cmaf": {},
            "9x16": {},
            "shield": {}
          },
          "HD": {
            "href": "https://media.video-cdn.espn.com/motion/wsc/2026/0508/ed49137f-122f-4bad-a14b-3e5bc904fd7e/ed49137f-122f-4bad-a14b-3e5bc904fd7e_720p30_2896k.mp4"
          },
          "full": {
            "href": "https://media.video-cdn.espn.com/motion/wsc/2026/0508/ed49137f-122f-4bad-a14b-3e5bc904fd7e/ed49137f-122f-4bad-a14b-3e5bc904fd7e_360p30_1464k.mp4"
          }
        },
        "sportscenter": {
          "href": "sportscenter://x-callback-url/showVideo?videoID=48708577&videoDSI=8dd990f895c00"
        }
      },
      "ad": {
        "sport": "nhl",
        "bundle": "sportscenter"
      },
      "tracking": {
        "sportName": "hockey",
        "leagueName": "NHL",
        "coverageType": "Final Game Highlight",
        "trackingName": "NHL_One-Play (Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights) 2026/05/08 ESHEET",
        "trackingId": "dm_20260508_NHL_carolina_hurricanes_vs_philadelphia_flyers_game_highlights",
        "program": "Turner Network Television, truTV"
      }
    },
    {
      "id": 48708705,
      "cerebroId": "69fd6106c35ca554fbca25e8",
      "source": "espn",
      "headline": "Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights",
      "description": "Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights",
      "lastModified": "2026-05-08T04:05:51Z",
      "originalPublishDate": "2026-05-08T04:05:26Z",
      "duration": 89,
      "timeRestrictions": {
        "embargoDate": "2026-05-08T04:05:17Z",
        "expirationDate": "2026-07-31T04:00:00Z"
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
          "PW"
        ]
      },
      "thumbnail": "https://a.espncdn.com/media/motion/wsc/2026/0508/1ebfd808-1a11-4c6d-ac8d-8754bc4db2f4/1ebfd808-1a11-4c6d-ac8d-8754bc4db2f4.jpg",
      "links": {
        "web": {
          "href": "https://www.espn.com/video/clip/_/id/48708705/game-highlights",
          "self": {
            "href": "https://www.espn.com/video/clip/_/id/48708705/game-highlights",
            "dsi": {}
          }
        },
        "mobile": {
          "source": {
            "href": "https://media.video-cdn.espn.com/motion/wsc/2026/0508/1ebfd808-1a11-4c6d-ac8d-8754bc4db2f4/1ebfd808-1a11-4c6d-ac8d-8754bc4db2f4.mp4"
          },
          "alert": {
            "href": "https://m.espn.com/general/video/videoAlert?vid=48708705"
          },
          "streaming": {
            "href": "https://watch.auth.api.espn.com/video/auth/brightcove/b858405a-b22e-4a32-9560-ac6b569b7c35/asset?UMADPARAMreferer=https://www.espn.com/video/clip/_/id/48708705/game-highlights"
          },
          "progressiveDownload": {
            "href": "https://watch.auth.api.espn.com/video/auth/brightcove/b858405a-b22e-4a32-9560-ac6b569b7c35/asset?UMADPARAMreferer=https://www.espn.com/video/clip/_/id/48708705/game-highlights"
          }
        },
        "api": {
          "self": {
            "href": "https://content.core.api.espn.com/v1/video/clips/48708705"
          },
          "artwork": {
            "href": "https://artwork.api.espn.com/artwork/collections/media/b858405a-b22e-4a32-9560-ac6b569b7c35"
          }
        },
        "source": {
          "href": "https://media.video-cdn.espn.com/motion/wsc/2026/0508/1ebfd808-1a11-4c6d-ac8d-8754bc4db2f4/1ebfd808-1a11-4c6d-ac8d-8754bc4db2f4_360p30_1464k.mp4",
          "mezzanine": {
            "href": "https://media.video-origin.espn.com/espnvideo/wsc/2026/0508/1ebfd808-1a11-4c6d-ac8d-8754bc4db2f4/1ebfd808-1a11-4c6d-ac8d-8754bc4db2f4.mp4"
          },
          "flash": {
            "href": "https://media.video-cdn.espn.com/motion/wsc/2026/0508/1ebfd808-1a11-4c6d-ac8d-8754bc4db2f4/1ebfd808-1a11-4c6d-ac8d-8754bc4db2f4.smil"
          },
          "hds": {
            "href": "https://hds.video-cdn.espn.com/z/motion/wsc/2026/0508/1ebfd808-1a11-4c6d-ac8d-8754bc4db2f4/1ebfd808-1a11-4c6d-ac8d-8754bc4db2f4_rel.smil/manifest.f4m"
          },
          "HLS": {
            "href": "https://service-pkgespn.akamaized.net/opp/hls/espn/wsc/2026/0508/1ebfd808-1a11-4c6d-ac8d-8754bc4db2f4/1ebfd808-1a11-4c6d-ac8d-8754bc4db2f4/playlist.m3u8",
            "HD": {},
            "cmaf": {},
            "9x16": {},
            "shield": {}
          },
          "HD": {
            "href": "https://media.video-cdn.espn.com/motion/wsc/2026/0508/1ebfd808-1a11-4c6d-ac8d-8754bc4db2f4/1ebfd808-1a11-4c6d-ac8d-8754bc4db2f4_720p30_2896k.mp4"
          },
          "full": {
            "href": "https://media.video-cdn.espn.com/motion/wsc/2026/0508/1ebfd808-1a11-4c6d-ac8d-8754bc4db2f4/1ebfd808-1a11-4c6d-ac8d-8754bc4db2f4_360p30_1464k.mp4"
          }
        },
        "sportscenter": {
          "href": "sportscenter://x-callback-url/showVideo?videoID=48708705&videoDSI=227e22a896efb"
        }
      },
      "ad": {
        "sport": "nhl",
        "bundle": "sportscenter"
      },
      "tracking": {
        "sportName": "hockey",
        "leagueName": "NHL",
        "coverageType": "Final Game Highlight",
        "trackingName": "NHL_One-Play (Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights) 2026/05/08 ESHEET",
        "trackingId": "dm_20260508_NHL_carolina_hurricanes_vs_philadelphia_flyers_game_highlights",
        "program": "Turner Network Television, truTV"
      }
    }
  ],
  "onIce": [
    {
      "entries": [
        {
          "whereabouts": {
            "id": "1",
            "description": "In Play",
            "name": "ROSTER_WHEREABOUTS_IN_PLAY"
          },
          "projections": {
            "$ref": "http://sports.core.api.espn.pvt/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15/roster/2562601/projections?lang=en&region=us"
          },
          "athleteid": "2562601"
        },
        {
          "whereabouts": {
            "id": "1",
            "description": "In Play",
            "name": "ROSTER_WHEREABOUTS_IN_PLAY"
          },
          "projections": {
            "$ref": "http://sports.core.api.espn.pvt/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15/roster/2335062/projections?lang=en&region=us"
          },
          "athleteid": "2335062"
        }
      ],
      "teamId": "15"
    },
    {
      "entries": [
        {
          "whereabouts": {
            "id": "1",
            "description": "In Play",
            "name": "ROSTER_WHEREABOUTS_IN_PLAY"
          },
          "projections": {
            "$ref": "http://sports.core.api.espn.pvt/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/7/roster/2517899/projections?lang=en&region=us"
          },
          "athleteid": "2517899"
        },
        {
          "whereabouts": {
            "id": "1",
            "description": "In Play",
            "name": "ROSTER_WHEREABOUTS_IN_PLAY"
          },
          "projections": {
            "$ref": "http://sports.core.api.espn.pvt/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/7/roster/2976849/projections?lang=en&region=us"
          },
          "athleteid": "2976849"
        }
      ],
      "teamId": "7"
    }
  ],
  "plays": [
    {
      "id": "401871412000000520",
      "sequenceNumber": "2",
      "type": {
        "id": "518",
        "text": "Period Start",
        "abbreviation": "period-start"
      },
      "text": "Start of 1st Period",
      "awayScore": 0,
      "homeScore": 0,
      "period": {
        "number": 1,
        "displayValue": "1st"
      },
      "clock": {
        "displayValue": "0:00"
      },
      "scoringPlay": false,
      "scoreValue": 0,
      "modified": "2026-05-08T15:11Z",
      "wallclock": "2026-05-08T00:11:08Z",
      "shootingPlay": true,
      "strength": {
        "id": "701",
        "text": "Even Strength",
        "abbreviation": "even-strength"
      }
    },
    {
      "id": "401871412000000510",
      "sequenceNumber": "3",
      "type": {
        "id": "502",
        "text": "Face Off",
        "abbreviation": "faceoff"
      },
      "text": "Luke Glendening faceoff won against Jordan Staal",
      "awayScore": 0,
      "homeScore": 0,
      "period": {
        "number": 1,
        "displayValue": "1st"
      },
      "clock": {
        "displayValue": "0:00"
      },
      "scoringPlay": false,
      "scoreValue": 0,
      "modified": "2026-05-08T15:11Z",
      "team": {
        "id": "15"
      },
      "participants": [
        {
          "athlete": {
            "id": "2335062",
            "displayName": "Luke Glendening",
            "shortName": "L. Glendening",
            "headshot": {}
          },
          "type": "face-off-winner",
          "ytdGoals": 1
        },
        {
          "athlete": {
            "id": "3541",
            "displayName": "Jordan Staal",
            "shortName": "J. Staal",
            "headshot": {}
          },
          "ytdAssists": 2
        }
      ],
      "wallclock": "2026-05-08T00:11:08Z",
      "shootingPlay": true,
      "strength": {
        "id": "701",
        "text": "Even Strength",
        "abbreviation": "even-strength"
      }
    }
  ],
  "wallclockAvailable": true,
  "meta": {
    "gp_topic": "gp-hockey-nhl-401871412",
    "gameSwitcherEnabled": true,
    "picker_topic": "picker-hockey-nhl",
    "lastUpdatedAt": "2026-05-08T02:52:03Z",
    "firstPlayWallClock": "2026-05-08T00:11:08Z",
    "lastPlayWallClock": "2026-05-08T02:52:03Z",
    "gameState": "post",
    "syncUrl": "https://client.espncdn.com/fauxcast/stats/90/401871412/en/us/"
  }
}
```
