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
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671834?lang=en&region=us",
      "id": "401671834",
      "uid": "s:20~l:28~e:401671834",
      "date": "2025-01-04T21:30Z",
      "name": "Cleveland Browns at Baltimore Ravens",
      "shortName": "CLE @ BAL",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671834",
          "guid": "a20b6202-648c-374c-ad73-859aaf40cd81",
          "uid": "s:20~l:28~e:401671834~c:401671834",
          "date": "2025-01-04T21:30Z",
          "attendance": 70562,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "onWatchESPN": true,
          "recent": false,
          "bracketAvailable": false,
          "wallclockAvailable": true,
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
            "id": "3814",
            "guid": "664f4464-f7c1-3d07-b1e5-71dc5a42f362",
            "fullName": "M&T Bank Stadium",
            "address": {
              "city": "Baltimore",
              "state": "MD",
              "zipCode": "21230",
              "country": "USA"
            },
            "grass": true,
            "indoor": false,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3814.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3814.jpg",
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
              "id": "33",
              "uid": "s:20~l:28~t:33",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "5",
              "uid": "s:20~l:28~t:5",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671834/browns-ravens",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671834",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671834",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671834",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671834",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671834",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671834",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671834&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671834",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671834",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671834",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671834",
              "text": "Videos",
              "shortText": "Videos",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "watchespn",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/watchespn/index?gameId=401671834&sourceLang=en",
              "text": "WatchESPN",
              "shortText": "WatchESPN",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "watchespn",
                "app",
                "event"
              ],
              "href": "watchespn://showEvent?gameId=401671834&sourceLang=en",
              "text": "WatchESPN",
              "shortText": "WatchESPN",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671834",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671834",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "periods": 1,
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 600.0
            },
            "suddenDeath": {
              "periods": 1,
              "clock": 600.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671834/browns-ravens",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671834",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671834",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671834",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671834",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671834",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671834",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671834&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671834",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671834",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671834",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671834",
          "text": "Videos",
          "shortText": "Videos",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "watchespn",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/watchespn/index?gameId=401671834&sourceLang=en",
          "text": "WatchESPN",
          "shortText": "WatchESPN",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "watchespn",
            "app",
            "event"
          ],
          "href": "watchespn://showEvent?gameId=401671834&sourceLang=en",
          "text": "WatchESPN",
          "shortText": "WatchESPN",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671834",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671834",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671836?lang=en&region=us",
      "id": "401671836",
      "uid": "s:20~l:28~e:401671836",
      "date": "2025-01-05T01:00Z",
      "name": "Cincinnati Bengals at Pittsburgh Steelers",
      "shortName": "CIN @ PIT",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671836",
          "guid": "1cb7c458-96af-32df-82ee-5764dc79a166",
          "uid": "s:20~l:28~e:401671836~c:401671836",
          "date": "2025-01-05T01:00Z",
          "attendance": 65631,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "onWatchESPN": true,
          "recent": false,
          "bracketAvailable": false,
          "wallclockAvailable": true,
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
            "id": "3752",
            "guid": "cba1fd96-fceb-3ccd-89bc-ec61231f9bd9",
            "fullName": "Acrisure Stadium",
            "address": {
              "city": "Pittsburgh",
              "state": "PA",
              "zipCode": "15212",
              "country": "USA"
            },
            "grass": true,
            "indoor": false,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3752.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3752.jpg",
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
              "id": "23",
              "uid": "s:20~l:28~t:23",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "4",
              "uid": "s:20~l:28~t:4",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671836/bengals-steelers",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671836",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671836",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671836",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671836",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671836",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671836",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671836&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671836",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671836",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671836",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671836",
              "text": "Videos",
              "shortText": "Videos",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "watchespn",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/watchespn/index?gameId=401671836&sourceLang=en",
              "text": "WatchESPN",
              "shortText": "WatchESPN",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "watchespn",
                "app",
                "event"
              ],
              "href": "watchespn://showEvent?gameId=401671836&sourceLang=en",
              "text": "WatchESPN",
              "shortText": "WatchESPN",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671836",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671836",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "periods": 1,
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 600.0
            },
            "suddenDeath": {
              "periods": 1,
              "clock": 600.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671836/bengals-steelers",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671836",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671836",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671836",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671836",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671836",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671836",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671836&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671836",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671836",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671836",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671836",
          "text": "Videos",
          "shortText": "Videos",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "watchespn",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/watchespn/index?gameId=401671836&sourceLang=en",
          "text": "WatchESPN",
          "shortText": "WatchESPN",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "watchespn",
            "app",
            "event"
          ],
          "href": "watchespn://showEvent?gameId=401671836&sourceLang=en",
          "text": "WatchESPN",
          "shortText": "WatchESPN",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671836",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671836",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671827?lang=en&region=us",
      "id": "401671827",
      "uid": "s:20~l:28~e:401671827",
      "date": "2025-01-05T18:00Z",
      "name": "Carolina Panthers at Atlanta Falcons",
      "shortName": "CAR @ ATL",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671827",
          "guid": "264d1ff0-c4fe-3934-9f5e-ad8a113dbd25",
          "uid": "s:20~l:28~e:401671827~c:401671827",
          "date": "2025-01-05T18:00Z",
          "attendance": 69581,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "5348",
            "guid": "4194f4ed-0558-33dc-8494-98eee2a0e4fc",
            "fullName": "Mercedes-Benz Stadium",
            "address": {
              "city": "Atlanta",
              "state": "GA",
              "zipCode": "30313",
              "country": "USA"
            },
            "grass": false,
            "indoor": true,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/5348.jpg",
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
              "id": "1",
              "uid": "s:20~l:28~t:1",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "29",
              "uid": "s:20~l:28~t:29",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671827/panthers-falcons",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671827",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671827",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671827",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671827",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671827",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671827",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671827&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671827",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671827",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671827",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671827",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671827",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671827",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "periods": 1,
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 600.0
            },
            "suddenDeath": {
              "periods": 1,
              "clock": 600.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671827/panthers-falcons",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671827",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671827",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671827",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671827",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671827",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671827",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671827&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671827",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671827",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671827",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671827",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671827",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671827",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671840?lang=en&region=us",
      "id": "401671840",
      "uid": "s:20~l:28~e:401671840",
      "date": "2025-01-05T18:00Z",
      "name": "Washington Commanders at Dallas Cowboys",
      "shortName": "WSH @ DAL",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671840",
          "guid": "ffc43a4a-68e4-334a-8f42-080bbd3868f0",
          "uid": "s:20~l:28~e:401671840~c:401671840",
          "date": "2025-01-05T18:00Z",
          "attendance": 91349,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "3687",
            "guid": "a6d02626-ee51-3e35-98f1-2b05c0b9b2b3",
            "fullName": "AT&T Stadium",
            "address": {
              "city": "Arlington",
              "state": "TX",
              "zipCode": "76011",
              "country": "USA"
            },
            "grass": false,
            "indoor": true,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3687.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3687.jpg",
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
              "id": "6",
              "uid": "s:20~l:28~t:6",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "28",
              "uid": "s:20~l:28~t:28",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671840/commanders-cowboys",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671840",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671840",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671840",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671840",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671840",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671840",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671840&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671840",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671840",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671840",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671840",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671840",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671840",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "periods": 1,
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 600.0
            },
            "suddenDeath": {
              "periods": 1,
              "clock": 600.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671840/commanders-cowboys",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671840",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671840",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671840",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671840",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671840",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671840",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671840&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671840",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671840",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671840",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671840",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671840",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671840",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671844?lang=en&region=us",
      "id": "401671844",
      "uid": "s:20~l:28~e:401671844",
      "date": "2025-01-05T18:00Z",
      "name": "Chicago Bears at Green Bay Packers",
      "shortName": "CHI @ GB",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671844",
          "guid": "3fabed93-c44a-3250-bd11-c1690c711bde",
          "uid": "s:20~l:28~e:401671844~c:401671844",
          "date": "2025-01-05T18:00Z",
          "attendance": 77862,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "3798",
            "guid": "cbb0ff53-c186-3d41-91cf-410bc23baaef",
            "fullName": "Lambeau Field",
            "address": {
              "city": "Green Bay",
              "state": "WI",
              "zipCode": "54304",
              "country": "USA"
            },
            "grass": true,
            "indoor": false,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3798.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3798.jpg",
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
              "id": "9",
              "uid": "s:20~l:28~t:9",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "3",
              "uid": "s:20~l:28~t:3",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671844/bears-packers",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671844",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671844",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671844",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671844",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671844",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671844",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671844&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671844",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671844",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671844",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671844",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671844",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671844",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "periods": 1,
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 600.0
            },
            "suddenDeath": {
              "periods": 1,
              "clock": 600.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671844/bears-packers",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671844",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671844",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671844",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671844",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671844",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671844",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671844&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671844",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671844",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671844",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671844",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671844",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671844",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671826?lang=en&region=us",
      "id": "401671826",
      "uid": "s:20~l:28~e:401671826",
      "date": "2025-01-05T18:00Z",
      "name": "Houston Texans at Tennessee Titans",
      "shortName": "HOU @ TEN",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671826",
          "guid": "af2365fd-a305-3610-91bd-df321eeb5d0a",
          "uid": "s:20~l:28~e:401671826~c:401671826",
          "date": "2025-01-05T18:00Z",
          "attendance": 61815,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "3810",
            "guid": "2ce67339-d170-3f5a-af8f-85c1869b9e23",
            "fullName": "Nissan Stadium",
            "address": {
              "city": "Nashville",
              "state": "TN",
              "zipCode": "37213",
              "country": "USA"
            },
            "grass": false,
            "indoor": false,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3810.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3810.jpg",
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
              "id": "10",
              "uid": "s:20~l:28~t:10",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "34",
              "uid": "s:20~l:28~t:34",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671826/texans-titans",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671826",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671826",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671826",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671826",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671826",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671826",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671826&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671826",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671826",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671826",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671826",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671826",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671826",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "periods": 1,
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 600.0
            },
            "suddenDeath": {
              "periods": 1,
              "clock": 600.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671826/texans-titans",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671826",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671826",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671826",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671826",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671826",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671826",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671826&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671826",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671826",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671826",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671826",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671826",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671826",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671837?lang=en&region=us",
      "id": "401671837",
      "uid": "s:20~l:28~e:401671837",
      "date": "2025-01-05T18:00Z",
      "name": "Jacksonville Jaguars at Indianapolis Colts",
      "shortName": "JAX @ IND",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671837",
          "guid": "2549e9bb-f9e0-314d-9595-c24fef6a376e",
          "uid": "s:20~l:28~e:401671837~c:401671837",
          "date": "2025-01-05T18:00Z",
          "attendance": 63041,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "3812",
            "guid": "f636eae9-da36-3954-83d5-3b24995d4e59",
            "fullName": "Lucas Oil Stadium",
            "address": {
              "city": "Indianapolis",
              "state": "IN",
              "zipCode": "46225",
              "country": "USA"
            },
            "grass": false,
            "indoor": true,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3812.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3812.jpg",
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
              "id": "11",
              "uid": "s:20~l:28~t:11",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "30",
              "uid": "s:20~l:28~t:30",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671837/jaguars-colts",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671837",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671837",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671837",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671837",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671837",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671837",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671837&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671837",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671837",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671837",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671837",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671837",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671837",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "periods": 1,
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 600.0
            },
            "suddenDeath": {
              "periods": 1,
              "clock": 600.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671837/jaguars-colts",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671837",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671837",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671837",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671837",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671837",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671837",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671837&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671837",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671837",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671837",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671837",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671837",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671837",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671831?lang=en&region=us",
      "id": "401671831",
      "uid": "s:20~l:28~e:401671831",
      "date": "2025-01-05T18:00Z",
      "name": "Buffalo Bills at New England Patriots",
      "shortName": "BUF @ NE",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671831",
          "guid": "199eb426-123c-3df9-a5c7-682036d04838",
          "uid": "s:20~l:28~e:401671831~c:401671831",
          "date": "2025-01-05T18:00Z",
          "attendance": 64626,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "3738",
            "guid": "ed71e174-10c3-3e30-bc8d-4134bb06188e",
            "fullName": "Gillette Stadium",
            "address": {
              "city": "Foxborough",
              "state": "MA",
              "zipCode": "02035",
              "country": "USA"
            },
            "grass": false,
            "indoor": false,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3738.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3738.jpg",
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
              "id": "17",
              "uid": "s:20~l:28~t:17",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "2",
              "uid": "s:20~l:28~t:2",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671831/bills-patriots",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671831",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671831",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671831",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671831",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671831",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671831",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671831&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671831",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671831",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671831",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671831",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671831",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671831",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "periods": 1,
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 600.0
            },
            "suddenDeath": {
              "periods": 1,
              "clock": 600.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671831/bills-patriots",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671831",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671831",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671831",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671831",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671831",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671831",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671831&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671831",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671831",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671831",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671831",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671831",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671831",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671841?lang=en&region=us",
      "id": "401671841",
      "uid": "s:20~l:28~e:401671841",
      "date": "2025-01-05T18:00Z",
      "name": "New York Giants at Philadelphia Eagles",
      "shortName": "NYG @ PHI",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671841",
          "guid": "82f19f44-5b03-3d1b-ac8b-707544e8f934",
          "uid": "s:20~l:28~e:401671841~c:401671841",
          "date": "2025-01-05T18:00Z",
          "attendance": 69879,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "3806",
            "guid": "a1e68291-b8a7-3428-9754-7bfc568e5d71",
            "fullName": "Lincoln Financial Field",
            "address": {
              "city": "Philadelphia",
              "state": "PA",
              "zipCode": "19148",
              "country": "USA"
            },
            "grass": true,
            "indoor": false,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3806.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3806.jpg",
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
              "id": "21",
              "uid": "s:20~l:28~t:21",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "19",
              "uid": "s:20~l:28~t:19",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671841/giants-eagles",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671841",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671841",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671841",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671841",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671841",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671841",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671841&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671841",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671841",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671841",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671841",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671841",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671841",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "periods": 1,
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 600.0
            },
            "suddenDeath": {
              "periods": 1,
              "clock": 600.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671841/giants-eagles",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671841",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671841",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671841",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671841",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671841",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671841",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671841&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671841",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671841",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671841",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671841",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671841",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671841",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671828?lang=en&region=us",
      "id": "401671828",
      "uid": "s:20~l:28~e:401671828",
      "date": "2025-01-05T18:00Z",
      "name": "New Orleans Saints at Tampa Bay Buccaneers",
      "shortName": "NO @ TB",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671828",
          "guid": "d5174093-778d-3b60-b9c0-fbaff0ad72fa",
          "uid": "s:20~l:28~e:401671828~c:401671828",
          "date": "2025-01-05T18:00Z",
          "attendance": 63001,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "3886",
            "guid": "aab423e5-c4ef-3dd1-a6fb-cf7efdbe0be8",
            "fullName": "Raymond James Stadium",
            "address": {
              "city": "Tampa",
              "state": "FL",
              "zipCode": "33607",
              "country": "USA"
            },
            "grass": true,
            "indoor": false,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3886.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3886.jpg",
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
              "id": "27",
              "uid": "s:20~l:28~t:27",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "18",
              "uid": "s:20~l:28~t:18",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671828/saints-buccaneers",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671828",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671828",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671828",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671828",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671828",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671828",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671828&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671828",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671828",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671828",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671828",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671828",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671828",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "periods": 1,
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 600.0
            },
            "suddenDeath": {
              "periods": 1,
              "clock": 600.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671828/saints-buccaneers",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671828",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671828",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671828",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671828",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671828",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671828",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671828&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671828",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671828",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671828",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671828",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671828",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671828",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671838?lang=en&region=us",
      "id": "401671838",
      "uid": "s:20~l:28~e:401671838",
      "date": "2025-01-05T21:25Z",
      "name": "Kansas City Chiefs at Denver Broncos",
      "shortName": "KC @ DEN",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671838",
          "guid": "0724ccc0-d671-39fa-90c6-4e863dd69d10",
          "uid": "s:20~l:28~e:401671838~c:401671838",
          "date": "2025-01-05T21:25Z",
          "attendance": 76489,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "3937",
            "guid": "aac80886-8e5d-31b2-8514-8cc859c358d0",
            "fullName": "Empower Field at Mile High",
            "address": {
              "city": "Denver",
              "state": "CO",
              "zipCode": "80204",
              "country": "USA"
            },
            "grass": true,
            "indoor": false,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3937.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3937.jpg",
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
              "id": "7",
              "uid": "s:20~l:28~t:7",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "12",
              "uid": "s:20~l:28~t:12",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671838/chiefs-broncos",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671838",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671838",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671838",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671838",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671838",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671838",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671838&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671838",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671838",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671838",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671838",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671838",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671838",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "periods": 1,
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 600.0
            },
            "suddenDeath": {
              "periods": 1,
              "clock": 600.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671838/chiefs-broncos",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671838",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671838",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671838",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671838",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671838",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671838",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671838&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671838",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671838",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671838",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671838",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671838",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671838",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671839?lang=en&region=us",
      "id": "401671839",
      "uid": "s:20~l:28~e:401671839",
      "date": "2025-01-05T21:25Z",
      "name": "Los Angeles Chargers at Las Vegas Raiders",
      "shortName": "LAC @ LV",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671839",
          "guid": "0e6f8d6a-4173-3d2f-95da-d8fb7a5b7367",
          "uid": "s:20~l:28~e:401671839~c:401671839",
          "date": "2025-01-05T21:25Z",
          "attendance": 61352,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "6501",
            "guid": "850ed279-ceca-3fbd-911b-eece9b2c23fe",
            "fullName": "Allegiant Stadium",
            "address": {
              "city": "Las Vegas",
              "state": "NV",
              "zipCode": "89118",
              "country": "USA"
            },
            "grass": true,
            "indoor": true,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/6501.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/6501.jpg",
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
              "id": "13",
              "uid": "s:20~l:28~t:13",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "24",
              "uid": "s:20~l:28~t:24",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671839/chargers-raiders",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671839",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671839",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671839",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671839",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671839",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671839",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671839&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671839",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671839",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671839",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671839",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671839",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671839",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "periods": 1,
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 600.0
            },
            "suddenDeath": {
              "periods": 1,
              "clock": 600.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671839/chargers-raiders",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671839",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671839",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671839",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671839",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671839",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671839",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671839&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671839",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671839",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671839",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671839",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671839",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671839",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671830?lang=en&region=us",
      "id": "401671830",
      "uid": "s:20~l:28~e:401671830",
      "date": "2025-01-05T21:25Z",
      "name": "Seattle Seahawks at Los Angeles Rams",
      "shortName": "SEA @ LAR",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671830",
          "guid": "5319dc7d-8032-3654-bd0b-09c0e582de0a",
          "uid": "s:20~l:28~e:401671830~c:401671830",
          "date": "2025-01-05T21:25Z",
          "attendance": 72610,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "7065",
            "guid": "08c7d61d-9771-3b5b-ba66-80a2297ba4a8",
            "fullName": "SoFi Stadium",
            "address": {
              "city": "Inglewood",
              "state": "CA",
              "zipCode": "90301",
              "country": "USA"
            },
            "grass": false,
            "indoor": false,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/7065.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/7065.jpg",
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
              "id": "14",
              "uid": "s:20~l:28~t:14",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "26",
              "uid": "s:20~l:28~t:26",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671830/seahawks-rams",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671830",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671830",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671830",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671830",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671830",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671830",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671830&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671830",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671830",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671830",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671830",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671830",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671830",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "periods": 1,
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 600.0
            },
            "suddenDeath": {
              "periods": 1,
              "clock": 600.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671830/seahawks-rams",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671830",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671830",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671830",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671830",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671830",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671830",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671830&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671830",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671830",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671830",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671830",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671830",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671830",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671833?lang=en&region=us",
      "id": "401671833",
      "uid": "s:20~l:28~e:401671833",
      "date": "2025-01-05T21:25Z",
      "name": "Miami Dolphins at New York Jets",
      "shortName": "MIA @ NYJ",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671833",
          "guid": "15fb614b-5f0c-3c00-a79e-5fc5905a11b3",
          "uid": "s:20~l:28~e:401671833~c:401671833",
          "date": "2025-01-05T21:25Z",
          "attendance": 68818,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "3839",
            "guid": "604c24d0-2641-3553-82bc-1b005ac94765",
            "fullName": "MetLife Stadium",
            "address": {
              "city": "East Rutherford",
              "state": "NJ",
              "zipCode": "07073",
              "country": "USA"
            },
            "grass": false,
            "indoor": false,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3839.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3839.jpg",
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
              "id": "20",
              "uid": "s:20~l:28~t:20",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "15",
              "uid": "s:20~l:28~t:15",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671833/dolphins-jets",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671833",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671833",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671833",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671833",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671833",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671833",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671833&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671833",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671833",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671833",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671833",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671833",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671833",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "periods": 1,
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 600.0
            },
            "suddenDeath": {
              "periods": 1,
              "clock": 600.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671833/dolphins-jets",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671833",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671833",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671833",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671833",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671833",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671833",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671833&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671833",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671833",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671833",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671833",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671833",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671833",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671845?lang=en&region=us",
      "id": "401671845",
      "uid": "s:20~l:28~e:401671845",
      "date": "2025-01-05T21:25Z",
      "name": "San Francisco 49ers at Arizona Cardinals",
      "shortName": "SF @ ARI",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671845",
          "guid": "56f788fc-2bf8-36e3-92f6-598a082ddc5a",
          "uid": "s:20~l:28~e:401671845~c:401671845",
          "date": "2025-01-05T21:25Z",
          "attendance": 63849,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "3970",
            "guid": "e5462558-fb41-3902-bb35-1f76c0807d76",
            "fullName": "State Farm Stadium",
            "address": {
              "city": "Glendale",
              "state": "AZ",
              "zipCode": "85305",
              "country": "USA"
            },
            "grass": true,
            "indoor": true,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3970.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3970.jpg",
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
              "id": "22",
              "uid": "s:20~l:28~t:22",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "25",
              "uid": "s:20~l:28~t:25",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671845/49ers-cardinals",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671845",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671845",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671845",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671845",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671845",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671845",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671845&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671845",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671845",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671845",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671845",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671845",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671845",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "periods": 1,
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 600.0
            },
            "suddenDeath": {
              "periods": 1,
              "clock": 600.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671845/49ers-cardinals",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671845",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671845",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671845",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671845",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671845",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671845",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671845&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671845",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671845",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671845",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671845",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671845",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671845",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671843?lang=en&region=us",
      "id": "401671843",
      "uid": "s:20~l:28~e:401671843",
      "date": "2025-01-06T01:20Z",
      "name": "Minnesota Vikings at Detroit Lions",
      "shortName": "MIN @ DET",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671843",
          "guid": "d92d8935-6099-3c2e-9710-a97511c2dd8b",
          "uid": "s:20~l:28~e:401671843~c:401671843",
          "date": "2025-01-06T01:20Z",
          "attendance": 64774,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "3727",
            "guid": "3a82e243-5f64-3ce4-bdaf-49367fb55391",
            "fullName": "Ford Field",
            "address": {
              "city": "Detroit",
              "state": "MI",
              "zipCode": "48226",
              "country": "USA"
            },
            "grass": false,
            "indoor": true,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3727.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3727.jpg",
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
              "id": "8",
              "uid": "s:20~l:28~t:8",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "16",
              "uid": "s:20~l:28~t:16",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671843/vikings-lions",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671843",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671843",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671843",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671843",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671843",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671843",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671843&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671843",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671843",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671843",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671843",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671843",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671843",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "periods": 1,
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 600.0
            },
            "suddenDeath": {
              "periods": 1,
              "clock": 600.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671843/vikings-lions",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671843",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671843",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671843",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671843",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671843",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671843",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671843&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671843",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671843",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671843",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671843",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671843",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671843",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671878?lang=en&region=us",
      "id": "401671878",
      "uid": "s:20~l:28~e:401671878",
      "date": "2025-01-11T21:30Z",
      "name": "Los Angeles Chargers at Houston Texans",
      "shortName": "LAC @ HOU",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671878",
          "guid": "ab55e1f1-e151-31a8-97ad-1ba21885f321",
          "uid": "s:20~l:28~e:401671878~c:401671878",
          "date": "2025-01-11T21:30Z",
          "attendance": 71408,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "3891",
            "guid": "8a2b5070-4bbd-3dd3-bc00-e867aea0b5be",
            "fullName": "NRG Stadium",
            "address": {
              "city": "Houston",
              "state": "TX",
              "zipCode": "77054",
              "country": "USA"
            },
            "grass": false,
            "indoor": true,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3891.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3891.jpg",
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
              "id": "34",
              "uid": "s:20~l:28~t:34",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "24",
              "uid": "s:20~l:28~t:24",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [
            {
              "type": "event",
              "headline": "AFC Wild Card Playoffs"
            }
          ],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671878/chargers-texans",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671878",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671878",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671878",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671878",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671878",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671878",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671878&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671878",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671878",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671878",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671878",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671878",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671878",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 900.0
            },
            "suddenDeath": {
              "periods": 0,
              "clock": 900.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671878/chargers-texans",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671878",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671878",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671878",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671878",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671878",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671878",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671878&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671878",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671878",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671878",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671878",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671878",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671878",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671879?lang=en&region=us",
      "id": "401671879",
      "uid": "s:20~l:28~e:401671879",
      "date": "2025-01-12T01:00Z",
      "name": "Pittsburgh Steelers at Baltimore Ravens",
      "shortName": "PIT @ BAL",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671879",
          "guid": "237e3380-422e-32dc-9518-589b567c789e",
          "uid": "s:20~l:28~e:401671879~c:401671879",
          "date": "2025-01-12T01:00Z",
          "attendance": 70546,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "3814",
            "guid": "664f4464-f7c1-3d07-b1e5-71dc5a42f362",
            "fullName": "M&T Bank Stadium",
            "address": {
              "city": "Baltimore",
              "state": "MD",
              "zipCode": "21230",
              "country": "USA"
            },
            "grass": true,
            "indoor": false,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3814.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3814.jpg",
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
              "id": "33",
              "uid": "s:20~l:28~t:33",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "23",
              "uid": "s:20~l:28~t:23",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [
            {
              "type": "event",
              "headline": "AFC Wild Card Playoffs"
            }
          ],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671879/steelers-ravens",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671879",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671879",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671879",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671879",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671879",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671879",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671879&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671879",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671879",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671879",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671879",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671879",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671879",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 900.0
            },
            "suddenDeath": {
              "periods": 0,
              "clock": 900.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671879/steelers-ravens",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671879",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671879",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671879",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671879",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671879",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671879",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671879&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671879",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671879",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671879",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671879",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671879",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671879",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671881?lang=en&region=us",
      "id": "401671881",
      "uid": "s:20~l:28~e:401671881",
      "date": "2025-01-12T18:00Z",
      "name": "Denver Broncos at Buffalo Bills",
      "shortName": "DEN @ BUF",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671881",
          "guid": "24fb2b07-9bca-3863-a5de-e641ce88d4f4",
          "uid": "s:20~l:28~e:401671881~c:401671881",
          "date": "2025-01-12T18:00Z",
          "attendance": 70332,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "3883",
            "guid": "3c80cf14-6837-377a-b30d-1b9559f2f6f4",
            "fullName": "Highmark Stadium",
            "address": {
              "city": "Orchard Park",
              "state": "NY",
              "zipCode": "14127",
              "country": "USA"
            },
            "grass": false,
            "indoor": false,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3883.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3883.jpg",
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
              "id": "2",
              "uid": "s:20~l:28~t:2",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "7",
              "uid": "s:20~l:28~t:7",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [
            {
              "type": "event",
              "headline": "AFC Wild Card Playoffs"
            }
          ],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671881/broncos-bills",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671881",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671881",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671881",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671881",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671881",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671881",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671881&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671881",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671881",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671881",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671881",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671881",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671881",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 900.0
            },
            "suddenDeath": {
              "periods": 0,
              "clock": 900.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671881/broncos-bills",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671881",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671881",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671881",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671881",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671881",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671881",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671881&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671881",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671881",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671881",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671881",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671881",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671881",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671880?lang=en&region=us",
      "id": "401671880",
      "uid": "s:20~l:28~e:401671880",
      "date": "2025-01-12T21:30Z",
      "name": "Green Bay Packers at Philadelphia Eagles",
      "shortName": "GB @ PHI",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671880",
          "guid": "dfb1666c-3f57-36d6-a657-b76cb84e1355",
          "uid": "s:20~l:28~e:401671880~c:401671880",
          "date": "2025-01-12T21:30Z",
          "attendance": 69879,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "3806",
            "guid": "a1e68291-b8a7-3428-9754-7bfc568e5d71",
            "fullName": "Lincoln Financial Field",
            "address": {
              "city": "Philadelphia",
              "state": "PA",
              "zipCode": "19148",
              "country": "USA"
            },
            "grass": true,
            "indoor": false,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3806.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3806.jpg",
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
              "id": "21",
              "uid": "s:20~l:28~t:21",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "9",
              "uid": "s:20~l:28~t:9",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [
            {
              "type": "event",
              "headline": "NFC Wild Card Playoffs"
            }
          ],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671880/packers-eagles",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671880",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671880",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671880",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671880",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671880",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671880",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671880&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671880",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671880",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671880",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671880",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671880",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671880",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 900.0
            },
            "suddenDeath": {
              "periods": 0,
              "clock": 900.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671880/packers-eagles",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671880",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671880",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671880",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671880",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671880",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671880",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671880&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671880",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671880",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671880",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671880",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671880",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671880",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671883?lang=en&region=us",
      "id": "401671883",
      "uid": "s:20~l:28~e:401671883",
      "date": "2025-01-13T01:00Z",
      "name": "Washington Commanders at Tampa Bay Buccaneers",
      "shortName": "WSH @ TB",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671883",
          "guid": "ae7fd529-266c-33ce-8f97-808cf90dcaae",
          "uid": "s:20~l:28~e:401671883~c:401671883",
          "date": "2025-01-13T01:00Z",
          "attendance": 64614,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "3886",
            "guid": "aab423e5-c4ef-3dd1-a6fb-cf7efdbe0be8",
            "fullName": "Raymond James Stadium",
            "address": {
              "city": "Tampa",
              "state": "FL",
              "zipCode": "33607",
              "country": "USA"
            },
            "grass": true,
            "indoor": false,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3886.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3886.jpg",
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
              "id": "27",
              "uid": "s:20~l:28~t:27",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "28",
              "uid": "s:20~l:28~t:28",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [
            {
              "type": "event",
              "headline": "NFC Wild Card Playoffs"
            }
          ],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671883/commanders-buccaneers",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671883",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671883",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671883",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671883",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671883",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671883",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671883&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671883",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671883",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671883",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671883",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671883",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671883",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 900.0
            },
            "suddenDeath": {
              "periods": 0,
              "clock": 900.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671883/commanders-buccaneers",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671883",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671883",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671883",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671883",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671883",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671883",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671883&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671883",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671883",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671883",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671883",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671883",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671883",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671882?lang=en&region=us",
      "id": "401671882",
      "uid": "s:20~l:28~e:401671882",
      "date": "2025-01-14T01:00Z",
      "name": "Minnesota Vikings at Los Angeles Rams",
      "shortName": "MIN @ LAR",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671882",
          "guid": "06e893f3-2919-3742-ae5c-8ed8c5783c40",
          "uid": "s:20~l:28~e:401671882~c:401671882",
          "date": "2025-01-14T01:00Z",
          "attendance": 64515,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "onWatchESPN": true,
          "recent": false,
          "bracketAvailable": false,
          "wallclockAvailable": true,
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
            "id": "3970",
            "guid": "e5462558-fb41-3902-bb35-1f76c0807d76",
            "fullName": "State Farm Stadium",
            "address": {
              "city": "Glendale",
              "state": "AZ",
              "zipCode": "85305",
              "country": "USA"
            },
            "grass": true,
            "indoor": true,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3970.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3970.jpg",
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
              "id": "14",
              "uid": "s:20~l:28~t:14",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "16",
              "uid": "s:20~l:28~t:16",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [
            {
              "type": "event",
              "headline": "NFC Wild Card Playoffs - Moved to Arizona due to wildfires in Los Angeles"
            }
          ],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671882/vikings-rams",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671882",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671882",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671882",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671882",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671882",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671882",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671882&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671882",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671882",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671882",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671882",
              "text": "Videos",
              "shortText": "Videos",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "watchespn",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/watchespn/index?gameId=401671882&sourceLang=en",
              "text": "WatchESPN",
              "shortText": "WatchESPN",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "watchespn",
                "app",
                "event"
              ],
              "href": "watchespn://showEvent?gameId=401671882&sourceLang=en",
              "text": "WatchESPN",
              "shortText": "WatchESPN",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671882",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671882",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 900.0
            },
            "suddenDeath": {
              "periods": 0,
              "clock": 900.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671882/vikings-rams",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671882",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671882",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671882",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671882",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671882",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671882",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671882&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671882",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671882",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671882",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671882",
          "text": "Videos",
          "shortText": "Videos",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "watchespn",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/watchespn/index?gameId=401671882&sourceLang=en",
          "text": "WatchESPN",
          "shortText": "WatchESPN",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "watchespn",
            "app",
            "event"
          ],
          "href": "watchespn://showEvent?gameId=401671882&sourceLang=en",
          "text": "WatchESPN",
          "shortText": "WatchESPN",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671882",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671882",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671884?lang=en&region=us",
      "id": "401671884",
      "uid": "s:20~l:28~e:401671884",
      "date": "2025-01-18T21:30Z",
      "name": "Houston Texans at Kansas City Chiefs",
      "shortName": "HOU @ KC",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671884",
          "guid": "11f6c1bd-53b3-3a81-908c-e16f82b35343",
          "uid": "s:20~l:28~e:401671884~c:401671884",
          "date": "2025-01-18T21:30Z",
          "attendance": 73458,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "onWatchESPN": true,
          "recent": false,
          "bracketAvailable": false,
          "wallclockAvailable": true,
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
            "id": "3622",
            "guid": "abb5abf3-0c80-3f8d-a148-dc5cf8be2e7a",
            "fullName": "GEHA Field at Arrowhead Stadium",
            "address": {
              "city": "Kansas City",
              "state": "MO",
              "zipCode": "64129",
              "country": "USA"
            },
            "grass": true,
            "indoor": false,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3622.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3622.jpg",
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
              "id": "12",
              "uid": "s:20~l:28~t:12",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "34",
              "uid": "s:20~l:28~t:34",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [
            {
              "type": "event",
              "headline": "AFC Divisional Playoffs"
            }
          ],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671884/texans-chiefs",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671884",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671884",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671884",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671884",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671884",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671884",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671884&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671884",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671884",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671884",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671884",
              "text": "Videos",
              "shortText": "Videos",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "watchespn",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/watchespn/index?gameId=401671884&sourceLang=en",
              "text": "WatchESPN",
              "shortText": "WatchESPN",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "watchespn",
                "app",
                "event"
              ],
              "href": "watchespn://showEvent?gameId=401671884&sourceLang=en",
              "text": "WatchESPN",
              "shortText": "WatchESPN",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671884",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671884",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 900.0
            },
            "suddenDeath": {
              "periods": 0,
              "clock": 900.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671884/texans-chiefs",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671884",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671884",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671884",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671884",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671884",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671884",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671884&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671884",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671884",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671884",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671884",
          "text": "Videos",
          "shortText": "Videos",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "watchespn",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/watchespn/index?gameId=401671884&sourceLang=en",
          "text": "WatchESPN",
          "shortText": "WatchESPN",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "watchespn",
            "app",
            "event"
          ],
          "href": "watchespn://showEvent?gameId=401671884&sourceLang=en",
          "text": "WatchESPN",
          "shortText": "WatchESPN",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671884",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671884",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671885?lang=en&region=us",
      "id": "401671885",
      "uid": "s:20~l:28~e:401671885",
      "date": "2025-01-19T01:00Z",
      "name": "Washington Commanders at Detroit Lions",
      "shortName": "WSH @ DET",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671885",
          "guid": "7da0ff50-5bb1-39ef-95b1-0e595d5ad53f",
          "uid": "s:20~l:28~e:401671885~c:401671885",
          "date": "2025-01-19T01:00Z",
          "attendance": 64774,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "3727",
            "guid": "3a82e243-5f64-3ce4-bdaf-49367fb55391",
            "fullName": "Ford Field",
            "address": {
              "city": "Detroit",
              "state": "MI",
              "zipCode": "48226",
              "country": "USA"
            },
            "grass": false,
            "indoor": true,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3727.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3727.jpg",
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
              "id": "8",
              "uid": "s:20~l:28~t:8",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "28",
              "uid": "s:20~l:28~t:28",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [
            {
              "type": "event",
              "headline": "NFC Divisional Playoffs"
            }
          ],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671885/commanders-lions",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671885",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671885",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671885",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671885",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671885",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671885",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671885&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671885",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671885",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671885",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671885",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671885",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671885",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 900.0
            },
            "suddenDeath": {
              "periods": 0,
              "clock": 900.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671885/commanders-lions",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671885",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671885",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671885",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671885",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671885",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671885",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671885&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671885",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671885",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671885",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671885",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671885",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671885",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401671937?lang=en&region=us",
      "id": "401671937",
      "uid": "s:20~l:28~e:401671937",
      "date": "2025-01-19T20:00Z",
      "name": "Los Angeles Rams at Philadelphia Eagles",
      "shortName": "LAR @ PHI",
      "season": {},
      "seasonType": {},
      "week": {},
      "timeValid": true,
      "competitions": [
        {
          "id": "401671937",
          "guid": "2cd7c88e-13be-359a-b84a-e4bb4cebb05c",
          "uid": "s:20~l:28~e:401671937~c:401671937",
          "date": "2025-01-19T20:00Z",
          "attendance": 69879,
          "type": {
            "id": "1",
            "text": "Standard",
            "abbreviation": "STD",
            "slug": "standard",
            "type": "standard"
          },
          "timeValid": true,
          "dateValid": true,
          "neutralSite": false,
          "divisionCompetition": false,
          "conferenceCompetition": false,
          "previewAvailable": false,
          "recapAvailable": true,
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
          "wallclockAvailable": true,
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
            "id": "3806",
            "guid": "a1e68291-b8a7-3428-9754-7bfc568e5d71",
            "fullName": "Lincoln Financial Field",
            "address": {
              "city": "Philadelphia",
              "state": "PA",
              "zipCode": "19148",
              "country": "USA"
            },
            "grass": true,
            "indoor": false,
            "images": [
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/3806.jpg",
                "width": 2000,
                "height": 1125,
                "alt": "",
                "rel": [
                  "full",
                  "day"
                ]
              },
              {
                "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3806.jpg",
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
              "id": "21",
              "uid": "s:20~l:28~t:21",
              "type": "team",
              "order": 0,
              "homeAway": "home",
              "winner": true,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            },
            {
              "id": "14",
              "uid": "s:20~l:28~t:14",
              "type": "team",
              "order": 1,
              "homeAway": "away",
              "winner": false,
              "team": {},
              "score": {},
              "linescores": {},
              "roster": {},
              "statistics": {},
              "leaders": {},
              "record": {}
            }
          ],
          "notes": [
            {
              "type": "event",
              "headline": "NFC Divisional Playoffs"
            }
          ],
          "situation": {},
          "status": {},
          "odds": {},
          "broadcasts": {},
          "officials": {},
          "details": {},
          "leaders": {},
          "links": [
            {
              "language": "en-US",
              "rel": [
                "summary",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/game/_/gameId/401671937/rams-eagles",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671937",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671937",
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
              "href": "https://www.espn.com/nfl/matchup?gameId=401671937",
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
              "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671937",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671937",
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
              "href": "https://www.espn.com/nfl/game/_/gameId/401671937",
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
              "href": "http://m.espn.com/nfl/gamecast?gameId=401671937&action=summary",
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
              "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671937",
              "text": "Gamecast",
              "shortText": "Gamecast",
              "isExternal": false,
              "isPremium": false
            },
            {
              "language": "en-US",
              "rel": [
                "recap",
                "desktop",
                "event"
              ],
              "href": "https://www.espn.com/nfl/recap?gameId=401671937",
              "text": "Recap",
              "shortText": "Recap",
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
              "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671937",
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
              "href": "https://www.espn.com/nfl/video?gameId=401671937",
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
              "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
              "href": "https://www.espn.com/nfl/odds/_/gameId/401671937",
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
              "href": "https://m.espn.com/nfl/odds/_/gameId/401671937",
              "text": "Odds",
              "shortText": "Odds",
              "isExternal": false,
              "isPremium": false
            }
          ],
          "predictor": {},
          "probabilities": {},
          "powerIndexes": {},
          "format": {
            "regulation": {
              "periods": 4,
              "displayName": "Quarter",
              "slug": "quarter",
              "clock": 900.0
            },
            "overtime": {
              "displayName": "sudden-death",
              "slug": "sudden-death",
              "clock": 900.0
            },
            "suddenDeath": {
              "periods": 0,
              "clock": 900.0
            }
          },
          "relevancy": {},
          "drives": {},
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671937/rams-eagles",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671937",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671937",
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
          "href": "https://www.espn.com/nfl/matchup?gameId=401671937",
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
          "href": "https://www.espn.com/nfl/boxscore/_/gameId/401671937",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671937",
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
          "href": "https://www.espn.com/nfl/game/_/gameId/401671937",
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
          "href": "http://m.espn.com/nfl/gamecast?gameId=401671937&action=summary",
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
          "href": "sportscenter://x-callback-url/showGame?sportName=football&leagueAbbrev=nfl&gameId=401671937",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "recap",
            "desktop",
            "event"
          ],
          "href": "https://www.espn.com/nfl/recap?gameId=401671937",
          "text": "Recap",
          "shortText": "Recap",
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
          "href": "https://www.espn.com/nfl/playbyplay/_/gameId/401671937",
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
          "href": "https://www.espn.com/nfl/video?gameId=401671937",
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
          "href": "https://fantasy.espn.com/football/welcome?addata=nfl_gamecast_ffl2024",
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
          "href": "https://www.espn.com/nfl/odds/_/gameId/401671937",
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
          "href": "https://m.espn.com/nfl/odds/_/gameId/401671937",
          "text": "Odds",
          "shortText": "Odds",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "venues": [
        {}
      ],
      "league": {}
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
  "seasonType": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/types/3?lang=en&region=us",
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
  "week": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/types/3/weeks/5?lang=en&region=us",
    "number": 5,
    "startDate": "2026-02-04T08:00Z",
    "endDate": "2026-02-12T07:59Z",
    "text": "Super Bowl",
    "rankings": {},
    "events": {},
    "talentpicks": {},
    "qbr": {}
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
          "id": "17",
          "uid": "s:20~l:28~t:17",
          "type": "team",
          "order": 0,
          "homeAway": "home",
          "winner": false,
          "team": {},
          "score": {},
          "linescores": {},
          "roster": {},
          "statistics": {},
          "leaders": {},
          "record": {}
        },
        {
          "id": "26",
          "uid": "s:20~l:28~t:26",
          "type": "team",
          "order": 1,
          "homeAway": "away",
          "winner": true,
          "team": {},
          "score": {},
          "linescores": {},
          "roster": {},
          "statistics": {},
          "leaders": {},
          "record": {}
        }
      ],
      "notes": [
        {
          "type": "event",
          "headline": "Super Bowl LX"
        }
      ],
      "situation": {},
      "status": {},
      "odds": {},
      "broadcasts": {},
      "officials": {},
      "details": {},
      "leaders": {},
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
      "predictor": {},
      "probabilities": {},
      "powerIndexes": {},
      "format": {
        "regulation": {
          "periods": 4,
          "displayName": "Quarter",
          "slug": "quarter",
          "clock": 900.0
        },
        "overtime": {
          "displayName": "sudden-death",
          "slug": "sudden-death",
          "clock": 900.0
        },
        "suddenDeath": {
          "periods": 0,
          "clock": 900.0
        }
      },
      "relevancy": {},
      "drives": {},
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
    }
  ],
  "league": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl?lang=en&region=us",
    "id": "28",
    "guid": "ad4c3bd2-ddb6-3f8c-8abf-744855a08fa4",
    "uid": "s:20~l:28",
    "name": "National Football League",
    "displayName": "NFL",
    "abbreviation": "NFL",
    "shortName": "NFL",
    "slug": "nfl",
    "isTournament": false,
    "season": {
      "year": 2025,
      "startDate": "2025-07-31T07:00Z",
      "endDate": "2026-02-12T07:59Z",
      "displayName": "2025",
      "type": {
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
    "seasons": {},
    "franchises": {},
    "teams": {},
    "group": {},
    "groups": {},
    "events": {},
    "notes": {},
    "rankings": {},
    "draft": {},
    "awards": {},
    "links": [
      {
        "language": "en-US",
        "rel": [
          "index",
          "desktop",
          "league"
        ],
        "href": "https://www.espn.com/nfl/",
        "text": "Index",
        "shortText": "Index",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "index",
          "sportscenter",
          "app",
          "league"
        ],
        "href": "sportscenter://x-callback-url/showClubhouse?uid=s:20~l:28",
        "text": "Index",
        "shortText": "Index",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "schedule",
          "desktop",
          "league"
        ],
        "href": "https://www.espn.com/nfl/schedule",
        "text": "Schedule",
        "shortText": "Schedule",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "schedule",
          "sportscenter",
          "app",
          "league"
        ],
        "href": "sportscenter://x-callback-url/showClubhouse?uid=s:20~l:28&section=scores",
        "text": "Schedule",
        "shortText": "Schedule",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "standings",
          "desktop",
          "league"
        ],
        "href": "https://www.espn.com/nfl/standings",
        "text": "Standings",
        "shortText": "Standings",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "standings",
          "sportscenter",
          "app",
          "league"
        ],
        "href": "sportscenter://x-callback-url/showClubhouse?uid=s:20~l:28&section=standings",
        "text": "Standings",
        "shortText": "Standings",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "rankings",
          "desktop",
          "league"
        ],
        "href": "https://www.espn.com/nfl/powerrankings",
        "text": "Power Rankings",
        "shortText": "Power Rankings",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "scores",
          "desktop",
          "league"
        ],
        "href": "https://www.espn.com/nfl/scoreboard",
        "text": "Scores",
        "shortText": "Scores",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "scores",
          "sportscenter",
          "app",
          "league"
        ],
        "href": "sportscenter://x-callback-url/showClubhouse?uid=s:20~l:28&section=scores",
        "text": "Scores",
        "shortText": "Scores",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "stats",
          "desktop",
          "league"
        ],
        "href": "https://www.espn.com/nfl/stats",
        "text": "Stats",
        "shortText": "Stats",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "teams",
          "desktop",
          "league"
        ],
        "href": "https://www.espn.com/nfl/teams",
        "text": "Teams",
        "shortText": "Teams",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "athletes",
          "desktop",
          "league"
        ],
        "href": "https://www.espn.com/nfl/players",
        "text": "Players",
        "shortText": "Players",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "injuries",
          "desktop",
          "league"
        ],
        "href": "https://www.espn.com/nfl/injuries",
        "text": "Injuries",
        "shortText": "Injuries",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "odds",
          "desktop",
          "league"
        ],
        "href": "https://www.espn.com/nfl/odds",
        "text": "Odds",
        "shortText": "Odds",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "odds",
          "sportscenter",
          "app",
          "league"
        ],
        "href": "sportscenter://x-callback-url/showClubhouse?uid=s:20~l:28&section=odds",
        "text": "Odds",
        "shortText": "Odds",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "freeagency",
          "desktop",
          "league"
        ],
        "href": "https://insider.espn.com/nfl/freeagency/",
        "text": "Freeagency",
        "shortText": "Freeagency",
        "isExternal": false,
        "isPremium": true
      }
    ],
    "logos": [
      {
        "href": "https://a.espncdn.com/i/teamlogos/leagues/500/nfl.png",
        "width": 500,
        "height": 500,
        "alt": "",
        "rel": [
          "full",
          "default"
        ],
        "lastUpdated": "2018-06-05T12:07Z"
      },
      {
        "href": "https://a.espncdn.com/i/teamlogos/leagues/500-dark/nfl.png",
        "width": 500,
        "height": 500,
        "alt": "",
        "rel": [
          "full",
          "dark"
        ],
        "lastUpdated": "2024-07-22T16:53Z"
      }
    ],
    "athletes": {},
    "freeAgents": {},
    "calendar": {},
    "transactions": {},
    "talentPicks": {},
    "leaders": {},
    "gender": "MALE"
  }
}
```
---

## Specific Event 