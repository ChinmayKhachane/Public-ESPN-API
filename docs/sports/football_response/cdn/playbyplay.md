# CDN Play By Play

## https://cdn.espn.com/core/nfl/playbyplay?xhr=1&gameId={event}

Notes:
- Verified with `gameId=401772988` on 2026-05-07.
- The tested completed game returned drives under `gamepackageJSON.drives.previous`.
- For this game, `gamepackageJSON.plays` was not present as a top-level list; play data should be read from drive packages when available.

## Example Response

```json
{
  "gameId": 401772988,
  "gamepackageJSON": {
    "drives": {
      "current": null,
      "previous": []
    },
    "scoringPlays": [],
    "winprobability": [],
    "boxscore": {},
    "header": {},
    "broadcasts": [],
    "videos": [],
    "standings": {}
  },
  "driveSummary": {
    "previousCount": 28
  }
}
```
