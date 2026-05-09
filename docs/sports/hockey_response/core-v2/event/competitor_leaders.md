# Competitor Leaders

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}/leaders

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15/leaders`

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15/leaders?lang=en&region=us",
  "id": "0",
  "name": "Total",
  "abbreviation": "TOTAL",
  "type": "total",
  "categories": [
    {
      "name": "goals",
      "displayName": "Goals",
      "shortDisplayName": "Goals",
      "abbreviation": "G",
      "leaders": [
        {
          "displayValue": "1",
          "value": 1.0,
          "rel": [],
          "athlete": {},
          "team": {},
          "statistics": {}
        }
      ]
    },
    {
      "name": "assists",
      "displayName": "Assists",
      "shortDisplayName": "Assists",
      "abbreviation": "A",
      "leaders": [
        {
          "displayValue": "1",
          "value": 1.0,
          "rel": [],
          "athlete": {},
          "team": {},
          "statistics": {}
        },
        {
          "displayValue": "1",
          "value": 1.0,
          "rel": [],
          "athlete": {},
          "team": {},
          "statistics": {}
        }
      ]
    }
  ]
}
```
