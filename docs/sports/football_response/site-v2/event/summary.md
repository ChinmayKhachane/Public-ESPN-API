# Summary

## https://site.api.espn.com/apis/site/v2/sports/football/{league}/summary?event={event}

Notes:
- Verified with `league=nfl`, `event=401772988` on 2026-05-08.
- This is the richest `site-v2` NFL game response and includes box score, drives, injuries, odds, pickcenter, and win probability.

## Example Response

```json
{
  "header": {
    "id": "401772988",
    "season": {
      "year": 2025,
      "type": 3
    },
    "competitions": [
      {
        "id": "401772988",
        "date": "2026-02-08T23:30Z"
      }
    ]
  },
  "boxscore": {
    "teams": [
      {
        "team": {
          "id": "26",
          "displayName": "Seattle Seahawks"
        },
        "homeAway": "away"
      },
      {
        "team": {
          "id": "17",
          "displayName": "New England Patriots"
        },
        "homeAway": "home"
      }
    ]
  },
  "scoringPlays": [
    {
      "text": "Jason Myers 33 Yd Field Goal"
    }
  ],
  "winprobability": [
    {
      "playId": "4017729881"
    }
  ],
  "leaders": [],
  "injuries": []
}
```
