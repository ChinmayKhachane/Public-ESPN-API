# Transactions

## https://site.api.espn.com/apis/site/v2/sports/basketball/{league}/transactions

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "season": {
    "name": "Postseason",
    "displayName": "2025-26",
    "year": 2026,
    "type": 3
  },
  "status": "success",
  "timestamp": "2026-05-09T03:15:03Z",
  "count": 152,
  "pageIndex": 1,
  "pageSize": 25,
  "pageCount": 7,
  "transactions": [
    {
      "team": {
        "id": "30",
        "name": "Hornets",
        "displayName": "Charlotte Hornets",
        "abbreviation": "CHA",
        "location": "Charlotte",
        "color": "008ca8",
        "alternateColor": "1d1060",
        "logos": [],
        "links": []
      },
      "date": "2026-05-07T07:00Z",
      "description": "Signed head coach Charles Lee to a contract extension."
    },
    {
      "team": {
        "id": "19",
        "name": "Magic",
        "displayName": "Orlando Magic",
        "abbreviation": "ORL",
        "location": "Orlando",
        "color": "0150b5",
        "alternateColor": "9ca0a3",
        "logos": [],
        "links": []
      },
      "date": "2026-05-04T07:00Z",
      "description": "Fired head coach Jamahl Mosley."
    }
  ],
  "requestedYear": {
    "displayName": "2026",
    "year": 2026
  }
}
```
