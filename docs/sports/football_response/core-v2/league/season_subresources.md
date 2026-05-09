# Season Subresources

Verified with `league=nfl` and `season=2025` on 2026-05-08.

---

## Season Athletes

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/seasons/{season}/athletes`

```json
{
  "count": 20155,
  "pageIndex": 1,
  "pageSize": 1,
  "pageCount": 20155,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/athletes/4246273?lang=en&region=us",
      "id": "4246273",
      "uid": "s:20~l:28~a:4246273",
      "guid": "64cf2ccf-8d24-d1be-3ff4-d46499bdd9a6",
      "type": "football",
      "alternateIds": {
        "sdr": "4246273"
      },
      "firstName": "",
      "lastName": "[35]",
      "fullName": " [35]",
      "displayName": " [35]",
      "shortName": "[35]",
      "links": [
        {
          "language": "en-US",
          "rel": [
            "playercard",
            "desktop",
            "athlete"
          ],
          "href": "https://www.espn.com/nfl/player/_/id/4246273/35",
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
          "href": "https://www.espn.com/nfl/player/stats/_/id/4246273/35",
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
          "href": "https://www.espn.com/nfl/player/splits/_/id/4246273/35",
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
          "href": "https://www.espn.com/nfl/player/gamelog/_/id/4246273/35",
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
          "href": "https://www.espn.com/nfl/player/news/_/id/4246273/35",
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
          "href": "https://www.espn.com/nfl/player/bio/_/id/4246273/35",
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
          "href": "https://www.espn.com/nfl/player/_/id/4246273/35",
          "text": "Overview",
          "shortText": "Overview",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "birthPlace": {},
      "slug": "35",
      "jersey": "35",
      "position": {
        "id": "4",
        "name": "Center",
        "displayName": "Center",
        "abbreviation": "C",
        "leaf": true,
        "parent": {}
      },
      "injuries": [],
      "linked": true,
      "team": {},
      "notes": {},
      "contracts": {},
      "active": false,
      "eventLog": {},
      "status": {
        "id": "2",
        "name": "Inactive",
        "type": "inactive",
        "abbreviation": "Inactive"
      }
    }
  ]
}
```

---

## Season Draft

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/seasons/{season}/draft`

Notes:
- `season=2025` returned a fully populated draft object.
- This is an object response, not a paginated collection.

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/draft?lang=en&region=us",
  "uid": "s:20~l:28~e:draft~y:2025",
  "year": 2025,
  "numberOfRounds": 7,
  "displayName": "2025 National Football League Draft",
  "shortDisplayName": "2025 NFL Draft",
  "status": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/draft/status?lang=en&region=us",
    "round": 7,
    "type": {
      "id": 3,
      "name": "COMPLETED",
      "state": "post",
      "description": "Completed"
    }
  },
  "athletes": {
    "count": 926,
    "pageIndex": 1,
    "pageSize": 25,
    "pageCount": 38,
    "items": [
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {},
      {}
    ]
  },
  "rounds": {
    "count": 7,
    "pageIndex": 1,
    "pageSize": 25,
    "pageCount": 1,
    "items": [
      {
        "number": 1,
        "displayName": "1st Round",
        "shortDisplayName": "1st",
        "picks": [
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 1,
            "overall": 1,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 2,
            "overall": 2,
            "round": 1,
            "traded": true,
            "tradeNote": "From CLE",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 3,
            "overall": 3,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 4,
            "overall": 4,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 5,
            "overall": 5,
            "round": 1,
            "traded": true,
            "tradeNote": "From JAX",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 6,
            "overall": 6,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 7,
            "overall": 7,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 8,
            "overall": 8,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 9,
            "overall": 9,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 10,
            "overall": 10,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 11,
            "overall": 11,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 12,
            "overall": 12,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 13,
            "overall": 13,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 14,
            "overall": 14,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 15,
            "overall": 15,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 16,
            "overall": 16,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 17,
            "overall": 17,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 18,
            "overall": 18,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 19,
            "overall": 19,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 20,
            "overall": 20,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 21,
            "overall": 21,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 22,
            "overall": 22,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 23,
            "overall": 23,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 24,
            "overall": 24,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 25,
            "overall": 25,
            "round": 1,
            "traded": true,
            "tradeNote": "From HOU",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 26,
            "overall": 26,
            "round": 1,
            "traded": true,
            "tradeNote": "From LAR",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 27,
            "overall": 27,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 28,
            "overall": 28,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 29,
            "overall": 29,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 30,
            "overall": 30,
            "round": 1,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 31,
            "overall": 31,
            "round": 1,
            "traded": true,
            "tradeNote": "From KC",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 32,
            "overall": 32,
            "round": 1,
            "traded": true,
            "tradeNote": "From PHI",
            "athlete": {},
            "team": {}
          }
        ],
        "status": {
          "round": 1,
          "type": {
            "id": 3,
            "name": "COMPLETED",
            "state": "post",
            "description": "Completed"
          }
        }
      },
      {
        "number": 2,
        "displayName": "2nd Round",
        "shortDisplayName": "2nd",
        "picks": [
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 1,
            "overall": 33,
            "round": 2,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 2,
            "overall": 34,
            "round": 2,
            "traded": true,
            "tradeNote": "From NYG",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 3,
            "overall": 35,
            "round": 2,
            "traded": true,
            "tradeNote": "From TEN",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 4,
            "overall": 36,
            "round": 2,
            "traded": true,
            "tradeNote": "From JAX",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 5,
            "overall": 37,
            "round": 2,
            "traded": true,
            "tradeNote": "From LV",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 6,
            "overall": 38,
            "round": 2,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 7,
            "overall": 39,
            "round": 2,
            "traded": true,
            "tradeNote": "From CAR",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 8,
            "overall": 40,
            "round": 2,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 9,
            "overall": 41,
            "round": 2,
            "traded": true,
            "tradeNote": "From CHI",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 10,
            "overall": 42,
            "round": 2,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 11,
            "overall": 43,
            "round": 2,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 12,
            "overall": 44,
            "round": 2,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 13,
            "overall": 45,
            "round": 2,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 14,
            "overall": 46,
            "round": 2,
            "traded": true,
            "tradeNote": "From ATL",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 15,
            "overall": 47,
            "round": 2,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 16,
            "overall": 48,
            "round": 2,
            "traded": true,
            "tradeNote": "From MIA through LV",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 17,
            "overall": 49,
            "round": 2,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 18,
            "overall": 50,
            "round": 2,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 19,
            "overall": 51,
            "round": 2,
            "traded": true,
            "tradeNote": "From DEN",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 20,
            "overall": 52,
            "round": 2,
            "traded": true,
            "tradeNote": "From PIT through SEA",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 21,
            "overall": 53,
            "round": 2,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 22,
            "overall": 54,
            "round": 2,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 23,
            "overall": 55,
            "round": 2,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 24,
            "overall": 56,
            "round": 2,
            "traded": true,
            "tradeNote": "From MIN through HOU and BUF",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 25,
            "overall": 57,
            "round": 2,
            "traded": true,
            "tradeNote": "From LAR through CAR and DEN",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 26,
            "overall": 58,
            "round": 2,
            "traded": true,
            "tradeNote": "From HOU",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 27,
            "overall": 59,
            "round": 2,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 28,
            "overall": 60,
            "round": 2,
            "traded": true,
            "tradeNote": "From DET",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 29,
            "overall": 61,
            "round": 2,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 30,
            "overall": 62,
            "round": 2,
            "traded": true,
            "tradeNote": "From BUF",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 31,
            "overall": 63,
            "round": 2,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 32,
            "overall": 64,
            "round": 2,
            "traded": false,
            "athlete": {},
            "team": {}
          }
        ],
        "status": {
          "round": 2,
          "type": {
            "id": 3,
            "name": "COMPLETED",
            "state": "post",
            "description": "Completed"
          }
        }
      },
      {
        "number": 3,
        "displayName": "3rd Round",
        "shortDisplayName": "3rd",
        "picks": [
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 1,
            "overall": 65,
            "round": 3,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 2,
            "overall": 66,
            "round": 3,
            "traded": true,
            "tradeNote": "From TEN",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 3,
            "overall": 67,
            "round": 3,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 4,
            "overall": 68,
            "round": 3,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 5,
            "overall": 69,
            "round": 3,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 6,
            "overall": 70,
            "round": 3,
            "traded": true,
            "tradeNote": "From JAX",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 7,
            "overall": 71,
            "round": 3,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 8,
            "overall": 72,
            "round": 3,
            "traded": true,
            "tradeNote": "From CHI",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 9,
            "overall": 73,
            "round": 3,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 10,
            "overall": 74,
            "round": 3,
            "traded": true,
            "tradeNote": "From CAR",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 11,
            "overall": 75,
            "round": 3,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 12,
            "overall": 76,
            "round": 3,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 13,
            "overall": 77,
            "round": 3,
            "traded": true,
            "tradeNote": "From ATL through NE",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 14,
            "overall": 78,
            "round": 3,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 15,
            "overall": 79,
            "round": 3,
            "traded": true,
            "tradeNote": "From MIA through PHI and WSH",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 16,
            "overall": 80,
            "round": 3,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 17,
            "overall": 81,
            "round": 3,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 18,
            "overall": 82,
            "round": 3,
            "traded": true,
            "tradeNote": "From SEA",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 19,
            "overall": 83,
            "round": 3,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 20,
            "overall": 84,
            "round": 3,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 21,
            "overall": 85,
            "round": 3,
            "traded": true,
            "tradeNote": "From DEN through CAR and NE",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 22,
            "overall": 86,
            "round": 3,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 23,
            "overall": 87,
            "round": 3,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 24,
            "overall": 88,
            "round": 3,
            "traded": true,
            "tradeNote": "From MIN",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 25,
            "overall": 89,
            "round": 3,
            "traded": true,
            "tradeNote": "From HOU",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 26,
            "overall": 90,
            "round": 3,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 27,
            "overall": 91,
            "round": 3,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 28,
            "overall": 92,
            "round": 3,
            "traded": true,
            "tradeNote": "From DET through NYJ and LV",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 29,
            "overall": 93,
            "round": 3,
            "traded": true,
            "tradeNote": "From WSH",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 30,
            "overall": 94,
            "round": 3,
            "traded": true,
            "tradeNote": "From BUF",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 31,
            "overall": 95,
            "round": 3,
            "traded": true,
            "tradeNote": "From KC",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 32,
            "overall": 96,
            "round": 3,
            "traded": true,
            "tradeNote": "From PHI",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 33,
            "overall": 97,
            "round": 3,
            "traded": true,
            "tradeNote": "(Compensatory Selection From MIN)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 34,
            "overall": 98,
            "round": 3,
            "traded": true,
            "tradeNote": "(Compensatory Selection From MIA)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 35,
            "overall": 99,
            "round": 3,
            "traded": true,
            "tradeNote": "(Compensatory Selection From NYG through HOU)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 36,
            "overall": 100,
            "round": 3,
            "traded": false,
            "tradeNote": "(Special Compensatory Selection)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 37,
            "overall": 101,
            "round": 3,
            "traded": true,
            "tradeNote": "(Special Compensatory Selection From LAR through ATL and PHI)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 38,
            "overall": 102,
            "round": 3,
            "traded": true,
            "tradeNote": "(Special Compensatory Selection From DET through JAX and HOU)",
            "athlete": {},
            "team": {}
          }
        ],
        "status": {
          "round": 3,
          "type": {
            "id": 3,
            "name": "COMPLETED",
            "state": "post",
            "description": "Completed"
          }
        }
      },
      {
        "number": 4,
        "displayName": "4th Round",
        "shortDisplayName": "4th",
        "picks": [
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 1,
            "overall": 103,
            "round": 4,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 2,
            "overall": 104,
            "round": 4,
            "traded": true,
            "tradeNote": "From CLE",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 3,
            "overall": 105,
            "round": 4,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 4,
            "overall": 106,
            "round": 4,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 5,
            "overall": 107,
            "round": 4,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 6,
            "overall": 108,
            "round": 4,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 7,
            "overall": 109,
            "round": 4,
            "traded": true,
            "tradeNote": "From CHI through BUF and CHI",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 8,
            "overall": 110,
            "round": 4,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 9,
            "overall": 111,
            "round": 4,
            "traded": true,
            "tradeNote": "From CAR through DEN",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 10,
            "overall": 112,
            "round": 4,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 11,
            "overall": 113,
            "round": 4,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 12,
            "overall": 114,
            "round": 4,
            "traded": true,
            "tradeNote": "From DAL",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 13,
            "overall": 115,
            "round": 4,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 14,
            "overall": 116,
            "round": 4,
            "traded": true,
            "tradeNote": "From MIA",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 15,
            "overall": 117,
            "round": 4,
            "traded": true,
            "tradeNote": "From IND",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 16,
            "overall": 118,
            "round": 4,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 17,
            "overall": 119,
            "round": 4,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 18,
            "overall": 120,
            "round": 4,
            "traded": true,
            "tradeNote": "From SEA",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 19,
            "overall": 121,
            "round": 4,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 20,
            "overall": 122,
            "round": 4,
            "traded": true,
            "tradeNote": "From DEN",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 21,
            "overall": 123,
            "round": 4,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 22,
            "overall": 124,
            "round": 4,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 23,
            "overall": 125,
            "round": 4,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 24,
            "overall": 126,
            "round": 4,
            "traded": true,
            "tradeNote": "From MIN through JAX",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 25,
            "overall": 127,
            "round": 4,
            "traded": true,
            "tradeNote": "From LAR",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 26,
            "overall": 128,
            "round": 4,
            "traded": true,
            "tradeNote": "From HOU",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 27,
            "overall": 129,
            "round": 4,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 28,
            "overall": 130,
            "round": 4,
            "traded": true,
            "tradeNote": "From DET through DEN and PHI",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 29,
            "overall": 131,
            "round": 4,
            "traded": true,
            "tradeNote": "From WSH",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 30,
            "overall": 132,
            "round": 4,
            "traded": true,
            "tradeNote": "From BUF",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 31,
            "overall": 133,
            "round": 4,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 32,
            "overall": 134,
            "round": 4,
            "traded": true,
            "tradeNote": "From PHI through DET and PHI",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 33,
            "overall": 135,
            "round": 4,
            "traded": true,
            "tradeNote": "(Compensatory Selection From MIA)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 34,
            "overall": 136,
            "round": 4,
            "traded": true,
            "tradeNote": "(Compensatory Selection From BAL)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 35,
            "overall": 137,
            "round": 4,
            "traded": true,
            "tradeNote": "(Compensatory Selection From SEA)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 36,
            "overall": 138,
            "round": 4,
            "traded": false,
            "tradeNote": "(Compensatory Selection)",
            "athlete": {},
            "team": {}
          }
        ],
        "status": {
          "round": 4,
          "type": {
            "id": 3,
            "name": "COMPLETED",
            "state": "post",
            "description": "Completed"
          }
        }
      },
      {
        "number": 5,
        "displayName": "5th Round",
        "shortDisplayName": "5th",
        "picks": [
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 1,
            "overall": 139,
            "round": 5,
            "traded": true,
            "tradeNote": "From CLE",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 2,
            "overall": 140,
            "round": 5,
            "traded": true,
            "tradeNote": "From NYG",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 3,
            "overall": 141,
            "round": 5,
            "traded": true,
            "tradeNote": "From TEN",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 4,
            "overall": 142,
            "round": 5,
            "traded": true,
            "tradeNote": "From JAX through HOU and MIN",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 5,
            "overall": 143,
            "round": 5,
            "traded": true,
            "tradeNote": "From LV",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 6,
            "overall": 144,
            "round": 5,
            "traded": true,
            "tradeNote": "From NE through SEA",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 7,
            "overall": 145,
            "round": 5,
            "traded": true,
            "tradeNote": "From NYJ",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 8,
            "overall": 146,
            "round": 5,
            "traded": true,
            "tradeNote": "From CAR",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 9,
            "overall": 147,
            "round": 5,
            "traded": true,
            "tradeNote": "From NO through WSH",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 10,
            "overall": 148,
            "round": 5,
            "traded": true,
            "tradeNote": "From CHI",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 12,
            "overall": 149,
            "round": 5,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 13,
            "overall": 150,
            "round": 5,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 14,
            "overall": 151,
            "round": 5,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 16,
            "overall": 152,
            "round": 5,
            "traded": true,
            "tradeNote": "From ARI",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 17,
            "overall": 153,
            "round": 5,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 18,
            "overall": 154,
            "round": 5,
            "traded": true,
            "tradeNote": "From SEA",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 19,
            "overall": 155,
            "round": 5,
            "traded": true,
            "tradeNote": "From DEN",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 20,
            "overall": 156,
            "round": 5,
            "traded": true,
            "tradeNote": "From PIT",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 21,
            "overall": 157,
            "round": 5,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 22,
            "overall": 158,
            "round": 5,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 23,
            "overall": 159,
            "round": 5,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 24,
            "overall": 160,
            "round": 5,
            "traded": true,
            "tradeNote": "From MIN",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 25,
            "overall": 161,
            "round": 5,
            "traded": true,
            "tradeNote": "From HOU",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 26,
            "overall": 162,
            "round": 5,
            "traded": true,
            "tradeNote": "From LAR through PIT",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 27,
            "overall": 163,
            "round": 5,
            "traded": true,
            "tradeNote": "From BAL",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 28,
            "overall": 164,
            "round": 5,
            "traded": true,
            "tradeNote": "From DET through CLE, PHI and KC",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 29,
            "overall": 165,
            "round": 5,
            "traded": true,
            "tradeNote": "From WSH through PHI",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 30,
            "overall": 166,
            "round": 5,
            "traded": true,
            "tradeNote": "From BUF through HOU and CLE",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 31,
            "overall": 167,
            "round": 5,
            "traded": true,
            "tradeNote": "From KC",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 32,
            "overall": 168,
            "round": 5,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 33,
            "overall": 169,
            "round": 5,
            "traded": true,
            "tradeNote": "(Compensatory Selection From BUF)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 34,
            "overall": 170,
            "round": 5,
            "traded": true,
            "tradeNote": "(Compensatory Selection From DAL)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 35,
            "overall": 171,
            "round": 5,
            "traded": true,
            "tradeNote": "(Compensatory Selection From DAL through NE)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 36,
            "overall": 172,
            "round": 5,
            "traded": true,
            "tradeNote": "(Compensatory Selection From SEA through MIN)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 37,
            "overall": 173,
            "round": 5,
            "traded": false,
            "tradeNote": "(Compensatory Selection)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 38,
            "overall": 174,
            "round": 5,
            "traded": true,
            "tradeNote": "(Compensatory Selection From DAL)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 39,
            "overall": 175,
            "round": 5,
            "traded": false,
            "tradeNote": "(Compensatory Selection)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 40,
            "overall": 176,
            "round": 5,
            "traded": true,
            "tradeNote": "(Compensatory Selection From BAL)",
            "athlete": {},
            "team": {}
          }
        ],
        "status": {
          "round": 5,
          "type": {
            "id": 3,
            "name": "COMPLETED",
            "state": "post",
            "description": "Completed"
          }
        }
      },
      {
        "number": 6,
        "displayName": "6th Round",
        "shortDisplayName": "6th",
        "picks": [
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 1,
            "overall": 177,
            "round": 6,
            "traded": true,
            "tradeNote": "From NYG",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 2,
            "overall": 178,
            "round": 6,
            "traded": true,
            "tradeNote": "From TEN",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 3,
            "overall": 179,
            "round": 6,
            "traded": true,
            "tradeNote": "From CLE through HOU",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 4,
            "overall": 180,
            "round": 6,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 5,
            "overall": 181,
            "round": 6,
            "traded": true,
            "tradeNote": "From NE through LAC",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 6,
            "overall": 182,
            "round": 6,
            "traded": true,
            "tradeNote": "From JAX through DET",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 7,
            "overall": 183,
            "round": 6,
            "traded": true,
            "tradeNote": "From CAR through BAL",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 8,
            "overall": 184,
            "round": 6,
            "traded": true,
            "tradeNote": "From NO through WSH",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 9,
            "overall": 185,
            "round": 6,
            "traded": true,
            "tradeNote": "From CHI through SEA",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 10,
            "overall": 186,
            "round": 6,
            "traded": true,
            "tradeNote": "From NYJ",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 11,
            "overall": 187,
            "round": 6,
            "traded": true,
            "tradeNote": "From SF through MIN",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 12,
            "overall": 188,
            "round": 6,
            "traded": true,
            "tradeNote": "From DAL",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 13,
            "overall": 189,
            "round": 6,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 14,
            "overall": 190,
            "round": 6,
            "traded": true,
            "tradeNote": "From ATL through LAR",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 15,
            "overall": 191,
            "round": 6,
            "traded": true,
            "tradeNote": "From ARI through DEN",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 16,
            "overall": 192,
            "round": 6,
            "traded": true,
            "tradeNote": "From MIA through CHI and CLE",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 17,
            "overall": 193,
            "round": 6,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 18,
            "overall": 194,
            "round": 6,
            "traded": true,
            "tradeNote": "From SEA",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 19,
            "overall": 195,
            "round": 6,
            "traded": true,
            "tradeNote": "From PIT through LAR",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 20,
            "overall": 196,
            "round": 6,
            "traded": true,
            "tradeNote": "From TB",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 21,
            "overall": 197,
            "round": 6,
            "traded": true,
            "tradeNote": "From DEN",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 22,
            "overall": 198,
            "round": 6,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 23,
            "overall": 199,
            "round": 6,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 24,
            "overall": 200,
            "round": 6,
            "traded": true,
            "tradeNote": "From MIN through CLE",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 25,
            "overall": 201,
            "round": 6,
            "traded": true,
            "tradeNote": "From LAR",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 26,
            "overall": 202,
            "round": 6,
            "traded": true,
            "tradeNote": "From HOU through PIT, CHI and LAR",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 27,
            "overall": 203,
            "round": 6,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 28,
            "overall": 204,
            "round": 6,
            "traded": true,
            "tradeNote": "From DET through CLE and BUF",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 29,
            "overall": 205,
            "round": 6,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 30,
            "overall": 206,
            "round": 6,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 31,
            "overall": 207,
            "round": 6,
            "traded": true,
            "tradeNote": "From KC through NYJ",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 32,
            "overall": 208,
            "round": 6,
            "traded": true,
            "tradeNote": "From PHI through DEN",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 33,
            "overall": 209,
            "round": 6,
            "traded": true,
            "tradeNote": "(Compensatory Selection From LAC)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 34,
            "overall": 210,
            "round": 6,
            "traded": false,
            "tradeNote": "(Compensatory Selection)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 35,
            "overall": 211,
            "round": 6,
            "traded": true,
            "tradeNote": "(Compensatory Selection From DAL)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 36,
            "overall": 212,
            "round": 6,
            "traded": false,
            "tradeNote": "(Compensatory Selection)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 37,
            "overall": 213,
            "round": 6,
            "traded": false,
            "tradeNote": "(Compensatory Selection)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 38,
            "overall": 214,
            "round": 6,
            "traded": false,
            "tradeNote": "(Compensatory Selection)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 39,
            "overall": 215,
            "round": 6,
            "traded": false,
            "tradeNote": "(Compensatory Selection)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 40,
            "overall": 216,
            "round": 6,
            "traded": true,
            "tradeNote": "(Compensatory Selection from CLE through HOU)",
            "athlete": {},
            "team": {}
          }
        ],
        "status": {
          "round": 6,
          "type": {
            "id": 3,
            "name": "COMPLETED",
            "state": "post",
            "description": "Completed"
          }
        }
      },
      {
        "number": 7,
        "displayName": "7th Round",
        "shortDisplayName": "7th",
        "picks": [
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 1,
            "overall": 217,
            "round": 7,
            "traded": true,
            "tradeNote": "From TEN through NE",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 2,
            "overall": 218,
            "round": 7,
            "traded": true,
            "tradeNote": "From CLE through LAC",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 3,
            "overall": 219,
            "round": 7,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 4,
            "overall": 220,
            "round": 7,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 5,
            "overall": 221,
            "round": 7,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 6,
            "overall": 222,
            "round": 7,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 7,
            "overall": 223,
            "round": 7,
            "traded": true,
            "tradeNote": "From NO through PHI and PIT",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 8,
            "overall": 224,
            "round": 7,
            "traded": true,
            "tradeNote": "From CHI through MIA",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 9,
            "overall": 225,
            "round": 7,
            "traded": true,
            "tradeNote": "From NYJ through KC",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 10,
            "overall": 226,
            "round": 7,
            "traded": true,
            "tradeNote": "From CAR through KC",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 11,
            "overall": 227,
            "round": 7,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 12,
            "overall": 228,
            "round": 7,
            "traded": true,
            "tradeNote": "From DAL through DET and NE",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 13,
            "overall": 229,
            "round": 7,
            "traded": true,
            "tradeNote": "From ATL through PHI",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 14,
            "overall": 230,
            "round": 7,
            "traded": true,
            "tradeNote": "From ARI through CAR and DEN",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 15,
            "overall": 231,
            "round": 7,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 16,
            "overall": 232,
            "round": 7,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 17,
            "overall": 233,
            "round": 7,
            "traded": true,
            "tradeNote": "From CIN",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 18,
            "overall": 234,
            "round": 7,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 19,
            "overall": 235,
            "round": 7,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 20,
            "overall": 236,
            "round": 7,
            "traded": true,
            "tradeNote": "From DEN through PHI, WSH, and HOU",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 21,
            "overall": 237,
            "round": 7,
            "traded": true,
            "tradeNote": "From PIT",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 22,
            "overall": 238,
            "round": 7,
            "traded": true,
            "tradeNote": "From LAC through NE",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 23,
            "overall": 239,
            "round": 7,
            "traded": true,
            "tradeNote": "From GB through TEN",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 24,
            "overall": 240,
            "round": 7,
            "traded": true,
            "tradeNote": "From MIN through CLE and CHI",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 25,
            "overall": 241,
            "round": 7,
            "traded": true,
            "tradeNote": "From HOU",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 26,
            "overall": 242,
            "round": 7,
            "traded": true,
            "tradeNote": "From LAR through ATL",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 27,
            "overall": 243,
            "round": 7,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 28,
            "overall": 244,
            "round": 7,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 29,
            "overall": 245,
            "round": 7,
            "traded": false,
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 30,
            "overall": 246,
            "round": 7,
            "traded": true,
            "tradeNote": "From BUF",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 31,
            "overall": 247,
            "round": 7,
            "traded": true,
            "tradeNote": "From KC through CAR",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 32,
            "overall": 248,
            "round": 7,
            "traded": true,
            "tradeNote": "From PHI through WSH",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 33,
            "overall": 249,
            "round": 7,
            "traded": false,
            "tradeNote": "(Compensatory Selection)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 34,
            "overall": 250,
            "round": 7,
            "traded": false,
            "tradeNote": "(Compensatory Selection)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 35,
            "overall": 251,
            "round": 7,
            "traded": true,
            "tradeNote": "(Compensatory Selection From KC)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 36,
            "overall": 252,
            "round": 7,
            "traded": false,
            "tradeNote": "(Compensatory Selection)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 37,
            "overall": 253,
            "round": 7,
            "traded": false,
            "tradeNote": "(Compensatory Selection)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 38,
            "overall": 254,
            "round": 7,
            "traded": false,
            "tradeNote": "(Compensatory Selection)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 39,
            "overall": 255,
            "round": 7,
            "traded": true,
            "tradeNote": "(Compensatory Selection from CLE)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 40,
            "overall": 256,
            "round": 7,
            "traded": false,
            "tradeNote": "(Compensatory Selection)",
            "athlete": {},
            "team": {}
          },
          {
            "status": {
              "id": 3,
              "name": "SELECTION_MADE",
              "description": "Selection Made"
            },
            "pick": 41,
            "overall": 257,
            "round": 7,
            "traded": true,
            "tradeNote": "(Compensatory Selection From KC)",
            "athlete": {},
            "team": {}
          }
        ],
        "status": {
          "round": 7,
          "type": {
            "id": 3,
            "name": "COMPLETED",
            "state": "post",
            "description": "Completed"
          }
        }
      }
    ]
  },
  "startDate": "2025-04-25T00:00Z",
  "endDate": "2025-04-28T03:59Z"
}
```

---

## Season Free Agents

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/seasons/{season}/freeagents`

Notes:
- `season=2025` returned an empty collection.

```json
{
  "count": 0,
  "pageIndex": 0,
  "pageSize": 25,
  "pageCount": 0,
  "items": []
}
```

---

## Season Manufacturers

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/seasons/{season}/manufacturers`

Notes:
- NFL currently returns HTTP 400 with `getManufacturers() not supported for football/nfl`.

```json
{
  "error": {
    "message": "getManufacturers() not supported for football/nfl",
    "code": 400
  }
}
```
