# Current Season

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/season

Notes:
- Verified with `league=nfl` on 2026-05-08.
- On 2026-05-08 this endpoint still resolved to the 2025 NFL season, with `type=4` (`Off Season`).

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
|  |  |  |

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025?lang=en&region=us",
  "year": 2025,
  "displayName": "2025",
  "startDate": "2025-07-31T07:00Z",
  "endDate": "2026-02-12T07:59Z",
  "type": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/types/4?lang=en&region=us",
    "id": "4",
    "type": 4,
    "name": "Off Season",
    "abbreviation": "off",
    "year": 2025,
    "startDate": "2026-02-12T08:00Z",
    "endDate": "2026-08-06T06:59Z",
    "hasGroups": false,
    "hasStandings": false,
    "hasLegs": false,
    "groups": {},
    "week": {
      "number": 1,
      "startDate": "2026-02-12T08:00Z",
      "endDate": "2026-08-01T06:59Z",
      "text": "Week 1",
      "rankings": {},
      "events": {}
    },
    "weeks": {},
    "slug": "off-season"
  },
  "types": {
    "count": 4,
    "pageIndex": 1,
    "pageSize": 25,
    "pageCount": 1,
    "items": [
      {},
      {},
      {},
      {}
    ]
  },
  "athletes": {
    "count": 20173,
    "pageIndex": 1,
    "pageSize": 25,
    "pageCount": 807,
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
  "rankings": {
    "count": 0,
    "pageIndex": 0,
    "pageSize": 25,
    "pageCount": 0,
    "items": []
  },
  "coaches": {
    "count": 32,
    "pageIndex": 1,
    "pageSize": 25,
    "pageCount": 2,
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
  "leaders": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/types/3/leaders?lang=en&region=us",
    "id": "0",
    "name": "TOTAL",
    "abbreviation": "Any",
    "type": "total",
    "categories": [
      {
        "name": "passingYards",
        "displayName": "Passing Yards",
        "shortDisplayName": "PYDS",
        "abbreviation": "YDS",
        "leaders": [
          {
            "displayValue": "936",
            "value": 936.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "828",
            "value": 828.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "672",
            "value": 672.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "618",
            "value": 618.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "556",
            "value": 556.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "462",
            "value": 462.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "402",
            "value": 402.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "323",
            "value": 323.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "279",
            "value": 279.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "264",
            "value": 264.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "207",
            "value": 207.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "168",
            "value": 168.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "159",
            "value": 159.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "146",
            "value": 146.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "133",
            "value": 133.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "29",
            "value": 29.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "4",
            "value": 4.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          }
        ]
      },
      {
        "name": "rushingYards",
        "displayName": "Rushing Yards",
        "shortDisplayName": "RYDS",
        "abbreviation": "YDS",
        "leaders": [
          {
            "displayValue": "313",
            "value": 313.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "217",
            "value": 217.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "183",
            "value": 183.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "178",
            "value": 178.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "163",
            "value": 163.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "130",
            "value": 130.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "129",
            "value": 129.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "119",
            "value": 119.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "106",
            "value": 106.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "99",
            "value": 99.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "83",
            "value": 83.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "76",
            "value": 76.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "67",
            "value": 67.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "63",
            "value": 63.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "62",
            "value": 62.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "61",
            "value": 61.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "60",
            "value": 60.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "57",
            "value": 57.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "57",
            "value": 57.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "55",
            "value": 55.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "51",
            "value": 51.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "46",
            "value": 46.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "43",
            "value": 43.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "32",
            "value": 32.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "31",
            "value": 31.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          }
        ]
      },
      {
        "name": "receivingYards",
        "displayName": "Receiving Yards",
        "shortDisplayName": "RECYDS",
        "abbreviation": "YDS",
        "leaders": [
          {
            "displayValue": "332",
            "value": 332.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "199",
            "value": 199.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "193",
            "value": 193.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "185",
            "value": 185.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "168",
            "value": 168.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "164",
            "value": 164.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "157",
            "value": 157.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "157",
            "value": 157.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "155",
            "value": 155.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "152",
            "value": 152.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "134",
            "value": 134.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "129",
            "value": 129.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "126",
            "value": 126.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "124",
            "value": 124.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "116",
            "value": 116.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "112",
            "value": 112.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "112",
            "value": 112.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "111",
            "value": 111.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "110",
            "value": 110.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "107",
            "value": 107.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "105",
            "value": 105.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "104",
            "value": 104.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "98",
            "value": 98.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "88",
            "value": 88.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "84",
            "value": 84.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          }
        ]
      },
      {
        "name": "totalTackles",
        "displayName": "Total Tackles",
        "shortDisplayName": "TACK",
        "abbreviation": "TOT",
        "leaders": [
          {
            "displayValue": "30",
            "value": 30.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "27",
            "value": 27.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "26",
            "value": 26.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "26",
            "value": 26.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "25",
            "value": 25.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "23",
            "value": 23.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "21",
            "value": 21.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "21",
            "value": 21.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "20",
            "value": 20.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "20",
            "value": 20.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "19",
            "value": 19.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "19",
            "value": 19.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "19",
            "value": 19.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "19",
            "value": 19.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "18",
            "value": 18.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "18",
            "value": 18.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "17",
            "value": 17.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "17",
            "value": 17.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "16",
            "value": 16.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "15",
            "value": 15.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "15",
            "value": 15.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "15",
            "value": 15.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "14",
            "value": 14.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "14",
            "value": 14.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "13",
            "value": 13.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          }
        ]
      },
      {
        "name": "sacks",
        "displayName": "Sacks",
        "shortDisplayName": "SACK",
        "abbreviation": "SACK",
        "leaders": [
          {
            "displayValue": "4",
            "value": 3.5,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "3",
            "value": 3.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "3",
            "value": 3.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "3",
            "value": 3.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "3",
            "value": 3.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 1.5,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 1.5,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 1.5,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          }
        ]
      },
      {
        "name": "kickoffYards",
        "displayName": "Kickoff Yards",
        "shortDisplayName": "KYDS",
        "abbreviation": "KYDS",
        "leaders": [
          {
            "displayValue": "1313",
            "value": 1313.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "986",
            "value": 986.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "950",
            "value": 950.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "856",
            "value": 856.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "692",
            "value": 692.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "665",
            "value": 665.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "503",
            "value": 503.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "498",
            "value": 498.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "331",
            "value": 331.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "319",
            "value": 319.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "311",
            "value": 311.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "291",
            "value": 291.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "183",
            "value": 183.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "120",
            "value": 120.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          }
        ]
      },
      {
        "name": "interceptions",
        "displayName": "Interceptions",
        "shortDisplayName": "INT",
        "abbreviation": "INT",
        "leaders": [
          {
            "displayValue": "3",
            "value": 3.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          }
        ]
      },
      {
        "name": "passingTouchdowns",
        "displayName": "Passing Touchdowns",
        "shortDisplayName": "TD",
        "abbreviation": "TD",
        "leaders": [
          {
            "displayValue": "6",
            "value": 6.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "6",
            "value": 6.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "5",
            "value": 5.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "4",
            "value": 4.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "4",
            "value": 4.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "4",
            "value": 4.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "3",
            "value": 3.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "3",
            "value": 3.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          }
        ]
      },
      {
        "name": "quarterbackRating",
        "displayName": "Quarterback Rating",
        "shortDisplayName": "RAT",
        "abbreviation": "RAT",
        "leaders": [
          {
            "displayValue": "104",
            "value": 103.80400085449219,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "102",
            "value": 102.44999694824219,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "100",
            "value": 99.83100128173828,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "94",
            "value": 94.37999725341797,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "87",
            "value": 87.13800048828125,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "86",
            "value": 86.38899993896484,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "82",
            "value": 82.22200012207031,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "79",
            "value": 79.22599792480469,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "75",
            "value": 74.52999877929688,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "71",
            "value": 71.25,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "68",
            "value": 68.31900024414062,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "66",
            "value": 65.87999725341797,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "63",
            "value": 62.970001220703125,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "52",
            "value": 51.766998291015625,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "51",
            "value": 50.82099914550781,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "158",
            "value": 158.33299255371094,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "79",
            "value": 79.16699981689453,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "56",
            "value": 56.25,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "40",
            "value": 39.58300018310547,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "40",
            "value": 39.58300018310547,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "40",
            "value": 39.58300018310547,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "40",
            "value": 39.58300018310547,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "40",
            "value": 39.58300018310547,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "40",
            "value": 39.58300018310547,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "40",
            "value": 39.58300018310547,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          }
        ]
      },
      {
        "name": "rushingTouchdowns",
        "displayName": "Rushing Touchdowns",
        "shortDisplayName": "TD",
        "abbreviation": "TD",
        "leaders": [
          {
            "displayValue": "4",
            "value": 4.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          }
        ]
      },
      {
        "name": "receptions",
        "displayName": "Receptions",
        "shortDisplayName": "REC",
        "abbreviation": "REC",
        "leaders": [
          {
            "displayValue": "24",
            "value": 24.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "19",
            "value": 19.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "17",
            "value": 17.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "15",
            "value": 15.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "14",
            "value": 14.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "12",
            "value": 12.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "12",
            "value": 12.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "12",
            "value": 12.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "11",
            "value": 11.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "11",
            "value": 11.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "11",
            "value": 11.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "10",
            "value": 10.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "9",
            "value": 9.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "9",
            "value": 9.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "9",
            "value": 9.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "9",
            "value": 9.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "9",
            "value": 9.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "9",
            "value": 9.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "9",
            "value": 9.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "8",
            "value": 8.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "8",
            "value": 8.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "8",
            "value": 8.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "8",
            "value": 8.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "8",
            "value": 8.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "7",
            "value": 7.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          }
        ]
      },
      {
        "name": "receivingTouchdowns",
        "displayName": "Receiving Touchdowns",
        "shortDisplayName": "TD",
        "abbreviation": "TD",
        "leaders": [
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          }
        ]
      },
      {
        "name": "totalPoints",
        "displayName": "Total Points",
        "shortDisplayName": "TP",
        "abbreviation": "TP",
        "leaders": [
          {
            "displayValue": "35",
            "value": 35.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "27",
            "value": 27.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "24",
            "value": 24.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "24",
            "value": 24.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "21",
            "value": 21.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "19",
            "value": 19.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "18",
            "value": 18.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "16",
            "value": 16.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "16",
            "value": 16.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "16",
            "value": 16.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "12",
            "value": 12.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "12",
            "value": 12.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "12",
            "value": 12.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "12",
            "value": 12.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "12",
            "value": 12.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "12",
            "value": 12.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "12",
            "value": 12.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "12",
            "value": 12.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "11",
            "value": 11.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "7",
            "value": 7.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "7",
            "value": 7.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "6",
            "value": 6.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "6",
            "value": 6.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "6",
            "value": 6.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "6",
            "value": 6.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          }
        ]
      },
      {
        "name": "totalTouchdowns",
        "displayName": "Total Touchdowns",
        "shortDisplayName": "TTD",
        "abbreviation": "TTD",
        "leaders": [
          {
            "displayValue": "4",
            "value": 4.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "4",
            "value": 4.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "3",
            "value": 3.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "1",
            "value": 1.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          }
        ]
      },
      {
        "name": "puntYards",
        "displayName": "Punt Yards",
        "shortDisplayName": "PYDS",
        "abbreviation": "PYDS",
        "leaders": [
          {
            "displayValue": "1122",
            "value": 1122.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "710",
            "value": 710.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "646",
            "value": 646.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "497",
            "value": 497.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "419",
            "value": 419.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "273",
            "value": 273.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "211",
            "value": 211.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "208",
            "value": 208.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "195",
            "value": 195.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "163",
            "value": 163.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "154",
            "value": 154.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "141",
            "value": 141.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "130",
            "value": 130.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "100",
            "value": 100.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          }
        ]
      },
      {
        "name": "passesDefended",
        "displayName": "Passes Defended",
        "shortDisplayName": "PD",
        "abbreviation": "PD",
        "leaders": [
          {
            "displayValue": "7",
            "value": 7.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "7",
            "value": 7.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "6",
            "value": 6.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "4",
            "value": 4.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "4",
            "value": 4.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "4",
            "value": 4.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "4",
            "value": 4.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "4",
            "value": 4.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "4",
            "value": 4.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "4",
            "value": 4.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "3",
            "value": 3.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "3",
            "value": 3.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "3",
            "value": 3.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          },
          {
            "displayValue": "2",
            "value": 2.0,
            "rel": [
              "athlete"
            ],
            "athlete": {},
            "team": {},
            "statistics": {}
          }
        ]
      }
    ]
  }
}
```
