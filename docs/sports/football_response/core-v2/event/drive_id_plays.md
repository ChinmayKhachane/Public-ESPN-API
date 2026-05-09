# Competition Drive Plays

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/drives/{drive}/plays

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988`, `drive=4017729881` on 2026-05-08.

## Example Response

```json
{
  "count": 10,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 5,
  "items": [
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
      "drive": {},
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
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/plays/40177298857?lang=en&region=us",
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
      "drive": {},
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
    }
  ]
}
```
