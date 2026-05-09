# MLB Statistics By Team

## https://site.web.api.espn.com/apis/common/v3/sports/baseball/mlb/statistics/byteam?season=2026&seasontype=2

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "categories": [
    {
      "name": "batting",
      "labels": [
        "GP",
        "AB"
      ],
      "names": [
        "gamesPlayed",
        "atBats"
      ],
      "displayNames": [
        "Games Played",
        "At Bats"
      ],
      "descriptions": [
        "Games Played",
        "The number of times a batter comes up to bat not counting walks and sacrifices."
      ]
    },
    {
      "name": "pitching",
      "labels": [
        "GP",
        "W"
      ],
      "names": [
        "gamesPlayed",
        "wins"
      ],
      "displayNames": [
        "Games Played",
        "Wins"
      ],
      "descriptions": [
        "Games Played",
        "The number of times the pitcher was attributed with a win"
      ]
    }
  ],
  "currentSeason": {
    "displayName": "2026",
    "year": 2026,
    "startDate": "2026-02-19T08:00:00.000+00:00",
    "endDate": "2026-11-12T07:59:00.000+00:00",
    "type": {
      "id": "2",
      "name": "Regular Season",
      "type": 2,
      "startDate": "2026-03-25T07:00:00.000+00:00",
      "endDate": "2026-09-30T06:59:00.000+00:00",
      "week": {
        "number": 7,
        "startDate": "2026-05-06T07:00:00.000+00:00",
        "endDate": "2026-05-13T06:59:00.000+00:00",
        "text": "Week 7"
      }
    }
  },
  "requestedSeason": {
    "displayName": "2026",
    "year": 2026,
    "startDate": "2026-02-19T08:00:00.000+00:00",
    "endDate": "2026-11-12T07:59:00.000+00:00",
    "type": {
      "id": "2",
      "name": "Regular Season",
      "type": 2,
      "startDate": "2026-03-25T07:00:00.000+00:00",
      "endDate": "2026-09-30T06:59:00.000+00:00",
      "week": {
        "number": 7,
        "startDate": "2026-05-06T07:00:00.000+00:00",
        "endDate": "2026-05-13T06:59:00.000+00:00",
        "text": "Week 7"
      }
    }
  },
  "teams": [
    {
      "team": {
        "id": "1",
        "uid": "s:1~l:10~t:1",
        "guid": "9ca473b8-e73e-a33d-8ea0-b4d160be4be7",
        "name": "Orioles",
        "displayName": "Baltimore Orioles",
        "abbreviation": "BAL",
        "slug": "baltimore-orioles",
        "shortDisplayName": "Orioles",
        "logos": [
          {
            "href": "https://a.espncdn.com/i/teamlogos/mlb/500/bal.png",
            "width": 500,
            "height": 500,
            "rel": []
          },
          {
            "href": "https://a.espncdn.com/i/teamlogos/mlb/500-dark/bal.png",
            "width": 500,
            "height": 500,
            "rel": []
          }
        ],
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/_/name/bal/baltimore-orioles",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/roster/_/name/bal/baltimore-orioles",
            "text": "Roster",
            "shortText": "Roster",
            "isExternal": false,
            "isPremium": false
          }
        ]
      },
      "categories": [
        {
          "name": "batting",
          "displayName": "Own Batting",
          "splitId": "0",
          "totals": [
            "38",
            "1262"
          ],
          "values": [
            38.0,
            1262.0
          ],
          "ranks": [
            "2",
            "14"
          ]
        },
        {
          "name": "pitching",
          "displayName": "Own Pitching",
          "splitId": "0",
          "totals": [
            "38",
            "17"
          ],
          "values": [
            38.0,
            17.0
          ],
          "ranks": [
            "2",
            "16"
          ]
        }
      ]
    },
    {
      "team": {
        "id": "2",
        "uid": "s:1~l:10~t:2",
        "guid": "c6df06f6-785d-3900-4935-5fd13742e2ee",
        "name": "Red Sox",
        "displayName": "Boston Red Sox",
        "abbreviation": "BOS",
        "slug": "boston-red-sox",
        "shortDisplayName": "Red Sox",
        "logos": [
          {
            "href": "https://a.espncdn.com/i/teamlogos/mlb/500/bos.png",
            "width": 500,
            "height": 500,
            "rel": []
          },
          {
            "href": "https://a.espncdn.com/i/teamlogos/mlb/500-dark/bos.png",
            "width": 500,
            "height": 500,
            "rel": []
          }
        ],
        "links": [
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/_/name/bos/boston-red-sox",
            "text": "Clubhouse",
            "shortText": "Clubhouse",
            "isExternal": false,
            "isPremium": false
          },
          {
            "language": "en",
            "rel": [],
            "href": "https://www.espn.com/mlb/team/roster/_/name/bos/boston-red-sox",
            "text": "Roster",
            "shortText": "Roster",
            "isExternal": false,
            "isPremium": false
          }
        ]
      },
      "categories": [
        {
          "name": "batting",
          "displayName": "Own Batting",
          "splitId": "0",
          "totals": [
            "38",
            "1276"
          ],
          "values": [
            38.0,
            1276.0
          ],
          "ranks": [
            "2",
            "9"
          ]
        },
        {
          "name": "pitching",
          "displayName": "Own Pitching",
          "splitId": "0",
          "totals": [
            "38",
            "16"
          ],
          "values": [
            38.0,
            16.0
          ],
          "ranks": [
            "2",
            "23"
          ]
        }
      ]
    }
  ],
  "glossary": [
    {
      "displayName": "Doubles",
      "abbreviation": "2B"
    },
    {
      "displayName": "Triples",
      "abbreviation": "3B"
    }
  ]
}
```
