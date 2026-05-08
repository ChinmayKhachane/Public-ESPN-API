# Venues

Verified with `league=nfl` on 2026-05-08.

---

## Collection

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/venues`

```json
{
  "count": 674,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 337,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/venues/16?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/venues/24?lang=en&region=us"
    }
  ]
}
```

---

## Venue ID

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/venues/{id}`

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/venues/4738?lang=en&region=us",
  "id": "4738",
  "guid": "ad9d3113-9b26-3c9a-98a9-250109205ef9",
  "fullName": "Levi's Stadium",
  "address": {
    "city": "Santa Clara",
    "state": "CA",
    "zipCode": "95054",
    "country": "USA"
  },
  "grass": true,
  "indoor": false
}
```
