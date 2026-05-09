# Team Power Index

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/{year}/powerindex/{teamId}

Notes:
- Verified with `league=nba`, `year=2026`, `teamId=25` on 2026-05-09.
- The tested `teamId=25` was taken from the first item in the live NBA `powerindex` collection.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/powerindex/25?lang=en&region=us",
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/teams/25?lang=en&region=us"
  },
  "league": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba?lang=en&region=us"
  },
  "season": 2026,
  "seasonType": 3,
  "lastUpdated": "2026-05-09T02:24Z",
  "runDateTimeKey": 20260508220250,
  "stats": [
    {
      "name": "bpi",
      "displayName": "BPI",
      "description": "Basketball Power Index that measures team’s true strength on net points scale; expected point margin vs average opponent on neutral court.",
      "displayValue": ""
    },
    {
      "name": "bpirank",
      "displayName": "BPI RANK",
      "description": "Basketball Power Index Rank vs all NBA teams",
      "value": 1.0,
      "displayValue": "1st"
    },
    {
      "name": "numwins",
      "displayName": "NUM WINS",
      "value": 64.0,
      "displayValue": "64"
    }
  ],
  "maxGameDate": "2026-05-09T04:35Z",
  "runCutoffDate": "2026-05-08T10:00Z"
}
```
