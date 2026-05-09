# MLB Site Calendar

## https://site.api.espn.com/apis/site/v2/sports/baseball/mlb/scoreboard?limit=1

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "events": [
    {
      "id": "401815256",
      "uid": "s:1~l:10~e:401815256",
      "name": "Houston Astros at Cincinnati Reds",
      "competitions": [
        {
          "id": "401815256",
          "uid": "s:1~l:10~e:401815256~c:401815256",
          "competitors": [
            {},
            {}
          ],
          "status": {
            "clock": 0.0,
            "displayClock": "0:00",
            "period": 9,
            "type": {},
            "featuredAthletes": []
          },
          "date": "2026-05-08T22:10Z",
          "attendance": 24347,
          "type": {
            "id": "1",
            "abbreviation": "STD"
          },
          "timeValid": true,
          "neutralSite": false,
          "conferenceCompetition": false
        }
      ],
      "status": {
        "clock": 0.0,
        "displayClock": "0:00",
        "period": 9,
        "type": {
          "id": "3",
          "name": "STATUS_FINAL",
          "state": "post",
          "completed": true,
          "detail": "Final",
          "shortDetail": "Final"
        }
      },
      "season": {
        "slug": "regular-season",
        "year": 2026,
        "type": 2
      },
      "date": "2026-05-08T22:10Z",
      "shortName": "HOU @ CIN",
      "links": [
        {
          "language": "en-US",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/mlb/game/_/gameId/401815256/astros-reds",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "boxscore",
            "desktop"
          ],
          "href": "https://www.espn.com/mlb/boxscore/_/gameId/401815256",
          "text": "Box Score",
          "shortText": "Box Score",
          "isExternal": false,
          "isPremium": false
        }
      ]
    }
  ],
  "season": {
    "type": 2,
    "year": 2026
  },
  "leagues": [
    {
      "id": "10",
      "uid": "s:1~l:10",
      "name": "Major League Baseball",
      "abbreviation": "MLB",
      "slug": "mlb",
      "season": {
        "displayName": "2026",
        "year": 2026,
        "startDate": "2026-02-19T08:00Z",
        "endDate": "2026-11-12T07:59Z",
        "type": {
          "id": "2",
          "name": "Regular Season",
          "abbreviation": "reg",
          "type": 2
        }
      },
      "midsizeName": "MLB",
      "logos": [
        {
          "href": "https://a.espncdn.com/i/teamlogos/leagues/500/mlb.png",
          "width": 500,
          "height": 500,
          "alt": "",
          "rel": [
            "full",
            "default"
          ],
          "lastUpdated": "2023-03-29T12:34Z"
        },
        {
          "href": "https://a.espncdn.com/combiner/i?img=/i/teamlogos/leagues/500-dark/mlb.png&w=500&h=500&transparent=true",
          "width": 500,
          "height": 500,
          "alt": "",
          "rel": [
            "full",
            "dark"
          ],
          "lastUpdated": "2026-05-08T20:30Z"
        }
      ],
      "calendarType": "day",
      "calendarIsWhitelist": false
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
        "rel": [
          "light"
        ]
      },
      {
        "href": "https://a.espncdn.com/i/betting/Draftkings_Dark.svg",
        "rel": [
          "dark"
        ]
      }
    ]
  }
}
```
