# Standings

## https://site.api.espn.com/apis/v2/sports/football/{league}/standings

Notes:
- Verified with `league=nfl` on 2026-05-08.
- This is the full standings tree. The `/apis/site/v2/.../standings` version only returns a stub link.

## Example Response

```json
{
  "name": "National Football League",
  "abbreviation": "NFL",
  "season": {
    "year": 2026,
    "displayName": "2026"
  },
  "children": [
    {
      "id": "8",
      "name": "American Football Conference",
      "abbreviation": "AFC",
      "standings": {
        "name": "overall",
        "seasonDisplayName": "2025",
        "entries": [
          {
            "team": {
              "id": "17",
              "displayName": "New England Patriots",
              "abbreviation": "NE"
            }
          }
        ]
      }
    }
  ]
}
```

---

## Site V2 Stub

`https://site.api.espn.com/apis/site/v2/sports/football/{league}/standings`

```json
{
  "fullViewLink": {
    "text": "Full Standings",
    "href": "https://www.espn.com/nfl/standings"
  }
}
```
