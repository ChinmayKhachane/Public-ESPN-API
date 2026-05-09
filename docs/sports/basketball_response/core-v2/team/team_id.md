# Team Detail

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/seasons/{season}/teams/{id}

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/teams/20?lang=en&region=us",
  "id": "20",
  "uid": "s:40~l:46~t:20",
  "guid": "ca1685ed-b799-53e4-7924-e58ea6eb8f3a",
  "name": "76ers",
  "displayName": "Philadelphia 76ers",
  "shortDisplayName": "76ers",
  "abbreviation": "PHI",
  "slug": "philadelphia-76ers",
  "events": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/teams/20/events?lang=en&region=us"
  },
  "athletes": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/teams/20/athletes?lang=en&region=us"
  },
  "statistics": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/types/2/teams/20/statistics?lang=en&region=us"
  }
}
```
