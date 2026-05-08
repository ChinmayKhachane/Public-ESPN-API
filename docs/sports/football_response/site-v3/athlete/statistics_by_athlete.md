# Statistics By Athlete

## https://site.web.api.espn.com/apis/common/v3/sports/football/{league}/statistics/byathlete

Notes:
- Verified with `league=nfl`, `season=2025`, `seasontype=2` on 2026-05-08.
- Response is paginated and returns `athletes[]` plus category metadata.
- A tested NFL query using `category=passing` returned HTTP `400` on 2026-05-07, so category filtering appears less reliable here than on `statistics/byteam`.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `season` | `int` | Season year |
| `seasontype` | `int` | `1=pre`, `2=regular`, `3=post` |
| `category` | `string` | Category filter |
| `sort` | `string` | Sort field |
| `limit` | `int` | Results per page |
| `page` | `int` | Page number |

## Example Response

```json
{
  "league": {
    "id": "28",
    "name": "National Football League",
    "abbreviation": "NFL"
  },
  "requestedSeason": {
    "year": 2025,
    "type": {
      "id": "2",
      "name": "Regular Season"
    }
  },
  "pagination": {
    "count": 53,
    "limit": 2,
    "page": 1,
    "pages": 27
  },
  "athletes": [
    {
      "athlete": {
        "id": "4431452",
        "displayName": "Drake Maye",
        "teamId": "17",
        "teamShortName": "NE"
      }
    }
  ],
  "categories": [
    {
      "name": "general",
      "displayName": "Own General"
    },
    {
      "name": "passing",
      "displayName": "Own Passing"
    }
  ]
}
```
