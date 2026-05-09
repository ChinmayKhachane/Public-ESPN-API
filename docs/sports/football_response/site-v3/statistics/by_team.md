# Statistics By Team

## https://site.web.api.espn.com/apis/common/v3/sports/football/{league}/statistics/byteam

Notes:
- Verified with `league=nfl`, `season=2025`, `seasontype=2` on 2026-05-07.
- This endpoint works for NFL and returns team leaderboard data grouped into category blocks.
- Each `teams[]` entry contains a `team` object plus a `categories[]` array.
- The top-level `categories[]` metadata contains label/name arrays. The useful
  totals and ranks live inside each team's category entries.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `season` | `int` | Season year |
| `seasontype` | `int` | `1=pre`, `2=regular`, `3=post` |
| `category` | `string` | Category filter such as `passing` |
| `sort` | `string` | Sort field such as `passing.netYardsPerGame:desc` |
| `limit` | `int` | Results per page |
| `page` | `int` | Page number |

## Example Response

```json
{
  "currentSeason": {
    "year": 2025,
    "type": {
      "id": "4",
      "name": "Off Season"
    }
  },
  "requestedSeason": {
    "year": 2025,
    "type": {
      "id": "2",
      "name": "Regular Season"
    }
  },
  "categories": [
    {
      "name": "general",
      "labels": ["GP", "FR", "LST"]
    },
    {
      "name": "passing",
      "labels": ["TYDS", "YDS/G", "NYDS"]
    },
    {
      "name": "rushing",
      "labels": ["YDS", "YDS/G", "ATT"]
    }
  ],
  "teams": [
    {
      "team": {
        "id": "22",
        "displayName": "Arizona Cardinals",
        "abbreviation": "ARI"
      },
      "categories": [
        {
          "name": "general",
          "displayName": "Own General",
          "splitId": "0",
          "totals": ["17", "9", "10"],
          "ranks": ["1", "9", "7"]
        },
        {
          "name": "passing",
          "displayName": "Own Passing",
          "splitId": "0",
          "totals": ["5,937", "347.6", "3,955"],
          "ranks": ["14", "14", "7"]
        }
      ]
    }
  ]
}
```
