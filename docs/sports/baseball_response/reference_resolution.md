# Reference Resolution Checks

Verified on 2026-05-08.

Representative non-self `$ref` links were fetched with `curl` after the primary endpoint pass.

- `http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb?lang=en&region=us` resolved as HTTP `200` with keys: $ref, id, guid, uid, name, displayName, abbreviation, shortName.
- `http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events?lang=en&region=us` resolved as HTTP `200` with keys: $meta, count, pageIndex, pageSize, pageCount, items.
- `http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/athletes?lang=en&region=us` resolved as HTTP `200` with keys: count, pageIndex, pageSize, pageCount, items.
- `http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026?lang=en&region=us` resolved as HTTP `200` with keys: $ref, year, startDate, endDate, displayName, type, types, rankings.
