# CDN Matchup

## https://cdn.espn.com/core/nfl/matchup?xhr=1&gameId={event}

Notes:
- Verified with `gameId=401772988` on 2026-05-07.
- Returns the standard CDN game wrapper with a matchup-oriented subset.
- The tested completed game had `pickcenter`, `gameInfo`, `leaders`, and `standings`, but no populated top-level `predictor` or `againstTheSpread` fields in `gamepackageJSON`.

## Example Response

```json
{
  "gameId": 401772988,
  "gamepackageJSON": {
    "news": {
      "articles": []
    },
    "pickcenter": [],
    "winprobability": [],
    "gameInfo": {},
    "boxscore": {},
    "header": {},
    "broadcasts": [],
    "leaders": [],
    "standings": {}
  },
  "content": {}
}
```
