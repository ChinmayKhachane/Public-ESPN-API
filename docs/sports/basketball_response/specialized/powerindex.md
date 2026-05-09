# Season Power Index

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/{year}/powerindex

Notes:
- Verified with `league=nba`, `year=2026` on 2026-05-09.
- The NBA season resource exposes this path via `powerIndexes`.
- The collection returned one entry per team/season-type context in live testing.
- NBA `powerindex/leaders` returned an empty collection.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `page` | `int` | Page number |
| `limit` | `int` | Page size |

## Example Response

```json
{
  "count": 90,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 45,
  "items": [
    {
      "team": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/teams/25?lang=en&region=us"
      },
      "season": 2026,
      "seasonType": 3,
      "lastUpdated": "2026-05-09T02:24Z",
      "runDateTimeKey": 20260508220250,
      "stats": [
        {
          "name": "bpi",
          "displayName": "BPI",
          "displayValue": "11.493"
        },
        {
          "name": "bpirank",
          "displayName": "BPI RANK",
          "displayValue": "1st"
        }
      ]
    },
    {
      "team": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/teams/25?lang=en&region=us"
      },
      "season": 2026,
      "seasonType": 5,
      "lastUpdated": "2026-04-14T10:03Z",
      "runDateTimeKey": 20260414050007
    }
  ]
}
```
