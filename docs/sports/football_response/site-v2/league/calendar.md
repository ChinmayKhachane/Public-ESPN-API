# Calendar

## https://site.api.espn.com/apis/site/v2/sports/football/{league}/calendar

Notes:
- Verified with `league=nfl` on 2026-05-08.
- NFL currently returns `404` here.
- Use the embedded `calendar` array from the scoreboard response instead.

## Example Response

```json
{
  "code": 404
}
```
