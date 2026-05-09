# Event Competition

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988` on 2026-05-08.
- This is the richest football core-v2 event object. Most useful subresources hang off this response.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `period` | `int` | Used by some competition child endpoints |
| `sort` | `string` | Used by some competition child endpoints |
| `source` | `string` | Used by some competition child endpoints |
| `showsubplays` | `bool` | Used by some competition child endpoints |

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988?lang=en&region=us",
  "id": "401772988",
  "uid": "s:20~l:28~e:401772988~c:401772988",
  "date": "2026-02-08T23:30Z",
  "attendance": 70823,
  "neutralSite": true,
  "boxscoreAvailable": true,
  "playByPlayAvailable": true,
  "summaryAvailable": true,
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
      "team": {},
      "score": {},
      "linescores": {},
      "roster": {},
      "statistics": {},
      "leaders": {},
      "record": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/26?lang=en&region=us",
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
  "broadcasts": {
    "count": 2,
    "pageIndex": 1,
    "pageSize": 25,
    "pageCount": 1,
    "items": [
      {
        "type": {
          "id": "1",
          "shortName": "TV",
          "longName": "Television",
          "slug": "tv"
        },
        "channel": 379,
        "station": "NBC",
        "slug": "nbc",
        "priority": 1,
        "market": {
          "id": "1",
          "type": "National"
        },
        "media": {
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
        },
        "lang": "en",
        "region": "us",
        "competition": {},
        "partnered": false
      },
      {
        "type": {
          "id": "4",
          "shortName": "Streaming",
          "longName": "Streaming",
          "slug": "streaming"
        },
        "channel": 789,
        "station": "Peacock",
        "slug": "peacock",
        "priority": 2,
        "market": {
          "id": "1",
          "type": "National"
        },
        "media": {
          "id": "789",
          "callLetters": "Peacock",
          "name": "Peacock",
          "shortName": "Peacock",
          "slug": "peacock",
          "logos": [
            {
              "href": "https://a.espncdn.com/guid/9855b9a2-1dc7-3a5e-804b-871b6a889e8d/logos/default.png",
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
              "href": "https://a.espncdn.com/guid/9855b9a2-1dc7-3a5e-804b-871b6a889e8d/logos/default-dark.png",
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
        },
        "lang": "en",
        "region": "us",
        "competition": {},
        "partnered": false
      }
    ]
  },
  "odds": {
    "count": 1,
    "pageIndex": 1,
    "pageSize": 25,
    "pageCount": 1,
    "items": [
      {
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
  },
  "officials": {
    "count": 7,
    "pageIndex": 1,
    "pageSize": 25,
    "pageCount": 1,
    "items": [
      {
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
      },
      {
        "id": "2615667",
        "firstName": "Greg",
        "lastName": "Steed",
        "fullName": "Greg Steed",
        "displayName": "Greg Steed",
        "position": {
          "name": "Back Judge",
          "displayName": "Back Judge",
          "id": "110"
        },
        "order": 2
      },
      {
        "id": "2615961",
        "firstName": "Roy",
        "lastName": "Ellison",
        "fullName": "Roy Ellison",
        "displayName": "Roy Ellison",
        "position": {
          "name": "Umpire",
          "displayName": "Umpire",
          "id": "116"
        },
        "order": 3
      },
      {
        "id": "2616811",
        "firstName": "Julian",
        "lastName": "Mapp",
        "fullName": "Julian Mapp",
        "displayName": "Julian Mapp",
        "position": {
          "name": "Line Judge",
          "displayName": "Line Judge",
          "id": "113"
        },
        "order": 4
      },
      {
        "id": "3055221",
        "firstName": "Eugene",
        "lastName": "Hall",
        "fullName": "Eugene Hall",
        "displayName": "Eugene Hall",
        "position": {
          "name": "Side Judge",
          "displayName": "Side Judge",
          "id": "115"
        },
        "order": 5
      },
      {
        "id": "3132039",
        "firstName": "Shawn",
        "lastName": "Smith",
        "fullName": "Shawn Smith",
        "displayName": "Shawn Smith",
        "position": {
          "name": "Referee",
          "displayName": "Referee",
          "id": "114"
        },
        "order": 6
      },
      {
        "id": "5157996",
        "firstName": "Jason",
        "lastName": "Ledet",
        "fullName": "Jason Ledet",
        "displayName": "Jason Ledet",
        "position": {
          "name": "Field Judge",
          "displayName": "Field Judge",
          "id": "111"
        },
        "order": 7
      }
    ]
  },
  "situation": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/situation?lang=en&region=us",
    "lastPlay": {},
    "down": 0,
    "yardLine": 71,
    "distance": 0,
    "isRedZone": false,
    "homeTimeouts": 0,
    "awayTimeouts": 0
  }
}
```
