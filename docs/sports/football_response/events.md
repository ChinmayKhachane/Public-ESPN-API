# ESPN Core V2 API Response Schemas (Events)

> Example JSON response structures for the most commonly used endpoints.  
> All responses are truncated for brevity — actual responses contain more fields.

---

## Events 

`https://sports.core.api.espn.com/v2/sports/{sport}/leagues/{league}/events`

```json
{
  "count": 335,
  "pageIndex": 1,
  "pageSize": 25,
  "pageCount": 14,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671834?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671836?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671827?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671840?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671844?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671826?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671837?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671831?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671841?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671828?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671838?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671839?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671830?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671833?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671845?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671843?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671878?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671879?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671881?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671880?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671883?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671882?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671884?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671885?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671937?lang=en&region=us"
    }
  ]
}
```

---

## Specific Event

`GET https://sports.core.api.espn.com/v2/sports/{sport}/leagues/{league}/events/{id}`

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988?lang=en&region=us",
  "id": "401772988",
  "uid": "s:20~l:28~e:401772988",
  "date": "2026-02-08T23:30Z",
  "name": "Seattle Seahawks at New England Patriots",
  "shortName": "SEA VS NE",
  "season": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025?lang=en&region=us"
  },
  "seasonType": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/types/3?lang=en&region=us"
  },
  "week": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/types/3/weeks/5?lang=en&region=us"
  },
  "timeValid": true,
  "competitions": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988?lang=en&region=us",
      "id": "401772988",
      "guid": "db65e2af-b39e-3d1a-9ccd-ea444480b1df",
      "uid": "s:20~l:28~e:401772988~c:401772988",
      "date": "2026-02-08T23:30Z",
      "attendance": 70823,
      "type": {
        "id": "1",
        "text": "Standard",
        "abbreviation": "STD",
        "slug": "standard",
        "type": "standard"
      },
      "timeValid": true,
      "dateValid": true,
      "neutralSite": true,
      "divisionCompetition": false,
      "conferenceCompetition": false,
      "previewAvailable": false,
      "recapAvailable": false,
      "boxscoreAvailable": true,
      "lineupAvailable": false,
      "gamecastAvailable": true,
      "playByPlayAvailable": true,
      "conversationAvailable": true,
      "commentaryAvailable": false,
      "pickcenterAvailable": true,
      "summaryAvailable": true,
      "liveAvailable": false,
      "ticketsAvailable": false,
      "shotChartAvailable": false,
      "timeoutsAvailable": false,
      "possessionArrowAvailable": false,
      "onWatchESPN": false,
      "recent": false,
      "bracketAvailable": false,
      "wallclockAvailable": false,
      "highlightsAvailable": true,
      "gameSource": {
        "id": "1",
        "description": "basic/manual",
        "state": "basic"
      },
      "boxscoreSource": {
        "id": "2",
        "description": "feed",
        "state": "full"
      },
      "playByPlaySource": {
        "id": "2",
        "description": "feed",
        "state": "full"
      },
      "linescoreSource": {
        "id": "1",
        "description": "basic/manual",
        "state": "basic"
      },
      "statsSource": {
        "id": "3",
        "description": "scrubbed",
        "state": "full"
      },
      "venue": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/venues/4738?lang=en&region=us",
        "id": "4738",
        "guid": "ad9d3113-9b26-3c9a-98a9-250109205ef9",
        "fullName": "Levi's Stadium",
        "address": {
          "city": "Santa Clara",
          "state": "CA",
          "zipCode": "95054",
          "country": "USA"
        },
        "grass": true,
        "indoor": false,
        "images": [
          {
            "href": "https://a.espncdn.com/i/venues/nfl/day/4738.jpg",
            "width": 2000,
            "height": 1125,
            "alt": "",
            "rel": [
              "full",
              "day"
            ]
          },
          {
            "href": "https://a.espncdn.com/i/venues/nfl/day/interior/4738.jpg",
            "width": 2000,
            "height": 1125,
            "alt": "",
            "rel": [
              "full",
              "day",
              "interior"
            ]
          }
        ]
      },
      "competitors": [
        {
          "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17?lang=en&region=us",
          "id": "17",
          "uid": "s:20~l:28~t:17",
          "type": "team",
          "order": 0,
          "homeAway": "home",
          "winner": false,
          "team": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/teams/17?lang=en&region=us"
          },
          "score": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17/score?lang=en&region=us"
          },
          "linescores": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17/linescores?lang=en&region=us"
          },
          "roster": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17/roster?lang=en&region=us"
          },
          "statistics": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17/statistics?lang=en&region=us"
          },
          "leaders": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17/leaders?lang=en&region=us"
          },
          "record": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/types/2/teams/17/record?lang=en&region=us"
          }
        },
        {
          "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/26?lang=en&region=us",
          "id": "26",
          "uid": "s:20~l:28~t:26",
          "type": "team",
          "order": 1,
          "homeAway": "away",
          "winner": true,
          "team": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/teams/26?lang=en&region=us"
          },
          "score": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/26/score?lang=en&region=us"
          },
          "linescores": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/26/linescores?lang=en&region=us"
          },
          "roster": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/26/roster?lang=en&region=us"
          },
          "statistics": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/26/statistics?lang=en&region=us"
          },
          "leaders": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/26/leaders?lang=en&region=us"
          },
          "record": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/types/2/teams/26/record?lang=en&region=us"
          }
        }
      ],
      "notes": [
        {
          "type": "event",
          "headline": "Super Bowl LX"
        }
      ],
      "situation": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/situation?lang=en&region=us"
      },
      "status": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/status?lang=en&region=us"
      },
      "odds": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/odds?lang=en&region=us"
      },
      "broadcasts": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/broadcasts?lang=en&region=us"
      },
      "officials": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/officials?lang=en&region=us"
      },
      "details": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/plays?lang=en&region=us"
      },
      "leaders": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/leaders?lang=en&region=us"
      },
      "links": [
        {
          "language": "en-US",
          "rel": [
            "summary",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/game/_/gameId/401772988/seahawks-patriots",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "summary",
            "sportscenter",
            "app",
            "event"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401772988",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "now",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/game/_/gameId/401772988",
          "text": "Now",
          "shortText": "Now",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "teamstats",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/matchup?gameId=401772988",
          "text": "Team Stats",
          "shortText": "Team Stats",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "boxscore",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401772988",
          "text": "Box Score",
          "shortText": "Box Score",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "boxscore",
            "sportscenter",
            "app",
            "event"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401772988",
          "text": "Box Score",
          "shortText": "Box Score",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "gamecast",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/game/_/gameId/401772988",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "gamecast",
            "mobile",
            "event"
          ],
          "href": "http://m.espn.com/nfl/gamecast?gameId=401772988&action=summary",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "gamecast",
            "sportscenter",
            "app",
            "event"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401772988",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "pbp",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401772988",
          "text": "Play-by-Play",
          "shortText": "Play-by-Play",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "videos",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/video?gameId=401772988",
          "text": "Videos",
          "shortText": "Videos",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "fantasy",
            "desktop",
            "event"
          ],
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2025",
          "text": "Play Fantasy Football",
          "shortText": "Play Fantasy Football",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "odds",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/odds/_/gameId/401772988",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "odds",
            "mobile",
            "event"
          ],
          "href": "https://m.espn.com/nfl/odds/_/gameId/401772988",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "predictor": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/predictor?lang=en&region=us"
      },
      "probabilities": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/probabilities?lang=en&region=us"
      },
      "powerIndexes": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/powerindex?lang=en&region=us"
      },
      "format": {
        "regulation": {
          "periods": 4,
          "displayName": "Quarter",
          "slug": "quarter",
          "clock": 900
        },
        "overtime": {
          "displayName": "sudden-death",
          "slug": "sudden-death",
          "clock": 900
        },
        "suddenDeath": {
          "periods": 0,
          "clock": 900
        }
      },
      "relevancy": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/relevancy?lang=en&region=us"
      },
      "drives": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/drives?lang=en&region=us"
      },
      "hasDefensiveStats": false
    }
  ],
  "links": [
    {
      "language": "en-US",
      "rel": [
        "summary",
        "desktop",
        "event"
      ],
      "href": "https://www.espn.com/nfl/game/_/gameId/401772988/seahawks-patriots",
      "text": "Gamecast",
      "shortText": "Summary",
      "isExternal": false,
      "isPremium": false
    },
    {
      "language": "en-US",
      "rel": [
        "summary",
        "sportscenter",
        "app",
        "event"
      ],
      "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401772988",
      "text": "Gamecast",
      "shortText": "Summary",
      "isExternal": false,
      "isPremium": false
    },
    {
      "language": "en-US",
      "rel": [
        "now",
        "desktop",
        "event"
      ],
      "href": "https://www.espn.com/nfl/game/_/gameId/401772988",
      "text": "Now",
      "shortText": "Now",
      "isExternal": false,
      "isPremium": false
    },
    {
      "language": "en-US",
      "rel": [
        "teamstats",
        "desktop",
        "event"
      ],
      "href": "https://www.espn.com/nfl/matchup?gameId=401772988",
      "text": "Team Stats",
      "shortText": "Team Stats",
      "isExternal": false,
      "isPremium": false
    },
    {
      "language": "en-US",
      "rel": [
        "boxscore",
        "desktop",
        "event"
      ],
      "href": "https://www.espn.com/nfl/boxscore/_/gameId/401772988",
      "text": "Box Score",
      "shortText": "Box Score",
      "isExternal": false,
      "isPremium": false
    },
    {
      "language": "en-US",
      "rel": [
        "boxscore",
        "sportscenter",
        "app",
        "event"
      ],
      "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401772988",
      "text": "Box Score",
      "shortText": "Box Score",
      "isExternal": false,
      "isPremium": false
    },
    {
      "language": "en-US",
      "rel": [
        "gamecast",
        "desktop",
        "event"
      ],
      "href": "https://www.espn.com/nfl/game/_/gameId/401772988",
      "text": "Gamecast",
      "shortText": "Gamecast",
      "isExternal": false,
      "isPremium": false
    },
    {
      "language": "en-US",
      "rel": [
        "gamecast",
        "mobile",
        "event"
      ],
      "href": "http://m.espn.com/nfl/gamecast?gameId=401772988&action=summary",
      "text": "Gamecast",
      "shortText": "Gamecast",
      "isExternal": false,
      "isPremium": false
    },
    {
      "language": "en-US",
      "rel": [
        "gamecast",
        "sportscenter",
        "app",
        "event"
      ],
      "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401772988",
      "text": "Gamecast",
      "shortText": "Gamecast",
      "isExternal": false,
      "isPremium": false
    },
    {
      "language": "en-US",
      "rel": [
        "pbp",
        "desktop",
        "event"
      ],
      "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401772988",
      "text": "Play-by-Play",
      "shortText": "Play-by-Play",
      "isExternal": false,
      "isPremium": false
    },
    {
      "language": "en-US",
      "rel": [
        "videos",
        "desktop",
        "event"
      ],
      "href": "https://www.espn.com/nfl/video?gameId=401772988",
      "text": "Videos",
      "shortText": "Videos",
      "isExternal": false,
      "isPremium": false
    },
    {
      "language": "en-US",
      "rel": [
        "fantasy",
        "desktop",
        "event"
      ],
      "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2025",
      "text": "Play Fantasy Football",
      "shortText": "Play Fantasy Football",
      "isExternal": false,
      "isPremium": false
    },
    {
      "language": "en-US",
      "rel": [
        "odds",
        "desktop",
        "event"
      ],
      "href": "https://www.espn.com/nfl/odds/_/gameId/401772988",
      "text": "Odds",
      "shortText": "Odds",
      "isExternal": false,
      "isPremium": false
    },
    {
      "language": "en-US",
      "rel": [
        "odds",
        "mobile",
        "event"
      ],
      "href": "https://m.espn.com/nfl/odds/_/gameId/401772988",
      "text": "Odds",
      "shortText": "Odds",
      "isExternal": false,
      "isPremium": false
    }
  ],
  "venues": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/venues/4738?lang=en&region=us"
    }
  ],
  "league": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl?lang=en&region=us"
  }
}
```
---

## Specific Event 