# MLB CDN Matchup

## https://cdn.espn.com/core/mlb/matchup?xhr=1&gameId=401815256

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200`, but response content type is `text/html;charset=UTF-8` rather than JSON.

## Example Response

```json
{
  "rawPreview": " <!DOCTYPE html> <html class=\"no-icon-fonts\" lang=\"en\"> <head> <meta http-equiv=\"content-type\" content=\"tex...",
  "curlError": null
}
```
