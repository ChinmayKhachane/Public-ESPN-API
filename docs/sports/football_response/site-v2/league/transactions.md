# League Transactions

## https://site.api.espn.com/apis/site/v2/sports/football/{league}/transactions

Notes:
- Verified with `league=nfl` on 2026-05-08.

## Example Response

```json
{
  "count": 611,
  "pageIndex": 1,
  "pageSize": 25,
  "transactions": [
    {
      "date": "2026-05-07T07:00Z",
      "description": "Signed CB Darren Hall to a contract. Signed TE Kyle Pitts to a franchise tag contract.",
      "team": {
        "id": "1",
        "displayName": "Atlanta Falcons",
        "abbreviation": "ATL"
      }
    }
  ]
}
```
