# Athlete Detail

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/athletes/{id}

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/athletes/3059318?lang=en&region=us",
  "id": "3059318",
  "uid": "s:40~l:46~a:3059318",
  "guid": "89410119-0011-5ba0-efed-7b0b6c756c23",
  "displayName": "Joel Embiid",
  "slug": "joel-embiid",
  "type": "basketball",
  "status": {
    "id": "1",
    "name": "Active",
    "abbreviation": "Active",
    "type": "active"
  },
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/teams/20?lang=en&region=us"
  },
  "statistics": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/athletes/3059318/statistics?lang=en&region=us"
  },
  "alternateIds": {
    "sdr": "3059318"
  },
  "firstName": "Joel"
}
```
