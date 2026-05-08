# Events

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events

Notes:
- Verified with `league=nfl` on 2026-05-08.
- During the 2025 offseason this default collection returned a single current-context event ref.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `page` | `int` | Page number |
| `limit` | `int` | Page size |

## Example Response

```json
{
  "count": 1,
  "pageIndex": 1,
  "pageSize": 1,
  "pageCount": 1,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988?lang=en&region=us"
    }
  ]
}
```
