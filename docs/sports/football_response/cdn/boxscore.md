# CDN Boxscore

## https://cdn.espn.com/core/nfl/boxscore?xhr=1&gameId={event}

Notes:
- Verified with `gameId=401772988` on 2026-05-07.
- This endpoint returns the same CDN wrapper as the full game view.
- The `gamepackageJSON` subset is centered on `boxscore`.
- The tested game returned two team boxscore entries and two player boxscore entries.

## Example Response

```json
{
  "gameId": 401772988,
  "gamepackageJSON": {
    "news": {
      "articles": []
    },
    "winprobability": [],
    "boxscore": {
      "teams": [{}, {}],
      "players": [{}, {}]
    },
    "header": {},
    "broadcasts": [],
    "videos": [],
    "standings": {}
  }
}
```
