# Core Standings

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/standings

Notes:
- Verified with `league=nba` on 2026-05-09.
- The NBA core standings route returns a `$ref` to the current season/type/group standings collection.
- Following the returned `$ref` resolved successfully and returned standings views such as `overall`, `expanded`, `division`, and `inseason`.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/types/2/groups/7/standings?lang=en&region=us"
}
```

## Resolved Reference Example

```json
{
  "count": 4,
  "pageIndex": 1,
  "pageSize": 25,
  "pageCount": 1,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/types/2/groups/7/standings/0?lang=en&region=us",
      "id": "0",
      "name": "overall",
      "displayName": "Standings"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/types/2/groups/7/standings/1?lang=en&region=us",
      "id": "1",
      "name": "expanded",
      "displayName": "Expanded Standings"
    }
  ]
}
```
