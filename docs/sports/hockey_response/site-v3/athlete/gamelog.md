# Athlete Game Log

## https://site.web.api.espn.com/apis/common/v3/sports/hockey/{league}/athletes/{athlete}/gamelog

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.web.api.espn.com/apis/common/v3/sports/hockey/nhl/athletes/4565230/gamelog`

## Example Response

```json
{
  "filters": [
    {
      "displayName": "Season",
      "name": "season",
      "value": "2026",
      "options": [
        {
          "value": "2026",
          "displayValue": "2025-26"
        },
        {
          "value": "2025",
          "displayValue": "2024-25"
        }
      ]
    }
  ],
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
  "events": {
    "401871412": {
      "id": "401871412",
      "links": [
        {
          "language": "en",
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
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
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
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/car/carolina-hurricanes",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=7",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
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
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=15",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
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
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/game/_/gameId/401871411/flyers-hurricanes",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
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
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/car/carolina-hurricanes",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=7",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
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
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=15",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
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
    "401871166": {
      "id": "401871166",
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/game/_/gameId/401871166/flyers-hurricanes",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
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
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/car/carolina-hurricanes",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=7",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
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
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=15",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
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
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/game/_/gameId/401869802/penguins-flyers",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
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
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/pit/pittsburgh-penguins",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=16",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
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
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=15",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
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
    "401869801": {
      "id": "401869801",
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/game/_/gameId/401869801/flyers-penguins",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
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
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/pit/pittsburgh-penguins",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=16",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
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
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=15",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
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
    "401869799": {
      "id": "401869799",
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/game/_/gameId/401869799/penguins-flyers",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=hockey&leagueAbbrev=nhl&gameId=401869799",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "atVs": "vs",
      "gameDate": "2026-04-26T00:00:00.000+00:00",
      "score": "4-2",
      "homeTeamId": "15",
      "awayTeamId": "16",
      "homeTeamScore": "2",
      "awayTeamScore": "4",
      "gameResult": "L",
      "opponent": {
        "id": "16",
        "uid": "s:70~l:90~t:16",
        "displayName": "Pittsburgh Penguins",
        "abbreviation": "PIT",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/pit/pittsburgh-penguins",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=16",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/pit.png"
      },
      "leagueName": "National Hockey League",
      "leagueAbbreviation": "NHL",
      "leagueShortName": "NHL",
      "eventNote": "East 1st Round - Game 4",
      "team": {
        "id": "15",
        "uid": "s:70~l:90~t:15",
        "abbreviation": "PHI",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=15",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
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
    "401869796": {
      "id": "401869796",
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/game/_/gameId/401869796/penguins-flyers",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=hockey&leagueAbbrev=nhl&gameId=401869796",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "atVs": "vs",
      "gameDate": "2026-04-22T23:00:00.000+00:00",
      "score": "5-2",
      "homeTeamId": "15",
      "awayTeamId": "16",
      "homeTeamScore": "5",
      "awayTeamScore": "2",
      "gameResult": "W",
      "opponent": {
        "id": "16",
        "uid": "s:70~l:90~t:16",
        "displayName": "Pittsburgh Penguins",
        "abbreviation": "PIT",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/pit/pittsburgh-penguins",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=16",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/pit.png"
      },
      "leagueName": "National Hockey League",
      "leagueAbbreviation": "NHL",
      "leagueShortName": "NHL",
      "eventNote": "East 1st Round - Game 3",
      "team": {
        "id": "15",
        "uid": "s:70~l:90~t:15",
        "abbreviation": "PHI",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=15",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
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
    "401869794": {
      "id": "401869794",
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/game/_/gameId/401869794/flyers-penguins",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=hockey&leagueAbbrev=nhl&gameId=401869794",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "atVs": "@",
      "gameDate": "2026-04-20T23:00:00.000+00:00",
      "score": "3-0",
      "homeTeamId": "16",
      "awayTeamId": "15",
      "homeTeamScore": "0",
      "awayTeamScore": "3",
      "gameResult": "W",
      "opponent": {
        "id": "16",
        "uid": "s:70~l:90~t:16",
        "displayName": "Pittsburgh Penguins",
        "abbreviation": "PIT",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/pit/pittsburgh-penguins",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=16",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/pit.png"
      },
      "leagueName": "National Hockey League",
      "leagueAbbreviation": "NHL",
      "leagueShortName": "NHL",
      "eventNote": "East 1st Round - Game 2",
      "team": {
        "id": "15",
        "uid": "s:70~l:90~t:15",
        "abbreviation": "PHI",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=15",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
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
    "401869717": {
      "id": "401869717",
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/game/_/gameId/401869717/flyers-penguins",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=hockey&leagueAbbrev=nhl&gameId=401869717",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "atVs": "@",
      "gameDate": "2026-04-19T00:00:00.000+00:00",
      "score": "3-2",
      "homeTeamId": "16",
      "awayTeamId": "15",
      "homeTeamScore": "2",
      "awayTeamScore": "3",
      "gameResult": "W",
      "opponent": {
        "id": "16",
        "uid": "s:70~l:90~t:16",
        "displayName": "Pittsburgh Penguins",
        "abbreviation": "PIT",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/pit/pittsburgh-penguins",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=16",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/pit.png"
      },
      "leagueName": "National Hockey League",
      "leagueAbbreviation": "NHL",
      "leagueShortName": "NHL",
      "eventNote": "East 1st Round - Game 1",
      "team": {
        "id": "15",
        "uid": "s:70~l:90~t:15",
        "abbreviation": "PHI",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=15",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
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
    "401803635": {
      "id": "401803635",
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/game/_/gameId/401803635/hurricanes-flyers",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=hockey&leagueAbbrev=nhl&gameId=401803635",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "atVs": "vs",
      "gameDate": "2026-04-13T23:00:00.000+00:00",
      "score": "3-2 SO",
      "homeTeamId": "15",
      "awayTeamId": "7",
      "homeTeamScore": "3",
      "awayTeamScore": "2",
      "gameResult": "W",
      "opponent": {
        "id": "7",
        "uid": "s:70~l:90~t:7",
        "displayName": "Carolina Hurricanes",
        "abbreviation": "CAR",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/car/carolina-hurricanes",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=7",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/car.png"
      },
      "leagueName": "National Hockey League",
      "leagueAbbreviation": "NHL",
      "leagueShortName": "NHL",
      "team": {
        "id": "15",
        "uid": "s:70~l:90~t:15",
        "abbreviation": "PHI",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=15",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
        "isAllStar": false
      },
      "type": {
        "id": "1",
        "text": "Standard",
        "abbreviation": "STD",
        "type": "standard",
        "slug": "standard"
      }
    },
    "401803623": {
      "id": "401803623",
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/game/_/gameId/401803623/flyers-jets",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=hockey&leagueAbbrev=nhl&gameId=401803623",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "atVs": "@",
      "gameDate": "2026-04-11T23:00:00.000+00:00",
      "score": "7-1",
      "homeTeamId": "28",
      "awayTeamId": "15",
      "homeTeamScore": "1",
      "awayTeamScore": "7",
      "gameResult": "W",
      "opponent": {
        "id": "28",
        "uid": "s:70~l:90~t:28",
        "displayName": "Winnipeg Jets",
        "abbreviation": "WPG",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/wpg/winnipeg-jets",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=28",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/wpg.png"
      },
      "leagueName": "National Hockey League",
      "leagueAbbreviation": "NHL",
      "leagueShortName": "NHL",
      "team": {
        "id": "15",
        "uid": "s:70~l:90~t:15",
        "abbreviation": "PHI",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=15",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
        "isAllStar": false
      },
      "type": {
        "id": "1",
        "text": "Standard",
        "abbreviation": "STD",
        "type": "standard",
        "slug": "standard"
      }
    },
    "401803601": {
      "id": "401803601",
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/game/_/gameId/401803601/flyers-red-wings",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=hockey&leagueAbbrev=nhl&gameId=401803601",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "atVs": "@",
      "gameDate": "2026-04-09T23:00:00.000+00:00",
      "score": "6-3",
      "homeTeamId": "5",
      "awayTeamId": "15",
      "homeTeamScore": "6",
      "awayTeamScore": "3",
      "gameResult": "L",
      "opponent": {
        "id": "5",
        "uid": "s:70~l:90~t:5",
        "displayName": "Detroit Red Wings",
        "abbreviation": "DET",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/det/detroit-red-wings",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=5",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/det.png"
      },
      "leagueName": "National Hockey League",
      "leagueAbbreviation": "NHL",
      "leagueShortName": "NHL",
      "team": {
        "id": "15",
        "uid": "s:70~l:90~t:15",
        "abbreviation": "PHI",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=15",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
        "isAllStar": false
      },
      "type": {
        "id": "1",
        "text": "Standard",
        "abbreviation": "STD",
        "type": "standard",
        "slug": "standard"
      }
    },
    "401803587": {
      "id": "401803587",
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/game/_/gameId/401803587/flyers-devils",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=hockey&leagueAbbrev=nhl&gameId=401803587",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "atVs": "@",
      "gameDate": "2026-04-07T23:00:00.000+00:00",
      "score": "5-1",
      "homeTeamId": "11",
      "awayTeamId": "15",
      "homeTeamScore": "1",
      "awayTeamScore": "5",
      "gameResult": "W",
      "opponent": {
        "id": "11",
        "uid": "s:70~l:90~t:11",
        "displayName": "New Jersey Devils",
        "abbreviation": "NJ",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/nj/new-jersey-devils",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=11",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/nj.png"
      },
      "leagueName": "National Hockey League",
      "leagueAbbreviation": "NHL",
      "leagueShortName": "NHL",
      "team": {
        "id": "15",
        "uid": "s:70~l:90~t:15",
        "abbreviation": "PHI",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=15",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
        "isAllStar": false
      },
      "type": {
        "id": "1",
        "text": "Standard",
        "abbreviation": "STD",
        "type": "standard",
        "slug": "standard"
      }
    },
    "401803575": {
      "id": "401803575",
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/game/_/gameId/401803575/bruins-flyers",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=hockey&leagueAbbrev=nhl&gameId=401803575",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "atVs": "vs",
      "gameDate": "2026-04-05T19:30:00.000+00:00",
      "score": "2-1 OT",
      "homeTeamId": "15",
      "awayTeamId": "1",
      "homeTeamScore": "2",
      "awayTeamScore": "1",
      "gameResult": "W",
      "opponent": {
        "id": "1",
        "uid": "s:70~l:90~t:1",
        "displayName": "Boston Bruins",
        "abbreviation": "BOS",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/bos/boston-bruins",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=1",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/bos.png"
      },
      "leagueName": "National Hockey League",
      "leagueAbbreviation": "NHL",
      "leagueShortName": "NHL",
      "team": {
        "id": "15",
        "uid": "s:70~l:90~t:15",
        "abbreviation": "PHI",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=15",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
        "isAllStar": false
      },
      "type": {
        "id": "1",
        "text": "Standard",
        "abbreviation": "STD",
        "type": "standard",
        "slug": "standard"
      }
    },
    "401803556": {
      "id": "401803556",
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/game/_/gameId/401803556/flyers-islanders",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=hockey&leagueAbbrev=nhl&gameId=401803556",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "atVs": "@",
      "gameDate": "2026-04-03T23:00:00.000+00:00",
      "score": "4-1",
      "homeTeamId": "12",
      "awayTeamId": "15",
      "homeTeamScore": "1",
      "awayTeamScore": "4",
      "gameResult": "W",
      "opponent": {
        "id": "12",
        "uid": "s:70~l:90~t:12",
        "displayName": "New York Islanders",
        "abbreviation": "NYI",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/nyi/new-york-islanders",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=12",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/nyi.png"
      },
      "leagueName": "National Hockey League",
      "leagueAbbreviation": "NHL",
      "leagueShortName": "NHL",
      "team": {
        "id": "15",
        "uid": "s:70~l:90~t:15",
        "abbreviation": "PHI",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=15",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
        "isAllStar": false
      },
      "type": {
        "id": "1",
        "text": "Standard",
        "abbreviation": "STD",
        "type": "standard",
        "slug": "standard"
      }
    },
    "401803547": {
      "id": "401803547",
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/game/_/gameId/401803547/red-wings-flyers",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=hockey&leagueAbbrev=nhl&gameId=401803547",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "atVs": "vs",
      "gameDate": "2026-04-02T23:00:00.000+00:00",
      "score": "4-2",
      "homeTeamId": "15",
      "awayTeamId": "5",
      "homeTeamScore": "2",
      "awayTeamScore": "4",
      "gameResult": "L",
      "opponent": {
        "id": "5",
        "uid": "s:70~l:90~t:5",
        "displayName": "Detroit Red Wings",
        "abbreviation": "DET",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/det/detroit-red-wings",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=5",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/det.png"
      },
      "leagueName": "National Hockey League",
      "leagueAbbreviation": "NHL",
      "leagueShortName": "NHL",
      "team": {
        "id": "15",
        "uid": "s:70~l:90~t:15",
        "abbreviation": "PHI",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=15",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
        "isAllStar": false
      },
      "type": {
        "id": "1",
        "text": "Standard",
        "abbreviation": "STD",
        "type": "standard",
        "slug": "standard"
      }
    },
    "401803535": {
      "id": "401803535",
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/game/_/gameId/401803535/flyers-capitals",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=hockey&leagueAbbrev=nhl&gameId=401803535",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "atVs": "@",
      "gameDate": "2026-03-31T23:00:00.000+00:00",
      "score": "6-4",
      "homeTeamId": "23",
      "awayTeamId": "15",
      "homeTeamScore": "6",
      "awayTeamScore": "4",
      "gameResult": "L",
      "opponent": {
        "id": "23",
        "uid": "s:70~l:90~t:23",
        "displayName": "Washington Capitals",
        "abbreviation": "WSH",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/wsh/washington-capitals",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=23",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/wsh.png"
      },
      "leagueName": "National Hockey League",
      "leagueAbbreviation": "NHL",
      "leagueShortName": "NHL",
      "team": {
        "id": "15",
        "uid": "s:70~l:90~t:15",
        "abbreviation": "PHI",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=15",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
        "isAllStar": false
      },
      "type": {
        "id": "1",
        "text": "Standard",
        "abbreviation": "STD",
        "type": "standard",
        "slug": "standard"
      }
    },
    "401803523": {
      "id": "401803523",
      "links": [
        {
          "language": "en",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/game/_/gameId/401803523/stars-flyers",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=hockey&leagueAbbrev=nhl&gameId=401803523",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "atVs": "vs",
      "gameDate": "2026-03-29T23:00:00.000+00:00",
      "score": "2-1 OT",
      "homeTeamId": "15",
      "awayTeamId": "9",
      "homeTeamScore": "2",
      "awayTeamScore": "1",
      "gameResult": "W",
      "opponent": {
        "id": "9",
        "uid": "s:70~l:90~t:9",
        "displayName": "Dallas Stars",
        "abbreviation": "DAL",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/dal/dallas-stars",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=9",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/dal.png"
      },
      "leagueName": "National Hockey League",
      "leagueAbbreviation": "NHL",
      "leagueShortName": "NHL",
      "team": {
        "id": "15",
        "uid": "s:70~l:90~t:15",
        "abbreviation": "PHI",
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=15",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
        "isAllStar": false
      },
      "type": {
        "id": "1",
        "text": "Standard",
        "abbreviation": "STD",
        "type": "standard",
        "slug": "standard"
      }
    }
  },
  "seasonTypes": [
    {
      "displayName": "2025-26 Postseason",
      "displayTeam": "PHI",
      "categories": [
        {
          "displayName": "Postseason",
          "type": "event",
          "splitType": "3",
          "events": [
            {},
            {}
          ],
          "totals": [
            "2",
            "3"
          ]
        }
      ],
      "summary": {
        "displayName": "Postseason",
        "stats": [
          {
            "displayName": "Totals",
            "stats": [],
            "type": "total"
          }
        ]
      }
    },
    {
      "displayName": "2025-26 Regular Season",
      "displayTeam": "PHI",
      "categories": [
        {
          "displayName": "April",
          "type": "event",
          "splitType": "april",
          "events": [
            {},
            {}
          ],
          "totals": [
            "3",
            "4"
          ]
        },
        {
          "displayName": "March",
          "type": "event",
          "splitType": "march",
          "events": [
            {},
            {}
          ],
          "totals": [
            "2",
            "8"
          ]
        }
      ],
      "summary": {
        "displayName": "Regular Season Stats",
        "stats": [
          {
            "displayName": "Totals",
            "stats": [],
            "type": "total"
          }
        ]
      }
    }
  ],
  "glossary": [
    {
      "abbreviation": "+/-",
      "displayName": "Plus/Minus Rating"
    },
    {
      "abbreviation": "A",
      "displayName": "Assists"
    }
  ]
}
```
