# Competition Competitor Leaders

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}/leaders

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988`, `competitor=17` on 2026-05-08.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17/leaders?lang=en&region=us",
  "name": "leaders",
  "abbreviation": "LDR",
  "categories": [
    {
      "name": "passingLeader",
      "displayName": "Passing Leader",
      "leaders": [
        {
          "displayValue": "27/43, 295 YDS, 2 TD, 2 INT",
          "athlete": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/athletes/4431452?lang=en&region=us",
            "id": "4431452",
            "uid": "s:20~l:28~a:4431452",
            "guid": "2fbd2f9e-624d-3e63-9bb5-9bb965782b68",
            "type": "football",
            "alternateIds": {
              "sdr": "4431452"
            },
            "firstName": "Drake",
            "lastName": "Maye",
            "fullName": "Drake Maye",
            "displayName": "Drake Maye",
            "shortName": "D. Maye",
            "weight": 225.0,
            "displayWeight": "225 lbs",
            "height": 76.0,
            "displayHeight": "6' 4\"",
            "age": 23,
            "dateOfBirth": "2002-08-30T07:00Z",
            "links": [
              {
                "language": "en-US",
                "rel": [
                  "playercard",
                  "desktop",
                  "athlete"
                ],
                "href": "https://www.espn.com/nfl/player/_/id/4431452/drake-maye",
                "text": "Player Card",
                "shortText": "Player Card",
                "isExternal": false,
                "isPremium": false
              },
              {
                "language": "en-US",
                "rel": [
                  "stats",
                  "desktop",
                  "athlete"
                ],
                "href": "https://www.espn.com/nfl/player/stats/_/id/4431452/drake-maye",
                "text": "Stats",
                "shortText": "Stats",
                "isExternal": false,
                "isPremium": false
              },
              {
                "language": "en-US",
                "rel": [
                  "splits",
                  "desktop",
                  "athlete"
                ],
                "href": "https://www.espn.com/nfl/player/splits/_/id/4431452/drake-maye",
                "text": "Splits",
                "shortText": "Splits",
                "isExternal": false,
                "isPremium": false
              },
              {
                "language": "en-US",
                "rel": [
                  "gamelog",
                  "desktop",
                  "athlete"
                ],
                "href": "https://www.espn.com/nfl/player/gamelog/_/id/4431452/drake-maye",
                "text": "Game Log",
                "shortText": "Game Log",
                "isExternal": false,
                "isPremium": false
              },
              {
                "language": "en-US",
                "rel": [
                  "news",
                  "desktop",
                  "athlete"
                ],
                "href": "https://www.espn.com/nfl/player/news/_/id/4431452/drake-maye",
                "text": "News",
                "shortText": "News",
                "isExternal": false,
                "isPremium": false
              },
              {
                "language": "en-US",
                "rel": [
                  "bio",
                  "desktop",
                  "athlete"
                ],
                "href": "https://www.espn.com/nfl/player/bio/_/id/4431452/drake-maye",
                "text": "Bio",
                "shortText": "Bio",
                "isExternal": false,
                "isPremium": false
              },
              {
                "language": "en-US",
                "rel": [
                  "overview",
                  "desktop",
                  "athlete"
                ],
                "href": "https://www.espn.com/nfl/player/_/id/4431452/drake-maye",
                "text": "Overview",
                "shortText": "Overview",
                "isExternal": false,
                "isPremium": false
              }
            ],
            "birthPlace": {
              "city": "Huntersville",
              "state": "NC",
              "country": "USA"
            },
            "college": {},
            "slug": "drake-maye",
            "headshot": {
              "href": "https://a.espncdn.com/i/headshots/nfl/players/full/4431452.png",
              "alt": "Drake Maye"
            },
            "jersey": "10",
            "position": {
              "id": "8",
              "name": "Quarterback",
              "displayName": "Quarterback",
              "abbreviation": "QB",
              "leaf": true,
              "parent": {}
            },
            "injuries": [],
            "linked": true,
            "team": {},
            "teams": [
              {}
            ],
            "statistics": {},
            "projections": {},
            "notes": {},
            "contracts": {},
            "experience": {
              "years": 3
            },
            "collegeAthlete": {},
            "active": true,
            "eventLog": {},
            "draft": {
              "displayText": "Year: 2024 Round: 1 Pick: 3",
              "round": 1,
              "year": 2024,
              "selection": 3,
              "team": {},
              "pick": {}
            },
            "status": {
              "id": "1",
              "name": "Active",
              "type": "active",
              "abbreviation": "Active"
            }
          }
        }
      ]
    },
    {
      "name": "rushingLeader",
      "displayName": "Rushing Leader"
    },
    {
      "name": "receivingLeader",
      "displayName": "Receiving Leader"
    }
  ]
}
```
