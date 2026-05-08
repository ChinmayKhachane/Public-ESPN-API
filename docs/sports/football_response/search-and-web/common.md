# Football Search And Web Response Docs

Verified with live requests on 2026-05-07.

- `search_v2.md` - global and sport-scoped ESPN search
- `scoreboard_header.md` - site web scoreboard header/navigation state
- `unsupported_or_quirks.md` - tested parameter quirks

Notes:
- Search results are grouped by result type. Individual records live under each result group's `contents[]`.
- Scoreboard header filtering requires both `sport=football` and `league=nfl` for the tested NFL-specific response.
