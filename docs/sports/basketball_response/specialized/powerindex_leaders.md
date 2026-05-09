# Season Power Index Leaders

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/{year}/powerindex/leaders

Notes:
- Verified with `league=nba`, `year=2026` on 2026-05-09.
- The endpoint exists for NBA and returned HTTP `200`.
- NBA currently returns an empty collection.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `page` | `int` | Page number |
| `limit` | `int` | Page size |

## Example Response

```json
{
  "count": 0,
  "pageIndex": 0,
  "pageSize": 25,
  "pageCount": 0,
  "items": []
}
```
