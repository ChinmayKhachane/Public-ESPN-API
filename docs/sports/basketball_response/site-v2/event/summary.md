# Summary

## https://site.api.espn.com/apis/site/v2/sports/basketball/{league}/summary?event={event}

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "leaders": [
    {
      "team": {
        "id": "20",
        "uid": "s:40~l:46~t:20",
        "displayName": "Philadelphia 76ers",
        "abbreviation": "PHI",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/phi.png",
        "logos": []
      },
      "leaders": [
        {},
        {}
      ]
    },
    {
      "team": {
        "id": "18",
        "uid": "s:40~l:46~t:18",
        "displayName": "New York Knicks",
        "abbreviation": "NY",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/ny.png",
        "logos": []
      },
      "leaders": [
        {},
        {}
      ]
    }
  ],
  "boxscore": {
    "teams": [
      {
        "team": {},
        "statistics": [],
        "homeAway": "away",
        "displayOrder": 1
      },
      {
        "team": {},
        "statistics": [],
        "homeAway": "home",
        "displayOrder": 2
      }
    ],
    "players": [
      {
        "team": {},
        "statistics": [],
        "displayOrder": 1
      },
      {
        "team": {},
        "statistics": [],
        "displayOrder": 2
      }
    ]
  },
  "header": {
    "id": "401871161",
    "uid": "s:40~l:46~e:401871161",
    "season": {
      "year": 2026,
      "type": 3,
      "current": true
    },
    "competitions": [
      {
        "id": "401871161",
        "uid": "s:40~l:46~e:401871161~c:401871161",
        "status": {},
        "competitors": [],
        "date": "2026-05-08T23:00Z",
        "broadcasts": [],
        "neutralSite": false,
        "conferenceCompetition": false,
        "boxscoreAvailable": true,
        "commentaryAvailable": false,
        "liveAvailable": false,
        "shotChartAvailable": true
      }
    ],
    "league": {
      "id": "46",
      "uid": "s:40~l:46",
      "name": "National Basketball Association",
      "abbreviation": "NBA",
      "slug": "nba",
      "isTournament": false,
      "links": [
        {},
        {}
      ],
      "logos": [
        {},
        {}
      ]
    },
    "timeValid": true,
    "links": [
      {
        "text": "Gamecast",
        "shortText": "Summary",
        "rel": [],
        "href": "https://www.espn.com/nba/game/_/gameId/401871161/knicks-76ers",
        "isExternal": false,
        "isPremium": false
      },
      {
        "text": "Recap",
        "shortText": "Recap",
        "rel": [],
        "href": "https://www.espn.com/nba/recap?gameId=401871161",
        "isExternal": false,
        "isPremium": false
      }
    ],
    "gameNote": "East Semifinals - Game 3"
  },
  "plays": [
    {
      "id": "4018711614",
      "type": {
        "id": "615",
        "text": "Jumpball"
      },
      "team": {
        "id": "20"
      },
      "text": "Karl-Anthony Towns vs. Joel Embiid (Kelly Oubre Jr. gains possession)",
      "period": {
        "displayValue": "1st Quarter",
        "number": 1
      },
      "clock": {
        "displayValue": "12:00"
      },
      "homeScore": 0,
      "awayScore": 0,
      "scoringPlay": false,
      "scoreValue": 0,
      "sequenceNumber": "4",
      "participants": [
        {},
        {}
      ]
    },
    {
      "id": "4018711617",
      "type": {
        "id": "110",
        "text": "Driving Layup Shot"
      },
      "team": {
        "id": "20"
      },
      "text": "Kelly Oubre Jr. makes driving layup",
      "period": {
        "displayValue": "1st Quarter",
        "number": 1
      },
      "clock": {
        "displayValue": "11:55"
      },
      "homeScore": 2,
      "awayScore": 0,
      "scoringPlay": true,
      "scoreValue": 2,
      "sequenceNumber": "7",
      "participants": [
        {}
      ]
    }
  ],
  "news": {
    "articles": [
      {
        "id": 48717062,
        "type": "HeadlineNews",
        "categories": [],
        "nowId": "1-48717062",
        "contentKey": "48717062-1-5-1",
        "dataSourceIdentifier": "e02cdd9cd0568",
        "headline": "Jalen Brunson shuts door on 76ers as Knicks go up 3-0",
        "description": "Behind 33 points from Jalen Brunson, including big buckets late, the Knicks held off the 76ers to take a commanding 3-0 lead in the semifinal series.",
        "lastModified": "2026-05-09T02:49:38Z",
        "published": "2026-05-09T02:49:38Z",
        "images": [],
        "premium": false
      },
      {
        "id": 48676289,
        "type": "Story",
        "categories": [],
        "nowId": "1-48676289",
        "contentKey": "48676289-1-6-1",
        "dataSourceIdentifier": "67dd576b8b9d4",
        "headline": "2026 NBA playoffs: Conference semifinals takeaways",
        "description": "Here's what we've learned -- and what's next -- for the 76ers-Knicks, Cavs-Pistons, Wolves-Spurs and Lakers-Thunder series.",
        "lastModified": "2026-05-09T02:30:20Z",
        "published": "2026-05-09T02:30:20Z",
        "images": [],
        "premium": false
      }
    ],
    "header": "NBA News",
    "link": {
      "text": "All NBA News",
      "shortText": "All News",
      "language": "en",
      "rel": [
        "index",
        "desktop"
      ],
      "href": "https://www.espn.com/nba/",
      "isExternal": false,
      "isPremium": false
    }
  },
  "injuries": [
    {
      "team": {
        "id": "20",
        "uid": "s:40~l:46~t:20",
        "displayName": "Philadelphia 76ers",
        "abbreviation": "PHI",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/phi.png",
        "logos": []
      },
      "injuries": []
    },
    {
      "team": {
        "id": "18",
        "uid": "s:40~l:46~t:18",
        "displayName": "New York Knicks",
        "abbreviation": "NY",
        "links": [],
        "logo": "https://a.espncdn.com/i/teamlogos/nba/500/ny.png",
        "logos": []
      },
      "injuries": [
        {}
      ]
    }
  ],
  "broadcasts": [
    {
      "type": {
        "id": "4",
        "slug": "streaming",
        "shortName": "Streaming",
        "longName": "Streaming"
      },
      "station": "Prime Video",
      "market": {
        "id": "1",
        "type": "National"
      },
      "media": {
        "name": "Prime Video",
        "callLetters": "Prime Video",
        "shortName": "Prime Video"
      },
      "lang": "en",
      "region": "us",
      "isNational": true
    },
    {
      "type": {
        "id": "5",
        "slug": "radio",
        "shortName": "Radio",
        "longName": "Radio"
      },
      "station": "13715732",
      "stationKey": "espn",
      "market": {
        "id": "1",
        "type": "National"
      },
      "media": {
        "name": "ESPN Radio",
        "callLetters": "ERADM",
        "shortName": "ERADM"
      },
      "lang": "en",
      "region": "us",
      "isNational": true
    }
  ],
  "odds": [],
  "format": {
    "regulation": {
      "displayName": "Quarter",
      "slug": "quarter",
      "clock": 720.0,
      "periods": 4
    },
    "overtime": {
      "displayName": "Quarter",
      "slug": "quarter",
      "clock": 300.0
    }
  },
  "gameInfo": {
    "venue": {
      "id": "1845",
      "guid": "5c91d4ea-46af-3b85-a87d-54422ed1e8c3",
      "fullName": "Xfinity Mobile Arena",
      "shortName": "Wells Fargo Center",
      "address": {
        "city": "Philadelphia",
        "state": "PA"
      },
      "grass": false,
      "images": [
        {}
      ]
    },
    "officials": [
      {
        "displayName": "Josh Tiven",
        "fullName": "Josh Tiven",
        "position": {},
        "order": 1
      },
      {
        "displayName": "Marc Davis",
        "fullName": "Marc Davis",
        "position": {},
        "order": 2
      }
    ],
    "attendance": 19746
  },
  "seasonseries": [
    {
      "type": "season",
      "events": [
        {},
        {}
      ],
      "title": "Regular Season Series",
      "description": "East Semifinals",
      "summary": "Series tied 2-2",
      "completed": true,
      "totalCompetitions": 4,
      "seriesLabel": "Regular Season",
      "seriesScore": "2-2",
      "shortSummary": "Season tied"
    },
    {
      "type": "playoff",
      "events": [
        {},
        {}
      ],
      "title": "Playoff Series",
      "description": "East Semifinals",
      "summary": "NY leads series 3-0",
      "completed": false,
      "totalCompetitions": 7,
      "seriesLabel": "Playoffs",
      "seriesScore": "3-0",
      "shortSummary": "NY leads series",
      "round": "East Semifinals"
    }
  ],
  "pickcenter": [
    {
      "header": {
        "text": "Game Odds",
        "logo": {}
      },
      "provider": {
        "id": "100",
        "name": "Draft Kings",
        "priority": 1,
        "logos": []
      },
      "details": "PHI -3.5",
      "overUnder": 214.5,
      "spread": -3.5,
      "overOdds": -110.0,
      "underOdds": -110.0,
      "awayTeamOdds": {
        "team": {},
        "favorite": false,
        "underdog": true,
        "moneyLine": 130,
        "spreadOdds": -115.0,
        "teamId": "18",
        "favoriteAtOpen": true
      },
      "homeTeamOdds": {
        "team": {},
        "favorite": true,
        "underdog": false,
        "moneyLine": -155,
        "spreadOdds": -105.0,
        "teamId": "20",
        "favoriteAtOpen": false
      },
      "links": [
        {},
        {}
      ],
      "moneyline": {
        "displayName": "Moneyline",
        "shortDisplayName": "ML",
        "home": {},
        "away": {}
      },
      "pointSpread": {
        "displayName": "Spread",
        "shortDisplayName": "Spread",
        "home": {},
        "away": {}
      }
    }
  ]
}
```
