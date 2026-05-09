# Competition Competitor Leaders

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}/leaders

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/competitors/20/leaders?lang=en&region=us",
  "id": "0",
  "name": "Total",
  "abbreviation": "Total",
  "type": "total",
  "categories": [
    {
      "name": "points",
      "displayName": "Points",
      "shortDisplayName": "Pts",
      "abbreviation": "Pts",
      "leaders": [
        {},
        {}
      ]
    },
    {
      "name": "assists",
      "displayName": "Assists",
      "shortDisplayName": "Ast",
      "abbreviation": "Ast",
      "leaders": [
        {},
        {}
      ]
    }
  ]
}
```
