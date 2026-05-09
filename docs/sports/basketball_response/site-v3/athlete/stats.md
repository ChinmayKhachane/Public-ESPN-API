# Athlete Stats

## https://site.web.api.espn.com/apis/common/v3/sports/basketball/{league}/athletes/{id}/stats

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "teams": {
    "philadelphia-76ers": {
      "id": "20",
      "uid": "s:40~l:46~t:20",
      "guid": "ca1685ed-b799-53e4-7924-e58ea6eb8f3a",
      "name": "76ers",
      "displayName": "Philadelphia 76ers",
      "shortDisplayName": "76ers",
      "abbreviation": "PHI",
      "slug": "philadelphia-76ers",
      "record": {},
      "venue": {
        "id": "1845",
        "fullName": "Xfinity Mobile Arena",
        "shortName": "Wells Fargo Center",
        "address": {},
        "grass": false,
        "indoor": true,
        "images": []
      },
      "groups": {},
      "location": "Philadelphia"
    }
  },
  "categories": [
    {
      "name": "averages",
      "displayName": "Regular Season Averages",
      "statistics": [
        {},
        {}
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
        "The number of games started by an athlete."
      ],
      "totals": [
        "490",
        "490"
      ],
      "sortKey": "averages"
    },
    {
      "name": "totals",
      "displayName": "Regular Season Totals",
      "statistics": [
        {},
        {}
      ],
      "labels": [
        "FG",
        "FG%"
      ],
      "names": [
        "fieldGoalsMade-fieldGoalsAttempted",
        "fieldGoalPct"
      ],
      "displayNames": [
        "Field Goals Made-Attempted",
        "Field Goal Percentage"
      ],
      "descriptions": [
        "The number of times a 2pt field goal was made.-The number of times a 2pt field goal was attempted.",
        "The ratio of field goals made to field goals attempted: FGM / FGA"
      ],
      "totals": [
        "4458-8909",
        "50.0"
      ],
      "sortKey": "totals"
    }
  ],
  "glossary": [
    {
      "displayName": "3-Point Field Goal Percentage",
      "abbreviation": "3P%"
    },
    {
      "displayName": "3-Point Field Goals Made-Attempted Per Game",
      "abbreviation": "3PT"
    }
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
      "name": "seasontype",
      "displayName": "Season",
      "value": "2",
      "options": [
        {},
        {}
      ]
    }
  ]
}
```
