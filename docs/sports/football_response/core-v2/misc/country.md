# Countries

Verified with `league=nfl` on 2026-05-08.

---

## Collection

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/countries`

```json
{
  "count": 255,
  "pageIndex": 1,
  "pageSize": 1,
  "pageCount": 255,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/countries/1?lang=en&region=us",
      "id": "1",
      "slug": "usa",
      "name": "USA",
      "abbreviation": "USA",
      "flag": {
        "href": "https://a.espncdn.com/i/teamlogos/countries/500/usa.png",
        "alt": "USA",
        "rel": [
          "country-flag"
        ]
      },
      "athletes": {}
    }
  ]
}
```

---

## Country ID

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/countries/{id}`

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/countries/1?lang=en&region=us",
  "id": "1",
  "slug": "usa",
  "name": "USA",
  "abbreviation": "USA",
  "flag": {
    "href": "https://a.espncdn.com/i/teamlogos/countries/500/usa.png",
    "alt": "USA"
  },
  "athletes": {
    "count": 0,
    "pageIndex": 0,
    "pageSize": 25,
    "pageCount": 0,
    "items": []
  }
}
```
