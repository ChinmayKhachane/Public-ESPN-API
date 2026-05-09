# Competition Play ID

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/plays/{play}

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988`, `play=40177298840` on 2026-05-08.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/plays/40177298840?lang=en&region=us",
  "id": "40177298840",
  "sequenceNumber": "4000",
  "type": {
    "id": "53",
    "text": "Kickoff",
    "abbreviation": "K"
  },
  "text": "A.Borregales kicks 65 yards from NE 35 to end zone, Touchback to the SEA 35.",
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
  "probability": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/probabilities/40177298840?lang=en&region=us",
    "competition": {},
    "play": {},
    "homeTeam": {},
    "awayTeam": {},
    "tiePercentage": 0.0,
    "homeWinPercentage": 0.3766,
    "awayWinPercentage": 0.6234,
    "lastModified": "2026-02-09T02:41Z",
    "sequenceNumber": "4000",
    "source": {
      "id": "2",
      "description": "feed",
      "state": "full"
    },
    "secondsLeft": 0,
    "spreadCoverProbHome": 0.5071,
    "spreadPushProb": 0.0,
    "totalOverProb": 0.5049,
    "totalPushProb": 0.0
  },
  "drive": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/drives/4017729881?lang=en&region=us",
    "id": "4017729881",
    "description": "8 plays, 51 yards, 3:02",
    "sequenceNumber": "1",
    "team": {},
    "endTeam": {},
    "start": {
      "period": {
        "type": "quarter",
        "number": 1
      },
      "clock": {
        "value": 900.0,
        "displayValue": "15:00"
      },
      "yardLine": 65,
      "text": "SEA 35"
    },
    "end": {
      "period": {
        "type": "quarter",
        "number": 1
      },
      "clock": {
        "value": 718.0,
        "displayValue": "11:58"
      },
      "yardLine": 14,
      "text": "NE 14"
    },
    "timeElapsed": {
      "value": 182.0,
      "displayValue": "3:02"
    },
    "yards": 51,
    "isScore": true,
    "offensivePlays": 8,
    "result": "FG",
    "shortDisplayResult": "FG",
    "displayResult": "Field Goal",
    "source": {
      "id": "2",
      "description": "Feed"
    },
    "plays": {
      "count": 10,
      "pageIndex": 1,
      "pageSize": 1000,
      "pageCount": 1,
      "items": [
        {
          "id": "40177298840",
          "sequenceNumber": "4000",
          "type": {
            "id": "53",
            "text": "Kickoff",
            "abbreviation": "K"
          },
          "text": "A.Borregales kicks 65 yards from NE 35 to end zone, Touchback to the SEA 35.",
          "shortText": "Andy Borregales 65 Yd Kickoff, Touchback",
          "alternativeText": "A.Borregales kicks 65 yards from NE 35 to end zone, Touchback to the SEA 35.",
          "shortAlternativeText": "Andy Borregales 65 Yd Kickoff, Touchback",
          "awayScore": 0,
          "homeScore": 0,
          "period": {
            "number": 1
          },
          "clock": {
            "value": 900.0,
            "displayValue": "15:00"
          },
          "scoringPlay": false,
          "priority": false,
          "scoreValue": 0,
          "modified": "2026-02-09T03:39Z",
          "team": {},
          "participants": [
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 1,
              "type": "kicker"
            }
          ],
          "probability": {},
          "wallclock": "2026-02-08T23:40:23Z",
          "teamParticipants": [
            {
              "team": {},
              "id": "26",
              "order": 1,
              "type": "offense"
            },
            {
              "team": {},
              "id": "17",
              "statistics": {},
              "playStatistics": {},
              "order": 2,
              "type": "defense"
            }
          ],
          "isPenalty": false,
          "statYardage": 0,
          "drive": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/drives/4017729881?lang=en&region=us"
          },
          "start": {
            "down": 0,
            "distance": 0,
            "yardLine": 35,
            "yardsToEndzone": 65,
            "team": {}
          },
          "end": {
            "down": 1,
            "distance": 10,
            "yardLine": 65,
            "yardsToEndzone": 65,
            "downDistanceText": "1st & 10 at SEA 35",
            "shortDownDistanceText": "1st & 10",
            "possessionText": "SEA 35",
            "team": {}
          },
          "isTurnover": false
        },
        {
          "id": "40177298857",
          "sequenceNumber": "5700",
          "type": {
            "id": "5",
            "text": "Rush",
            "abbreviation": "RUSH"
          },
          "text": "K.Walker left end pushed ob at SEA 45 for 10 yards (C.Davis).",
          "shortText": "Kenneth Walker III 10 Yd Rush",
          "alternativeText": "K.Walker left end pushed ob at SEA 45 for 10 yards (C.Davis).",
          "shortAlternativeText": "Kenneth Walker III 10 Yd Rush",
          "awayScore": 0,
          "homeScore": 0,
          "period": {
            "number": 1
          },
          "clock": {
            "value": 900.0,
            "displayValue": "15:00"
          },
          "scoringPlay": false,
          "priority": false,
          "scoreValue": 0,
          "modified": "2026-02-09T03:39Z",
          "team": {},
          "participants": [
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 1,
              "type": "rusher"
            },
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 2,
              "type": "tackler"
            }
          ],
          "probability": {},
          "wallclock": "2026-02-08T23:41:07Z",
          "teamParticipants": [
            {
              "team": {},
              "id": "26",
              "statistics": {},
              "playStatistics": {},
              "order": 1,
              "type": "offense"
            },
            {
              "team": {},
              "id": "17",
              "statistics": {},
              "playStatistics": {},
              "order": 2,
              "type": "defense"
            }
          ],
          "isPenalty": false,
          "statYardage": 10,
          "drive": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/drives/4017729881?lang=en&region=us"
          },
          "start": {
            "down": 1,
            "distance": 10,
            "yardLine": 65,
            "yardsToEndzone": 65,
            "downDistanceText": "1st & 10 at SEA 35",
            "shortDownDistanceText": "1st & 10",
            "possessionText": "SEA 35",
            "team": {}
          },
          "end": {
            "down": 1,
            "distance": 10,
            "yardLine": 55,
            "yardsToEndzone": 55,
            "downDistanceText": "1st & 10 at SEA 45",
            "shortDownDistanceText": "1st & 10",
            "possessionText": "SEA 45",
            "team": {}
          },
          "isTurnover": false
        },
        {
          "id": "40177298884",
          "sequenceNumber": "8400",
          "type": {
            "id": "5",
            "text": "Rush",
            "abbreviation": "RUSH"
          },
          "text": "K.Walker right tackle to SEA 45 for no gain (C.Barmore).",
          "shortText": "Kenneth Walker III Rush, No Gain",
          "alternativeText": "K.Walker right tackle to SEA 45 for no gain (C.Barmore).",
          "shortAlternativeText": "Kenneth Walker III Rush, No Gain",
          "awayScore": 0,
          "homeScore": 0,
          "period": {
            "number": 1
          },
          "clock": {
            "value": 869.0,
            "displayValue": "14:29"
          },
          "scoringPlay": false,
          "priority": false,
          "scoreValue": 0,
          "modified": "2026-02-09T03:39Z",
          "team": {},
          "participants": [
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 1,
              "type": "rusher"
            },
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 2,
              "type": "tackler"
            },
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 3,
              "type": "other"
            }
          ],
          "probability": {},
          "wallclock": "2026-02-08T23:41:42Z",
          "teamParticipants": [
            {
              "team": {},
              "id": "26",
              "statistics": {},
              "playStatistics": {},
              "order": 1,
              "type": "offense"
            },
            {
              "team": {},
              "id": "17",
              "statistics": {},
              "playStatistics": {},
              "order": 2,
              "type": "defense"
            }
          ],
          "isPenalty": false,
          "statYardage": 0,
          "drive": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/drives/4017729881?lang=en&region=us"
          },
          "start": {
            "down": 1,
            "distance": 10,
            "yardLine": 55,
            "yardsToEndzone": 55,
            "downDistanceText": "1st & 10 at SEA 45",
            "shortDownDistanceText": "1st & 10",
            "possessionText": "SEA 45",
            "team": {}
          },
          "end": {
            "down": 2,
            "distance": 10,
            "yardLine": 55,
            "yardsToEndzone": 55,
            "downDistanceText": "2nd & 10 at SEA 45",
            "shortDownDistanceText": "2nd & 10",
            "possessionText": "SEA 45",
            "team": {}
          },
          "isTurnover": false
        },
        {
          "id": "401772988106",
          "sequenceNumber": "10600",
          "type": {
            "id": "24",
            "text": "Pass Reception",
            "abbreviation": "REC"
          },
          "text": "(Shotgun) S.Darnold pass short left to A.Barner to NE 40 for 15 yards (C.Woodson).",
          "shortText": "Sam Darnold Pass Complete for 15 Yds to AJ Barner",
          "alternativeText": "(Shotgun) S.Darnold pass short left to A.Barner to NE 40 for 15 yards (C.Woodson).",
          "shortAlternativeText": "Sam Darnold Pass Complete for 15 Yds to AJ Barner",
          "awayScore": 0,
          "homeScore": 0,
          "period": {
            "number": 1
          },
          "clock": {
            "value": 830.0,
            "displayValue": "13:50"
          },
          "scoringPlay": false,
          "priority": false,
          "scoreValue": 0,
          "modified": "2026-02-09T03:39Z",
          "team": {},
          "participants": [
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 1,
              "type": "passer"
            },
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 2,
              "type": "receiver"
            },
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 3,
              "type": "tackler"
            }
          ],
          "probability": {},
          "wallclock": "2026-02-08T23:42:21Z",
          "teamParticipants": [
            {
              "team": {},
              "id": "26",
              "statistics": {},
              "playStatistics": {},
              "order": 1,
              "type": "offense"
            },
            {
              "team": {},
              "id": "17",
              "statistics": {},
              "playStatistics": {},
              "order": 2,
              "type": "defense"
            }
          ],
          "isPenalty": false,
          "statYardage": 15,
          "drive": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/drives/4017729881?lang=en&region=us"
          },
          "start": {
            "down": 2,
            "distance": 10,
            "yardLine": 55,
            "yardsToEndzone": 55,
            "downDistanceText": "2nd & 10 at SEA 45",
            "shortDownDistanceText": "2nd & 10",
            "possessionText": "SEA 45",
            "team": {}
          },
          "end": {
            "down": 1,
            "distance": 10,
            "yardLine": 40,
            "yardsToEndzone": 40,
            "downDistanceText": "1st & 10 at NE 40",
            "shortDownDistanceText": "1st & 10",
            "possessionText": "NE 40",
            "team": {}
          },
          "yardsAfterCatch": -39,
          "isTurnover": false
        },
        {
          "id": "401772988131",
          "sequenceNumber": "13100",
          "type": {
            "id": "24",
            "text": "Pass Reception",
            "abbreviation": "REC"
          },
          "text": "(No Huddle) S.Darnold pass deep left to C.Kupp pushed ob at NE 17 for 23 yards (M.Jones).",
          "shortText": "Sam Darnold Pass Complete for 23 Yds to Cooper Kupp",
          "alternativeText": "(No Huddle) S.Darnold pass deep left to C.Kupp pushed ob at NE 17 for 23 yards (M.Jones).",
          "shortAlternativeText": "Sam Darnold Pass Complete for 23 Yds to Cooper Kupp",
          "awayScore": 0,
          "homeScore": 0,
          "period": {
            "number": 1
          },
          "clock": {
            "value": 805.0,
            "displayValue": "13:25"
          },
          "scoringPlay": false,
          "priority": false,
          "scoreValue": 0,
          "modified": "2026-02-09T03:39Z",
          "team": {},
          "participants": [
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 1,
              "type": "passer"
            },
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 2,
              "type": "receiver"
            },
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 3,
              "type": "tackler"
            }
          ],
          "probability": {},
          "wallclock": "2026-02-08T23:42:46Z",
          "teamParticipants": [
            {
              "team": {},
              "id": "26",
              "statistics": {},
              "playStatistics": {},
              "order": 1,
              "type": "offense"
            },
            {
              "team": {},
              "id": "17",
              "statistics": {},
              "playStatistics": {},
              "order": 2,
              "type": "defense"
            }
          ],
          "isPenalty": false,
          "statYardage": 23,
          "drive": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/drives/4017729881?lang=en&region=us"
          },
          "start": {
            "down": 1,
            "distance": 10,
            "yardLine": 40,
            "yardsToEndzone": 40,
            "downDistanceText": "1st & 10 at NE 40",
            "shortDownDistanceText": "1st & 10",
            "possessionText": "NE 40",
            "team": {}
          },
          "end": {
            "down": 1,
            "distance": 10,
            "yardLine": 17,
            "yardsToEndzone": 17,
            "downDistanceText": "1st & 10 at NE 17",
            "shortDownDistanceText": "1st & 10",
            "possessionText": "NE 17",
            "team": {}
          },
          "yardsAfterCatch": 1,
          "isTurnover": false
        },
        {
          "id": "401772988161",
          "sequenceNumber": "16100",
          "type": {
            "id": "3",
            "text": "Pass Incompletion"
          },
          "text": "(No Huddle, Shotgun) S.Darnold pass incomplete deep left.",
          "shortText": "Sam Darnold Incomplete Pass",
          "alternativeText": "(No Huddle, Shotgun) S.Darnold pass incomplete deep left.",
          "shortAlternativeText": "Sam Darnold Incomplete Pass",
          "awayScore": 0,
          "homeScore": 0,
          "period": {
            "number": 1
          },
          "clock": {
            "value": 778.0,
            "displayValue": "12:58"
          },
          "scoringPlay": false,
          "priority": false,
          "scoreValue": 0,
          "modified": "2026-02-09T03:39Z",
          "team": {},
          "participants": [
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 1,
              "type": "passer"
            }
          ],
          "probability": {},
          "wallclock": "2026-02-08T23:43:17Z",
          "teamParticipants": [
            {
              "team": {},
              "id": "26",
              "statistics": {},
              "playStatistics": {},
              "order": 1,
              "type": "offense"
            },
            {
              "team": {},
              "id": "17",
              "order": 2,
              "type": "defense"
            }
          ],
          "isPenalty": false,
          "statYardage": 0,
          "drive": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/drives/4017729881?lang=en&region=us"
          },
          "start": {
            "down": 1,
            "distance": 10,
            "yardLine": 17,
            "yardsToEndzone": 17,
            "downDistanceText": "1st & 10 at NE 17",
            "shortDownDistanceText": "1st & 10",
            "possessionText": "NE 17",
            "team": {}
          },
          "end": {
            "down": 2,
            "distance": 10,
            "yardLine": 17,
            "yardsToEndzone": 17,
            "downDistanceText": "2nd & 10 at NE 17",
            "shortDownDistanceText": "2nd & 10",
            "possessionText": "NE 17",
            "team": {}
          },
          "isTurnover": false
        },
        {
          "id": "401772988184",
          "sequenceNumber": "18400",
          "type": {
            "id": "5",
            "text": "Rush",
            "abbreviation": "RUSH"
          },
          "text": "K.Walker right tackle to NE 14 for 3 yards (A.Jennings).",
          "shortText": "Kenneth Walker III 3 Yd Rush",
          "alternativeText": "K.Walker right tackle to NE 14 for 3 yards (A.Jennings).",
          "shortAlternativeText": "Kenneth Walker III 3 Yd Rush",
          "awayScore": 0,
          "homeScore": 0,
          "period": {
            "number": 1
          },
          "clock": {
            "value": 771.0,
            "displayValue": "12:51"
          },
          "scoringPlay": false,
          "priority": false,
          "scoreValue": 0,
          "modified": "2026-02-09T03:39Z",
          "team": {},
          "participants": [
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 1,
              "type": "rusher"
            },
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 2,
              "type": "tackler"
            },
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 3,
              "type": "other"
            }
          ],
          "probability": {},
          "wallclock": "2026-02-08T23:44:04Z",
          "teamParticipants": [
            {
              "team": {},
              "id": "26",
              "statistics": {},
              "playStatistics": {},
              "order": 1,
              "type": "offense"
            },
            {
              "team": {},
              "id": "17",
              "statistics": {},
              "playStatistics": {},
              "order": 2,
              "type": "defense"
            }
          ],
          "isPenalty": false,
          "statYardage": 3,
          "drive": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/drives/4017729881?lang=en&region=us"
          },
          "start": {
            "down": 2,
            "distance": 10,
            "yardLine": 17,
            "yardsToEndzone": 17,
            "downDistanceText": "2nd & 10 at NE 17",
            "shortDownDistanceText": "2nd & 10",
            "possessionText": "NE 17",
            "team": {}
          },
          "end": {
            "down": 3,
            "distance": 7,
            "yardLine": 14,
            "yardsToEndzone": 14,
            "downDistanceText": "3rd & 7 at NE 14",
            "shortDownDistanceText": "3rd & 7",
            "possessionText": "NE 14",
            "team": {}
          },
          "isTurnover": false
        },
        {
          "id": "401772988206",
          "sequenceNumber": "20600",
          "type": {
            "id": "3",
            "text": "Pass Incompletion"
          },
          "text": "(Shotgun) S.Darnold pass incomplete short left to R.Shaheed.",
          "shortText": "Sam Darnold Incomplete Pass, Intended For Rashid Shaheed",
          "alternativeText": "(Shotgun) S.Darnold pass incomplete short left to R.Shaheed.",
          "shortAlternativeText": "Sam Darnold Incomplete Pass, Intended For Rashid Shaheed",
          "awayScore": 0,
          "homeScore": 0,
          "period": {
            "number": 1
          },
          "clock": {
            "value": 726.0,
            "displayValue": "12:06"
          },
          "scoringPlay": false,
          "priority": false,
          "scoreValue": 0,
          "modified": "2026-02-09T03:39Z",
          "team": {},
          "participants": [
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 1,
              "type": "passer"
            },
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 2,
              "type": "receiver"
            },
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 3,
              "type": "receiver"
            }
          ],
          "probability": {},
          "wallclock": "2026-02-08T23:44:49Z",
          "teamParticipants": [
            {
              "team": {},
              "id": "26",
              "statistics": {},
              "playStatistics": {},
              "order": 1,
              "type": "offense"
            },
            {
              "team": {},
              "id": "17",
              "order": 2,
              "type": "defense"
            }
          ],
          "isPenalty": false,
          "statYardage": 0,
          "drive": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/drives/4017729881?lang=en&region=us"
          },
          "start": {
            "down": 3,
            "distance": 7,
            "yardLine": 14,
            "yardsToEndzone": 14,
            "downDistanceText": "3rd & 7 at NE 14",
            "shortDownDistanceText": "3rd & 7",
            "possessionText": "NE 14",
            "team": {}
          },
          "end": {
            "down": 4,
            "distance": 7,
            "yardLine": 14,
            "yardsToEndzone": 14,
            "downDistanceText": "4th & 7 at NE 14",
            "shortDownDistanceText": "4th & 7",
            "possessionText": "NE 14",
            "team": {}
          },
          "isTurnover": false
        },
        {
          "id": "401772988229",
          "sequenceNumber": "22900",
          "type": {
            "id": "59",
            "text": "Field Goal Good",
            "abbreviation": "FG"
          },
          "text": "J.Myers 33 yard field goal is GOOD, Center-C.Stoll, Holder-M.Dickson.",
          "shortText": "Jason Myers 33 Yd Field Goal",
          "alternativeText": "J.Myers 33 yard field goal is GOOD, Center-C.Stoll, Holder-M.Dickson.",
          "shortAlternativeText": "Jason Myers 33 Yd Field Goal",
          "awayScore": 3,
          "homeScore": 0,
          "period": {
            "number": 1
          },
          "clock": {
            "value": 718.0,
            "displayValue": "11:58"
          },
          "scoringPlay": true,
          "priority": false,
          "scoreValue": 3,
          "modified": "2026-02-09T03:39Z",
          "team": {},
          "participants": [
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 1,
              "type": "kicker"
            },
            {
              "athlete": {},
              "position": {},
              "statistics": {},
              "playStatistics": {},
              "order": 2,
              "type": "scorer"
            },
            {
              "athlete": {},
              "position": {},
              "order": 3,
              "type": "snapper"
            },
            {
              "athlete": {},
              "position": {},
              "order": 4,
              "type": "holder"
            }
          ],
          "probability": {},
          "wallclock": "2026-02-08T23:45:31Z",
          "teamParticipants": [
            {
              "team": {},
              "id": "26",
              "statistics": {},
              "playStatistics": {},
              "order": 1,
              "type": "offense"
            },
            {
              "team": {},
              "id": "17",
              "order": 2,
              "type": "defense"
            }
          ],
          "isPenalty": false,
          "statYardage": 33,
          "drive": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/drives/4017729881?lang=en&region=us"
          },
          "start": {
            "down": 4,
            "distance": 7,
            "yardLine": 14,
            "yardsToEndzone": 14,
            "downDistanceText": "4th & 7 at NE 14",
            "shortDownDistanceText": "4th & 7",
            "possessionText": "NE 14",
            "team": {}
          },
          "end": {
            "down": -1,
            "distance": 10,
            "yardLine": 0,
            "yardsToEndzone": 0,
            "team": {}
          },
          "scoringType": {
            "name": "field-goal",
            "displayName": "Field Goal",
            "abbreviation": "FG"
          },
          "isTurnover": false
        },
        {
          "id": "401772988239",
          "sequenceNumber": "23900",
          "type": {
            "id": "74",
            "text": "Official Timeout",
            "abbreviation": "Off TO"
          },
          "text": "Official Timeout at 11:58.",
          "shortText": "Official Timeout At 11:58. ",
          "alternativeText": "Official Timeout at 11:58.",
          "shortAlternativeText": "Official Timeout At 11:58. ",
          "awayScore": 3,
          "homeScore": 0,
          "period": {
            "number": 1
          },
          "clock": {
            "value": 718.0,
            "displayValue": "11:58"
          },
          "scoringPlay": false,
          "priority": false,
          "scoreValue": 0,
          "modified": "2026-02-09T03:39Z",
          "team": {},
          "probability": {},
          "wallclock": "2026-02-08T23:45:47Z",
          "teamParticipants": [
            {
              "team": {},
              "id": "26",
              "order": 1,
              "type": "offense"
            },
            {
              "team": {},
              "id": "17",
              "order": 2,
              "type": "defense"
            }
          ],
          "isPenalty": false,
          "statYardage": 0,
          "drive": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/drives/4017729881?lang=en&region=us"
          },
          "start": {
            "down": -1,
            "distance": 0,
            "yardLine": 65,
            "yardsToEndzone": 0,
            "downDistanceText": " & Goal at SEA 35",
            "shortDownDistanceText": " & Goal",
            "possessionText": "SEA 35",
            "team": {}
          },
          "end": {
            "down": -1,
            "distance": 0,
            "yardLine": 65,
            "yardsToEndzone": 0,
            "team": {}
          },
          "isTurnover": false
        }
      ]
    }
  },
  "participants": [
    {
      "athlete": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/athletes/4569923?lang=en&region=us",
        "id": "4569923",
        "uid": "s:20~l:28~a:4569923",
        "guid": "fa8ee11d-9490-303f-95ff-bab0c55161d5",
        "type": "football",
        "alternateIds": {
          "sdr": "4569923"
        },
        "firstName": "Andy",
        "lastName": "Borregales",
        "fullName": "Andy Borregales",
        "displayName": "Andy Borregales",
        "shortName": "A. Borregales",
        "weight": 202.0,
        "displayWeight": "202 lbs",
        "height": 71.0,
        "displayHeight": "5' 11\"",
        "age": 23,
        "dateOfBirth": "2003-01-02T08:00Z",
        "links": [
          {
            "language": "en-US",
            "rel": [
              "playercard",
              "desktop",
              "athlete"
            ],
            "href": "https://www.espn.com/nfl/player/_/id/4569923/andy-borregales",
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
            "href": "https://www.espn.com/nfl/player/stats/_/id/4569923/andy-borregales",
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
            "href": "https://www.espn.com/nfl/player/splits/_/id/4569923/andy-borregales",
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
            "href": "https://www.espn.com/nfl/player/gamelog/_/id/4569923/andy-borregales",
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
            "href": "https://www.espn.com/nfl/player/news/_/id/4569923/andy-borregales",
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
            "href": "https://www.espn.com/nfl/player/bio/_/id/4569923/andy-borregales",
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
            "href": "https://www.espn.com/nfl/player/_/id/4569923/andy-borregales",
            "text": "Overview",
            "shortText": "Overview",
            "isExternal": false,
            "isPremium": false
          }
        ],
        "birthPlace": {
          "city": "Miami",
          "state": "FL",
          "country": "USA"
        },
        "college": {},
        "slug": "andy-borregales",
        "headshot": {
          "href": "https://a.espncdn.com/i/headshots/nfl/players/full/4569923.png",
          "alt": "Andy Borregales"
        },
        "jersey": "36",
        "position": {
          "id": "22",
          "name": "Place kicker",
          "displayName": "Place Kicker",
          "abbreviation": "PK",
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
          "years": 2
        },
        "collegeAthlete": {},
        "active": true,
        "eventLog": {},
        "draft": {
          "displayText": "Year: 2025 Round: 6 Pick: 182",
          "round": 6,
          "year": 2025,
          "selection": 182,
          "team": {}
        },
        "status": {
          "id": "1",
          "name": "Active",
          "type": "active",
          "abbreviation": "Active"
        }
      },
      "type": "kicker"
    }
  ]
}
```
