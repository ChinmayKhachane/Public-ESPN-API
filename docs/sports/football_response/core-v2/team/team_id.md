# Team ID

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/seasons/{season}/teams/{id}

Notes:
- Verified with `league=nfl`, `season=2025`, `id=22` on 2026-05-08.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
|  |  |  |

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/teams/22?lang=en&region=us",
  "id": "22",
  "uid": "s:20~l:28~t:22",
  "slug": "arizona-cardinals",
  "displayName": "Arizona Cardinals",
  "abbreviation": "ARI",
  "location": "Arizona",
  "name": "Cardinals",
  "venue": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/venues/3970?lang=en&region=us"
  },
  "record": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/types/2/teams/22/record?lang=en&region=us"
  },
  "statistics": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/types/2/teams/22/statistics?lang=en&region=us"
  },
  "athletes": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/teams/22/athletes?lang=en&region=us"
  },
  "injuries": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/teams/22/injuries?lang=en&region=us"
  },
  "events": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/teams/22/events?lang=en&region=us"
  }
}
```
