# League Statistics

## https://site.api.espn.com/apis/site/v2/sports/football/{league}/statistics

Notes:
- Verified with `league=nfl` on 2026-05-08.
- This is a very large object. The useful content lives under `stats`, with category metadata and ranked athlete/team tables.

## Example Response

```json
{
  "league": {
    "id": "28",
    "name": "National Football League",
    "abbreviation": "NFL",
    "slug": "nfl"
  },
  "season": {
    "year": 2025
  },
  "stats": {
    "categories": [
      {
        "name": "passing",
        "displayName": "Passing"
      }
    ],
    "labels": [
      "CMP",
      "ATT",
      "YDS",
      "CMP%",
      "AVG",
      "TD"
    ]
  },
  "status": "success"
}
```
