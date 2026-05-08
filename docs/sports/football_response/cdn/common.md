# Football CDN Response Docs

Verified with live NFL requests on 2026-05-07.

- `scoreboard.md` - CDN scoreboard shell and embedded scoreboard data
- `game.md` - full game package view
- `boxscore.md` - boxscore-focused game package
- `playbyplay.md` - play-by-play-focused game package
- `matchup.md` - matchup-focused game package

Notes:
- CDN endpoints require `xhr=1` for JSON.
- Game-specific CDN views all return a page shell plus `gamepackageJSON`.
- `game`, `boxscore`, `playbyplay`, and `matchup` share the same top-level wrapper, but each view includes a different subset of `gamepackageJSON` keys.
