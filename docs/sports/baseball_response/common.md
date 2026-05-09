# Baseball Response Docs

Verified with live MLB requests on 2026-05-08.

Primary fixtures used across these examples:
- `league=mlb`
- `season=2026`
- `event=401815256` (Houston Astros at Cincinnati Reds, May 8, 2026)
- `competition=401815256`
- `competitor=17` (Cincinnati Reds)
- `team id=17` (Cincinnati Reds)
- `athlete id=4414528` (Andrew Abbott)
- `play=4018152560000000059`

Examples are intentionally truncated, but fields shown come from live response payloads.

Endpoint coverage was built from previous football/basketball response docs, then extended with MLB-specific routes listed in `docs/sports/baseball.md`, `README.md`, and `docs/README.md`.

Status summary from validation: `86` HTTP 200 responses, `4` HTTP 400
responses, and `4` HTTP 404 responses.
