# Statistics By Team

## https://site.web.api.espn.com/apis/common/v3/sports/hockey/{league}/statistics/byteam

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.web.api.espn.com/apis/common/v3/sports/hockey/nhl/statistics/byteam`

## Example Response

```json
{
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
  "teams": [
    {
      "team": {
        "id": "1",
        "uid": "s:70~l:90~t:1",
        "guid": "585503ab-1462-2afe-7a70-aafea4f6b1b6",
        "name": "Bruins",
        "nickname": "Bruins",
        "abbreviation": "BOS",
        "displayName": "Boston Bruins",
        "shortDisplayName": "Bruins",
        "logos": [
          {
            "href": "https://a.espncdn.com/i/teamlogos/nhl/500/bos.png",
            "width": 500,
            "height": 500,
            "rel": []
          },
          {
            "href": "https://a.espncdn.com/i/teamlogos/nhl/500-dark/bos.png",
            "width": 500,
            "height": 500,
            "rel": []
          }
        ],
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/bos/boston-bruins",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=1",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "slug": "boston-bruins",
        "group": {}
      },
      "categories": [
        {
          "name": "general",
          "displayName": "Own General",
          "splitId": "0",
          "totals": [
            "6",
            "2"
          ],
          "values": [
            6.0,
            2.0
          ],
          "ranks": [
            "6",
            "10"
          ]
        },
        {
          "name": "offensive",
          "displayName": "Own Offensive",
          "splitId": "0",
          "totals": [
            "2.00",
            "12"
          ],
          "values": [
            2.0,
            12.0
          ],
          "ranks": [
            "-",
            "13"
          ]
        }
      ]
    },
    {
      "team": {
        "id": "2",
        "uid": "s:70~l:90~t:2",
        "guid": "06b99a18-3096-3b1f-d339-adab9ac51614",
        "name": "Sabres",
        "nickname": "Sabres",
        "abbreviation": "BUF",
        "displayName": "Buffalo Sabres",
        "shortDisplayName": "Sabres",
        "logos": [
          {
            "href": "https://a.espncdn.com/i/teamlogos/nhl/500/buf.png",
            "width": 500,
            "height": 500,
            "rel": []
          },
          {
            "href": "https://a.espncdn.com/i/teamlogos/nhl/500-dark/buf.png",
            "width": 500,
            "height": 500,
            "rel": []
          }
        ],
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/nhl/team/_/name/buf/buffalo-sabres",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://m.espn.com/nhl/clubhouse?teamId=2",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "slug": "buffalo-sabres",
        "group": {}
      },
      "categories": [
        {
          "name": "general",
          "displayName": "Own General",
          "splitId": "0",
          "totals": [
            "8",
            "5"
          ],
          "values": [
            8.0,
            5.0
          ],
          "ranks": [
            "8",
            "4"
          ]
        },
        {
          "name": "offensive",
          "displayName": "Own Offensive",
          "splitId": "0",
          "totals": [
            "3.13",
            "25"
          ],
          "values": [
            3.125,
            25.0
          ],
          "ranks": [
            "-",
            "5"
          ]
        }
      ]
    }
  ],
  "glossary": [
    {
      "abbreviation": "A",
      "displayName": "Assists"
    },
    {
      "abbreviation": "G",
      "displayName": "Goals"
    }
  ],
  "categories": [
    {
      "name": "general",
      "labels": [
        "GP",
        "WINS"
      ],
      "names": [
        "games",
        "wins"
      ],
      "displayNames": [
        "Games Played",
        "Wins"
      ],
      "descriptions": [
        "Total games played.",
        "Wins"
      ]
    },
    {
      "name": "offensive",
      "labels": [
        "GF/G",
        "G"
      ],
      "names": [
        "avgGoals",
        "goals"
      ],
      "displayNames": [
        "Goals For per Game",
        "Goals"
      ],
      "descriptions": [
        "Average goals scored per game.",
        "Total goals scored."
      ]
    }
  ]
}
```
