# CDN Game

## https://cdn.espn.com/core/nfl/game?xhr=1&gameId={event}

Notes:
- Verified with `gameId=401772988` on 2026-05-07.
- This is the broadest CDN game view.
- The useful game data is inside `gamepackageJSON`.
- The tested completed game included drives, scoring plays, win probability, game info, boxscore, broadcasts, leaders, videos, and standings.

## Example Response

```json
{
  "gameId": 401772988,
  "gamepackageJSON": {
    "header": {
      "id": "401772988",
      "season": {
        "current": true,
        "year": 2025,
        "type": 3
      },
      "week": 5,
      "competitions": [{}]
    },
    "drives": {
      "previous": []
    },
    "boxscore": {},
    "leaders": [],
    "standings": {}
  },
  "packageKeys": [
    "news",
    "pickcenter",
    "drives",
    "scoringPlays",
    "winprobability",
    "gameInfo",
    "boxscore",
    "header",
    "broadcasts",
    "leaders",
    "videos",
    "standings"
  ]
}
```
