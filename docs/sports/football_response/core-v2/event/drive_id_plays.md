# Competition Drive Plays

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/drives/{drive}/plays

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988`, `drive=4017729881` on 2026-05-08.

## Example Response

```json
{
  "count": 10,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 5,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/plays/40177298840?lang=en&region=us",
      "id": "40177298840",
      "type": {
        "text": "Kickoff"
      }
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/plays/40177298857?lang=en&region=us",
      "id": "40177298857",
      "type": {
        "text": "Rush"
      }
    }
  ]
}
```
