# Hockey Core V2 Response Docs

Verified with live NHL requests on 2026-05-09.

Many core-v2 collection responses return `$ref`-only `items`. Representative
references were resolved in the athlete, team, event, play, and venue detail
docs so the reachable response shapes are visible.

- `athlete/` - athlete collection, athlete detail, statistics, statistics log, and contracts
- `league/` - league root, calendar, seasons, standings, and season child resources
- `team/` - season-scoped team collection and team detail
- `event/` - event, competition, plays, probabilities, odds, officials, and competitor subresources
- `venue/` - venue collection and venue detail
- `misc/` - casinos, countries, franchises, positions, providers, recruiting, and tournaments
