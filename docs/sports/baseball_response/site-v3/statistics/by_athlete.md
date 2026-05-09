# MLB Statistics By Athlete

## https://site.web.api.espn.com/apis/common/v3/sports/baseball/mlb/statistics/byathlete?category=batting&sort=batting.homeRuns:desc&season=2026&seasontype=2

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "athletes": [
    {
      "categories": [
        {
          "name": "batting",
          "displayName": "Own Batting",
          "totals": [
            "39",
            "141"
          ],
          "values": [
            39.0,
            141.0
          ],
          "ranks": [
            "2",
            "49"
          ]
        },
        {
          "name": "pitching",
          "displayName": "Own Pitching",
          "totals": [
            "-",
            "-"
          ],
          "values": [
            null,
            null
          ],
          "ranks": [
            "-",
            "-"
          ]
        }
      ],
      "athlete": {
        "id": "33192",
        "uid": "s:1~l:10~a:33192",
        "guid": "e3e39e69-2861-f5b5-49be-b0880534c802",
        "displayName": "Aaron Judge",
        "slug": "aaron-judge",
        "status": {
          "id": "1",
          "name": "Active",
          "abbreviation": "Active",
          "type": "active"
        },
        "type": "baseball",
        "firstName": "Aaron",
        "lastName": "Judge",
        "shortName": "A. Judge"
      }
    },
    {
      "categories": [
        {
          "name": "batting",
          "displayName": "Own Batting",
          "splitId": "0",
          "totals": [
            "38",
            "135"
          ],
          "values": [
            38.0,
            135.0
          ],
          "ranks": [
            "29",
            "89"
          ]
        },
        {
          "name": "pitching",
          "displayName": "Own Pitching",
          "splitId": "0",
          "totals": [
            "-",
            "-"
          ],
          "values": [
            null,
            null
          ],
          "ranks": [
            "-",
            "-"
          ]
        }
      ],
      "athlete": {
        "id": "4872595",
        "uid": "s:1~l:10~a:4872595",
        "guid": "37da30e0-9b6a-31a8-b783-083261547d7a",
        "displayName": "Munetaka Murakami",
        "slug": "munetaka-murakami",
        "status": {
          "id": "1",
          "name": "Active",
          "abbreviation": "Active",
          "type": "active"
        },
        "type": "baseball",
        "firstName": "Munetaka",
        "lastName": "Murakami",
        "shortName": "M. Murakami"
      }
    }
  ],
  "categories": [
    {
      "name": "batting",
      "displayName": "Own Batting",
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
      "displayName": "Own Pitching",
      "labels": [
        "GP",
        "GS"
      ],
      "names": [
        "gamesPlayed",
        "gamesStarted"
      ],
      "displayNames": [
        "Games Played",
        "Games Started"
      ],
      "descriptions": [
        "Games Played",
        "The number of games started by a pitcher."
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
  "pagination": {
    "count": 177,
    "limit": 50,
    "page": 1,
    "pages": 4,
    "first": "http://site.api.espn.com:80/apis/common/v3/sports/baseball/mlb/statistics/byathlete?category=batting&season...",
    "next": "http://site.api.espn.com:80/apis/common/v3/sports/baseball/mlb/statistics/byathlete?category=batting&season...",
    "last": "http://site.api.espn.com:80/apis/common/v3/sports/baseball/mlb/statistics/byathlete?category=batting&season..."
  },
  "league": {
    "id": "10",
    "uid": "s:1~l:10",
    "name": "Major League Baseball",
    "abbreviation": "MLB",
    "slug": "mlb",
    "midsizeName": "MLB",
    "shortName": "MLB"
  },
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
