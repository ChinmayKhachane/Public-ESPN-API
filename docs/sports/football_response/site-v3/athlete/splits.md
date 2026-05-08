# Athlete Splits

## https://site.web.api.espn.com/apis/common/v3/sports/football/{league}/athletes/{id}/splits

Notes:
- Verified with `league=nfl`, `id=4431452` on 2026-05-08.
- Split data is organized under `splitCategories`, with categories such as home/away and outcome.

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
      "name": "season",
      "value": "2025"
    }
  ],
  "categories": [
    {
      "name": "passing",
      "displayName": "Passing",
      "count": 10
    },
    {
      "name": "rushing",
      "displayName": "Rushing",
      "count": 5
    }
  ],
  "splitCategories": [
    {
      "name": "split",
      "displayName": "split",
      "splits": [
        {
          "displayName": "All Splits",
          "abbreviation": "Any"
        },
        {
          "displayName": "Home",
          "abbreviation": "Home"
        },
        {
          "displayName": "Away",
          "abbreviation": "Away"
        }
      ]
    },
    {
      "name": "byOutcome",
      "displayName": "Outcome"
    }
  ]
}
```
