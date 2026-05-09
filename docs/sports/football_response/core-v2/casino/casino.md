# Casinos

Verified with `league=nfl` on 2026-05-08.

---

## Collection

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/casinos`

```json
{
  "count": 63,
  "pageIndex": 1,
  "pageSize": 1,
  "pageCount": 63,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/casinos/52?lang=en&region=us",
      "id": "52",
      "name": "Caesars Sportsbook (Colorado)",
      "active": 1,
      "priority": 0
    }
  ]
}
```

---

## Casino ID

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/casinos/{id}`

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/casinos/52?lang=en&region=us",
  "id": "52",
  "name": "Caesars Sportsbook (Colorado)",
  "active": true,
  "priority": 19
}
```
