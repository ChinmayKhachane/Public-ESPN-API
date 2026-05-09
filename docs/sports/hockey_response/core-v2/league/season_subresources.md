# Season Type Leaders

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/seasons/{season}/types/3/leaders

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.
- Represents a resolved season child resource from the season/type graph.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/3/leaders?limit=2`

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/3/leaders?lang=en&region=us",
  "id": "0",
  "name": "Season",
  "abbreviation": "Season",
  "type": "season",
  "categories": [
    {
      "name": "goals",
      "displayName": "Goals",
      "shortDisplayName": "G",
      "abbreviation": "G",
      "leaders": [
        {
          "displayValue": "6",
          "value": 6.0,
          "rel": [],
          "athlete": {},
          "team": {},
          "statistics": {}
        },
        {
          "displayValue": "6",
          "value": 6.0,
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
      "shortDisplayName": "A",
      "abbreviation": "A",
      "leaders": [
        {
          "displayValue": "10",
          "value": 10.0,
          "rel": [],
          "athlete": {},
          "team": {},
          "statistics": {}
        },
        {
          "displayValue": "9",
          "value": 9.0,
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
