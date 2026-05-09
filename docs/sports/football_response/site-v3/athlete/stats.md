# Athlete Stats

## https://site.web.api.espn.com/apis/common/v3/sports/football/{league}/athletes/{id}/stats

Notes:
- Verified with `league=nfl`, `id=4431452` on 2026-05-08.
- The response is filter-driven and organized by stat categories, with a `teams` map and glossary metadata.

## Example Response

```json
{
  "filters": [
    {
      "displayName": "League",
      "name": "league",
      "value": "nfl"
    },
    {
      "displayName": "Season",
      "name": "seasontype",
      "value": "2"
    }
  ],
  "categories": [
    {
      "name": "passing",
      "displayName": "Passing",
      "labels": ["GP", "CMP", "ATT", "CMP%"],
      "statistics": []
    },
    {
      "name": "rushing",
      "displayName": "Rushing",
      "labels": ["GP", "CAR", "YDS", "AVG"],
      "statistics": []
    }
  ],
  "teams": {
    "new-england-patriots": {
      "id": "17",
      "displayName": "New England Patriots",
      "abbreviation": "NE"
    }
  },
  "glossary": []
}
```
