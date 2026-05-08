# Competition Competitor Leaders

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}/leaders

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988`, `competitor=17` on 2026-05-08.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17/leaders?lang=en&region=us",
  "name": "leaders",
  "abbreviation": "LDR",
  "categories": [
    {
      "name": "passingLeader",
      "displayName": "Passing Leader",
      "leaders": [
        {
          "displayValue": "27/43, 295 YDS, 2 TD, 2 INT",
          "athlete": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/athletes/4431452?lang=en&region=us"
          }
        }
      ]
    },
    {
      "name": "rushingLeader",
      "displayName": "Rushing Leader"
    },
    {
      "name": "receivingLeader",
      "displayName": "Receiving Leader"
    }
  ]
}
```
