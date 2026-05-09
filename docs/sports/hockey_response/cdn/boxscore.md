# CDN Boxscore

## https://cdn.espn.com/core/nhl/boxscore?xhr=1&gameId={event}

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `404`.
- Content type: `text/html; charset=iso-8859-1`.
- Endpoint returned an error payload; the response is documented as observed.
- Response was not parsed as JSON; raw body is wrapped for documentation.

Tested URL:
`https://cdn.espn.com/core/nhl/boxscore?xhr=1&gameId=401871412`

## Example Response

```json
{
  "raw": "<!DOCTYPE HTML PUBLIC \"-//IETF//DTD HTML 2.0//EN\">\n<html><head>\n<title>404 Not Found</title>\n</head><body>\n<h1>Not Found</h1>\n<p>The requested URL /nhl/boxscore/ was not found on this server.</p>\n</body></html>\n"
}
```
