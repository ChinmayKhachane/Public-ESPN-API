# League

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}

Notes:
- Verified with `league=nfl` on 2026-05-08.
- The league root currently points at the 2025 NFL season and its offseason state.
- `draft` points to the upcoming draft resource, while `teams` points to a season-scoped team collection.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
|  |  |  |

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl?lang=en&region=us",
  "id": "28",
  "uid": "s:20~l:28",
  "name": "National Football League",
  "abbreviation": "NFL",
  "slug": "nfl",
  "season": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025?lang=en&region=us",
    "year": 2025,
    "displayName": "2025",
    "type": {
      "id": "4",
      "name": "Off Season",
      "slug": "off-season"
    }
  },
  "teams": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/teams?lang=en&region=us"
  },
  "athletes": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/athletes?lang=en&region=us"
  },
  "events": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events?lang=en&region=us"
  },
  "franchises": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/franchises?lang=en&region=us"
  },
  "calendar": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/calendar?lang=en&region=us"
  },
  "rankings": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/rankings?lang=en&region=us"
  },
  "draft": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2026/draft?lang=en&region=us"
  },
  "notes": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/notes?lang=en&region=us"
  }
}
```
