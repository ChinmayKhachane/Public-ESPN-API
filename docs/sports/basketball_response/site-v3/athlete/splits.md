# Athlete Splits

## https://site.web.api.espn.com/apis/common/v3/sports/basketball/{league}/athletes/{id}/splits

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "displayName": "2025-26 Splits",
  "labels": [
    "GP",
    "MIN"
  ],
  "names": [
    "gamesPlayed",
    "avgMinutes"
  ],
  "displayNames": [
    "Games Played",
    "Minutes Per Game"
  ],
  "descriptions": [
    "Games Played",
    "The average number of minutes per game."
  ],
  "filters": [
    {
      "name": "league",
      "displayName": "League",
      "value": "nba",
      "options": [
        {},
        {}
      ]
    },
    {
      "name": "season",
      "displayName": "Season",
      "value": "2026",
      "options": [
        {},
        {}
      ]
    }
  ],
  "splitCategories": [
    {
      "name": "split",
      "displayName": "split",
      "splits": [
        {},
        {}
      ]
    },
    {
      "name": "byMonth",
      "displayName": "Month",
      "splits": [
        {},
        {}
      ]
    }
  ]
}
```
