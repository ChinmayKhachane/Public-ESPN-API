# Competition Subresources

Verified with `league=nfl`, `event=401772988`, and `competition=401772988` on 2026-05-08.

---

## Broadcasts

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/broadcasts`

```json
{
  "count": 2,
  "pageIndex": 1,
  "pageSize": 25,
  "items": [
    {
      "type": {
        "id": "1",
        "shortName": "TV",
        "longName": "Television",
        "slug": "tv"
      },
      "station": "NBC",
      "slug": "nbc",
      "market": {
        "id": "1",
        "type": "National"
      },
      "media": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/media/379?lang=en&region=us",
        "id": "379",
        "callLetters": "NBC",
        "name": "NBC",
        "shortName": "NBC",
        "slug": "nbc",
        "logos": [
          {
            "href": "https://a.espncdn.com/guid/682b31a2-7a3d-39fe-8d3f-f95358fbafdf/logos/default.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [
              "full",
              "default"
            ],
            "lastUpdated": "2025-03-05T17:28Z"
          },
          {
            "href": "https://a.espncdn.com/guid/682b31a2-7a3d-39fe-8d3f-f95358fbafdf/logos/default-dark.png",
            "width": 500,
            "height": 500,
            "alt": "",
            "rel": [
              "full",
              "dark"
            ],
            "lastUpdated": "2025-03-05T17:28Z"
          }
        ]
      }
    }
  ]
}
```

---

## Competitor

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}`

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17?lang=en&region=us",
  "id": "17",
  "uid": "s:20~l:28~t:17",
  "type": "team",
  "order": 0,
  "homeAway": "home",
  "winner": false,
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
  "score": {
    "value": 13.0,
    "displayValue": "13",
    "winner": false,
    "source": {
      "id": "1",
      "description": "Basic/Manual"
    }
  },
  "statistics": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17/statistics/0?lang=en&region=us",
    "competition": {},
    "team": {},
    "splits": {
      "id": "0",
      "name": "All Splits",
      "abbreviation": "Any",
      "categories": [
        {
          "name": "general",
          "displayName": "General",
          "shortDisplayName": "General",
          "abbreviation": "gen",
          "summary": "",
          "stats": [
            {
              "name": "fumbles",
              "displayName": "Fumbles",
              "shortDisplayName": "F",
              "description": "The number of times a player/team has fumbled the ball",
              "abbreviation": "FUM",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "fumblesLost",
              "displayName": "Fumbles Lost",
              "shortDisplayName": "FL",
              "description": "The number of times a fumble is recovered by the opposing team",
              "abbreviation": "LST",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "fumblesForced",
              "displayName": "Forced Fumbles",
              "shortDisplayName": "Forced Fumbles",
              "description": "The total number of forced fumbles.",
              "abbreviation": "FF",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fumblesForcedPrimary",
              "displayName": "Forced Fumbles Primary",
              "shortDisplayName": "Forced Fumbles Primary",
              "description": "The number of forced fumbles excluding misc and special teams.",
              "abbreviation": "FFP",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fumblesRecovered",
              "displayName": "Fumbles Recovered",
              "shortDisplayName": "FR",
              "description": "The number of fumbles recovered.",
              "abbreviation": "FR",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fumblesRecoveredYards",
              "displayName": "Fumbles Recovered Yards",
              "shortDisplayName": "FRYDS",
              "description": "The number of yards gained after a fumble is recovered.",
              "abbreviation": "YDS",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fumblesTouchdowns",
              "displayName": "Fumbles Touchdowns",
              "shortDisplayName": "FTD",
              "description": "The number of times a fumbles is recovered and returned for a touchdown.",
              "abbreviation": "FTD",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "gamesPlayed",
              "displayName": "Games Played",
              "shortDisplayName": "GP",
              "description": "Games Played",
              "abbreviation": "GP",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "offensiveTwoPtReturns",
              "displayName": "Two Point Returns",
              "shortDisplayName": "2PTR",
              "description": "The number of two point attempts that is returned by the offense.",
              "abbreviation": "2PTR",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "offensiveFumblesTouchdowns",
              "displayName": "Fumbles Touchdowns",
              "shortDisplayName": "OFTD",
              "description": "The number of times a fumbles is recovered and returned for a touchdown by the offense.",
              "abbreviation": "OFTD",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "defensiveFumblesTouchdowns",
              "displayName": "Fumbles Touchdowns",
              "shortDisplayName": "DFTD",
              "description": "The number of times a fumbles is recovered and returned for a touchdown by the defense.",
              "abbreviation": "DFTD",
              "value": 0.0,
              "displayValue": "0"
            }
          ],
          "athletes": [
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            }
          ]
        },
        {
          "name": "passing",
          "displayName": "Passing",
          "shortDisplayName": "Passing",
          "abbreviation": "pass",
          "summary": "",
          "stats": [
            {
              "name": "avgGain",
              "displayName": "Average Gain",
              "shortDisplayName": "AG",
              "description": "The average gained yards per position.",
              "abbreviation": "AG",
              "value": 0.0,
              "displayValue": "0.0"
            },
            {
              "name": "completionPct",
              "displayName": "Completion Percentage",
              "shortDisplayName": "CMP%",
              "description": "The percentage of completed passes.",
              "abbreviation": "CMP%",
              "value": 62.791,
              "displayValue": "62.8"
            },
            {
              "name": "completions",
              "displayName": "Completions",
              "shortDisplayName": "CMP",
              "description": "The times a player completes a pass to another player who is eligible to catch a pass.",
              "abbreviation": "CMP",
              "value": 27.0,
              "displayValue": "27"
            },
            {
              "name": "ESPNQBRating",
              "displayName": "ESPN Quarterback Rating",
              "shortDisplayName": "ESPN Quarterback Rating",
              "description": "The quarterback rating used to compare overall offensive performance relative to other positions.",
              "abbreviation": "EQBR",
              "value": 375.0,
              "displayValue": "375.0"
            },
            {
              "name": "interceptionPct",
              "displayName": "Interception Percentage",
              "shortDisplayName": "INT%",
              "description": "The percentage of passes thrown that were intercepted by the opposing team.",
              "abbreviation": "INT%",
              "value": 4.65,
              "displayValue": "4.65"
            },
            {
              "name": "interceptions",
              "displayName": "Interceptions",
              "shortDisplayName": "INT",
              "description": "The number of passes thrown that were intercepted by the opposing team.",
              "abbreviation": "INT",
              "value": 2.0,
              "displayValue": "2"
            },
            {
              "name": "longPassing",
              "displayName": "Longest Pass",
              "shortDisplayName": "LONG",
              "description": "The longest pass play completed in terms of yards.",
              "abbreviation": "LNG",
              "value": 35.0,
              "displayValue": "35"
            },
            {
              "name": "miscYards",
              "displayName": "Miscellaneous Yards",
              "shortDisplayName": "MISC",
              "description": "The number of miscellaneous yards.",
              "abbreviation": "MISC",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "netPassingYards",
              "displayName": "Net Passing Yards",
              "shortDisplayName": "NYDS",
              "description": "The amount of net passing yards.",
              "abbreviation": "NYDS",
              "value": 252.0,
              "displayValue": "252"
            },
            {
              "name": "netPassingYardsPerGame",
              "displayName": "Net Passing Yards Per Game",
              "shortDisplayName": "NYDS/G",
              "description": "The amount of net passing yards per game.",
              "abbreviation": "NYDS/G",
              "value": 252.0,
              "displayValue": "252.0"
            },
            {
              "name": "netTotalYards",
              "displayName": "Net Total Yards",
              "shortDisplayName": "NTYDS",
              "description": "The amount of net total yards.",
              "abbreviation": "NTYDS",
              "value": 331.0,
              "displayValue": "331"
            },
            {
              "name": "netYardsPerGame",
              "displayName": "Net Yards Per Game",
              "shortDisplayName": "NTYDS/G",
              "description": "The amount of net total yards per game.",
              "abbreviation": "NTYDS/G",
              "value": 331.0,
              "displayValue": "331.0"
            },
            {
              "name": "passingAttempts",
              "displayName": "Passing Attempts",
              "shortDisplayName": "ATT",
              "description": "The number of times a pass is attempted.",
              "abbreviation": "ATT",
              "value": 43.0,
              "displayValue": "43"
            },
            {
              "name": "passingBigPlays",
              "displayName": "Passing Big Plays",
              "shortDisplayName": "BIG",
              "description": "The number of times a pass results in a big play.",
              "abbreviation": "BIGP",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "passingFirstDowns",
              "displayName": "Passing 1st downs",
              "shortDisplayName": "FIRST",
              "description": "The number of times a pass results in a first down.",
              "abbreviation": "FIRST",
              "value": 14.0,
              "displayValue": "14"
            },
            {
              "name": "passingFumbles",
              "displayName": "Passing Fumbles",
              "shortDisplayName": "F",
              "description": "The number of times the ball has been fumbled after a completed pass.",
              "abbreviation": "FUM",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "passingFumblesLost",
              "displayName": "Passing Fumbles Lost",
              "shortDisplayName": "FL",
              "description": "The number of times the ball has been fumbled and lost to the opposing team after a completed pass.",
              "abbreviation": "FL",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "passingTouchdownPct",
              "displayName": "Passing Touchdown %",
              "shortDisplayName": "TD%",
              "description": "The percentage of passes that result in a touchdown.",
              "abbreviation": "TD%",
              "value": 4.6511626,
              "displayValue": "4.65"
            },
            {
              "name": "passingTouchdowns",
              "displayName": "Passing Touchdowns",
              "shortDisplayName": "Touchdowns",
              "description": "The total number of passing touchdowns.",
              "abbreviation": "TD",
              "value": 2.0,
              "displayValue": "2"
            },
            {
              "name": "passingYards",
              "displayName": "Passing Yards",
              "shortDisplayName": "Pass Yards",
              "description": "The total passing yards.",
              "abbreviation": "YDS",
              "value": 295.0,
              "displayValue": "295"
            },
            {
              "name": "passingYardsAfterCatch",
              "displayName": "Passing Yards After Catch",
              "shortDisplayName": "PYAC",
              "description": "The amount of passing yards after catch.",
              "abbreviation": "PYAC",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "passingYardsAtCatch",
              "displayName": "Passing Yards At Catch",
              "shortDisplayName": "PY@C",
              "description": "The amount of passing yards when the catch is made.",
              "abbreviation": "PY@C",
              "value": 295.0,
              "displayValue": "295"
            },
            {
              "name": "passingYardsPerGame",
              "displayName": "Passing Yards Per Game",
              "shortDisplayName": "PYDS/G",
              "description": "The number of passing yards per game.",
              "abbreviation": "YDS/G",
              "value": 295.0,
              "displayValue": "295.0"
            },
            {
              "name": "QBRating",
              "displayName": "Passer Rating",
              "shortDisplayName": "RTG",
              "description": "Passer Rating",
              "abbreviation": "RTG",
              "value": 79.121376,
              "displayValue": "79.1"
            },
            {
              "name": "sacks",
              "displayName": "Total Sacks",
              "shortDisplayName": "Sacks",
              "description": "The total number of sacks.",
              "abbreviation": "SACK",
              "value": 6.0,
              "displayValue": "6"
            },
            {
              "name": "sackYardsLost",
              "displayName": "Sack Yards Lost",
              "shortDisplayName": "SYL",
              "description": "The amount of yards lost due to sacks.",
              "abbreviation": "SYL",
              "value": 43.0,
              "displayValue": "43"
            },
            {
              "name": "netPassingAttempts",
              "displayName": "Net Passing Attempts",
              "shortDisplayName": "NATT",
              "description": "Net Passing Attempts.",
              "abbreviation": "NATT",
              "value": 49.0,
              "displayValue": "49"
            },
            {
              "name": "teamGamesPlayed",
              "displayName": "Team Games Played",
              "shortDisplayName": "TGP",
              "description": "The amount of games played.",
              "abbreviation": "TGP",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "totalOffensivePlays",
              "displayName": "Total Offensive Plays",
              "shortDisplayName": "TOP",
              "description": "The number of total offensive plays.",
              "abbreviation": "TOP",
              "value": 67.0,
              "displayValue": "67"
            },
            {
              "name": "totalPoints",
              "displayName": "Total Points",
              "shortDisplayName": "TP",
              "description": "The number of total points scored.",
              "abbreviation": "PTS",
              "value": 13.0,
              "displayValue": "13"
            },
            {
              "name": "totalPointsPerGame",
              "displayName": "Total Points Per Game",
              "shortDisplayName": "TP/G",
              "description": "The number of points scored per game.",
              "abbreviation": "TP",
              "value": 13.0,
              "displayValue": "13.0"
            },
            {
              "name": "totalTouchdowns",
              "displayName": "Total Touchdowns",
              "shortDisplayName": "TTD",
              "description": "The number of touchdowns scored in total.",
              "abbreviation": "TD",
              "value": 2.0,
              "displayValue": "2"
            },
            {
              "name": "totalYards",
              "displayName": "Total Yards",
              "shortDisplayName": "TYDS",
              "description": "The number of yards accumulated in total.",
              "abbreviation": "TYDS",
              "value": 331.0,
              "displayValue": "331"
            },
            {
              "name": "totalYardsFromScrimmage",
              "displayName": "Total Yards From Scrimmage",
              "shortDisplayName": "SCRIM",
              "description": "The amound of total yards from the line of scrimmage.",
              "abbreviation": "CMP",
              "value": 374.0,
              "displayValue": "374"
            },
            {
              "name": "twoPointPassConvs",
              "displayName": "Two Point Pass Conversions",
              "shortDisplayName": "2PTPC",
              "description": "The number of times a 2-point conversion is successful with a pass.",
              "abbreviation": "2PTPC",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "twoPtPass",
              "displayName": "Two Point Pass",
              "shortDisplayName": "2PTP",
              "description": "The number of times a pass was thrown for 2 points.",
              "abbreviation": "2PTP",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "twoPtPassAttempts",
              "displayName": "Two Point Pass Attempts",
              "shortDisplayName": "2PTPA",
              "description": "The number of times a pass was attempted on a 2-point conversion.",
              "abbreviation": "2PTPA",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "yardsFromScrimmagePerGame",
              "displayName": "Yards From Scrimmage Per Game",
              "shortDisplayName": "SCRIM/G",
              "description": "The yards gained from scrimmage per game.",
              "abbreviation": "SCRIM/G",
              "value": 374.0,
              "displayValue": "374"
            },
            {
              "name": "yardsPerCompletion",
              "displayName": "Yards Per Completion",
              "shortDisplayName": "YDS/CMP",
              "description": "The average number of yards per pass completion.",
              "abbreviation": "CMP",
              "value": 9.333,
              "displayValue": "9.3"
            },
            {
              "name": "yardsPerGame",
              "displayName": "Yards Per Game",
              "shortDisplayName": "YDS/G",
              "description": "The average number of yards per game.",
              "abbreviation": "YDS/G",
              "value": 331.0,
              "displayValue": "331.0"
            },
            {
              "name": "yardsPerPassAttempt",
              "displayName": "Yards Per Pass Attempt",
              "shortDisplayName": "AVG",
              "description": "The average number yards per pass attempt.",
              "abbreviation": "AVG",
              "value": 6.86,
              "displayValue": "6.9"
            },
            {
              "name": "netYardsPerPassAttempt",
              "displayName": "Net Yards Per Pass Attempt",
              "shortDisplayName": "NYDS/PA",
              "description": "The average number net yards (taking sack yardage into account) per pass attempt.",
              "abbreviation": "NYDS/PA",
              "value": 5.142857142857143,
              "displayValue": "5.1"
            },
            {
              "name": "quarterbackRating",
              "displayName": "Quarterback Rating",
              "shortDisplayName": "Quarterback Rating",
              "description": "Quarterback Rating.",
              "abbreviation": "RAT",
              "value": 79.121376,
              "displayValue": "79.1"
            }
          ],
          "athletes": [
            {
              "athlete": {},
              "statistics": {}
            }
          ]
        },
        {
          "name": "rushing",
          "displayName": "Rushing",
          "shortDisplayName": "Rushing",
          "abbreviation": "rush",
          "summary": "",
          "stats": [
            {
              "name": "avgGain",
              "displayName": "Average Gain",
              "shortDisplayName": "AG",
              "description": "The average gained yards per position.",
              "abbreviation": "AG",
              "value": 0.0,
              "displayValue": "0.0"
            },
            {
              "name": "ESPNRBRating",
              "displayName": "ESPN RB Rating",
              "shortDisplayName": "ESPNRB",
              "description": "The ESPN Widereceiver Rating.",
              "abbreviation": "ESPNRB",
              "value": 64.0,
              "displayValue": "64.00"
            },
            {
              "name": "longRushing",
              "displayName": "Long Rushing",
              "shortDisplayName": "LONG",
              "description": "The amount of yards for the longest run.",
              "abbreviation": "LNG",
              "value": 16.0,
              "displayValue": "16"
            },
            {
              "name": "miscYards",
              "displayName": "Miscellaneous Yards",
              "shortDisplayName": "MISC",
              "description": "The amount of miscellaneous yards.",
              "abbreviation": "MISC",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "netTotalYards",
              "displayName": "Net Total Yards",
              "shortDisplayName": "NTYDS",
              "description": "The amount of total net yards.",
              "abbreviation": "NTYDS",
              "value": 331.0,
              "displayValue": "331"
            },
            {
              "name": "netYardsPerGame",
              "displayName": "Net Yards Per Game",
              "shortDisplayName": "NYDS/G",
              "description": "The total net yards gained per game.",
              "abbreviation": "CMP",
              "value": 331.0,
              "displayValue": "331.0"
            },
            {
              "name": "rushingAttempts",
              "displayName": "Rushing Attempts",
              "shortDisplayName": "Rushing Attempts",
              "description": "The total number of attempted rushes.",
              "abbreviation": "CAR",
              "value": 18.0,
              "displayValue": "18"
            },
            {
              "name": "rushingBigPlays",
              "displayName": "20+ Yard Rushing Plays",
              "shortDisplayName": "BIG",
              "description": "The number of rushing big plays gained.",
              "abbreviation": "BIG",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "rushingFirstDowns",
              "displayName": "Rushing 1st downs",
              "shortDisplayName": "FIRST",
              "description": "The number of times a first down is picked up by a run.",
              "abbreviation": "FD",
              "value": 2.0,
              "displayValue": "2"
            },
            {
              "name": "rushingFumbles",
              "displayName": "Rushing Fumbles",
              "shortDisplayName": "F",
              "description": "The number of times there is a run and then the ball is fumbled.",
              "abbreviation": "FUM",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "rushingFumblesLost",
              "displayName": "Rushing Fumbles Lost",
              "shortDisplayName": "FL",
              "description": "The number of times there is a run and then the ball is fumbled.",
              "abbreviation": "LST",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "rushingTouchdowns",
              "displayName": "Rushing Touchdowns",
              "shortDisplayName": "Rushing Touchdowns",
              "description": "The total number of rushing touchdowns.",
              "abbreviation": "TD",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "rushingYards",
              "displayName": "Rushing Yards",
              "shortDisplayName": "Rush Yards",
              "description": "The total rushing yards.",
              "abbreviation": "YDS",
              "value": 79.0,
              "displayValue": "79"
            },
            {
              "name": "rushingYardsPerGame",
              "displayName": "Rushing Yards Per Game",
              "shortDisplayName": "YDS/G",
              "description": "The average number of rushing yards per game.",
              "abbreviation": "YDS/G",
              "value": 79.0,
              "displayValue": "79.0"
            },
            {
              "name": "stuffs",
              "displayName": "Stuffs",
              "shortDisplayName": "STF",
              "description": "The number of times a run is stopped at or behind the line of scrimmage.",
              "abbreviation": "STF",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "stuffYardsLost",
              "displayName": "Stuff Yards Lost",
              "shortDisplayName": "SYDSL",
              "description": "The number of yards lost when a run is stuffed.",
              "abbreviation": "SYDSL",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "teamGamesPlayed",
              "displayName": "Team Games Played",
              "shortDisplayName": "GP",
              "description": "The numbers of team games played.",
              "abbreviation": "GP",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "totalOffensivePlays",
              "displayName": "Total Offensive Plays",
              "shortDisplayName": "OP",
              "description": "The number of offenseive plays run.",
              "abbreviation": "OP",
              "value": 67.0,
              "displayValue": "67"
            },
            {
              "name": "totalPoints",
              "displayName": "Total Points",
              "shortDisplayName": "TP",
              "description": "The number of total points earned.",
              "abbreviation": "PTS",
              "value": 13.0,
              "displayValue": "13"
            },
            {
              "name": "totalPointsPerGame",
              "displayName": "Total Points Per Game",
              "shortDisplayName": "TP/G",
              "description": "The average number of points scored per game.",
              "abbreviation": "TP/G",
              "value": 13.0,
              "displayValue": "13.0"
            },
            {
              "name": "totalTouchdowns",
              "displayName": "Total Touchdowns",
              "shortDisplayName": "TTD",
              "description": "The number of total Touchdowns scored.",
              "abbreviation": "TD",
              "value": 2.0,
              "displayValue": "2"
            },
            {
              "name": "totalYards",
              "displayName": "Total Yards",
              "shortDisplayName": "TYDS",
              "description": "The number of total yards gained.",
              "abbreviation": "TYDS",
              "value": 331.0,
              "displayValue": "331"
            },
            {
              "name": "totalYardsFromScrimmage",
              "displayName": "Total Yards From Scrimmage",
              "shortDisplayName": "SCRIM",
              "description": "The number of yards gained from the line of scrimmage.",
              "abbreviation": "SCRIM",
              "value": 374.0,
              "displayValue": "374"
            },
            {
              "name": "twoPointRushConvs",
              "displayName": "Two Point Rush Conversion",
              "shortDisplayName": "2PTC",
              "description": "The number of times a 2-point is converted with a run.",
              "abbreviation": "2PTC",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "twoPtRush",
              "displayName": "Two Point Rush",
              "shortDisplayName": "2PTR",
              "description": "The number of rushes on 2-point attempts.",
              "abbreviation": "2PTR",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "twoPtRushAttempts",
              "displayName": "Two Point Rush Attempts",
              "shortDisplayName": "2PTA",
              "description": "The number of times a 2-point conversion is attempted with a rush.",
              "abbreviation": "CMP",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "yardsFromScrimmagePerGame",
              "displayName": "Yards From Scrimmage Per Game",
              "shortDisplayName": "SCRIM/G",
              "description": "The average number of yards gained from scrimmage per game.",
              "abbreviation": "SCRIM/G",
              "value": 374.0,
              "displayValue": "374.0"
            },
            {
              "name": "yardsPerGame",
              "displayName": "Yards Per Game",
              "shortDisplayName": "YDS/G",
              "description": "The average number of yards per game.",
              "abbreviation": "YDS/G",
              "value": 331.0,
              "displayValue": "331.0"
            },
            {
              "name": "yardsPerRushAttempt",
              "displayName": "Yards Per Rush Attempt",
              "shortDisplayName": "YDS/RA",
              "description": "The average number of yards per rush attempt.",
              "abbreviation": "AVG",
              "value": 4.389,
              "displayValue": "4.4"
            }
          ],
          "athletes": [
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            }
          ]
        },
        {
          "name": "receiving",
          "displayName": "Receiving",
          "shortDisplayName": "Receiving",
          "abbreviation": "rec",
          "summary": "",
          "stats": [
            {
              "name": "avgGain",
              "displayName": "Average Gain",
              "shortDisplayName": "AG",
              "description": "The average gained yards per position.",
              "abbreviation": "AG",
              "value": 0.0,
              "displayValue": "0.0"
            },
            {
              "name": "ESPNWRRating",
              "displayName": "ESPN Widereceiver Rating",
              "shortDisplayName": "ESPNWR",
              "description": "The ESPN Widereceiver Rating.",
              "abbreviation": "ESPNWR",
              "value": 500.0,
              "displayValue": "500.00"
            },
            {
              "name": "longReception",
              "displayName": "Long Reception",
              "shortDisplayName": "LONG",
              "description": "The amount of yards for the longest reception.",
              "abbreviation": "LNG",
              "value": 35.0,
              "displayValue": "35"
            },
            {
              "name": "miscYards",
              "displayName": "Miscellaneous Yards",
              "shortDisplayName": "MISC",
              "description": "The amount of miscellaneous yards.",
              "abbreviation": "MISC",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "netTotalYards",
              "displayName": "Net Total Yards",
              "shortDisplayName": "NTYDS",
              "description": "The amount of total net yards.",
              "abbreviation": "NTYDS",
              "value": 331.0,
              "displayValue": "331"
            },
            {
              "name": "netYardsPerGame",
              "displayName": "Net Yards Per Game",
              "shortDisplayName": "NYDS/G",
              "description": "The total net yards gained per game.",
              "abbreviation": "CMP",
              "value": 331.0,
              "displayValue": "331.0"
            },
            {
              "name": "receivingBigPlays",
              "displayName": "20+ Yard Receiving Plays",
              "shortDisplayName": "BIG",
              "description": "The number of receiving big plays gained.",
              "abbreviation": "BIG",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "receivingFirstDowns",
              "displayName": "Receiving First Downs",
              "shortDisplayName": "FIRST",
              "description": "The number of times a first down is picked up by a reception.",
              "abbreviation": "FD",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "receivingFumbles",
              "displayName": "Receiving Fumbles",
              "shortDisplayName": "F",
              "description": "The number of times a reception is made and then the ball is fumbled.",
              "abbreviation": "FUM",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "receivingFumblesLost",
              "displayName": "Receiving Fumbles Lost",
              "shortDisplayName": "FL",
              "description": "The number of times a reception is made and the balled is fumbled and recovered by the opposing team.",
              "abbreviation": "LST",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "receivingTargets",
              "displayName": "Receiving Targets",
              "shortDisplayName": "TGT",
              "description": "The number of times a receiver is thrown to.",
              "abbreviation": "TGTS",
              "value": 41.0,
              "displayValue": "41"
            },
            {
              "name": "receivingTouchdowns",
              "displayName": "Receiving Touchdowns",
              "shortDisplayName": "Receiving Touchdowns",
              "description": "The total number of receiving touchdowns.",
              "abbreviation": "TD",
              "value": 2.0,
              "displayValue": "2"
            },
            {
              "name": "receivingYards",
              "displayName": "Receiving Yards",
              "shortDisplayName": "Rec. Yards",
              "description": "The total receiving yards.",
              "abbreviation": "YDS",
              "value": 295.0,
              "displayValue": "295"
            },
            {
              "name": "receivingYardsAfterCatch",
              "displayName": "Receiving Yards After Catch",
              "shortDisplayName": "YAC",
              "description": "The number of yards gained after the catch has been made.",
              "abbreviation": "YAC",
              "value": 129.0,
              "displayValue": "129"
            },
            {
              "name": "receivingYardsAtCatch",
              "displayName": "Receiving Yards At Catch",
              "shortDisplayName": "Y@C",
              "description": "The number of yards gained at the time the catch is made.",
              "abbreviation": "Y@C",
              "value": 166.0,
              "displayValue": "166"
            },
            {
              "name": "receivingYardsPerGame",
              "displayName": "Receiving Yards Per Game",
              "shortDisplayName": "YDS/G",
              "description": "The average number of receiving yards per game.",
              "abbreviation": "YDS/G",
              "value": 295.0,
              "displayValue": "295.0"
            },
            {
              "name": "receptions",
              "displayName": "Receptions",
              "shortDisplayName": "Receptions",
              "description": "The total number of receptions.",
              "abbreviation": "REC",
              "value": 27.0,
              "displayValue": "27"
            },
            {
              "name": "teamGamesPlayed",
              "displayName": "Team Games Played",
              "shortDisplayName": "GP",
              "description": "The numbers of team games played.",
              "abbreviation": "GP",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "totalOffensivePlays",
              "displayName": "Total Offensive Plays",
              "shortDisplayName": "OP",
              "description": "The number of offenseive plays run.",
              "abbreviation": "OP",
              "value": 67.0,
              "displayValue": "67"
            },
            {
              "name": "totalPoints",
              "displayName": "Total Points",
              "shortDisplayName": "TP",
              "description": "The number of total points earned.",
              "abbreviation": "PTS",
              "value": 13.0,
              "displayValue": "13"
            },
            {
              "name": "totalPointsPerGame",
              "displayName": "Total Points Per Game",
              "shortDisplayName": "TP/G",
              "description": "The average number of points scored per game.",
              "abbreviation": "TP/G",
              "value": 13.0,
              "displayValue": "13.0"
            },
            {
              "name": "totalTouchdowns",
              "displayName": "Total Touchdowns",
              "shortDisplayName": "TTD",
              "description": "The number of total Touchdowns scored.",
              "abbreviation": "TD",
              "value": 2.0,
              "displayValue": "2"
            },
            {
              "name": "totalYards",
              "displayName": "Total Yards",
              "shortDisplayName": "TYDS",
              "description": "The number of total yards gained.",
              "abbreviation": "TYDS",
              "value": 331.0,
              "displayValue": "331"
            },
            {
              "name": "totalYardsFromScrimmage",
              "displayName": "Total Yards From Scrimmage",
              "shortDisplayName": "SCRIM",
              "description": "The number of yards gained from the line of scrimmage.",
              "abbreviation": "SCRIM",
              "value": 374.0,
              "displayValue": "374"
            },
            {
              "name": "twoPointRecConvs",
              "displayName": "Two Point Receiving Conversion",
              "shortDisplayName": "2PTC",
              "description": "The number of times a 2-point is converted with a reception.",
              "abbreviation": "2PTC",
              "value": 41.0,
              "displayValue": "41"
            },
            {
              "name": "twoPtReception",
              "displayName": "Two Point Receptions",
              "shortDisplayName": "2PTR",
              "description": "The number of receptions on 2-point attempts.",
              "abbreviation": "2PTR",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "twoPtReceptionAttempts",
              "displayName": "Two Point Reception Attempts",
              "shortDisplayName": "2PTA",
              "description": "The number of times a 2-point conversion is attempted with a reception.",
              "abbreviation": "CMP",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "yardsFromScrimmagePerGame",
              "displayName": "Yards From Scrimmage Per Game",
              "shortDisplayName": "SCRIM/G",
              "description": "The average number of yards gained from scrimmage per game.",
              "abbreviation": "SCRIM/G",
              "value": 374.0,
              "displayValue": "374.0"
            },
            {
              "name": "yardsPerGame",
              "displayName": "Yards Per Game",
              "shortDisplayName": "YDS/G",
              "description": "The average number of yards per game.",
              "abbreviation": "YDS/G",
              "value": 331.0,
              "displayValue": "331.0"
            },
            {
              "name": "yardsPerReception",
              "displayName": "Yards Per Reception",
              "shortDisplayName": "YDS/R",
              "description": "The average number of yards per reception.",
              "abbreviation": "AVG",
              "value": 10.926,
              "displayValue": "10.9"
            }
          ],
          "athletes": [
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            }
          ]
        },
        {
          "name": "defensive",
          "displayName": "Defense",
          "shortDisplayName": "Defensive",
          "abbreviation": "def",
          "summary": "",
          "stats": [
            {
              "name": "assistTackles",
              "displayName": "Assist Tackles",
              "shortDisplayName": "AST",
              "description": "The number of assists on tackles.",
              "abbreviation": "AST",
              "value": 19.0,
              "displayValue": "19"
            },
            {
              "name": "avgInterceptionYards",
              "displayName": "Average Interception Yards",
              "shortDisplayName": "YDS/INT",
              "description": "The average number of return yards per interception.",
              "abbreviation": "AVG",
              "value": 0.0,
              "displayValue": "0.0"
            },
            {
              "name": "avgSackYards",
              "displayName": "Average Sack Yards",
              "shortDisplayName": "YDS/SACK",
              "description": "The average number of yards lost per sack.",
              "abbreviation": "YDS/SACK",
              "value": 8.0,
              "displayValue": "8.0"
            },
            {
              "name": "avgStuffYards",
              "displayName": "Average Stuff Yards",
              "shortDisplayName": "YDS/STF",
              "description": "The average number of yards lost per stuff.",
              "abbreviation": "YDS/STF",
              "value": 0.0,
              "displayValue": "0.0"
            },
            {
              "name": "blockedFieldGoalTouchdowns",
              "displayName": "Block Field Goal Touchdown",
              "shortDisplayName": "BFGTD",
              "description": "The number of blocked field goals that were returned for a touchdown.",
              "abbreviation": "BFGTD",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "blockedPuntTouchdowns",
              "displayName": "Blocked Punt Touchdowns",
              "shortDisplayName": "BPTD",
              "description": "The number of blocked punts that were returned for touchdowns.",
              "abbreviation": "BPTD",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "defensiveTouchdowns",
              "displayName": "Defensive Touchdown",
              "shortDisplayName": "TD",
              "description": "The number of touchdowns recorded on defense.",
              "abbreviation": "TD",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "hurries",
              "displayName": "Hurries",
              "shortDisplayName": "HUR",
              "description": "The number of times that the quarterback is hurried.",
              "abbreviation": "HUR",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "kicksBlocked",
              "displayName": "Kicks Blocked",
              "shortDisplayName": "KBLK",
              "description": "The number of times a kick was blocked.",
              "abbreviation": "KB",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "longInterception",
              "displayName": "Long Interception",
              "shortDisplayName": "LNGINT",
              "description": "The amount of yards of the longest interception return.",
              "abbreviation": "LNG",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "miscTouchdowns",
              "displayName": "Miscellaneous Touchdowns",
              "shortDisplayName": "MISCTD",
              "description": "The number of miscellaneous touchdowns.",
              "abbreviation": "MISCTD",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "passesBattedDown",
              "displayName": "Passes Batted Down",
              "shortDisplayName": "BATD",
              "description": "The number of passes that were batted down.",
              "abbreviation": "BATD",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "passesDefended",
              "displayName": "Passes Defended",
              "shortDisplayName": "Passes Defended",
              "description": "The total number of passes defended.",
              "abbreviation": "PD",
              "value": 7.0,
              "displayValue": "7"
            },
            {
              "name": "QBHits",
              "displayName": "Quarterback Hits",
              "shortDisplayName": "QB HTS",
              "description": "The times the quarterback is hit.",
              "abbreviation": "QB HTS",
              "value": 6.0,
              "displayValue": "6"
            },
            {
              "name": "twoPtReturns",
              "displayName": "Two Point Returns",
              "shortDisplayName": "2PTR",
              "description": "The number of two point attempts that is returned by the defense.",
              "abbreviation": "2PTR",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "sacks",
              "displayName": "Sacks",
              "shortDisplayName": "Sacks",
              "description": "The total number of sacks.",
              "abbreviation": "SACK",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "sacksAssisted",
              "displayName": "Sacks Assisted",
              "shortDisplayName": "SCKAST",
              "description": "The number of sacks that were assisted.",
              "abbreviation": "SCKAST",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "sacksUnassisted",
              "displayName": "Sacks Unassisted",
              "shortDisplayName": "UASACK",
              "description": "The number of recorded sacks that were unassisted.",
              "abbreviation": "UASACK",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "sackYards",
              "displayName": "Sack Yards",
              "shortDisplayName": "SCKYDS",
              "description": "The number of yards lost from sacks.",
              "abbreviation": "SCKYDS",
              "value": 8.0,
              "displayValue": "8"
            },
            {
              "name": "safeties",
              "displayName": "Safeties",
              "shortDisplayName": "SAFE",
              "description": "The number of times the offense is tackled in their own endzone.",
              "abbreviation": "SAFE",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "soloTackles",
              "displayName": "Solo Tackles",
              "shortDisplayName": "SOLO",
              "description": "The number of times a tackle was made unassisted.",
              "abbreviation": "SOLO",
              "value": 46.0,
              "displayValue": "46"
            },
            {
              "name": "stuffs",
              "displayName": "Stuffs",
              "shortDisplayName": "STF",
              "description": "The number of times that a runner is stuffed at or behind the line of scrimmage.",
              "abbreviation": "STF",
              "value": 7.0,
              "displayValue": "7"
            },
            {
              "name": "stuffYards",
              "displayName": "Stuff Yards",
              "shortDisplayName": "STFYDS",
              "description": "The number of yards lost from stuffs.",
              "abbreviation": "STFYDS",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "tacklesForLoss",
              "displayName": "Tackles For Loss",
              "shortDisplayName": "TFL",
              "description": "The number of tackles that result in a loss of yardage.",
              "abbreviation": "TFL",
              "value": 7.0,
              "displayValue": "7"
            },
            {
              "name": "tacklesYardsLost",
              "displayName": "Tackles Yards Lost",
              "shortDisplayName": "TYDSL",
              "description": "The number of yards lost from tackles for lost.",
              "abbreviation": "CMP",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "teamGamesPlayed",
              "displayName": "Team Games Played",
              "shortDisplayName": "GP",
              "description": "The numbers of team games played.",
              "abbreviation": "GP",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "totalTackles",
              "displayName": "Total Tackles",
              "shortDisplayName": "Tackles",
              "description": "The total number of tackles.",
              "abbreviation": "TOT",
              "value": 65.0,
              "displayValue": "65"
            },
            {
              "name": "yardsAllowed",
              "displayName": "Yards Allowed",
              "shortDisplayName": "YA",
              "description": "Yards allowed",
              "abbreviation": "YA",
              "value": 335.0,
              "displayValue": "335"
            },
            {
              "name": "pointsAllowed",
              "displayName": "Points Allowed",
              "shortDisplayName": "PA",
              "description": "Points allowed",
              "abbreviation": "PA",
              "value": 23.0,
              "displayValue": "23"
            },
            {
              "name": "onePtSafetiesMade",
              "displayName": "One Point Safeties Made",
              "shortDisplayName": "OPSM",
              "description": "One Point Safeties Made",
              "abbreviation": "OPSM",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "missedFieldGoalReturnTd",
              "displayName": "Missed Field Goal Return TD",
              "shortDisplayName": "MFGRTD",
              "description": "The number of Missed Field Goal Return TD.",
              "abbreviation": "MFGRTD",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "blockedPuntEzRecTd",
              "displayName": "Blocked Punt Ez Rec TD",
              "shortDisplayName": "BPERTD",
              "description": "The number of Blocked Punt EZ Rec TD",
              "abbreviation": "BPERTD",
              "value": 0.0,
              "displayValue": "0"
            }
          ],
          "athletes": [
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            }
          ]
        },
        {
          "name": "defensiveInterceptions",
          "displayName": "Defensive Interceptions",
          "shortDisplayName": "Defensive Interceptions",
          "abbreviation": "defint",
          "summary": "",
          "stats": [
            {
              "name": "interceptions",
              "displayName": "Interceptions",
              "shortDisplayName": "Interceptions",
              "description": "The total number of interceptions.",
              "abbreviation": "INT",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "interceptionTouchdowns",
              "displayName": "Interception Touchdowns",
              "shortDisplayName": "INTTD",
              "description": "The number of times an interception is returned for a touchdown.",
              "abbreviation": "TD",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "interceptionYards",
              "displayName": "Interception Yards",
              "shortDisplayName": "INTYDS",
              "description": "The number of yards gained after an interception.",
              "abbreviation": "YDS",
              "value": 0.0,
              "displayValue": "0"
            }
          ],
          "athletes": [
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            }
          ]
        },
        {
          "name": "kicking",
          "displayName": "Kicking",
          "shortDisplayName": "Kicking",
          "abbreviation": "kick",
          "summary": "",
          "stats": [
            {
              "name": "avgKickoffReturnYards",
              "displayName": "Average Kickoff Return Yards",
              "shortDisplayName": "YDS/KR",
              "description": "The average number of yards per kickoff return.",
              "abbreviation": "YDS/KR",
              "value": 28.2,
              "displayValue": "28.2"
            },
            {
              "name": "avgKickoffYards",
              "displayName": "Average Kickoff Yards",
              "shortDisplayName": "YDS/K",
              "description": "The average number of yards per kickoff.",
              "abbreviation": "YDS/K",
              "value": 43.333,
              "displayValue": "43.3"
            },
            {
              "name": "extraPointAttempts",
              "displayName": "Extra Point Attempts",
              "shortDisplayName": "XPA",
              "description": "The number of extra point attempts.",
              "abbreviation": "XPA",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "extraPointPct",
              "displayName": "Extra Point Percentage",
              "shortDisplayName": "XP%",
              "description": "The percentage of extra point attempts that are converted.",
              "abbreviation": "XP%",
              "value": 100.0,
              "displayValue": "100.0"
            },
            {
              "name": "extraPointsBlocked",
              "displayName": "Extra Point Blocked",
              "shortDisplayName": "XPB",
              "description": "The number of extra points blocked.",
              "abbreviation": "XPB",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "extraPointsBlockedPct",
              "displayName": "Extra Points Blocked Percentage",
              "shortDisplayName": "XPB%",
              "description": "The percentage of extra points attempts that are blocked.",
              "abbreviation": "XPB%",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "extraPointsMade",
              "displayName": "Extra Points Made",
              "shortDisplayName": "XPM",
              "description": "The number of extra point attempts made.",
              "abbreviation": "XPM",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "fairCatches",
              "displayName": "Fair Catches",
              "shortDisplayName": "FC",
              "description": "The number of fair catches.",
              "abbreviation": "FC",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fairCatchPct",
              "displayName": "Fair Catch Percentage",
              "shortDisplayName": "FC%",
              "description": "The percentage of kicks that are fair catches.",
              "abbreviation": "FC%",
              "value": 54.55,
              "displayValue": "54.6"
            },
            {
              "name": "fieldGoalAttempts",
              "displayName": "Field Goal Attempts",
              "shortDisplayName": "FGA",
              "description": "The number of times a field goal is attempted.",
              "abbreviation": "FGA",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fieldGoalAttempts1_19",
              "displayName": "Field Goal Attempts 1-19",
              "shortDisplayName": "FGA 1-19",
              "description": "The number of field goals attempted from 1-19 yards.",
              "abbreviation": "FGA 1-19",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fieldGoalAttempts20_29",
              "displayName": "Field Goal Attempts 20-29",
              "shortDisplayName": "FGA 20-29",
              "description": "The number of field goals attempted from 20-29 yards.",
              "abbreviation": "FGA 20-29",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fieldGoalAttempts30_39",
              "displayName": "Field Goal Attempts 30-39",
              "shortDisplayName": "FGA 30-39",
              "description": "The number of field goals attempted from 30-39 yards.",
              "abbreviation": "FGA 30-39",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fieldGoalAttempts40_49",
              "displayName": "Field Goal Attempts 40-49",
              "shortDisplayName": "FGA 40-49",
              "description": "The number of field goals attempted from 40-49 yards.",
              "abbreviation": "FGA 40-49",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fieldGoalAttempts50_59",
              "displayName": "Field Goal Attempts 50-59",
              "shortDisplayName": "FGA 50-59",
              "description": "The number of field goals attempted from 50-59 yards.",
              "abbreviation": "FGA 50-59",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fieldGoalAttempts60_99",
              "displayName": "Field Goal Attempts 60-99",
              "shortDisplayName": "FGA 60-99",
              "description": "The number of field goals attempted from 60-99 yards.",
              "abbreviation": "FGA 60-99",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fieldGoalAttempts50",
              "displayName": "Field Goal Attempts 50+",
              "shortDisplayName": "FGA 50+",
              "description": "The number of field goals attempted from 50+ yards.",
              "abbreviation": "FGA 50+",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fieldGoalAttemptYards",
              "displayName": "Field Goal Attempt Yards",
              "shortDisplayName": "FGAYDS",
              "description": "The number of yards for field goal attempts.",
              "abbreviation": "FGAYDS",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fieldGoalPct",
              "displayName": "Field Goal Percentage",
              "shortDisplayName": "FG%",
              "description": "The percentage of field goals attempted that are made.",
              "abbreviation": "FG%",
              "value": 0.0,
              "displayValue": "0.0"
            },
            {
              "name": "fieldGoalsBlocked",
              "displayName": "Field Goals Blocked",
              "shortDisplayName": "FGB",
              "description": "The number of field goals that are blocked.",
              "abbreviation": "FGB",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fieldGoalsBlockedPct",
              "displayName": "Field Goals Blocked Percentage",
              "shortDisplayName": "FGB%",
              "description": "The percentage of field goals attempted that are blocked.",
              "abbreviation": "CMP",
              "value": 0.0,
              "displayValue": "0.0"
            },
            {
              "name": "fieldGoalsMade",
              "displayName": "Field Goal Made",
              "shortDisplayName": "FGM",
              "description": "The number of times a field goal is made.",
              "abbreviation": "FGM",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fieldGoalsMade1_19",
              "displayName": "Field Goals Made 1-19",
              "shortDisplayName": "FGM 1-19",
              "description": "The number of field goals made from 1-19 yards.",
              "abbreviation": "FGM 1-19",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fieldGoalsMade20_29",
              "displayName": "Field Goals Made 20-29",
              "shortDisplayName": "FGM 20-29",
              "description": "The number of field goals made from 20-29 yards.",
              "abbreviation": "FGM 20-29",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fieldGoalsMade30_39",
              "displayName": "Field Goals Made 30-39",
              "shortDisplayName": "FGM 30-39",
              "description": "The number of field goals made from 30-39 yards.",
              "abbreviation": "FGM 30-39",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fieldGoalsMade40_49",
              "displayName": "Field Goals Made 40-49",
              "shortDisplayName": "FGM 40-49",
              "description": "The number of field goals made from 40-49 yards.",
              "abbreviation": "FGM 40-49",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fieldGoalsMade50_59",
              "displayName": "Field Goals Made 50-59",
              "shortDisplayName": "FGM 50-59",
              "description": "The number of field goals made from 50-59 yards.",
              "abbreviation": "FGM 50-59",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fieldGoalsMade60_99",
              "displayName": "Field Goals Made 60-99",
              "shortDisplayName": "FGM 60-99",
              "description": "The number of field goals made from 60-99 yards.",
              "abbreviation": "FGM 60-99",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fieldGoalsMade50",
              "displayName": "Field Goals Made 50+",
              "shortDisplayName": "FGM 50+",
              "description": "The number of field goals made from 50+ yards.",
              "abbreviation": "50+",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fieldGoalsMadeYards",
              "displayName": "Field Goals Made Yards",
              "shortDisplayName": "FGMYDS",
              "description": "The total amount of yardage of field goals made.",
              "abbreviation": "FGMYDS",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fieldGoalsMissedYards",
              "displayName": "Field Goals Missed Yards",
              "shortDisplayName": "FGMSYD",
              "description": "The total amount of yardage of missed field goals.",
              "abbreviation": "FGMSYD",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "kickoffOB",
              "displayName": "Kickoff Out of Bounds",
              "shortDisplayName": "KOB",
              "description": "The number of times the ball is kicked out of bounds on a kickoff.",
              "abbreviation": "KOB",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "kickoffReturns",
              "displayName": "Kickoff Returns",
              "shortDisplayName": "KR",
              "description": "The number of times a kickoff was returned.",
              "abbreviation": "KR",
              "value": 5.0,
              "displayValue": "5"
            },
            {
              "name": "kickoffReturnTouchdowns",
              "displayName": "Kickoff Return Touchdowns",
              "shortDisplayName": "KRTD",
              "description": "The number of times a kickoff was returned for a touchdown.",
              "abbreviation": "KRTD",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "kickoffReturnYards",
              "displayName": "Kickoff Return Yards",
              "shortDisplayName": "KRYDS",
              "description": "The total amount of yardage of kick returns.",
              "abbreviation": "KRYDS",
              "value": 141.0,
              "displayValue": "141"
            },
            {
              "name": "kickoffs",
              "displayName": "Kickoffs",
              "shortDisplayName": "K",
              "description": "The number of kickoffs taken.",
              "abbreviation": "K",
              "value": 3.0,
              "displayValue": "3"
            },
            {
              "name": "kickoffYards",
              "displayName": "Kickoff Yards",
              "shortDisplayName": "KYDS",
              "description": "The total amount of yardage of kickoffs.",
              "abbreviation": "KYDS",
              "value": 130.0,
              "displayValue": "130"
            },
            {
              "name": "longFieldGoalAttempt",
              "displayName": "Long Field Goal Attempt",
              "shortDisplayName": "LFGA",
              "description": "The distance of the longest field goal attempt.",
              "abbreviation": "LFGA",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "longFieldGoalMade",
              "displayName": "Long Field Goal Made",
              "shortDisplayName": "LFGM",
              "description": "The distance of the longest field goal made.",
              "abbreviation": "LNG",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "longKickoff",
              "displayName": "Long Kickoff",
              "shortDisplayName": "LK",
              "description": "The distance of the longest kickoff.",
              "abbreviation": "LK",
              "value": 54.0,
              "displayValue": "54"
            },
            {
              "name": "teamGamesPlayed",
              "displayName": "Team Games Played",
              "shortDisplayName": "GP",
              "description": "The numbers of team games played.",
              "abbreviation": "GP",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "totalKickingPoints",
              "displayName": "Total Kicking Points",
              "shortDisplayName": "TP",
              "description": "The number of total points earned kicking.",
              "abbreviation": "PTS",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "touchbackPct",
              "displayName": "Touchback Percentage",
              "shortDisplayName": "TB%",
              "description": "The percentage of kickoffs that are touchbacks.",
              "abbreviation": "TB%",
              "value": 33.33,
              "displayValue": "33"
            },
            {
              "name": "touchbacks",
              "displayName": "Touchbacks",
              "shortDisplayName": "TB",
              "description": "The number of kickoffs that result in touchbacks.",
              "abbreviation": "TB",
              "value": 1.0,
              "displayValue": "1"
            }
          ],
          "athletes": [
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            }
          ]
        },
        {
          "name": "returning",
          "displayName": "Returning",
          "shortDisplayName": "Returning",
          "abbreviation": "ret",
          "summary": "",
          "stats": [
            {
              "name": "defFumbleReturns",
              "displayName": "Defensive Fumbles Returns",
              "shortDisplayName": "DFR",
              "description": "The number of fumble returns.",
              "abbreviation": "DFR",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "defFumbleReturnYards",
              "displayName": "Defensive Fumble Return Yards",
              "shortDisplayName": "DFRYDS",
              "description": "The amount of total yardage from fumble returns.",
              "abbreviation": "DFRYDS",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fumbleRecoveries",
              "displayName": "Fumble Recoveries",
              "shortDisplayName": "FR",
              "description": "The number of fumble recoveries.",
              "abbreviation": "FR",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fumbleRecoveryYards",
              "displayName": "Fumble Recovery Yards",
              "shortDisplayName": "FRYDS",
              "description": "The amount of total yardage from fumble recoveries.",
              "abbreviation": "YDS",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "kickReturnFairCatches",
              "displayName": "Kick Return Fair Catches",
              "shortDisplayName": "KRFC",
              "description": "The number of kick return fair catches.",
              "abbreviation": "KRFC",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "kickReturnFairCatchPct",
              "displayName": "Kick Return Fair Catch Percentage",
              "shortDisplayName": "KRFC%",
              "description": "The percentage of kick returns that are fair caught.",
              "abbreviation": "KRFC%",
              "value": 0.0,
              "displayValue": "0.00"
            },
            {
              "name": "kickReturnFumbles",
              "displayName": "Kick Return Fumbles",
              "shortDisplayName": "KRF",
              "description": "The number of kick return fumbles.",
              "abbreviation": "KRF",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "kickReturnFumblesLost",
              "displayName": "Kick Return Fumbles Lost",
              "shortDisplayName": "KRFL",
              "description": "The number of kick return fumbles that are lost.",
              "abbreviation": "KRFL",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "kickReturns",
              "displayName": "Kick Returns",
              "shortDisplayName": "KR",
              "description": "The number of kick returns.",
              "abbreviation": "ATT",
              "value": 5.0,
              "displayValue": "5"
            },
            {
              "name": "kickReturnTouchdowns",
              "displayName": "Kick Return Touchdowns",
              "shortDisplayName": "KRTD",
              "description": "The number of kick return touchdowns.",
              "abbreviation": "TD",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "kickReturnYards",
              "displayName": "Kick Return Yards",
              "shortDisplayName": "KRYDS",
              "description": "The total yardage of kick returns.",
              "abbreviation": "YDS",
              "value": 141.0,
              "displayValue": "141"
            },
            {
              "name": "longKickReturn",
              "displayName": "Long Kick Return",
              "shortDisplayName": "LGKR",
              "description": "The longest kick return in yards.",
              "abbreviation": "LNG",
              "value": 38.0,
              "displayValue": "38"
            },
            {
              "name": "longPuntReturn",
              "displayName": "Long Punt Return",
              "shortDisplayName": "LGPR",
              "description": "The longest punt return in yards.",
              "abbreviation": "LNG",
              "value": 4.0,
              "displayValue": "4"
            },
            {
              "name": "miscFumbleReturns",
              "displayName": "Miscellaneous Fumbles Returns",
              "shortDisplayName": "MFR",
              "description": "The number of fumble returns.",
              "abbreviation": "MFR",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "miscFumbleReturnYards",
              "displayName": "Miscellaneous Fumble Return Yards",
              "shortDisplayName": "MFRYDS",
              "description": "The amount of total yardage from fumble returns.",
              "abbreviation": "MFRYDS",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "oppFumbleRecoveries",
              "displayName": "Opposition Fumble Recoveries",
              "shortDisplayName": "OPFR",
              "description": "The number of fumble recoveries by the opposition.",
              "abbreviation": "OPFR",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "oppFumbleRecoveryYards",
              "displayName": "Opposition Fumble Recovery Yards",
              "shortDisplayName": "OFRYDS",
              "description": "The amount of total yardage from fumble recoveries by the opposition.",
              "abbreviation": "OFRYDS",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "oppSpecialTeamFumbleReturns",
              "displayName": "Opposition Special Team Fumble Returns",
              "shortDisplayName": "OPSTFR",
              "description": "The number of fumble returns by the opposition's speical teams.",
              "abbreviation": "OPSTFR",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "oppSpecialTeamFumbleReturnYards",
              "displayName": "Opposition Special Team Fumble Return Yards",
              "shortDisplayName": "OSFRYD",
              "description": "The amount of total yardage from fumble return by the opposition's special teams.",
              "abbreviation": "OFRYDS",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "puntReturnFairCatches",
              "displayName": "Punt Return Fair Catches",
              "shortDisplayName": "PRFC",
              "description": "The number of punt return fair catches.",
              "abbreviation": "FC",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "puntReturnFairCatchPct",
              "displayName": "Punt Return Fair Catch Percentage",
              "shortDisplayName": "PRFC%",
              "description": "The percentage of kick returns that are fair caught.",
              "abbreviation": "PRFC%",
              "value": 33.333336,
              "displayValue": "33.33"
            },
            {
              "name": "puntReturnFumbles",
              "displayName": "Punt Return Fumbles",
              "shortDisplayName": "PRF",
              "description": "The number of punt return fumbles.",
              "abbreviation": "PRF",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "puntReturnFumblesLost",
              "displayName": "Punt Return Fumbles Lost",
              "shortDisplayName": "PRFL",
              "description": "The number of punt return fumbles that are lost.",
              "abbreviation": "PRFL",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "puntReturns",
              "displayName": "Punt Returns",
              "shortDisplayName": "PR",
              "description": "The number of punt returns.",
              "abbreviation": "ATT",
              "value": 2.0,
              "displayValue": "2"
            },
            {
              "name": "puntReturnsStartedInsideThe10",
              "displayName": "Punt Returns Started Inside the 10",
              "shortDisplayName": "PR 10",
              "description": "The number of punt returns started inside the 10.",
              "abbreviation": "PR 10",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "puntReturnsStartedInsideThe20",
              "displayName": "Punt Returns Started Inside the 20",
              "shortDisplayName": "PR 20",
              "description": "The number of punt returns started inside the 20.",
              "abbreviation": "PR 20",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "puntReturnTouchdowns",
              "displayName": "Punt Return Touchdowns",
              "shortDisplayName": "PRTD",
              "description": "The number of punt return touchdowns.",
              "abbreviation": "TD",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "puntReturnYards",
              "displayName": "Punt Return Yards",
              "shortDisplayName": "PRYDS",
              "description": "The total yardage of punt returns.",
              "abbreviation": "YDS",
              "value": 4.0,
              "displayValue": "4"
            },
            {
              "name": "specialTeamFumbleReturns",
              "displayName": "Special Team Fumbles Returns",
              "shortDisplayName": "STFR",
              "description": "The number of special team fumble returns.",
              "abbreviation": "STFR",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "specialTeamFumbleReturnYards",
              "displayName": "Special Team Fumble Return Yards",
              "shortDisplayName": "STFRYD",
              "description": "The amount of total yardage from special team fumble returns.",
              "abbreviation": "STFRYD",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "teamGamesPlayed",
              "displayName": "Team Games Played",
              "shortDisplayName": "GP",
              "description": "The numbers of team games played.",
              "abbreviation": "GP",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "yardsPerKickReturn",
              "displayName": "Yards Per Kick Return",
              "shortDisplayName": "YDS/KR",
              "description": "The average number of yards per kick return.",
              "abbreviation": "AVG",
              "value": 28.2,
              "displayValue": "28.2"
            },
            {
              "name": "yardsPerPuntReturn",
              "displayName": "Yards Per Punt Return",
              "shortDisplayName": "YDS/PR",
              "description": "The average number of yards per punt return.",
              "abbreviation": "AVG",
              "value": 2.0,
              "displayValue": "2.0"
            },
            {
              "name": "yardsPerReturn",
              "displayName": "Yards Per Return",
              "shortDisplayName": "YDS/R",
              "description": "The average number of yards per return.",
              "abbreviation": "YDS/R",
              "value": 20.714285,
              "displayValue": "20.7"
            }
          ],
          "athletes": [
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            },
            {
              "athlete": {},
              "statistics": {}
            }
          ]
        },
        {
          "name": "punting",
          "displayName": "Punting",
          "shortDisplayName": "Punting",
          "abbreviation": "punt",
          "summary": "",
          "stats": [
            {
              "name": "avgPuntReturnYards",
              "displayName": "Average Punt Return Yards",
              "shortDisplayName": "YDS/PPR",
              "description": "The average number of yards per punt return.",
              "abbreviation": "AVG",
              "value": 2.0,
              "displayValue": "2.0"
            },
            {
              "name": "fairCatches",
              "displayName": "Fair Catches",
              "shortDisplayName": "FC",
              "description": "The number of fair catches.",
              "abbreviation": "FC",
              "value": 6.0,
              "displayValue": "6"
            },
            {
              "name": "grossAvgPuntYards",
              "displayName": "Gross Average Punt Yards",
              "shortDisplayName": "GYDS/P",
              "description": "The average gross number of yards per punt.",
              "abbreviation": "AVG",
              "value": 44.5,
              "displayValue": "44.5"
            },
            {
              "name": "longPunt",
              "displayName": "Long Punt",
              "shortDisplayName": "LP",
              "description": "The distance of the longest punt.",
              "abbreviation": "LNG",
              "value": 55.0,
              "displayValue": "55"
            },
            {
              "name": "netAvgPuntYards",
              "displayName": "Net Average Punt Yards",
              "shortDisplayName": "NYDS/P",
              "description": "The average net number of yards per punt.",
              "abbreviation": "NET",
              "value": 42.5,
              "displayValue": "42.5"
            },
            {
              "name": "puntReturns",
              "displayName": "Punt Returns",
              "shortDisplayName": "PR",
              "description": "The number of times a punt was returned.",
              "abbreviation": "ATT",
              "value": 2.0,
              "displayValue": "2"
            },
            {
              "name": "puntReturnYards",
              "displayName": "Punt Return Yards",
              "shortDisplayName": "PRYDS",
              "description": "The total amount of yardage of punt returns.",
              "abbreviation": "YDS",
              "value": 16.0,
              "displayValue": "16"
            },
            {
              "name": "punts",
              "displayName": "Punts",
              "shortDisplayName": "P",
              "description": "The number of punts taken.",
              "abbreviation": "PUNTS",
              "value": 8.0,
              "displayValue": "8"
            },
            {
              "name": "puntsBlocked",
              "displayName": "Punts Blocked",
              "shortDisplayName": "PBLK",
              "description": "The number of times a punt was blocked.",
              "abbreviation": "PBLK",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "puntsBlockedPct",
              "displayName": "Punts Blocked Percentage",
              "shortDisplayName": "PBLK%",
              "description": "The percentage of punts that are blocked.",
              "abbreviation": "PBLK%",
              "value": 0.0,
              "displayValue": "0.00"
            },
            {
              "name": "puntsInside10",
              "displayName": "Punts Inside 10",
              "shortDisplayName": "P 10",
              "description": "The number of times a punt is downed inside the 10.",
              "abbreviation": "P 10",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "puntsInside10Pct",
              "displayName": "Punts Inside 10 Percentage",
              "shortDisplayName": "P 10%",
              "description": "The percentage of punts that end up inside the 10.",
              "abbreviation": "P 10%",
              "value": 0.0,
              "displayValue": "0.00"
            },
            {
              "name": "puntsInside20",
              "displayName": "Punts Inside 20",
              "shortDisplayName": "P 20",
              "description": "The number of times a punt is downed inside the 20.",
              "abbreviation": "IN20",
              "value": 4.0,
              "displayValue": "4"
            },
            {
              "name": "puntsInside20Pct",
              "displayName": "Punts Inside 20 Percentage",
              "shortDisplayName": "P 10%",
              "description": "The percentage of punts that end up inside the 10.",
              "abbreviation": "IN20%",
              "value": 50.0,
              "displayValue": "50.00"
            },
            {
              "name": "puntsOver50",
              "displayName": "Punts Over 50",
              "shortDisplayName": "P 50+",
              "description": "The number of times a punt goes over 50 yards.",
              "abbreviation": "P 50+",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "puntYards",
              "displayName": "Punt Yards",
              "shortDisplayName": "PYDS",
              "description": "The total amount of yardage of punts.",
              "abbreviation": "YDS",
              "value": 356.0,
              "displayValue": "356"
            },
            {
              "name": "teamGamesPlayed",
              "displayName": "Team Games Played",
              "shortDisplayName": "GP",
              "description": "The numbers of team games played.",
              "abbreviation": "GP",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "touchbackPct",
              "displayName": "Touchback Percentage",
              "shortDisplayName": "TB%",
              "description": "The percentage of punts that are touchbacks.",
              "abbreviation": "TB%",
              "value": 0.0,
              "displayValue": "0.00"
            },
            {
              "name": "touchbacks",
              "displayName": "Touchbacks",
              "shortDisplayName": "TB",
              "description": "The number of kickoffs that result in touchbacks.",
              "abbreviation": "TB",
              "value": 0.0,
              "displayValue": "0"
            }
          ],
          "athletes": [
            {
              "athlete": {},
              "statistics": {}
            }
          ]
        },
        {
          "name": "scoring",
          "displayName": "Scoring",
          "shortDisplayName": "Scoring",
          "abbreviation": "s",
          "summary": "",
          "stats": [
            {
              "name": "defensivePoints",
              "displayName": "Defensive Points",
              "shortDisplayName": "DP",
              "description": "The number of points scored on defense.",
              "abbreviation": "DP",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fieldGoals",
              "displayName": "Field Goals",
              "shortDisplayName": "FG",
              "description": "The number of field goals made.",
              "abbreviation": "FG",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "kickExtraPoints",
              "displayName": "Kick Extra Points",
              "shortDisplayName": "XP",
              "description": "The number of extra points made.",
              "abbreviation": "PAT",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "kickExtraPointsMade",
              "displayName": "Kick Extra Points Made",
              "shortDisplayName": "XPM",
              "description": "The number of extra points made.",
              "abbreviation": "XPM",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "miscPoints",
              "displayName": "Miscellaneous Points",
              "shortDisplayName": "MISP",
              "description": "The number of miscellaneous points scored.",
              "abbreviation": "MISP",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "passingTouchdowns",
              "displayName": "Passing Touchdowns",
              "shortDisplayName": "Touchdowns",
              "description": "The total number of passing touchdowns.",
              "abbreviation": "PASS",
              "value": 2.0,
              "displayValue": "2"
            },
            {
              "name": "receivingTouchdowns",
              "displayName": "Receiving Touchdowns",
              "shortDisplayName": "RECTD",
              "description": "The total number of receiving touchdowns.",
              "abbreviation": "REC",
              "value": 2.0,
              "displayValue": "2"
            },
            {
              "name": "returnTouchdowns",
              "displayName": "Return Touchdowns",
              "shortDisplayName": "RETTD",
              "description": "The total number of return touchdowns.",
              "abbreviation": "RET",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "rushingTouchdowns",
              "displayName": "Rushing Touchdowns",
              "shortDisplayName": "RUSTD",
              "description": "The total number of rushing touchdowns.",
              "abbreviation": "RUSH",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "totalPoints",
              "displayName": "Total Points",
              "shortDisplayName": "TP",
              "description": "The number of total points scored.",
              "abbreviation": "PTS",
              "value": 13.0,
              "displayValue": "13"
            },
            {
              "name": "totalPointsPerGame",
              "displayName": "Total Points Per Game",
              "shortDisplayName": "TP/G",
              "description": "The number of points scored per game.",
              "abbreviation": "TP/G",
              "value": 13.0,
              "displayValue": "13.0"
            },
            {
              "name": "totalTouchdowns",
              "displayName": "Total Touchdowns",
              "shortDisplayName": "TTD",
              "description": "The number of touchdowns scored in total.",
              "abbreviation": "TD",
              "value": 2.0,
              "displayValue": "2"
            },
            {
              "name": "totalTwoPointConvs",
              "displayName": "Total Two Point Conversions",
              "shortDisplayName": "2PTC",
              "description": "The number of times a 2-point conversion is successful.",
              "abbreviation": "2PT",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "twoPointPassConvs",
              "displayName": "Two Point Pass Conversions",
              "shortDisplayName": "2PTPC",
              "description": "The number of times a 2-point conversion is successful with a pass.",
              "abbreviation": "2PTPC",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "twoPointRecConvs",
              "displayName": "Two Point Receiving Conversion",
              "shortDisplayName": "2PTRCC",
              "description": "The number of times a 2-point is converted with a reception.",
              "abbreviation": "2PTRCC",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "twoPointRushConvs",
              "displayName": "Two Point Rush Conversion",
              "shortDisplayName": "2PTRUC",
              "description": "The number of times a 2-point is converted with a run.",
              "abbreviation": "2PTRUC",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "onePtSafetiesMade",
              "displayName": "One Point Safeties Made",
              "shortDisplayName": "OPSM",
              "description": "One Point Safeties Made",
              "abbreviation": "OPSM",
              "value": 0.0,
              "displayValue": "0"
            }
          ]
        },
        {
          "name": "miscellaneous",
          "displayName": "Miscellaneous",
          "shortDisplayName": "Miscellaneous",
          "abbreviation": "misc",
          "summary": "",
          "stats": [
            {
              "name": "firstDowns",
              "displayName": "Total 1st downs",
              "shortDisplayName": "FIRST",
              "description": "The the number of first downs.",
              "abbreviation": "CMP",
              "value": 18.0,
              "displayValue": "18"
            },
            {
              "name": "firstDownsPassing",
              "displayName": "Passing 1st downs",
              "shortDisplayName": "FDP",
              "description": "The number of times a pass results in a first down.",
              "abbreviation": "FDP",
              "value": 14.0,
              "displayValue": "14"
            },
            {
              "name": "firstDownsPenalty",
              "displayName": "1st downs by penalty",
              "shortDisplayName": "FDPEN",
              "description": "The number of times a penalty results in a first down.",
              "abbreviation": "FDPEN",
              "value": 2.0,
              "displayValue": "2"
            },
            {
              "name": "firstDownsPerGame",
              "displayName": "First Downs Per Game",
              "shortDisplayName": "FD/G",
              "description": "The average number of first downs per game.",
              "abbreviation": "FD/G",
              "value": 18.0,
              "displayValue": "18.00"
            },
            {
              "name": "firstDownsRushing",
              "displayName": "Rushing 1st downs",
              "shortDisplayName": "FDR",
              "description": "The number of times a rush results in a first down.",
              "abbreviation": "FDR",
              "value": 2.0,
              "displayValue": "2"
            },
            {
              "name": "fourthDownAttempts",
              "displayName": "First Downs Attempts",
              "shortDisplayName": "FDA",
              "description": "The number of attempts at getting a first down.",
              "abbreviation": "FDA",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "fourthDownConvPct",
              "displayName": "4th down %",
              "shortDisplayName": "4THC%",
              "description": "The percentage of fourth down attempts that are converted.",
              "abbreviation": "4THC%",
              "value": 0.0,
              "displayValue": "0.00"
            },
            {
              "name": "fourthDownConvs",
              "displayName": "Fourth Down Conversions",
              "shortDisplayName": "4THC",
              "description": "The number of times a fourth down is converted.",
              "abbreviation": "4THC",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "possessionTimeSeconds",
              "displayName": "Possession Time Seconds",
              "shortDisplayName": "POSS",
              "description": "The total time of possession in seconds.",
              "abbreviation": "POSS",
              "value": 1609.0,
              "displayValue": "1609"
            },
            {
              "name": "redzoneAttemptPoints",
              "displayName": "Redzone Attempt Points",
              "shortDisplayName": "RZP",
              "description": "The number of points scored from red zone attempts.",
              "abbreviation": "RZP",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "redzoneAttempts",
              "displayName": "Red Zone Attempts",
              "shortDisplayName": "RZA",
              "description": "The number of times the team gets in the red zone.",
              "abbreviation": "RZA",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "redzoneConversion",
              "displayName": "Red Zone Conversion",
              "shortDisplayName": "RZC",
              "description": "The number of times the team scores in the red zone.",
              "abbreviation": "RZC",
              "value": 1.0,
              "displayValue": "1"
            },
            {
              "name": "redzoneEfficiencyPct",
              "displayName": "Red Zone Efficiency Percentage",
              "shortDisplayName": "RZ%",
              "description": "The percentage of red zone attempts that are converted.",
              "abbreviation": "CMP",
              "value": 0.0,
              "displayValue": "0.00"
            },
            {
              "name": "redzoneEndDowns",
              "displayName": "Red Zone End Downs",
              "shortDisplayName": "RZD",
              "description": "The times a red zone trip ends on downs.",
              "abbreviation": "RZD",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "redzoneEndGame",
              "displayName": "Red Zone End Game",
              "shortDisplayName": "RZG",
              "description": "The times a red zone trip ends with the end of the game.",
              "abbreviation": "RZG",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "redzoneEndHalf",
              "displayName": "Red Zone End Half",
              "shortDisplayName": "RZH",
              "description": "The times a red zone trip ends with the end of the half.",
              "abbreviation": "RZH",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "redzoneFieldGoalPct",
              "displayName": "Red Zone Field Goal Percentage",
              "shortDisplayName": "RZFG%",
              "description": "The percentage of red zone trips that end in a field goal.",
              "abbreviation": "RZFG%",
              "value": 0.0,
              "displayValue": "0.00"
            },
            {
              "name": "redzoneFieldGoalPoints",
              "displayName": "Red Zone Field Goal Points",
              "shortDisplayName": "RZFGP",
              "description": "The number of points scored in the red zone from field goals.",
              "abbreviation": "RZFGP",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "redzoneFieldGoals",
              "displayName": "Red Zone Field Goals",
              "shortDisplayName": "RZFG",
              "description": "The number of red zone trips that resulted in field goals.",
              "abbreviation": "RZFG",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "redzoneFieldGoalsMissed",
              "displayName": "Red Zone Field Goals Missed",
              "shortDisplayName": "RZFGM",
              "description": "The number of red zone trips that resulted in a missed field goal.",
              "abbreviation": "RZFGM",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "redzoneFumble",
              "displayName": "Red Zone Fumble",
              "shortDisplayName": "RZF",
              "description": "The number of red zone fumbles.",
              "abbreviation": "RZF",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "redzoneInterception",
              "displayName": "Red Zone Interception",
              "shortDisplayName": "RZINT",
              "description": "The number of red zone interceptions.",
              "abbreviation": "RZINT",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "redzoneScoringPct",
              "displayName": "Red Zone Scoring Percentage",
              "shortDisplayName": "RZ%",
              "description": "The percentage of red zone trips that result in a score.",
              "abbreviation": "RZ%",
              "value": 100.0,
              "displayValue": "100.00"
            },
            {
              "name": "redzoneTotalPoints",
              "displayName": "Red Total Points",
              "shortDisplayName": "RZTP",
              "description": "The number of total points scored in the red zone.",
              "abbreviation": "RZTP",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "redzoneTouchdownPass",
              "displayName": "Red Zone Touchdown Pass",
              "shortDisplayName": "RZTDP",
              "description": "The number of red zone touchdown passes.",
              "abbreviation": "RZTDP",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "redzoneTouchdownPct",
              "displayName": "Red Zone Touchdown Percentage",
              "shortDisplayName": "RZTD%",
              "description": "The percentage of red zone trips that result in a touchdown.",
              "abbreviation": "RZTD%",
              "value": 0.0,
              "displayValue": "0.00"
            },
            {
              "name": "redzoneTouchdownPoints",
              "displayName": "Red Zone Touchdown Points",
              "shortDisplayName": "RZTDPT",
              "description": "The number of points scored from red zone touchdowns.",
              "abbreviation": "RZTDPT",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "redzoneTouchdownRush",
              "displayName": "Red Zone Touchdown Rush",
              "shortDisplayName": "RZTDR",
              "description": "The number of red zone touchdowns scored from a run.",
              "abbreviation": "RZTDR",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "redzoneTouchdowns",
              "displayName": "Red Zone Touchdowns",
              "shortDisplayName": "RZTD",
              "description": "The number of touchdown scored in the red zone.",
              "abbreviation": "RZTD",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "thirdDownAttempts",
              "displayName": "Third Down Attempts",
              "shortDisplayName": "3RDA",
              "description": "The number of 3rd down attempts.",
              "abbreviation": "3RDA",
              "value": 15.0,
              "displayValue": "15"
            },
            {
              "name": "thirdDownConvPct",
              "displayName": "3rd down %",
              "shortDisplayName": "3RDC%",
              "description": "The percentage of 3rd downs that are converted.",
              "abbreviation": "3RDC%",
              "value": 40.0,
              "displayValue": "40.00"
            },
            {
              "name": "thirdDownConvs",
              "displayName": "Third Down Conversions",
              "shortDisplayName": "3RDC",
              "description": "The number of 3rd down conversions.",
              "abbreviation": "3RDC",
              "value": 6.0,
              "displayValue": "6"
            },
            {
              "name": "timeoutsUsed",
              "displayName": "Timeouts Used",
              "shortDisplayName": "TOU",
              "description": "The number of timeouts used.",
              "abbreviation": "TOU",
              "value": 3.0,
              "displayValue": "3"
            },
            {
              "name": "totalGiveaways",
              "displayName": "Total Giveaways",
              "shortDisplayName": "TGV",
              "description": "The total number of giveaways.",
              "abbreviation": "TGV",
              "value": 3.0,
              "displayValue": "3"
            },
            {
              "name": "totalPenalties",
              "displayName": "Total Penalties",
              "shortDisplayName": "TPEN",
              "description": "The total number of total penalties.",
              "abbreviation": "TPEN",
              "value": 3.0,
              "displayValue": "3"
            },
            {
              "name": "totalPenaltyYards",
              "displayName": "Total Penalty Yards",
              "shortDisplayName": "TPY",
              "description": "The total number of penalty yards.",
              "abbreviation": "TPY",
              "value": 25.0,
              "displayValue": "25"
            },
            {
              "name": "totalPlays",
              "displayName": "Total Plays",
              "shortDisplayName": "TPLY",
              "description": "The total number of plays run.",
              "abbreviation": "TPLY",
              "value": 67.0,
              "displayValue": "67"
            },
            {
              "name": "totalTakeaways",
              "displayName": "Total Takeaways",
              "shortDisplayName": "TT",
              "description": "The total number of takeaways.",
              "abbreviation": "TT",
              "value": 0.0,
              "displayValue": "0"
            },
            {
              "name": "totalDrives",
              "displayName": "Total Drives",
              "shortDisplayName": "D",
              "description": "The total number of drives.",
              "abbreviation": "D",
              "value": 15.0,
              "displayValue": "15"
            },
            {
              "name": "turnOverDifferential",
              "displayName": "Turnover Ratio",
              "shortDisplayName": "DIFF",
              "description": "Difference between take-aways and give-aways",
              "abbreviation": "DIFF",
              "value": -3.0,
              "displayValue": "-3"
            }
          ]
        }
      ]
    }
  }
}
```

---

## Odds

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/odds`

```json
{
  "count": 1,
  "pageIndex": 1,
  "pageSize": 25,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/odds/100?lang=en&region=us",
      "provider": {
        "id": "100",
        "name": "Draft Kings",
        "priority": 1
      },
      "details": "SEA -4.5",
      "overUnder": 45.5,
      "spread": 4.5,
      "overOdds": -108.0,
      "underOdds": -112.0,
      "awayTeamOdds": {
        "favorite": true,
        "underdog": false,
        "moneyLine": -230,
        "spreadOdds": -112.0,
        "open": {
          "favorite": true,
          "pointSpread": {
            "alternateDisplayValue": "-3.5",
            "american": "-3.5"
          },
          "spread": {
            "value": 1.9,
            "displayValue": "10/11",
            "alternateDisplayValue": "-110",
            "decimal": 1.9,
            "fraction": "10/11",
            "american": "-110"
          },
          "moneyLine": {
            "value": 1.5,
            "displayValue": "50/99",
            "alternateDisplayValue": "-198",
            "decimal": 1.5,
            "fraction": "50/99",
            "american": "-198"
          }
        },
        "close": {
          "pointSpread": {
            "alternateDisplayValue": "-4.5",
            "american": "-4.5"
          },
          "spread": {
            "value": 1.89,
            "displayValue": "25/28",
            "alternateDisplayValue": "-112",
            "decimal": 1.89,
            "fraction": "25/28",
            "american": "-112"
          },
          "moneyLine": {
            "value": 1.43,
            "displayValue": "10/23",
            "alternateDisplayValue": "-230",
            "decimal": 1.43,
            "fraction": "10/23",
            "american": "-230"
          }
        },
        "current": {
          "pointSpread": {
            "alternateDisplayValue": "-4.5",
            "american": "-4.5"
          },
          "spread": {
            "value": 1.89,
            "displayValue": "25/28",
            "alternateDisplayValue": "-112",
            "decimal": 1.89,
            "fraction": "25/28",
            "american": "-112"
          },
          "moneyLine": {
            "value": 1.43,
            "displayValue": "10/23",
            "alternateDisplayValue": "-230",
            "decimal": 1.43,
            "fraction": "10/23",
            "american": "-230"
          }
        },
        "team": {}
      },
      "homeTeamOdds": {
        "favorite": false,
        "underdog": true,
        "moneyLine": 190,
        "spreadOdds": -108.0,
        "open": {
          "favorite": false,
          "pointSpread": {
            "alternateDisplayValue": "+3.5",
            "american": "+3.5"
          },
          "spread": {
            "value": 1.9,
            "displayValue": "10/11",
            "alternateDisplayValue": "-110",
            "decimal": 1.9,
            "fraction": "10/11",
            "american": "-110"
          },
          "moneyLine": {
            "value": 2.64,
            "displayValue": "41/25",
            "alternateDisplayValue": "+164",
            "decimal": 2.64,
            "fraction": "41/25",
            "american": "+164"
          }
        },
        "close": {
          "pointSpread": {
            "alternateDisplayValue": "+4.5",
            "american": "+4.5"
          },
          "spread": {
            "value": 1.92,
            "displayValue": "25/27",
            "alternateDisplayValue": "-108",
            "decimal": 1.92,
            "fraction": "25/27",
            "american": "-108"
          },
          "moneyLine": {
            "value": 2.9,
            "displayValue": "19/10",
            "alternateDisplayValue": "+190",
            "decimal": 2.9,
            "fraction": "19/10",
            "american": "+190"
          }
        },
        "current": {
          "pointSpread": {
            "alternateDisplayValue": "+4.5",
            "american": "+4.5"
          },
          "spread": {
            "value": 1.92,
            "displayValue": "25/27",
            "alternateDisplayValue": "-108",
            "decimal": 1.92,
            "fraction": "25/27",
            "american": "-108"
          },
          "moneyLine": {
            "value": 2.9,
            "displayValue": "19/10",
            "alternateDisplayValue": "+190",
            "decimal": 2.9,
            "fraction": "19/10",
            "american": "+190"
          }
        },
        "team": {}
      },
      "links": [
        {
          "language": "en-US",
          "rel": [
            "home",
            "desktop",
            "bets",
            "draft-kings"
          ],
          "href": "https://sportsbook.draftkings.com/gateway?s=__s__&wpcid=__wpcid__&wpsrc=413&wpcn=ESPN&wpscn=Widget&wpcrn=BetSlipDeepLink&wpscid=__wpscid__&wpcrid=xx&preurl=https%3A%2F%2Fsportsbook.draftkings.com%2Fevent%2F33488086%3Foutcomes%3D0ML83048305_1",
          "text": "Home Bet",
          "shortText": "Home Bet",
          "isExternal": true,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "away",
            "desktop",
            "bets",
            "draft-kings"
          ],
          "href": "https://sportsbook.draftkings.com/gateway?s=__s__&wpcid=__wpcid__&wpsrc=413&wpcn=ESPN&wpscn=Widget&wpcrn=BetSlipDeepLink&wpscid=__wpscid__&wpcrid=xx&preurl=https%3A%2F%2Fsportsbook.draftkings.com%2Fevent%2F33488086%3Foutcomes%3D0ML83048305_3",
          "text": "Away Bet",
          "shortText": "Away Bet",
          "isExternal": true,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "homeSpread",
            "desktop",
            "bets"
          ],
          "href": "https://sportsbook.draftkings.com/gateway?s=__s__&wpcid=__wpcid__&wpsrc=413&wpcn=ESPN&wpscn=Widget&wpcrn=BetSlipDeepLink&wpscid=__wpscid__&wpcrid=xx&preurl=https%3A%2F%2Fsportsbook.draftkings.com%2Fevent%2F33488086%3Foutcomes%3D0HC83048305P450_1",
          "text": "Home Point Spread",
          "shortText": "Home Point Spread",
          "isExternal": true,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "awaySpread",
            "desktop",
            "bets"
          ],
          "href": "https://sportsbook.draftkings.com/gateway?s=__s__&wpcid=__wpcid__&wpsrc=413&wpcn=ESPN&wpscn=Widget&wpcrn=BetSlipDeepLink&wpscid=__wpscid__&wpcrid=xx&preurl=https%3A%2F%2Fsportsbook.draftkings.com%2Fevent%2F33488086%3Foutcomes%3D0HC83048305N450_3",
          "text": "Away Point Spread",
          "shortText": "Away Point Spread",
          "isExternal": true,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "over",
            "desktop",
            "bets"
          ],
          "href": "https://sportsbook.draftkings.com/gateway?s=__s__&wpcid=__wpcid__&wpsrc=413&wpcn=ESPN&wpscn=Widget&wpcrn=BetSlipDeepLink&wpscid=__wpscid__&wpcrid=xx&preurl=https%3A%2F%2Fsportsbook.draftkings.com%2Fevent%2F33488086%3Foutcomes%3D0OU83048305O4550_1",
          "text": "Over Odds",
          "shortText": "Over Odds",
          "isExternal": true,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "under",
            "desktop",
            "bets"
          ],
          "href": "https://sportsbook.draftkings.com/gateway?s=__s__&wpcid=__wpcid__&wpsrc=413&wpcn=ESPN&wpscn=Widget&wpcrn=BetSlipDeepLink&wpscid=__wpscid__&wpcrid=xx&preurl=https%3A%2F%2Fsportsbook.draftkings.com%2Fevent%2F33488086%3Foutcomes%3D0OU83048305U4550_3",
          "text": "Under Odds",
          "shortText": "Under Odds",
          "isExternal": true,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "game",
            "desktop",
            "bets"
          ],
          "href": "https://sportsbook.draftkings.com/gateway?s=__s__&wpcid=__wpcid__&wpsrc=413&wpcn=ESPN&wpscn=Widget&wpcrn=BetSlipDeepLink&wpscid=__wpscid__&wpcrid=xx&preurl=https%3A%2F%2Fsportsbook.draftkings.com%2Fevent%2F33488086",
          "text": "Game",
          "shortText": "Game",
          "isExternal": true,
          "isPremium": false
        }
      ],
      "moneylineWinner": false,
      "spreadWinner": false,
      "open": {
        "over": {
          "value": 1.9,
          "displayValue": "10/11",
          "alternateDisplayValue": "-110",
          "decimal": 1.9,
          "fraction": "10/11",
          "american": "-110"
        },
        "under": {
          "value": 1.9,
          "displayValue": "10/11",
          "alternateDisplayValue": "-110",
          "decimal": 1.9,
          "fraction": "10/11",
          "american": "-110"
        },
        "total": {
          "alternateDisplayValue": "46.5",
          "american": "46.5"
        }
      },
      "close": {
        "over": {
          "value": 1.92,
          "displayValue": "25/27",
          "alternateDisplayValue": "-108",
          "decimal": 1.92,
          "fraction": "25/27",
          "american": "-108"
        },
        "under": {
          "value": 1.89,
          "displayValue": "25/28",
          "alternateDisplayValue": "-112",
          "decimal": 1.89,
          "fraction": "25/28",
          "american": "-112"
        },
        "total": {
          "alternateDisplayValue": "45.5",
          "american": "45.5"
        }
      },
      "current": {
        "over": {
          "value": 1.92,
          "displayValue": "25/27",
          "alternateDisplayValue": "-108",
          "decimal": 1.92,
          "fraction": "25/27",
          "american": "-108"
        },
        "under": {
          "value": 1.89,
          "displayValue": "25/28",
          "alternateDisplayValue": "-112",
          "decimal": 1.89,
          "fraction": "25/28",
          "american": "-112"
        },
        "total": {
          "alternateDisplayValue": "45.5",
          "american": "45.5"
        }
      }
    }
  ]
}
```

---

## Officials

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/officials`

```json
{
  "count": 7,
  "pageIndex": 1,
  "pageSize": 25,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/officials/17655?lang=en&region=us",
      "id": "17655",
      "firstName": "Dana",
      "lastName": "McKenzie",
      "fullName": "Dana McKenzie",
      "displayName": "Dana McKenzie",
      "position": {
        "name": "Down Judge",
        "displayName": "Down Judge",
        "id": "112"
      },
      "order": 1
    }
  ]
}
```

---

## Play Personnel

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/plays/{play}/personnel`

Notes:
- Using the last play ref from the competition situation (`play=4017729884651`), NFL currently returned HTTP 500.

```json
{
  "error": {
    "message": "application error",
    "code": 500
  }
}
```
