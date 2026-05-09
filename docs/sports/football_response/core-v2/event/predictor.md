# Competition Predictor

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/predictor

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988` on 2026-05-08.
- NFL returns `homeTeam` and `awayTeam` blocks with predictor statistics rather than a single `gameProjection` field.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/predictor?lang=en&region=us",
  "name": "Matchup Predictor",
  "homeTeam": {
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
    "statistics": [
      {
        "name": "gameProjection",
        "displayName": "WIN PROB",
        "value": 40.572,
        "displayValue": "40.6"
      },
      {
        "name": "matchupQuality",
        "displayName": "Matchup Quality",
        "value": 83.451,
        "displayValue": "83.5"
      }
    ]
  },
  "awayTeam": {
    "team": {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/teams/26?lang=en&region=us",
      "id": "26",
      "guid": "3767a23c-a17d-294f-288a-9086ad1f0680",
      "uid": "s:20~l:28~t:26",
      "alternateIds": {
        "sdr": "8827"
      },
      "slug": "seattle-seahawks",
      "location": "Seattle",
      "name": "Seahawks",
      "nickname": "Seahawks",
      "abbreviation": "SEA",
      "displayName": "Seattle Seahawks",
      "shortDisplayName": "Seahawks",
      "color": "002a5c",
      "alternateColor": "69be28",
      "isActive": true,
      "isAllStar": false,
      "logos": [
        {
          "href": "https://a.espncdn.com/i/teamlogos/nfl/500/sea.png",
          "width": 500,
          "height": 500,
          "alt": "",
          "rel": [
            "full",
            "default"
          ],
          "lastUpdated": "2024-06-25T18:57Z"
        },
        {
          "href": "https://a.espncdn.com/i/teamlogos/nfl/500-dark/sea.png",
          "width": 500,
          "height": 500,
          "alt": "",
          "rel": [
            "full",
            "dark"
          ],
          "lastUpdated": "2024-06-25T18:57Z"
        },
        {
          "href": "https://a.espncdn.com/i/teamlogos/nfl/500/scoreboard/sea.png",
          "width": 500,
          "height": 500,
          "alt": "",
          "rel": [
            "full",
            "scoreboard"
          ],
          "lastUpdated": "2024-06-25T18:57Z"
        },
        {
          "href": "https://a.espncdn.com/i/teamlogos/nfl/500-dark/scoreboard/sea.png",
          "width": 500,
          "height": 500,
          "alt": "",
          "rel": [
            "full",
            "scoreboard",
            "dark"
          ],
          "lastUpdated": "2024-06-25T18:57Z"
        },
        {
          "href": "https://a.espncdn.com/guid/3767a23c-a17d-294f-288a-9086ad1f0680/logos/grayscale.png",
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
          "href": "https://a.espncdn.com/guid/3767a23c-a17d-294f-288a-9086ad1f0680/logos/primary_logo_on_white_color.png",
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
          "href": "https://a.espncdn.com/guid/3767a23c-a17d-294f-288a-9086ad1f0680/logos/primary_logo_on_black_color.png",
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
          "href": "https://a.espncdn.com/guid/3767a23c-a17d-294f-288a-9086ad1f0680/logos/primary_logo_on_primary_color.png",
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
          "href": "https://a.espncdn.com/guid/3767a23c-a17d-294f-288a-9086ad1f0680/logos/primary_logo_on_secondary_color.png",
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
          "href": "https://a.espncdn.com/guid/3767a23c-a17d-294f-288a-9086ad1f0680/logos/primary_logo_black.png",
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
          "href": "https://a.espncdn.com/guid/3767a23c-a17d-294f-288a-9086ad1f0680/logos/primary_logo_white.png",
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
          "href": "https://a.espncdn.com/guid/3767a23c-a17d-294f-288a-9086ad1f0680/logos/secondary_logo_on_white_color.png",
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
          "href": "https://a.espncdn.com/guid/3767a23c-a17d-294f-288a-9086ad1f0680/logos/secondary_logo_on_black_color.png",
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
          "href": "https://a.espncdn.com/guid/3767a23c-a17d-294f-288a-9086ad1f0680/logos/secondary_logo_on_primary_color.png",
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
          "href": "https://a.espncdn.com/guid/3767a23c-a17d-294f-288a-9086ad1f0680/logos/secondary_logo_on_secondary_color.png",
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
          "href": "https://a.espncdn.com/guid/3767a23c-a17d-294f-288a-9086ad1f0680/logos/secondary_logo_black.png",
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
          "href": "https://a.espncdn.com/guid/3767a23c-a17d-294f-288a-9086ad1f0680/logos/secondary_logo_white.png",
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
        "id": "3673",
        "guid": "35e528a7-d270-3182-b51f-612c9ad0045d",
        "fullName": "Lumen Field",
        "address": {
          "city": "Seattle",
          "state": "WA",
          "zipCode": "98134",
          "country": "USA"
        },
        "grass": false,
        "indoor": false,
        "images": [
          {
            "href": "https://a.espncdn.com/i/venues/nfl/day/3673.jpg",
            "width": 2000,
            "height": 1125,
            "alt": "",
            "rel": [
              "full",
              "day"
            ]
          },
          {
            "href": "https://a.espncdn.com/i/venues/nfl/day/interior/3673.jpg",
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
          "href": "https://www.espn.com/nfl/team/_/name/sea/seattle-seahawks",
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
          "href": "https://www.espn.com/nfl/team/roster/_/name/sea/seattle-seahawks",
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
          "href": "https://www.espn.com/nfl/team/stats/_/name/sea/seattle-seahawks",
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
          "href": "https://www.espn.com/nfl/team/schedule/_/name/sea",
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
          "href": "https://www.espn.com/nfl/team/photos/_/name/sea",
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
          "href": "https://www.vividseats.com/seattle-seahawks-tickets--sports-nfl-football/performer/772?wsUser=717",
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
          "href": "https://www.espn.com/nfl/draft/teams/_/name/sea/seattle-seahawks",
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
          "href": "https://www.espn.com/nfl/team/transactions/_/name/sea",
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
          "href": "https://www.espn.com/nfl/team/injuries/_/name/sea",
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
          "href": "https://www.espn.com/nfl/team/depth/_/name/sea",
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
    "statistics": [
      {
        "name": "gameProjection",
        "displayName": "WIN PROB",
        "value": 59.428000000000004,
        "displayValue": "59.4"
      }
    ]
  }
}
```
