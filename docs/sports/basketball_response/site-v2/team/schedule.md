# Team Schedule

## https://site.api.espn.com/apis/site/v2/sports/basketball/{league}/teams/{id}/schedule

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "season": {
    "name": "Postseason",
    "displayName": "2025-26",
    "year": 2026,
    "type": 3,
    "half": 1
  },
  "status": "success",
  "timestamp": "2026-05-09T03:15:58Z",
  "events": [
    {
      "id": "401869191",
      "name": "Philadelphia 76ers at Boston Celtics",
      "season": {
        "displayName": "2025-26",
        "year": 2026
      },
      "competitions": [
        {}
      ],
      "date": "2026-04-19T17:00Z",
      "timeValid": true,
      "shortName": "PHI @ BOS",
      "seasonType": {
        "id": "3",
        "name": "Postseason",
        "abbreviation": "post",
        "type": 3
      },
      "links": [
        {},
        {}
      ]
    },
    {
      "id": "401869396",
      "name": "Philadelphia 76ers at Boston Celtics",
      "season": {
        "displayName": "2025-26",
        "year": 2026
      },
      "competitions": [
        {}
      ],
      "date": "2026-04-21T23:00Z",
      "timeValid": true,
      "shortName": "PHI @ BOS",
      "seasonType": {
        "id": "3",
        "name": "Postseason",
        "abbreviation": "post",
        "type": 3
      },
      "links": [
        {},
        {}
      ]
    }
  ],
  "team": {
    "id": "20",
    "name": "76ers",
    "displayName": "Philadelphia 76ers",
    "abbreviation": "PHI",
    "groups": {
      "id": "1",
      "parent": {
        "id": "5"
      },
      "isConference": false
    },
    "location": "Philadelphia",
    "clubhouse": "https://www.espn.com/nba/team/_/name/phi/philadelphia-76ers",
    "color": "1d428a",
    "logo": "https://a.espncdn.com/i/teamlogos/nba/500/phi.png",
    "recordSummary": "45-37",
    "seasonSummary": "2025-26",
    "standingSummary": "4th in Atlantic Division"
  },
  "requestedSeason": {
    "name": "Postseason",
    "displayName": "2025-26",
    "year": 2026,
    "type": 3
  }
}
```
