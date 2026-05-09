# MLB Athlete Stats

## https://site.web.api.espn.com/apis/common/v3/sports/baseball/mlb/athletes/4414528/stats

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "categories": [
    {
      "name": "pitching",
      "displayName": "pitching",
      "statistics": [
        {
          "season": {
            "displayName": "2023",
            "year": 2023
          },
          "teamId": "17",
          "teamSlug": "cincinnati-reds",
          "stats": [
            "21",
            "21"
          ],
          "position": "SP"
        },
        {
          "season": {
            "displayName": "2024",
            "year": 2024
          },
          "teamId": "17",
          "teamSlug": "cincinnati-reds",
          "stats": [
            "25",
            "25"
          ],
          "position": "SP"
        }
      ],
      "labels": [
        "GP",
        "GS"
      ],
      "names": [
        "gamesPlayed",
        "gamesStarted"
      ],
      "displayNames": [
        "Games Played",
        "Games Started"
      ],
      "descriptions": [
        "Games Played",
        "The number of games started by a pitcher."
      ],
      "totals": [
        "83",
        "83"
      ],
      "averages": [
        "21",
        "21"
      ],
      "sortKey": "pitching"
    },
    {
      "name": "opponent-batting",
      "displayName": "Opponent Batting",
      "statistics": [
        {
          "season": {
            "displayName": "2023",
            "year": 2023
          },
          "teamId": "17",
          "teamSlug": "cincinnati-reds",
          "stats": [
            "1897",
            "459"
          ],
          "position": "SP"
        },
        {
          "season": {
            "displayName": "2024",
            "year": 2024
          },
          "teamId": "17",
          "teamSlug": "cincinnati-reds",
          "stats": [
            "2333",
            "586"
          ],
          "position": "SP"
        }
      ],
      "labels": [
        "P",
        "TBF"
      ],
      "names": [
        "pitches",
        "battersFaced"
      ],
      "displayNames": [
        "Pitches",
        "Batters Faced"
      ],
      "descriptions": [
        "The number of pitches a pitcher throws.",
        "The number of batters faced by a pitcher."
      ],
      "totals": [
        "7610",
        "1911"
      ],
      "sortKey": "opponent-batting"
    }
  ],
  "filters": [
    {
      "name": "category",
      "displayName": "Category",
      "value": "pitching",
      "options": [
        {
          "value": "pitching",
          "displayValue": "Pitching",
          "shortDisplayName": "Pitching"
        },
        {
          "value": "fielding",
          "displayValue": "Fielding",
          "shortDisplayName": "Fielding"
        }
      ]
    }
  ],
  "teams": {
    "cincinnati-reds": {
      "id": "17",
      "uid": "s:1~l:10~t:17",
      "guid": "04b65a0b-3cca-d795-0e21-23606470418a",
      "name": "Reds",
      "displayName": "Cincinnati Reds",
      "abbreviation": "CIN",
      "slug": "cincinnati-reds",
      "location": "Cincinnati",
      "shortDisplayName": "Reds",
      "color": "c6011f"
    }
  },
  "glossary": [
    {
      "displayName": "Doubles",
      "abbreviation": "2B"
    },
    {
      "displayName": "Triples",
      "abbreviation": "3B"
    }
  ]
}
```
