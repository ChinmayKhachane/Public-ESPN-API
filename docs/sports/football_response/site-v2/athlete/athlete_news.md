# Athlete News

## https://site.api.espn.com/apis/site/v2/sports/football/{league}/athletes/{id}/news

Notes:
- Verified with `league=nfl`, `id=4429202` on 2026-05-08.
- Even when no articles are present, the endpoint still returns the expected wrapper.

## Example Response

```json
{
  "header": "{0} News",
  "articles": []
}
```
