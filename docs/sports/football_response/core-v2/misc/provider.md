# Providers

Verified with `league=nfl` on 2026-05-08.

---

## Collection

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/providers`

```json
{
  "count": 71,
  "pageIndex": 1,
  "pageSize": 1,
  "pageCount": 71,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/providers/0?lang=en&region=us"
    }
  ]
}
```

---

## Provider ID

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/providers/{id}`

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/providers/100?lang=en&region=us",
  "id": "100",
  "name": "Draft Kings",
  "priority": 1
}
```
