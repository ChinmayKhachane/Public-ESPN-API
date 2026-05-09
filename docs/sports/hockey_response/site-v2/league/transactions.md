# Transactions

## https://site.api.espn.com/apis/site/v2/sports/hockey/{league}/transactions

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/transactions?limit=2`

## Example Response

```json
{
  "timestamp": "2026-05-09T17:01:04Z",
  "status": "success",
  "season": {
    "year": 2026,
    "type": 3,
    "name": "Postseason",
    "displayName": "2025-26"
  },
  "requestedYear": {
    "year": 2026,
    "displayName": "2026"
  },
  "count": 765,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 383,
  "transactions": [
    {
      "date": "2026-05-08T07:00Z",
      "description": "Assigned D Jack Ahcan to Colorado (AHL).",
      "team": {
        "id": "17",
        "location": "Colorado",
        "name": "Avalanche",
        "abbreviation": "COL",
        "displayName": "Colorado Avalanche",
        "color": "860038",
        "alternateColor": "005ea3",
        "logos": [
          {},
          {}
        ],
        "links": [
          {},
          {}
        ]
      }
    },
    {
      "date": "2026-05-06T07:00Z",
      "description": "Assigned D Adam Ji\u0159\u00ed\u010dek and F Justin Carbonneau to Springfield (AHL).",
      "team": {
        "id": "19",
        "location": "St. Louis",
        "name": "Blues",
        "abbreviation": "STL",
        "displayName": "St. Louis Blues",
        "color": "0070b9",
        "alternateColor": "fdb71a",
        "logos": [
          {},
          {}
        ],
        "links": [
          {},
          {}
        ]
      }
    }
  ]
}
```
