# Season Subresources

Verified with `league=nfl` and `season=2025` on 2026-05-08.

---

## Season Athletes

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/seasons/{season}/athletes`

```json
{
  "count": 20155,
  "pageIndex": 1,
  "pageSize": 1,
  "pageCount": 20155,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/athletes/4246273?lang=en&region=us"
    }
  ]
}
```

---

## Season Draft

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/seasons/{season}/draft`

Notes:
- `season=2025` returned a fully populated draft object.
- This is an object response, not a paginated collection.

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/draft?lang=en&region=us",
  "uid": "s:20~l:28~e:draft~y:2025",
  "year": 2025,
  "numberOfRounds": 7,
  "displayName": "2025 National Football League Draft",
  "shortDisplayName": "2025 NFL Draft",
  "status": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/draft/status?lang=en&region=us"
  },
  "athletes": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/draft/athletes?lang=en&region=us"
  },
  "rounds": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/draft/rounds?lang=en&region=us"
  },
  "startDate": "2025-04-25T00:00Z",
  "endDate": "2025-04-28T03:59Z"
}
```

---

## Season Free Agents

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/seasons/{season}/freeagents`

Notes:
- `season=2025` returned an empty collection.

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

## Season Manufacturers

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/seasons/{season}/manufacturers`

Notes:
- NFL currently returns HTTP 400 with `getManufacturers() not supported for football/nfl`.

```json
{
  "error": {
    "message": "getManufacturers() not supported for football/nfl",
    "code": 400
  }
}
```
