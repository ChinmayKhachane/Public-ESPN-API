# Basketball Core V2 Response Docs

Verified with live NBA requests on 2026-05-09.

Many core-v2 collection responses return `$ref`-only `items`. Keep self-referencing `$ref` values and follow non-self refs only when a child endpoint is explicitly documented.

- `league/` - league root, calendar, seasons, current season, draft, and unsupported season subresources
- `team/` - team collection and season-scoped team detail
- `athlete/` - athlete collection, athlete detail, stats, logs, and contracts
- `event/` - events, competition detail, status, situation, plays, probabilities, and competitor subresources
- `venue/` - venue collection and venue detail
- `misc/` - casinos, circuits, countries, franchises, positions, providers, recruiting, and tournaments
