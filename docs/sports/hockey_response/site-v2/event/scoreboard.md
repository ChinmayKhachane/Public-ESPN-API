# Scoreboard

## https://site.api.espn.com/apis/site/v2/sports/hockey/{league}/scoreboard

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/scoreboard?dates=20260507`

## Example Response

```json
{
  "leagues": [
    {
      "id": "90",
      "uid": "s:70~l:90",
      "name": "National Hockey League",
      "abbreviation": "NHL",
      "slug": "nhl",
      "season": {
        "year": 2026,
        "startDate": "2025-09-20T07:00Z",
        "endDate": "2026-07-01T06:59Z",
        "displayName": "2025-26",
        "type": {
          "id": "3",
          "type": 3,
          "name": "Postseason",
          "abbreviation": "post"
        }
      },
      "logos": [
        {
          "href": "https://a.espncdn.com/i/teamlogos/leagues/500/nhl.png",
          "width": 500,
          "height": 500,
          "alt": "",
          "rel": [
            "full",
            "default"
          ],
          "lastUpdated": "2018-06-05T12:07Z"
        },
        {
          "href": "https://a.espncdn.com/i/teamlogos/leagues/500-dark/nhl.png",
          "width": 500,
          "height": 500,
          "alt": "",
          "rel": [
            "full",
            "dark"
          ],
          "lastUpdated": "2021-08-10T20:37Z"
        }
      ],
      "calendarType": "day",
      "calendarIsWhitelist": true,
      "calendarStartDate": "2025-09-20T07:00Z",
      "calendarEndDate": "2026-07-01T06:59Z",
      "calendar": [
        "2025-09-20T07:00Z",
        "2025-09-21T07:00Z"
      ]
    }
  ],
  "events": [
    {
      "id": "401871412",
      "uid": "s:70~l:90~e:401871412",
      "date": "2026-05-08T00:00Z",
      "name": "Carolina Hurricanes at Philadelphia Flyers",
      "shortName": "CAR @ PHI",
      "season": {
        "year": 2026,
        "type": 3,
        "slug": "post-season"
      },
      "competitions": [
        {
          "id": "401871412",
          "uid": "s:70~l:90~e:401871412~c:401871412",
          "date": "2026-05-08T00:00Z",
          "attendance": 19970,
          "type": {
            "id": "15",
            "abbreviation": "QTR"
          },
          "timeValid": true,
          "neutralSite": false,
          "playByPlayAvailable": true,
          "recent": false,
          "venue": {
            "id": "1845",
            "fullName": "Xfinity Mobile Arena",
            "address": {},
            "indoor": true
          },
          "competitors": [
            {},
            {}
          ],
          "notes": [
            {}
          ],
          "status": {
            "clock": 0.0,
            "displayClock": "0:00",
            "period": 3,
            "type": {},
            "featuredAthletes": []
          },
          "broadcasts": [
            {}
          ],
          "format": {
            "regulation": {}
          },
          "startDate": "2026-05-08T00:00Z",
          "series": {
            "type": "playoff",
            "title": "Playoff Series",
            "summary": "CAR leads series 3-0",
            "completed": false,
            "totalCompetitions": 7,
            "competitors": []
          },
          "broadcast": "TNT/truTV/HBO Max"
        }
      ],
      "links": [
        {
          "language": "en-US",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/game/_/gameId/401871412/hurricanes-flyers",
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
          "href": "https://www.espn.com/nhl/boxscore/_/gameId/401871412",
          "text": "Box Score",
          "shortText": "Box Score",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "status": {
        "clock": 0.0,
        "displayClock": "0:00",
        "period": 3,
        "type": {
          "id": "3",
          "name": "STATUS_FINAL",
          "state": "post",
          "completed": true,
          "description": "Final",
          "detail": "Final",
          "shortDetail": "Final"
        }
      }
    }
  ],
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
