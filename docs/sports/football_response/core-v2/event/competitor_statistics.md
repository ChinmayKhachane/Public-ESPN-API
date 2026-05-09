# Competition Competitor Statistics

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}/statistics

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988`, `competitor=17` on 2026-05-08.
- This is team-level game stats with athlete refs nested under each stat category.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17/statistics?lang=en&region=us",
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/teams/17?lang=en&region=us",
    "id": "17",
    "guid": "0078f353-fe3e-67ed-a42c-43cca0568e21",
    "uid": "s:20~l:28~t:17",
    "alternateIds": {
      "sdr": "8818"
    },
    "slug": "new-england-patriots",
    "location": "New England",
    "name": "Patriots",
    "nickname": "Patriots",
    "abbreviation": "NE",
    "displayName": "New England Patriots",
    "shortDisplayName": "Patriots",
    "color": "002a5c",
    "alternateColor": "c60c30",
    "isActive": true,
    "isAllStar": false,
    "logos": [
      {
        "href": "https://a.espncdn.com/i/teamlogos/nfl/500/ne.png",
        "width": 500,
        "height": 500,
        "alt": "",
        "rel": [
          "full",
          "default"
        ],
        "lastUpdated": "2024-06-25T18:54Z"
      },
      {
        "href": "https://a.espncdn.com/i/teamlogos/nfl/500-dark/ne.png",
        "width": 500,
        "height": 500,
        "alt": "",
        "rel": [
          "full",
          "dark"
        ],
        "lastUpdated": "2024-06-25T18:54Z"
      },
      {
        "href": "https://a.espncdn.com/i/teamlogos/nfl/500/scoreboard/ne.png",
        "width": 500,
        "height": 500,
        "alt": "",
        "rel": [
          "full",
          "scoreboard"
        ],
        "lastUpdated": "2024-06-25T18:54Z"
      },
      {
        "href": "https://a.espncdn.com/i/teamlogos/nfl/500-dark/scoreboard/ne.png",
        "width": 500,
        "height": 500,
        "alt": "",
        "rel": [
          "full",
          "scoreboard",
          "dark"
        ],
        "lastUpdated": "2024-06-25T18:54Z"
      },
      {
        "href": "https://a.espncdn.com/guid/0078f353-fe3e-67ed-a42c-43cca0568e21/logos/grayscale.png",
        "width": 500,
        "height": 500,
        "alt": "",
        "rel": [
          "full",
          "grayscale"
        ],
        "lastUpdated": "2026-03-31T12:53Z"
      },
      {
        "href": "https://a.espncdn.com/guid/0078f353-fe3e-67ed-a42c-43cca0568e21/logos/primary_logo_on_white_color.png",
        "width": 4096,
        "height": 4096,
        "alt": "",
        "rel": [
          "full",
          "primary_logo_on_white_color"
        ],
        "lastUpdated": "2026-02-13T03:12Z"
      },
      {
        "href": "https://a.espncdn.com/guid/0078f353-fe3e-67ed-a42c-43cca0568e21/logos/primary_logo_on_black_color.png",
        "width": 4096,
        "height": 4096,
        "alt": "",
        "rel": [
          "full",
          "primary_logo_on_black_color"
        ],
        "lastUpdated": "2026-02-13T03:12Z"
      },
      {
        "href": "https://a.espncdn.com/guid/0078f353-fe3e-67ed-a42c-43cca0568e21/logos/primary_logo_on_primary_color.png",
        "width": 4096,
        "height": 4096,
        "alt": "",
        "rel": [
          "full",
          "primary_logo_on_primary_color"
        ],
        "lastUpdated": "2026-02-13T03:12Z"
      },
      {
        "href": "https://a.espncdn.com/guid/0078f353-fe3e-67ed-a42c-43cca0568e21/logos/primary_logo_on_secondary_color.png",
        "width": 4096,
        "height": 4096,
        "alt": "",
        "rel": [
          "full",
          "primary_logo_on_secondary_color"
        ],
        "lastUpdated": "2026-02-13T03:12Z"
      },
      {
        "href": "https://a.espncdn.com/guid/0078f353-fe3e-67ed-a42c-43cca0568e21/logos/primary_logo_black.png",
        "width": 4096,
        "height": 4096,
        "alt": "",
        "rel": [
          "full",
          "primary_logo_black"
        ],
        "lastUpdated": "2026-02-13T03:12Z"
      },
      {
        "href": "https://a.espncdn.com/guid/0078f353-fe3e-67ed-a42c-43cca0568e21/logos/primary_logo_white.png",
        "width": 4096,
        "height": 4096,
        "alt": "",
        "rel": [
          "full",
          "primary_logo_white"
        ],
        "lastUpdated": "2026-02-13T03:12Z"
      },
      {
        "href": "https://a.espncdn.com/guid/0078f353-fe3e-67ed-a42c-43cca0568e21/logos/secondary_logo_on_white_color.png",
        "width": 4096,
        "height": 4096,
        "alt": "",
        "rel": [
          "full",
          "secondary_logo_on_white_color"
        ],
        "lastUpdated": "2026-02-13T03:12Z"
      },
      {
        "href": "https://a.espncdn.com/guid/0078f353-fe3e-67ed-a42c-43cca0568e21/logos/secondary_logo_on_black_color.png",
        "width": 4096,
        "height": 4096,
        "alt": "",
        "rel": [
          "full",
          "secondary_logo_on_black_color"
        ],
        "lastUpdated": "2026-02-13T03:12Z"
      },
      {
        "href": "https://a.espncdn.com/guid/0078f353-fe3e-67ed-a42c-43cca0568e21/logos/secondary_logo_on_primary_color.png",
        "width": 4096,
        "height": 4096,
        "alt": "",
        "rel": [
          "full",
          "secondary_logo_on_primary_color"
        ],
        "lastUpdated": "2026-02-13T03:12Z"
      },
      {
        "href": "https://a.espncdn.com/guid/0078f353-fe3e-67ed-a42c-43cca0568e21/logos/secondary_logo_on_secondary_color.png",
        "width": 4096,
        "height": 4096,
        "alt": "",
        "rel": [
          "full",
          "secondary_logo_on_secondary_color"
        ],
        "lastUpdated": "2026-02-13T03:12Z"
      },
      {
        "href": "https://a.espncdn.com/guid/0078f353-fe3e-67ed-a42c-43cca0568e21/logos/secondary_logo_black.png",
        "width": 4096,
        "height": 4096,
        "alt": "",
        "rel": [
          "full",
          "secondary_logo_black"
        ],
        "lastUpdated": "2026-02-13T03:12Z"
      },
      {
        "href": "https://a.espncdn.com/guid/0078f353-fe3e-67ed-a42c-43cca0568e21/logos/secondary_logo_white.png",
        "width": 4096,
        "height": 4096,
        "alt": "",
        "rel": [
          "full",
          "secondary_logo_white"
        ],
        "lastUpdated": "2026-02-13T03:12Z"
      }
    ],
    "record": {},
    "oddsRecords": {},
    "athletes": {},
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
    "groups": {},
    "ranks": {},
    "statistics": {},
    "leaders": {},
    "links": [
      {
        "language": "en-US",
        "rel": [
          "clubhouse",
          "desktop",
          "team"
        ],
        "href": "https://www.espn.com/nfl/team/_/name/ne/new-england-patriots",
        "text": "Clubhouse",
        "shortText": "Clubhouse",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "roster",
          "desktop",
          "team"
        ],
        "href": "https://www.espn.com/nfl/team/roster/_/name/ne/new-england-patriots",
        "text": "Roster",
        "shortText": "Roster",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "stats",
          "desktop",
          "team"
        ],
        "href": "https://www.espn.com/nfl/team/stats/_/name/ne/new-england-patriots",
        "text": "Statistics",
        "shortText": "Statistics",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "schedule",
          "desktop",
          "team"
        ],
        "href": "https://www.espn.com/nfl/team/schedule/_/name/ne",
        "text": "Schedule",
        "shortText": "Schedule",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "photos",
          "desktop",
          "team"
        ],
        "href": "https://www.espn.com/nfl/team/photos/_/name/ne",
        "text": "photos",
        "shortText": "photos",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "tickets",
          "desktop",
          "team"
        ],
        "href": "https://www.vividseats.com/new-england-patriots-tickets--sports-nfl-football/performer/592?wsUser=717",
        "text": "Tickets",
        "shortText": "Tickets",
        "isExternal": true,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "draftpicks",
          "desktop",
          "team"
        ],
        "href": "https://www.espn.com/nfl/draft/teams/_/name/ne/new-england-patriots",
        "text": "Draft Picks",
        "shortText": "Draft Picks",
        "isExternal": false,
        "isPremium": true
      },
      {
        "language": "en-US",
        "rel": [
          "transactions",
          "desktop",
          "team"
        ],
        "href": "https://www.espn.com/nfl/team/transactions/_/name/ne",
        "text": "Transactions",
        "shortText": "Transactions",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "injuries",
          "desktop",
          "team"
        ],
        "href": "https://www.espn.com/nfl/team/injuries/_/name/ne",
        "text": "Injuries",
        "shortText": "Injuries",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "depthchart",
          "desktop",
          "team"
        ],
        "href": "https://www.espn.com/nfl/team/depth/_/name/ne",
        "text": "Depth Chart",
        "shortText": "Depth Chart",
        "isExternal": false,
        "isPremium": false
      }
    ],
    "injuries": {},
    "notes": {},
    "againstTheSpreadRecords": {},
    "awards": {},
    "franchise": {},
    "depthCharts": {},
    "projection": {},
    "events": {},
    "transactions": {},
    "coaches": {}
  },
  "competition": {
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
  },
  "splits": {
    "categories": [
      {
        "name": "general",
        "displayName": "General"
      },
      {
        "name": "passing",
        "displayName": "Passing"
      },
      {
        "name": "rushing",
        "displayName": "Rushing"
      }
    ]
  }
}
```
