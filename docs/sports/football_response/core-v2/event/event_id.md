# Event ID

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}

Notes:
- Verified with `league=nfl`, `event=401772988` on 2026-05-08.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
|  |  |  |

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988?lang=en&region=us",
  "id": "401772988",
  "uid": "s:20~l:28~e:401772988",
  "date": "2026-02-08T23:30Z",
  "name": "Seattle Seahawks at New England Patriots",
  "shortName": "SEA VS NE",
  "season": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025?lang=en&region=us"
  },
  "seasonType": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/types/3?lang=en&region=us"
  },
  "week": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/types/3/weeks/5?lang=en&region=us"
  },
  "competitions": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988?lang=en&region=us",
      "id": "401772988"
    }
  ]
}
```
