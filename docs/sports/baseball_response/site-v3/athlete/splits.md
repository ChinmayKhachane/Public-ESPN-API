# MLB Athlete Splits

## https://site.web.api.espn.com/apis/common/v3/sports/baseball/mlb/athletes/4414528/splits

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "displayName": "2026 Splits",
  "filters": [
    {
      "name": "season",
      "displayName": "Season",
      "value": "2026",
      "options": [
        {
          "value": "2026",
          "displayValue": "2026"
        },
        {
          "value": "2025",
          "displayValue": "2025"
        }
      ]
    },
    {
      "name": "category",
      "displayName": "Category",
      "value": "pitching",
      "options": [
        {
          "value": "pitching",
          "displayValue": "Pitching",
          "shortDisplayName": "pitching"
        }
      ]
    }
  ],
  "labels": [
    "ERA",
    "W"
  ],
  "names": [
    "ERA",
    "wins"
  ],
  "displayNames": [
    "Earned Run Average",
    "Wins"
  ],
  "descriptions": [
    "The average number of runs a pitcher yields per 9 innings thrown",
    "The number of times the pitcher was attributed with a win"
  ],
  "splitCategories": [
    {
      "name": "split",
      "displayName": "Overall",
      "splits": [
        {
          "displayName": "All Splits",
          "abbreviation": "Total",
          "stats": [
            "5.13",
            "1"
          ]
        }
      ]
    },
    {
      "name": "byOpponentBatting",
      "displayName": "Opponent Batting",
      "splits": [
        {
          "displayName": "All Splits",
          "abbreviation": "Total",
          "stats": [
            "162",
            "24"
          ]
        }
      ],
      "extraAthleteSplitsType": "batting"
    }
  ],
  "extraPlayerPageAthleteSplits": {
    "batting": {
      "labels": [
        "AB",
        "R"
      ],
      "names": [
        "atBats",
        "runs"
      ],
      "displayNames": [
        "At Bats",
        "Runs"
      ]
    }
  }
}
```
