# Event ID

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}

Notes:
- Verified with `league=nfl`, `event=401772988` on 2026-05-08.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
|  |  |  |

## Example Response

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
  ]
}
```
