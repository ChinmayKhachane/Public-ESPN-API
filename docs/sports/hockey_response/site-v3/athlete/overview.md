# Athlete Overview

## https://site.web.api.espn.com/apis/common/v3/sports/hockey/{league}/athletes/{athlete}/overview

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.web.api.espn.com/apis/common/v3/sports/hockey/nhl/athletes/4565230/overview`

## Example Response

```json
{
  "statistics": {
    "displayName": "2025-26 General",
    "labels": [
      "GP",
      "G"
    ],
    "names": [
      "games",
      "goals"
    ],
    "displayNames": [
      "Games Played",
      "Goals"
    ],
    "splits": [
      {
        "displayName": "Regular Season",
        "stats": [
          "81",
          "26"
        ]
      },
      {
        "displayName": "Postseason",
        "stats": [
          "9",
          "2"
        ]
      }
    ]
  },
  "news": [
    {
      "headline": "Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights",
      "lastModified": "2026-05-08T04:05:51.000+00:00",
      "root": "nhl",
      "premium": false,
      "links": {
        "api": {
          "artwork": {
            "href": "https://artwork.api.espn.com/artwork/collections/media/b858405a-b22e-4a32-9560-ac6b569b7c35"
          },
          "self": {
            "href": "https://content.core.api.espn.com/v1/video/clips/48708705"
          }
        },
        "web": {
          "href": "https://www.espn.com/video/clip/_/id/48708705/game-highlights",
          "self": {
            "href": "https://www.espn.com/video/clip/_/id/48708705/game-highlights"
          }
        }
      },
      "type": "Media",
      "section": "NHL",
      "id": 48708705,
      "categorized": "2026-05-08T04:05:55Z",
      "description": "Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights",
      "nowId": "1-48708705",
      "allowComments": true,
      "images": [
        {
          "width": 576,
          "height": 324,
          "url": "https://a.espncdn.com/media/motion/wsc/2026/0508/1ebfd808-1a11-4c6d-ac8d-8754bc4db2f4/1ebfd808-1a11-4c6d-ac8d-8754bc4db2f4.jpg",
          "name": "Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights",
          "caption": "Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights"
        }
      ],
      "categories": [
        {
          "id": 106576,
          "uid": "s:70~l:90~a:2517899",
          "description": "Frederik Andersen",
          "type": "athlete",
          "sportId": 90,
          "athleteId": 2517899,
          "athlete": {
            "id": 2517899,
            "links": {}
          }
        },
        {
          "id": 26419,
          "uid": "s:70~l:90~a:3541",
          "description": "Jordan Staal",
          "type": "athlete",
          "sportId": 90,
          "athleteId": 3541,
          "athlete": {
            "id": 3541,
            "links": {}
          }
        }
      ],
      "published": "2026-05-08T04:05:51.000+00:00",
      "video": [
        {
          "source": "espn",
          "id": 48708705,
          "headline": "Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights",
          "title": "Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights",
          "caption": "Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights",
          "description": "Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights",
          "premium": false,
          "ad": {
            "sport": "nhl",
            "bundle": "sportscenter"
          },
          "tracking": {
            "sportName": "hockey",
            "leagueName": "NHL",
            "coverageType": "Final Game Highlight",
            "trackingName": "NHL_One-Play (Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights) 2026/05/08 ESHEET",
            "trackingId": "dm_20260508_NHL_carolina_hurricanes_vs_philadelphia_flyers_game_highlights"
          },
          "cerebroId": "69fd6106c35ca554fbca25e8",
          "lastModified": "2026-05-08T04:05:51.000+00:00",
          "originalPublishDate": "2026-05-08T04:05:26.000+00:00",
          "timeRestrictions": {
            "embargoDate": "2026-05-08T04:05:17.000+00:00",
            "expirationDate": "2026-07-31T04:00:00.000+00:00"
          },
          "deviceRestrictions": {
            "type": "whitelist",
            "devices": []
          },
          "geoRestrictions": {
            "type": "whitelist",
            "countries": []
          },
          "syndicatable": true,
          "duration": 89,
          "categories": [
            {},
            {}
          ]
        }
      ],
      "dataSourceIdentifier": "227e22a896efb"
    },
    {
      "headline": "Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights",
      "lastModified": "2026-05-08T03:42:06.000+00:00",
      "root": "nhl",
      "premium": false,
      "links": {
        "api": {
          "artwork": {
            "href": "https://artwork.api.espn.com/artwork/collections/media/064f7bd0-b287-4019-be2c-f4e07dd50322"
          },
          "self": {
            "href": "https://content.core.api.espn.com/v1/video/clips/48708577"
          }
        },
        "web": {
          "href": "https://www.espn.com/video/clip/_/id/48708577/game-highlights",
          "self": {
            "href": "https://www.espn.com/video/clip/_/id/48708577/game-highlights"
          }
        }
      },
      "type": "Media",
      "section": "NHL",
      "id": 48708577,
      "categorized": "2026-05-08T03:42:09Z",
      "description": "Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights",
      "nowId": "1-48708577",
      "allowComments": true,
      "images": [
        {
          "width": 576,
          "height": 324,
          "url": "https://a.espncdn.com/media/motion/wsc/2026/0508/ed49137f-122f-4bad-a14b-3e5bc904fd7e/ed49137f-122f-4bad-a14b-3e5bc904fd7e.jpg",
          "name": "Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights",
          "caption": "Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights"
        }
      ],
      "categories": [
        {
          "id": 106576,
          "uid": "s:70~l:90~a:2517899",
          "description": "Frederik Andersen",
          "type": "athlete",
          "sportId": 90,
          "athleteId": 2517899,
          "athlete": {
            "id": 2517899,
            "links": {}
          }
        },
        {
          "id": 384310,
          "uid": "s:70~l:90~a:4565230",
          "description": "Trevor Zegras",
          "type": "athlete",
          "sportId": 90,
          "athleteId": 4565230,
          "athlete": {
            "id": 4565230,
            "links": {}
          }
        }
      ],
      "published": "2026-05-08T03:42:06.000+00:00",
      "video": [
        {
          "source": "espn",
          "id": 48708577,
          "headline": "Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights",
          "title": "Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights",
          "caption": "Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights",
          "description": "Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights",
          "premium": false,
          "ad": {
            "sport": "nhl",
            "bundle": "sportscenter"
          },
          "tracking": {
            "sportName": "hockey",
            "leagueName": "NHL",
            "coverageType": "Final Game Highlight",
            "trackingName": "NHL_One-Play (Carolina Hurricanes vs. Philadelphia Flyers: Game Highlights) 2026/05/08 ESHEET",
            "trackingId": "dm_20260508_NHL_carolina_hurricanes_vs_philadelphia_flyers_game_highlights"
          },
          "cerebroId": "69fd5b78c35ca554fbca235f",
          "lastModified": "2026-05-08T03:42:06.000+00:00",
          "originalPublishDate": "2026-05-08T03:41:45.000+00:00",
          "timeRestrictions": {
            "embargoDate": "2026-05-08T03:41:35.000+00:00",
            "expirationDate": "2026-07-31T04:00:00.000+00:00"
          },
          "deviceRestrictions": {
            "type": "whitelist",
            "devices": []
          },
          "geoRestrictions": {
            "type": "whitelist",
            "countries": []
          },
          "syndicatable": true,
          "duration": 85,
          "categories": [
            {},
            {}
          ]
        }
      ],
      "dataSourceIdentifier": "8dd990f895c00"
    }
  ],
  "nextGame": {
    "displayName": "Next Game",
    "league": {
      "id": "90",
      "uid": "s:70~l:90",
      "name": "National Hockey League",
      "shortName": "NHL",
      "abbreviation": "NHL",
      "slug": "nhl",
      "events": [
        {
          "id": "401871413",
          "competitionId": "401871413",
          "uid": "s:70~l:90~e:401871413",
          "date": "2026-05-09T22:00:00.000+00:00",
          "timeValid": true,
          "name": "Carolina Hurricanes at Philadelphia Flyers",
          "shortName": "CAR @ PHI",
          "location": "Xfinity Mobile Arena",
          "season": 2026,
          "seasonStartDate": "2025-09-20T07:00:00.000+00:00",
          "seasonEndDate": "2026-07-01T06:59:00.000+00:00",
          "seasonType": 3,
          "seasonTypeHasGroups": false,
          "period": 0,
          "clock": "0.0",
          "links": [
            {},
            {}
          ],
          "status": "pre",
          "fullStatus": {
            "clock": "0.0",
            "displayClock": "0:00",
            "period": 0,
            "type": {}
          }
        }
      ]
    },
    "summaryStatistics": [
      {
        "splitId": "7",
        "splitAbbreviation": "vs Carolina Hurricanes",
        "splitName": "vs Carolina Hurricanes",
        "shortDisplayName": "Goals",
        "displayValue": "3",
        "splitBy": "byOpponent",
        "name": "goals",
        "displayName": "Goals",
        "abbreviation": "G"
      },
      {
        "splitId": "7",
        "splitAbbreviation": "vs Carolina Hurricanes",
        "splitName": "vs Carolina Hurricanes",
        "shortDisplayName": "Assists",
        "displayValue": "3",
        "splitBy": "byOpponent",
        "name": "assists",
        "displayName": "Assists",
        "abbreviation": "A"
      }
    ],
    "statistics": {
      "labels": [
        "GP",
        "G"
      ],
      "names": [
        "games",
        "goals"
      ],
      "displayNames": [
        "Games Played",
        "Goals"
      ],
      "splits": [
        {
          "displayName": "In losses",
          "stats": [
            "39",
            "7"
          ],
          "type": "losses"
        },
        {
          "displayName": "vs Carolina Hurricanes",
          "stats": [
            "4",
            "3"
          ],
          "type": "byOpponent"
        }
      ]
    }
  },
  "gameLog": {
    "displayName": "Recent Games",
    "statistics": [
      {
        "displayName": "center",
        "labels": [
          "G",
          "A"
        ],
        "names": [
          "goals",
          "assists"
        ],
        "displayNames": [
          "Goals",
          "Assists"
        ],
        "events": [
          {
            "eventId": "401871412",
            "stats": []
          },
          {
            "eventId": "401871411",
            "stats": []
          }
        ]
      }
    ],
    "events": {
      "401869801": {
        "id": "401869801",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/game/_/gameId/401869801/flyers-penguins",
            "text": "Gamecast",
            "shortText": "Gamecast",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "sportscenter://x-callback-url/showGame?sportName=hockey&leagueAbbrev=nhl&gameId=401869801",
            "text": "Gamecast",
            "shortText": "Gamecast",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "atVs": "@",
        "gameDate": "2026-04-27T23:00:00.000+00:00",
        "score": "3-2",
        "homeTeamId": "16",
        "awayTeamId": "15",
        "homeTeamScore": "3",
        "awayTeamScore": "2",
        "gameResult": "L",
        "opponent": {
          "id": "16",
          "uid": "s:70~l:90~t:16",
          "displayName": "Pittsburgh Penguins",
          "abbreviation": "PIT",
          "links": [
            {},
            {}
          ],
          "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/pit.png"
        },
        "leagueName": "National Hockey League",
        "leagueAbbreviation": "NHL",
        "leagueShortName": "NHL",
        "eventNote": "East 1st Round - Game 5",
        "team": {
          "id": "15",
          "uid": "s:70~l:90~t:15",
          "abbreviation": "PHI",
          "links": [
            {},
            {}
          ],
          "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
          "isAllStar": false
        },
        "type": {
          "id": "14",
          "text": "Round of 16",
          "abbreviation": "RD16",
          "type": "round-of-16",
          "slug": "round-of-16"
        }
      },
      "401871412": {
        "id": "401871412",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/game/_/gameId/401871412/hurricanes-flyers",
            "text": "Gamecast",
            "shortText": "Gamecast",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "sportscenter://x-callback-url/showGame?sportName=hockey&leagueAbbrev=nhl&gameId=401871412",
            "text": "Gamecast",
            "shortText": "Gamecast",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "atVs": "vs",
        "gameDate": "2026-05-08T00:00:00.000+00:00",
        "score": "4-1",
        "homeTeamId": "15",
        "awayTeamId": "7",
        "homeTeamScore": "1",
        "awayTeamScore": "4",
        "gameResult": "L",
        "opponent": {
          "id": "7",
          "uid": "s:70~l:90~t:7",
          "displayName": "Carolina Hurricanes",
          "abbreviation": "CAR",
          "links": [
            {},
            {}
          ],
          "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/car.png"
        },
        "leagueName": "National Hockey League",
        "leagueAbbreviation": "NHL",
        "leagueShortName": "NHL",
        "eventNote": "East 2nd Round - Game 3",
        "team": {
          "id": "15",
          "uid": "s:70~l:90~t:15",
          "abbreviation": "PHI",
          "links": [
            {},
            {}
          ],
          "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
          "isAllStar": false
        },
        "type": {
          "id": "15",
          "text": "Quarterfinal",
          "abbreviation": "QTR",
          "type": "quarterfinal",
          "slug": "quarterfinal"
        }
      },
      "401869802": {
        "id": "401869802",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/game/_/gameId/401869802/penguins-flyers",
            "text": "Gamecast",
            "shortText": "Gamecast",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "sportscenter://x-callback-url/showGame?sportName=hockey&leagueAbbrev=nhl&gameId=401869802",
            "text": "Gamecast",
            "shortText": "Gamecast",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "atVs": "vs",
        "gameDate": "2026-04-29T23:30:00.000+00:00",
        "score": "1-0 OT",
        "homeTeamId": "15",
        "awayTeamId": "16",
        "homeTeamScore": "1",
        "awayTeamScore": "0",
        "gameResult": "W",
        "opponent": {
          "id": "16",
          "uid": "s:70~l:90~t:16",
          "displayName": "Pittsburgh Penguins",
          "abbreviation": "PIT",
          "links": [
            {},
            {}
          ],
          "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/pit.png"
        },
        "leagueName": "National Hockey League",
        "leagueAbbreviation": "NHL",
        "leagueShortName": "NHL",
        "eventNote": "East 1st Round - Game 6",
        "team": {
          "id": "15",
          "uid": "s:70~l:90~t:15",
          "abbreviation": "PHI",
          "links": [
            {},
            {}
          ],
          "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
          "isAllStar": false
        },
        "type": {
          "id": "14",
          "text": "Round of 16",
          "abbreviation": "RD16",
          "type": "round-of-16",
          "slug": "round-of-16"
        }
      },
      "401871166": {
        "id": "401871166",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/game/_/gameId/401871166/flyers-hurricanes",
            "text": "Gamecast",
            "shortText": "Gamecast",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "sportscenter://x-callback-url/showGame?sportName=hockey&leagueAbbrev=nhl&gameId=401871166",
            "text": "Gamecast",
            "shortText": "Gamecast",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "atVs": "@",
        "gameDate": "2026-05-03T00:00:00.000+00:00",
        "score": "3-0",
        "homeTeamId": "7",
        "awayTeamId": "15",
        "homeTeamScore": "3",
        "awayTeamScore": "0",
        "gameResult": "L",
        "opponent": {
          "id": "7",
          "uid": "s:70~l:90~t:7",
          "displayName": "Carolina Hurricanes",
          "abbreviation": "CAR",
          "links": [
            {},
            {}
          ],
          "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/car.png"
        },
        "leagueName": "National Hockey League",
        "leagueAbbreviation": "NHL",
        "leagueShortName": "NHL",
        "eventNote": "East 2nd Round - Game 1",
        "team": {
          "id": "15",
          "uid": "s:70~l:90~t:15",
          "abbreviation": "PHI",
          "links": [
            {},
            {}
          ],
          "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
          "isAllStar": false
        },
        "type": {
          "id": "15",
          "text": "Quarterfinal",
          "abbreviation": "QTR",
          "type": "quarterfinal",
          "slug": "quarterfinal"
        }
      },
      "401871411": {
        "id": "401871411",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/game/_/gameId/401871411/flyers-hurricanes",
            "text": "Gamecast",
            "shortText": "Gamecast",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "sportscenter://x-callback-url/showGame?sportName=hockey&leagueAbbrev=nhl&gameId=401871411",
            "text": "Gamecast",
            "shortText": "Gamecast",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "atVs": "@",
        "gameDate": "2026-05-04T23:00:00.000+00:00",
        "score": "3-2 OT",
        "homeTeamId": "7",
        "awayTeamId": "15",
        "homeTeamScore": "3",
        "awayTeamScore": "2",
        "gameResult": "L",
        "opponent": {
          "id": "7",
          "uid": "s:70~l:90~t:7",
          "displayName": "Carolina Hurricanes",
          "abbreviation": "CAR",
          "links": [
            {},
            {}
          ],
          "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/car.png"
        },
        "leagueName": "National Hockey League",
        "leagueAbbreviation": "NHL",
        "leagueShortName": "NHL",
        "eventNote": "East 2nd Round - Game 2",
        "team": {
          "id": "15",
          "uid": "s:70~l:90~t:15",
          "abbreviation": "PHI",
          "links": [
            {},
            {}
          ],
          "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
          "isAllStar": false
        },
        "type": {
          "id": "15",
          "text": "Quarterfinal",
          "abbreviation": "QTR",
          "type": "quarterfinal",
          "slug": "quarterfinal"
        }
      }
    }
  },
  "rotowire": {
    "headline": "Zegras scored a goal on two shots and added two PIM in Thursday's 4-1 loss to the Hurricanes in Game 3.",
    "story": "Zegras' tally tied the game at 1-1 early in the second period, but that was all Frederik Andersen would give up. The 25-year-old Zegras snapped a four-game point drought with the goal. He's earned two goals, three assists, 12 shots on net, 22 hits, 26 PIM and an even plus-minus rating over nine outings in a top-six role this postseason.",
    "description": "Zegras scored a goal on two shots and added two PIM in Thursday's 4-1 loss to the Hurricanes in Game 3.",
    "published": "Thu May 07 20:41:08 PDT 2026"
  },
  "fantasy": {
    "draftRank": "160",
    "positionRank": "15",
    "percentOwned": "73.94",
    "last7Days": "-1.07",
    "projection": "For a second straight year, Zegras' season was marred by injury. He managed just four goals and 10 points in his first 24 games before suffering a knee ailment, which sidelined him for six weeks. All told, the 2019 No. 9 overall pick finished with 12 goals and 32 points in 57 games. The biggest concern of all was the lack of production with the man advantage. Zegras had just four power-play points despite averaging a whopping 2:30 worth of power-play ice time. Still just 24 years of age, Zegras was traded from Anaheim to Philadelphia in June. The change of scenery may have been what was needed for Zegras to once again reach the 60-point mark."
  }
}
```
