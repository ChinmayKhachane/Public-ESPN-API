# Athlete Game Log

## https://site.web.api.espn.com/apis/common/v3/sports/basketball/{league}/athletes/{id}/gamelog

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "events": {
    "401871161": {
      "id": "401871161",
      "team": {
        "id": "20",
        "uid": "s:40~l:46~t:20",
        "abbreviation": "PHI",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/phi.png",
        "isAllStar": false
      },
      "score": "108-94",
      "links": [
        {},
        {}
      ],
      "atVs": "vs",
      "gameDate": "2026-05-08T23:00:00.000+00:00",
      "homeTeamId": "20",
      "awayTeamId": "18",
      "homeTeamScore": "94",
      "awayTeamScore": "108",
      "gameResult": "L",
      "opponent": {
        "id": "18",
        "uid": "s:40~l:46~t:18",
        "displayName": "New York Knicks",
        "abbreviation": "NY",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/ny.png"
      }
    },
    "401871159": {
      "id": "401871159",
      "team": {
        "id": "20",
        "uid": "s:40~l:46~t:20",
        "abbreviation": "PHI",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/phi.png",
        "isAllStar": false
      },
      "score": "137-98",
      "links": [
        {},
        {}
      ],
      "atVs": "@",
      "gameDate": "2026-05-05T00:00:00.000+00:00",
      "homeTeamId": "18",
      "awayTeamId": "20",
      "homeTeamScore": "137",
      "awayTeamScore": "98",
      "gameResult": "L",
      "opponent": {
        "id": "18",
        "uid": "s:40~l:46~t:18",
        "displayName": "New York Knicks",
        "abbreviation": "NY",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/ny.png"
      }
    },
    "401869412": {
      "id": "401869412",
      "team": {
        "id": "20",
        "uid": "s:40~l:46~t:20",
        "abbreviation": "PHI",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/phi.png",
        "isAllStar": false
      },
      "score": "109-100",
      "links": [
        {},
        {}
      ],
      "atVs": "@",
      "gameDate": "2026-05-02T23:30:00.000+00:00",
      "homeTeamId": "2",
      "awayTeamId": "20",
      "homeTeamScore": "100",
      "awayTeamScore": "109",
      "gameResult": "W",
      "opponent": {
        "id": "2",
        "uid": "s:40~l:46~t:2",
        "displayName": "Boston Celtics",
        "abbreviation": "BOS",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/bos.png"
      }
    },
    "401869410": {
      "id": "401869410",
      "team": {
        "id": "20",
        "uid": "s:40~l:46~t:20",
        "abbreviation": "PHI",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/phi.png",
        "isAllStar": false
      },
      "score": "106-93",
      "links": [
        {},
        {}
      ],
      "atVs": "vs",
      "gameDate": "2026-05-01T00:00:00.000+00:00",
      "homeTeamId": "20",
      "awayTeamId": "2",
      "homeTeamScore": "106",
      "awayTeamScore": "93",
      "gameResult": "W",
      "opponent": {
        "id": "2",
        "uid": "s:40~l:46~t:2",
        "displayName": "Boston Celtics",
        "abbreviation": "BOS",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/bos.png"
      }
    },
    "401869408": {
      "id": "401869408",
      "team": {
        "id": "20",
        "uid": "s:40~l:46~t:20",
        "abbreviation": "PHI",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/phi.png",
        "isAllStar": false
      },
      "score": "113-97",
      "links": [
        {},
        {}
      ],
      "atVs": "@",
      "gameDate": "2026-04-28T23:00:00.000+00:00",
      "homeTeamId": "2",
      "awayTeamId": "20",
      "homeTeamScore": "97",
      "awayTeamScore": "113",
      "gameResult": "W",
      "opponent": {
        "id": "2",
        "uid": "s:40~l:46~t:2",
        "displayName": "Boston Celtics",
        "abbreviation": "BOS",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/bos.png"
      }
    },
    "401869406": {
      "id": "401869406",
      "team": {
        "id": "20",
        "uid": "s:40~l:46~t:20",
        "abbreviation": "PHI",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/phi.png",
        "isAllStar": false
      },
      "score": "128-96",
      "links": [
        {},
        {}
      ],
      "atVs": "vs",
      "gameDate": "2026-04-26T23:00:00.000+00:00",
      "homeTeamId": "20",
      "awayTeamId": "2",
      "homeTeamScore": "96",
      "awayTeamScore": "128",
      "gameResult": "L",
      "opponent": {
        "id": "2",
        "uid": "s:40~l:46~t:2",
        "displayName": "Boston Celtics",
        "abbreviation": "BOS",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/bos.png"
      }
    },
    "401811001": {
      "id": "401811001",
      "team": {
        "id": "20",
        "uid": "s:40~l:46~t:20",
        "abbreviation": "PHI",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/phi.png",
        "isAllStar": false
      },
      "score": "115-102",
      "links": [
        {},
        {}
      ],
      "atVs": "@",
      "gameDate": "2026-04-07T00:00:00.000+00:00",
      "homeTeamId": "24",
      "awayTeamId": "20",
      "homeTeamScore": "115",
      "awayTeamScore": "102",
      "gameResult": "L",
      "opponent": {
        "id": "24",
        "uid": "s:40~l:46~t:24",
        "displayName": "San Antonio Spurs",
        "abbreviation": "SA",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/sa.png"
      }
    },
    "401810976": {
      "id": "401810976",
      "team": {
        "id": "20",
        "uid": "s:40~l:46~t:20",
        "abbreviation": "PHI",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/phi.png",
        "isAllStar": false
      },
      "score": "115-103",
      "links": [
        {},
        {}
      ],
      "atVs": "vs",
      "gameDate": "2026-04-03T23:00:00.000+00:00",
      "homeTeamId": "20",
      "awayTeamId": "16",
      "homeTeamScore": "115",
      "awayTeamScore": "103",
      "gameResult": "W",
      "opponent": {
        "id": "16",
        "uid": "s:40~l:46~t:16",
        "displayName": "Minnesota Timberwolves",
        "abbreviation": "MIN",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/min.png"
      }
    },
    "401810946": {
      "id": "401810946",
      "team": {
        "id": "20",
        "uid": "s:40~l:46~t:20",
        "abbreviation": "PHI",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/phi.png",
        "isAllStar": false
      },
      "score": "119-109",
      "links": [
        {},
        {}
      ],
      "atVs": "@",
      "gameDate": "2026-03-30T23:00:00.000+00:00",
      "homeTeamId": "14",
      "awayTeamId": "20",
      "homeTeamScore": "119",
      "awayTeamScore": "109",
      "gameResult": "L",
      "opponent": {
        "id": "14",
        "uid": "s:40~l:46~t:14",
        "displayName": "Miami Heat",
        "abbreviation": "MIA",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/mia.png"
      }
    },
    "401810932": {
      "id": "401810932",
      "team": {
        "id": "20",
        "uid": "s:40~l:46~t:20",
        "abbreviation": "PHI",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/phi.png",
        "isAllStar": false
      },
      "score": "118-114",
      "links": [
        {},
        {}
      ],
      "atVs": "@",
      "gameDate": "2026-03-28T22:00:00.000+00:00",
      "homeTeamId": "30",
      "awayTeamId": "20",
      "homeTeamScore": "114",
      "awayTeamScore": "118",
      "gameResult": "W",
      "opponent": {
        "id": "30",
        "uid": "s:40~l:46~t:30",
        "displayName": "Charlotte Hornets",
        "abbreviation": "CHA",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/cha.png"
      }
    },
    "401810908": {
      "id": "401810908",
      "team": {
        "id": "20",
        "uid": "s:40~l:46~t:20",
        "abbreviation": "PHI",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/phi.png",
        "isAllStar": false
      },
      "score": "157-137",
      "links": [
        {},
        {}
      ],
      "atVs": "vs",
      "gameDate": "2026-03-25T23:00:00.000+00:00",
      "homeTeamId": "20",
      "awayTeamId": "4",
      "homeTeamScore": "157",
      "awayTeamScore": "137",
      "gameResult": "W",
      "opponent": {
        "id": "4",
        "uid": "s:40~l:46~t:4",
        "displayName": "Chicago Bulls",
        "abbreviation": "CHI",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/chi.png"
      }
    },
    "401810704": {
      "id": "401810704",
      "team": {
        "id": "20",
        "uid": "s:40~l:46~t:20",
        "abbreviation": "PHI",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/phi.png",
        "isAllStar": false
      },
      "score": "124-117",
      "links": [
        {},
        {}
      ],
      "atVs": "vs",
      "gameDate": "2026-02-27T00:00:00.000+00:00",
      "homeTeamId": "20",
      "awayTeamId": "14",
      "homeTeamScore": "124",
      "awayTeamScore": "117",
      "gameResult": "W",
      "opponent": {
        "id": "14",
        "uid": "s:40~l:46~t:14",
        "displayName": "Miami Heat",
        "abbreviation": "MIA",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/mia.png"
      }
    }
  },
  "labels": [
    "MIN",
    "FG"
  ],
  "names": [
    "minutes",
    "fieldGoalsMade-fieldGoalsAttempted"
  ],
  "displayNames": [
    "Minutes",
    "Field Goals Made-Attempted"
  ],
  "glossary": [
    {
      "displayName": "3-Point Field Goal Percentage",
      "abbreviation": "3P%"
    },
    {
      "displayName": "3-Point Field Goals Made-Attempted",
      "abbreviation": "3PT"
    }
  ],
  "filters": [
    {
      "name": "league",
      "displayName": "League",
      "value": "nba",
      "options": [
        {},
        {}
      ]
    },
    {
      "name": "season",
      "displayName": "Season",
      "value": "2026",
      "options": [
        {},
        {}
      ]
    }
  ],
  "seasonTypes": [
    {
      "displayName": "2025-26 Postseason",
      "categories": [
        {},
        {}
      ],
      "displayTeam": "PHI",
      "summary": {
        "displayName": "Postseason",
        "stats": []
      }
    },
    {
      "displayName": "2025-26 Regular Season",
      "categories": [
        {},
        {}
      ],
      "displayTeam": "PHI",
      "summary": {
        "displayName": "Regular Season Stats",
        "stats": []
      }
    }
  ]
}
```
