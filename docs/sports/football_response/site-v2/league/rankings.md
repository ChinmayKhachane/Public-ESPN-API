# Rankings

## https://site.api.espn.com/apis/site/v2/sports/football/{league}/rankings

Notes:
- Verified with `league=nfl` on 2026-05-08.
- NFL currently returns HTTP-style error payload data rather than rankings.

## Example Response

```json
{
  "code": 404,
  "message": "Unable to get sport data for sport: football and league: nfl"
}
```
