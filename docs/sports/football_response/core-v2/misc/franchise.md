# Franchises

Verified with `league=nfl` on 2026-05-08.

---

## Collection

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/franchises`

```json
{
  "count": 32,
  "pageIndex": 1,
  "pageSize": 1,
  "pageCount": 32,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/franchises/1?lang=en&region=us"
    }
  ]
}
```

---

## Franchise ID

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/franchises/{id}`

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/franchises/1?lang=en&region=us",
  "id": "1",
  "uid": "s:20~l:28~f:1",
  "name": "Falcons",
  "displayName": "Atlanta Falcons",
  "abbreviation": "ATL",
  "location": "Atlanta",
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/teams/1?lang=en&region=us"
  },
  "venue": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/venues/5348?lang=en&region=us"
  }
}
```
