# Athletes

## https://sports.core.api.espn.com/v3/sports/football/{league}/athletes

Notes:
- Verified with `league=nfl`, `limit=2` on 2026-05-07.
- Response is paginated with `count`, `pageIndex`, `pageSize`, `pageCount`, and `items`.
- The default collection may include placeholder or inactive athlete records near the front of the result set.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `limit` | `int` | Results per page |
| `page` | `int` | Page number |
| `active` | `boolean` | Accepted, but tested `active=true` still returned inactive placeholder records near the front |
| `enable` | `string` | Enables additional fields when supported |

## Example Response

```json
{
  "count": 20159,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 10080,
  "items": [
    {
      "id": "4246273",
      "displayName": " [35]",
      "active": false
    },
    {
      "id": "4246281",
      "displayName": " [Downed]",
      "active": false
    }
  ]
}
```
