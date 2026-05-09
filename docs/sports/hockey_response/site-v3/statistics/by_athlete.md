# Statistics By Athlete

## https://site.web.api.espn.com/apis/common/v3/sports/hockey/{league}/statistics/byathlete

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.web.api.espn.com/apis/common/v3/sports/hockey/nhl/statistics/byathlete`

## Example Response

```json
{
  "pagination": {
    "count": 359,
    "limit": 50,
    "page": 1,
    "pages": 8,
    "first": "http://site.api.espn.com:80/apis/common/v3/sports/hockey/nhl/statistics/byathlete?limit=50",
    "next": "http://site.api.espn.com:80/apis/common/v3/sports/hockey/nhl/statistics/byathlete?limit=50&page=2",
    "last": "http://site.api.espn.com:80/apis/common/v3/sports/hockey/nhl/statistics/byathlete?limit=50&page=8"
  },
  "league": {
    "id": "90",
    "uid": "s:70~l:90",
    "name": "National Hockey League",
    "abbreviation": "NHL",
    "slug": "nhl",
    "shortName": "NHL"
  },
  "athletes": [
    {
      "athlete": {
        "id": "3899937",
        "uid": "s:70~l:90~a:3899937",
        "guid": "7abcd6a1-150b-8797-bc35-db51cfe8939b",
        "type": "hockey",
        "firstName": "Mitch",
        "lastName": "Marner",
        "displayName": "Mitch Marner",
        "shortName": "M. Marner",
        "debutYear": 2016,
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/player/_/id/3899937",
            "text": "Player Card",
            "shortText": "Player Card",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/player/stats/_/id/3899937/mitch-marner",
            "text": "Stats",
            "shortText": "Stats",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "headshot": {
          "href": "https://a.espncdn.com/i/headshots/nhl/players/full/3899937.png",
          "alt": "Mitch Marner"
        },
        "position": {
          "id": "3",
          "name": "Right Wing",
          "displayName": "Right Wing",
          "abbreviation": "RW",
          "leaf": true,
          "parent": {
            "leaf": false
          },
          "slug": "right-wing"
        },
        "status": {
          "id": "1",
          "name": "Active",
          "type": "active",
          "abbreviation": "Active"
        },
        "age": 29,
        "teamName": "Golden Knights",
        "teamShortName": "VGK",
        "teams": [
          {
            "name": "Golden Knights",
            "abbreviation": "VGK"
          }
        ],
        "slug": "mitch-marner"
      },
      "categories": [
        {
          "name": "general",
          "displayName": "Own General",
          "totals": [
            "9",
            "8"
          ],
          "values": [
            9.0,
            8.0
          ],
          "ranks": [
            "9",
            "3"
          ]
        },
        {
          "name": "offensive",
          "displayName": "Own Offensive",
          "totals": [
            "6",
            "7"
          ],
          "values": [
            6.0,
            7.0
          ],
          "ranks": [
            "1",
            "5"
          ]
        }
      ]
    },
    {
      "athlete": {
        "id": "3648002",
        "uid": "s:70~l:90~a:3648002",
        "guid": "656e0c4e-5914-3a54-d2de-1242b6f15b5e",
        "type": "hockey",
        "firstName": "Jack",
        "lastName": "Eichel",
        "displayName": "Jack Eichel",
        "shortName": "J. Eichel",
        "debutYear": 2015,
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/player/_/id/3648002",
            "text": "Player Card",
            "shortText": "Player Card",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/player/stats/_/id/3648002/jack-eichel",
            "text": "Stats",
            "shortText": "Stats",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "headshot": {
          "href": "https://a.espncdn.com/i/headshots/nhl/players/full/3648002.png",
          "alt": "Jack Eichel"
        },
        "position": {
          "id": "1",
          "name": "Center",
          "displayName": "Center",
          "abbreviation": "C",
          "leaf": true,
          "parent": {
            "leaf": false
          },
          "slug": "center"
        },
        "status": {
          "id": "1",
          "name": "Active",
          "type": "active",
          "abbreviation": "Active"
        },
        "age": 29,
        "teamName": "Golden Knights",
        "teamShortName": "VGK",
        "teams": [
          {
            "name": "Golden Knights",
            "abbreviation": "VGK"
          }
        ],
        "slug": "jack-eichel"
      },
      "categories": [
        {
          "name": "general",
          "displayName": "Own General",
          "splitId": "0",
          "totals": [
            "9",
            "2"
          ],
          "values": [
            9.0,
            2.0
          ],
          "ranks": [
            "9",
            "59"
          ]
        },
        {
          "name": "offensive",
          "displayName": "Own Offensive",
          "splitId": "0",
          "totals": [
            "1",
            "10"
          ],
          "values": [
            1.0,
            10.0
          ],
          "ranks": [
            "88",
            "1"
          ]
        }
      ]
    }
  ],
  "currentSeason": {
    "year": 2026,
    "displayName": "2025-26",
    "startDate": "2025-09-20T07:00:00.000+00:00",
    "endDate": "2026-07-01T06:59:00.000+00:00",
    "type": {
      "id": "3",
      "type": 3,
      "name": "Postseason",
      "startDate": "2026-04-18T07:00:00.000+00:00",
      "endDate": "2026-07-01T06:59:00.000+00:00",
      "week": {
        "number": 4,
        "startDate": "2026-05-09T07:00:00.000+00:00",
        "endDate": "2026-05-16T06:59:00.000+00:00",
        "text": "Week 4"
      }
    }
  },
  "requestedSeason": {
    "year": 2026,
    "displayName": "2025-26",
    "startDate": "2025-09-20T07:00:00.000+00:00",
    "endDate": "2026-07-01T06:59:00.000+00:00",
    "type": {
      "id": "3",
      "type": 3,
      "name": "Postseason",
      "startDate": "2026-04-18T07:00:00.000+00:00",
      "endDate": "2026-07-01T06:59:00.000+00:00",
      "week": {
        "number": 4,
        "startDate": "2026-05-09T07:00:00.000+00:00",
        "endDate": "2026-05-16T06:59:00.000+00:00",
        "text": "Week 4"
      }
    }
  },
  "glossary": [
    {
      "abbreviation": "+/-",
      "displayName": "Plus/Minus Rating"
    },
    {
      "abbreviation": "A",
      "displayName": "Assists"
    }
  ],
  "categories": [
    {
      "name": "general",
      "displayName": "Own General",
      "labels": [
        "GP",
        "+/-"
      ],
      "names": [
        "games",
        "plusMinus"
      ],
      "displayNames": [
        "Games Played",
        "Plus/Minus Rating"
      ],
      "descriptions": [
        "Total games played.",
        "Plus/Minus Rating"
      ]
    },
    {
      "name": "offensive",
      "displayName": "Own Offensive",
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
      "descriptions": [
        "Total goals scored.",
        "Total goal assists."
      ]
    }
  ]
}
```
