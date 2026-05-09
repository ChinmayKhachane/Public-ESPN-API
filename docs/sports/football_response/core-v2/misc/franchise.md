# Franchises

Verified with `league=nfl` on 2026-05-08.

---

## Collection

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/franchises`

```json
{
  "count": 32,
  "pageIndex": 1,
  "pageSize": 1,
  "pageCount": 32,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/franchises/1?lang=en&region=us",
      "id": "1",
      "uid": "s:20~l:28~f:1",
      "slug": "atlanta-falcons",
      "location": "Atlanta",
      "name": "Falcons",
      "nickname": "Falcons",
      "abbreviation": "ATL",
      "displayName": "Atlanta Falcons",
      "shortDisplayName": "Falcons",
      "color": "a71930",
      "isActive": true,
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
      "team": {}
    }
  ]
}
```

---

## Franchise ID

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/franchises/{id}`

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/franchises/1?lang=en&region=us",
  "id": "1",
  "uid": "s:20~l:28~f:1",
  "name": "Falcons",
  "displayName": "Atlanta Falcons",
  "abbreviation": "ATL",
  "location": "Atlanta",
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/teams/1?lang=en&region=us",
    "id": "1",
    "guid": "49fd392a-86fe-4df3-1b77-9bbfa18b2ad5",
    "uid": "s:20~l:28~t:1",
    "alternateIds": {
      "sdr": "8802"
    },
    "slug": "atlanta-falcons",
    "location": "Atlanta",
    "name": "Falcons",
    "nickname": "Falcons",
    "abbreviation": "ATL",
    "displayName": "Atlanta Falcons",
    "shortDisplayName": "Falcons",
    "color": "a71930",
    "alternateColor": "000000",
    "isActive": true,
    "isAllStar": false,
    "logos": [
      {
        "href": "https://a.espncdn.com/i/teamlogos/nfl/500/atl.png",
        "width": 500,
        "height": 500,
        "alt": "",
        "rel": [
          "full",
          "default"
        ],
        "lastUpdated": "2024-06-25T18:44Z"
      },
      {
        "href": "https://a.espncdn.com/i/teamlogos/nfl/500-dark/atl.png",
        "width": 500,
        "height": 500,
        "alt": "",
        "rel": [
          "full",
          "dark"
        ],
        "lastUpdated": "2024-06-25T18:44Z"
      },
      {
        "href": "https://a.espncdn.com/i/teamlogos/nfl/500/scoreboard/atl.png",
        "width": 500,
        "height": 500,
        "alt": "",
        "rel": [
          "full",
          "scoreboard"
        ],
        "lastUpdated": "2024-06-25T18:44Z"
      },
      {
        "href": "https://a.espncdn.com/i/teamlogos/nfl/500-dark/scoreboard/atl.png",
        "width": 500,
        "height": 500,
        "alt": "",
        "rel": [
          "full",
          "scoreboard",
          "dark"
        ],
        "lastUpdated": "2024-06-25T18:44Z"
      },
      {
        "href": "https://a.espncdn.com/guid/49fd392a-86fe-4df3-1b77-9bbfa18b2ad5/logos/grayscale.png",
        "width": 500,
        "height": 500,
        "alt": "",
        "rel": [
          "full",
          "grayscale"
        ],
        "lastUpdated": "2026-03-31T12:54Z"
      },
      {
        "href": "https://a.espncdn.com/guid/49fd392a-86fe-4df3-1b77-9bbfa18b2ad5/logos/primary_logo_on_white_color.png",
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
        "href": "https://a.espncdn.com/guid/49fd392a-86fe-4df3-1b77-9bbfa18b2ad5/logos/primary_logo_on_black_color.png",
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
        "href": "https://a.espncdn.com/guid/49fd392a-86fe-4df3-1b77-9bbfa18b2ad5/logos/primary_logo_on_primary_color.png",
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
        "href": "https://a.espncdn.com/guid/49fd392a-86fe-4df3-1b77-9bbfa18b2ad5/logos/primary_logo_on_secondary_color.png",
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
        "href": "https://a.espncdn.com/guid/49fd392a-86fe-4df3-1b77-9bbfa18b2ad5/logos/primary_logo_black.png",
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
        "href": "https://a.espncdn.com/guid/49fd392a-86fe-4df3-1b77-9bbfa18b2ad5/logos/primary_logo_white.png",
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
        "href": "https://a.espncdn.com/guid/49fd392a-86fe-4df3-1b77-9bbfa18b2ad5/logos/secondary_logo_on_white_color.png",
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
        "href": "https://a.espncdn.com/guid/49fd392a-86fe-4df3-1b77-9bbfa18b2ad5/logos/secondary_logo_on_black_color.png",
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
        "href": "https://a.espncdn.com/guid/49fd392a-86fe-4df3-1b77-9bbfa18b2ad5/logos/secondary_logo_on_primary_color.png",
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
        "href": "https://a.espncdn.com/guid/49fd392a-86fe-4df3-1b77-9bbfa18b2ad5/logos/secondary_logo_on_secondary_color.png",
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
        "href": "https://a.espncdn.com/guid/49fd392a-86fe-4df3-1b77-9bbfa18b2ad5/logos/secondary_logo_black.png",
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
        "href": "https://a.espncdn.com/guid/49fd392a-86fe-4df3-1b77-9bbfa18b2ad5/logos/secondary_logo_white.png",
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
        "href": "https://www.espn.com/nfl/team/_/name/atl/atlanta-falcons",
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
        "href": "https://www.espn.com/nfl/team/roster/_/name/atl/atlanta-falcons",
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
        "href": "https://www.espn.com/nfl/team/stats/_/name/atl/atlanta-falcons",
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
        "href": "https://www.espn.com/nfl/team/schedule/_/name/atl",
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
        "href": "https://www.espn.com/nfl/team/photos/_/name/atl",
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
        "href": "https://www.vividseats.com/atlanta-falcons-tickets--sports-nfl-football/performer/51?wsUser=717",
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
        "href": "https://www.espn.com/nfl/draft/teams/_/name/atl/atlanta-falcons",
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
        "href": "https://www.espn.com/nfl/team/transactions/_/name/atl",
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
        "href": "https://www.espn.com/nfl/team/injuries/_/name/atl",
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
        "href": "https://www.espn.com/nfl/team/depth/_/name/atl",
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
  "venue": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/venues/5348?lang=en&region=us",
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
  }
}
```
