# Football Now API Response Docs

Verified with live requests on 2026-05-07.

- `sports_news.md` - global and sport-filtered news feed
- `filter_behavior.md` - behavior of `sport`, `leagues`, and `team` filters

Notes:
- The active response key is `headlines`, not `feed`.
- `limit=2` returned `resultsLimit=10` in tested responses, so ESPN may clamp or ignore very small limits.
