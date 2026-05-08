# Empty Or Unsupported League Endpoints

Verified with `league=nfl` on 2026-05-08.

---

## Rankings

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/rankings`

```json
{
  "count": 0,
  "pageIndex": 0,
  "pageSize": 25,
  "pageCount": 0,
  "items": []
}
```

---

## Recruiting

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/recruiting`

Notes:
- Present for football in general, but NFL currently returns no recruiting data.

```json
{
  "count": 0,
  "pageIndex": 0,
  "pageSize": 25,
  "pageCount": 0,
  "items": []
}
```

---

## Tournaments

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/tournaments`

Notes:
- NFL currently returns HTTP 400 with `getTournaments() not supported for football/nfl`.

```json
{
  "error": {
    "message": "getTournaments() not supported for football/nfl",
    "code": 400
  }
}
```
