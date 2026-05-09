# Event Competition

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/events/{event}/competitions/{competition}

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.
- This response exposes most useful child refs for boxscore, team stats, plays, odds, broadcasts, status, and situation.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161?lang=en&region=us",
  "id": "401871161",
  "uid": "s:40~l:46~e:401871161~c:401871161",
  "guid": "55df9197-9744-326e-b9f2-ef8e591d40d5",
  "type": {
    "id": "15",
    "abbreviation": "QTR",
    "slug": "quarterfinal",
    "type": "quarterfinal",
    "text": "Quarterfinal"
  },
  "status": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/status?lang=en&region=us"
  },
  "competitors": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20?lang=en&region=us",
      "id": "20",
      "uid": "s:40~l:46~t:20",
      "type": "team",
      "team": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/teams/20?lang=en&region=us"
      },
      "statistics": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20/statistics?lang=en&region=us"
      },
      "leaders": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20/leaders?lang=en&region=us"
      },
      "score": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20/score?lang=en&region=us"
      },
      "linescores": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20/linescores?lang=en&region=us"
      },
      "roster": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20/roster?lang=en&region=us"
      },
      "record": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/types/2/teams/20/record?lang=en&region=us"
      },
      "winner": false
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/18?lang=en&region=us",
      "id": "18",
      "uid": "s:40~l:46~t:18",
      "type": "team",
      "team": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/teams/18?lang=en&region=us"
      },
      "statistics": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/18/statistics?lang=en&region=us"
      },
      "leaders": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/18/leaders?lang=en&region=us"
      },
      "score": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/18/score?lang=en&region=us"
      },
      "linescores": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/18/linescores?lang=en&region=us"
      },
      "roster": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/18/roster?lang=en&region=us"
      },
      "record": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/types/2/teams/18/record?lang=en&region=us"
      },
      "winner": true
    }
  ],
  "venue": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/venues/1845?lang=en&region=us",
    "id": "1845",
    "guid": "5c91d4ea-46af-3b85-a87d-54422ed1e8c3",
    "fullName": "Xfinity Mobile Arena",
    "shortName": "Wells Fargo Center",
    "address": {
      "city": "Philadelphia",
      "state": "PA"
    },
    "grass": false,
    "indoor": true,
    "images": [
      {
        "href": "https://a.espncdn.com/i/venues/nba/day/1845.jpg",
        "width": 2000,
        "height": 1125,
        "alt": "",
        "rel": []
      }
    ]
  },
  "date": "2026-05-08T23:00Z",
  "timeValid": true,
  "broadcasts": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/broadcasts?lang=en&region=us"
  },
  "officials": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/officials?lang=en&region=us"
  }
}
```
