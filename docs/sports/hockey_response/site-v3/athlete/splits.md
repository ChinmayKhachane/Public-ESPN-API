# Athlete Splits

## https://site.web.api.espn.com/apis/common/v3/sports/hockey/{league}/athletes/{athlete}/splits

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.web.api.espn.com/apis/common/v3/sports/hockey/nhl/athletes/4565230/splits`

## Example Response

```json
{
  "filters": [
    {
      "displayName": "Season",
      "name": "season",
      "value": "2026",
      "options": [
        {
          "value": "2026",
          "displayValue": "2025-26"
        },
        {
          "value": "2025",
          "displayValue": "2024-25"
        }
      ]
    }
  ],
  "displayName": "2025-26 Splits",
  "labels": [
    "GP",
    "G"
  ],
  "names": [
    "games",
    "goals"
  ],
  "displayNames": [
    "Games Played",
    "Goals"
  ],
  "descriptions": [
    "Total games played.",
    "Total goals scored."
  ],
  "splitCategories": [
    {
      "name": "split",
      "displayName": "split",
      "splits": [
        {
          "displayName": "All Splits",
          "stats": [
            "81",
            "26"
          ],
          "abbreviation": "Total"
        },
        {
          "displayName": "Home",
          "stats": [
            "40",
            "15"
          ],
          "abbreviation": "Home"
        }
      ]
    },
    {
      "name": "byOpponent",
      "displayName": "Opponent",
      "splits": [
        {
          "displayName": "Boston Bruins",
          "stats": [
            "3",
            "0"
          ],
          "abbreviation": "Boston Bruins"
        },
        {
          "displayName": "vs Buffalo Sabres",
          "stats": [
            "3",
            "2"
          ],
          "abbreviation": "vs Buffalo Sabres"
        }
      ]
    }
  ]
}
```
