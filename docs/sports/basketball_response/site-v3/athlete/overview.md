# Athlete Overview

## https://site.web.api.espn.com/apis/common/v3/sports/basketball/{league}/athletes/{id}/overview

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "statistics": {
    "displayName": "Stats",
    "splits": [
      {
        "displayName": "Regular Season",
        "stats": []
      },
      {
        "displayName": "Postseason",
        "stats": []
      }
    ],
    "labels": [
      "GP",
      "MIN"
    ],
    "names": [
      "gamesPlayed",
      "avgMinutes"
    ],
    "displayNames": [
      "Games Played",
      "Minutes Per Game"
    ]
  },
  "news": [
    {
      "id": 48716507,
      "type": "Media",
      "categories": [
        {},
        {}
      ],
      "headline": "Embiid recovers to stuff Josh Hart at the rim",
      "lastModified": "2026-05-09T01:46:27.000+00:00",
      "root": "nba",
      "premium": false,
      "links": {
        "api": {},
        "web": {}
      },
      "section": "NBA",
      "categorized": "2026-05-09T01:46:29Z",
      "description": "Embiid recovers to stuff Josh Hart at the rim",
      "nowId": "1-48716507"
    },
    {
      "id": 48715863,
      "type": "Story",
      "categories": [
        {},
        {}
      ],
      "headline": "VJ Edgecombe, Tyrese Maxey shine early; Mitchell Robinson puts Joel Embiid on poster",
      "lastModified": "2026-05-09T00:24:50.000+00:00",
      "root": "nba",
      "premium": false,
      "links": {
        "api": {},
        "mobile": {},
        "web": {}
      },
      "section": "NBA",
      "linkText": "'VJ Maxx' delivers early in Game 3, Mitchell Robinson answers with poster on Embiid",
      "categorized": "2026-05-09T02:10:53Z",
      "description": "Maxey and Edgecombe ignited the Sixers early, but Mitchell Robinson answered with a second-quarter poster on Joel Embiid."
    }
  ],
  "nextGame": {
    "displayName": "Previous Game",
    "statistics": {
      "splits": [
        {},
        {}
      ],
      "labels": [
        "GP",
        "MIN"
      ],
      "names": [
        "gamesPlayed",
        "avgMinutes"
      ],
      "displayNames": [
        "Games Played",
        "Minutes Per Game"
      ]
    },
    "league": {
      "id": "46",
      "uid": "s:40~l:46",
      "name": "National Basketball Association",
      "abbreviation": "NBA",
      "slug": "nba",
      "events": [
        {}
      ],
      "shortName": "NBA"
    },
    "summaryStatistics": [
      {
        "name": "points",
        "displayName": "Points",
        "shortDisplayName": "PTS",
        "abbreviation": "PTS",
        "displayValue": "18",
        "splitId": "0",
        "splitAbbreviation": "Total",
        "splitName": "All Splits",
        "splitBy": "lastTenGames"
      },
      {
        "name": "rebounds",
        "displayName": "Rebounds",
        "shortDisplayName": "REB",
        "abbreviation": "REB",
        "displayValue": "6",
        "splitId": "0",
        "splitAbbreviation": "Total",
        "splitName": "All Splits",
        "splitBy": "lastTenGames"
      }
    ]
  },
  "gameLog": {
    "displayName": "Recent Games",
    "events": {
      "401871161": {
        "id": "401871161",
        "team": {},
        "score": "108-94",
        "links": [],
        "atVs": "vs",
        "gameDate": "2026-05-08T23:00:00.000+00:00",
        "homeTeamId": "20",
        "awayTeamId": "18",
        "homeTeamScore": "94",
        "awayTeamScore": "108",
        "gameResult": "L",
        "opponent": {}
      },
      "401871159": {
        "id": "401871159",
        "team": {},
        "score": "137-98",
        "links": [],
        "atVs": "@",
        "gameDate": "2026-05-05T00:00:00.000+00:00",
        "homeTeamId": "18",
        "awayTeamId": "20",
        "homeTeamScore": "137",
        "awayTeamScore": "98",
        "gameResult": "L",
        "opponent": {}
      },
      "401869408": {
        "id": "401869408",
        "team": {},
        "score": "113-97",
        "links": [],
        "atVs": "@",
        "gameDate": "2026-04-28T23:00:00.000+00:00",
        "homeTeamId": "2",
        "awayTeamId": "20",
        "homeTeamScore": "97",
        "awayTeamScore": "113",
        "gameResult": "W",
        "opponent": {}
      },
      "401869412": {
        "id": "401869412",
        "team": {},
        "score": "109-100",
        "links": [],
        "atVs": "@",
        "gameDate": "2026-05-02T23:30:00.000+00:00",
        "homeTeamId": "2",
        "awayTeamId": "20",
        "homeTeamScore": "100",
        "awayTeamScore": "109",
        "gameResult": "W",
        "opponent": {}
      },
      "401869410": {
        "id": "401869410",
        "team": {},
        "score": "106-93",
        "links": [],
        "atVs": "vs",
        "gameDate": "2026-05-01T00:00:00.000+00:00",
        "homeTeamId": "20",
        "awayTeamId": "2",
        "homeTeamScore": "106",
        "awayTeamScore": "93",
        "gameResult": "W",
        "opponent": {}
      }
    },
    "statistics": [
      {
        "displayName": "totals",
        "events": [],
        "labels": [],
        "names": [],
        "displayNames": []
      }
    ]
  },
  "rotowire": {
    "headline": "Embiid (ankle) is available for Friday's Game 3 of the Eastern Conference Semifinals against New York, Tim Bontemps of ESPN.com reports.",
    "story": "Embiid didn't suit up for Game 2 and was considered day-to-day. However, he was spotted on the court during warmups and now officially has the green light to suit up. The 32-yea...",
    "description": "Embiid (ankle) is available for Friday's Game 3 of the Eastern Conference Semifinals against New York, Tim Bontemps of ESPN.com reports.",
    "published": "Fri May 08 15:35:29 PDT 2026"
  },
  "awards": [
    {
      "id": "33",
      "name": "MVP",
      "league": "nba",
      "displayCount": "1x",
      "seasons": [
        "2023"
      ]
    },
    {
      "id": "44",
      "name": "All-NBA 1st Team",
      "league": "nba",
      "displayCount": "1x",
      "seasons": [
        "2023"
      ]
    }
  ],
  "fantasy": {
    "draftRank": "38",
    "positionRank": "24",
    "percentOwned": "89.13",
    "last7Days": "-0.5",
    "projection": "Embiid was the MVP of the NBA just three seasons ago, but injuries have limited him to only 58 of a possible 164 games in the two seasons since. The injuries were so bad last se..."
  }
}
```
