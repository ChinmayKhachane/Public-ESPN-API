# Athlete Statistics Log

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/athletes/{athlete}/statisticslog

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/athletes/4565230/statisticslog?limit=2`

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/athletes/4565230/statisticslog?lang=en&region=us",
  "entries": [
    {
      "season": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026?lang=en&region=us"
      },
      "statistics": [
        {
          "type": "total",
          "statistics": {}
        },
        {
          "type": "team",
          "team": {},
          "statistics": {}
        }
      ]
    },
    {
      "season": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2025?lang=en&region=us"
      },
      "statistics": [
        {
          "type": "total",
          "statistics": {}
        },
        {
          "type": "team",
          "team": {},
          "statistics": {}
        }
      ]
    }
  ]
}
```
