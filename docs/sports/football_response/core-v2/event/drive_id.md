# Competition Drive ID

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/drives/{drive}

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988`, `drive=4017729881` on 2026-05-08.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/drives/4017729881?lang=en&region=us",
  "id": "4017729881",
  "description": "8 plays, 51 yards, 3:02",
  "sequenceNumber": "1",
  "result": "FG",
  "start": {
    "yardLine": 65,
    "text": "SEA 35"
  },
  "end": {
    "yardLine": 14,
    "text": "NE 14"
  },
  "timeElapsed": {
    "displayValue": "3:02"
  },
  "yards": 51,
  "offensivePlays": 8,
  "isScore": true,
  "plays": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/drives/4017729881/plays?lang=en&region=us"
  }
}
```
