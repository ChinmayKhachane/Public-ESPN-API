# Athlete Statistics Log

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/athletes/{id}/statisticslog

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/athletes/3059318/statisticslog?lang=en&region=us",
  "entries": [
    {
      "season": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026?lang=en&region=us"
      },
      "statistics": [
        {},
        {}
      ]
    },
    {
      "season": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2025?lang=en&region=us"
      },
      "statistics": [
        {},
        {}
      ]
    }
  ]
}
```
