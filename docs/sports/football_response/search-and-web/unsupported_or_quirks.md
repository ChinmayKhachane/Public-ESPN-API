# Unsupported Or Quirky Search/Web Paths

Verified with live requests on 2026-05-07.

## Scoreboard Header

```text
https://site.web.api.espn.com/apis/v2/scoreboard/header?sport=football
```

Returned HTTP `400`.

Use this instead:

```text
https://site.web.api.espn.com/apis/v2/scoreboard/header?sport=football&league=nfl
```

## Search Shape

Search results are grouped by type. A result group has fields like `type`, `totalFound`, `page`, `limit`, `displayName`, and `contents`.

Do not expect `displayName`, `description`, or `link` on the result group itself. Read those from `results[].contents[]`.
