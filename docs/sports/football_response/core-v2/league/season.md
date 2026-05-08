# Current Season

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/season

Notes:
- Verified with `league=nfl` on 2026-05-08.
- On 2026-05-08 this endpoint still resolved to the 2025 NFL season, with `type=4` (`Off Season`).

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
|  |  |  |

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025?lang=en&region=us",
  "year": 2025,
  "displayName": "2025",
  "startDate": "2025-07-31T07:00Z",
  "endDate": "2026-02-12T07:59Z",
  "type": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/types/4?lang=en&region=us",
    "id": "4",
    "type": 4,
    "name": "Off Season",
    "abbreviation": "off",
    "slug": "off-season"
  },
  "types": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/types?lang=en&region=us"
  },
  "athletes": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/athletes?lang=en&region=us"
  },
  "rankings": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/rankings?lang=en&region=us"
  },
  "coaches": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/coaches?lang=en&region=us"
  },
  "leaders": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/types/3/leaders?lang=en&region=us"
  }
}
```
