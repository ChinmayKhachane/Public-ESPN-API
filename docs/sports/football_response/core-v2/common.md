# Football Core V2 Response Docs

Verified with live NFL requests on 2026-05-07 and 2026-05-08.

Many core-v2 collection responses return `$ref`-only `items`. Some examples in
these docs show the referenced resource after a follow-up curl so the useful
response shape is visible; the original `$ref` is kept when present.

- `athlete/` - athlete collection, athlete detail, and athlete stat endpoints
- `league/` - league root, calendar, season, seasons, and season subresources
- `team/` - season-scoped team collection and team detail
- `event/` - events, event detail, competition detail, and leaf docs for status, situation, plays, drives, probabilities, predictor, powerindex, notes, and competitor subresources
- `media/` - media collection and media detail
- `venue/` - venue collection and venue detail
- `casino/` - casino collection and casino detail
- `misc/` - circuits, countries, franchises, positions, and providers
