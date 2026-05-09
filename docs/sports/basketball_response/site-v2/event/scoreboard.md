# Scoreboard

## https://site.api.espn.com/apis/site/v2/sports/basketball/{league}/scoreboard

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "season": {
    "year": 2026,
    "type": 3
  },
  "events": [
    {
      "id": "401871161",
      "uid": "s:40~l:46~e:401871161",
      "name": "New York Knicks at Philadelphia 76ers",
      "season": {
        "slug": "post-season",
        "year": 2026,
        "type": 3
      },
      "status": {
        "type": {},
        "period": 4,
        "clock": 720.0,
        "displayClock": "12:00"
      },
      "competitions": [
        {}
      ],
      "date": "2026-05-08T23:00Z",
      "shortName": "NY @ PHI",
      "links": [
        {},
        {}
      ]
    },
    {
      "id": "401871154",
      "uid": "s:40~l:46~e:401871154",
      "name": "San Antonio Spurs at Minnesota Timberwolves",
      "season": {
        "slug": "post-season",
        "year": 2026,
        "type": 3
      },
      "status": {
        "type": {},
        "period": 2,
        "clock": 0.0,
        "displayClock": "0.0"
      },
      "competitions": [
        {}
      ],
      "date": "2026-05-09T01:30Z",
      "shortName": "SA @ MIN",
      "links": [
        {},
        {}
      ]
    }
  ],
  "leagues": [
    {
      "id": "46",
      "uid": "s:40~l:46",
      "name": "National Basketball Association",
      "abbreviation": "NBA",
      "slug": "nba",
      "season": {
        "displayName": "2025-26",
        "year": 2026,
        "type": {},
        "startDate": "2025-10-01T07:00Z",
        "endDate": "2026-06-27T06:59Z"
      },
      "logos": [
        {},
        {}
      ],
      "calendarType": "day",
      "calendarIsWhitelist": true,
      "calendarStartDate": "2025-10-01T07:00Z",
      "calendarEndDate": "2026-06-27T06:59Z",
      "calendar": [
        "2025-10-02T07:00Z",
        "2025-10-03T07:00Z"
      ]
    }
  ],
  "day": {
    "date": "2026-05-08"
  },
  "provider": {
    "id": "100",
    "name": "Draft Kings",
    "displayName": "Draft Kings",
    "priority": 1,
    "logos": [
      {
        "href": "https://a.espncdn.com/i/betting/Draftkings_Light.svg",
        "rel": []
      },
      {
        "href": "https://a.espncdn.com/i/betting/Draftkings_Dark.svg",
        "rel": []
      }
    ]
  }
}
```
