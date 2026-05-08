# Athlete Bio

## https://site.web.api.espn.com/apis/common/v3/sports/football/{league}/athletes/{id}/bio

Notes:
- Verified with `league=nfl`, `id=4431452` on 2026-05-07.
- The NFL response is much thinner than `overview`, `stats`, or `gamelog`.
- For the tested player, the entire payload was just `teamHistory`.

## Example Response

```json
{
  "teamHistory": [
    {
      "displayName": "New England Patriots",
      "items": []
    }
  ]
}
```
