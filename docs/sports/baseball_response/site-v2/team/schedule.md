# MLB Team Schedule

## https://site.api.espn.com/apis/site/v2/sports/baseball/mlb/teams/17/schedule

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "events": [
    {
      "id": "401814689",
      "name": "Boston Red Sox at Cincinnati Reds",
      "competitions": [
        {
          "id": "401814689",
          "competitors": [
            {},
            {}
          ],
          "status": {
            "clock": 0.0,
            "displayClock": "0:00",
            "period": 9,
            "type": {},
            "featuredAthletes": [],
            "halfInning": 17,
            "periodPrefix": "End"
          },
          "date": "2026-03-26T20:10Z",
          "attendance": 43897,
          "type": {
            "id": "1",
            "abbreviation": "STD",
            "slug": "standard",
            "text": "Standard",
            "type": "standard"
          },
          "timeValid": true,
          "neutralSite": false,
          "boxscoreAvailable": true,
          "ticketsAvailable": false
        }
      ],
      "season": {
        "displayName": "2026",
        "year": 2026
      },
      "date": "2026-03-26T20:10Z",
      "shortName": "BOS @ CIN",
      "seasonType": {
        "id": "2",
        "name": "Regular Season",
        "abbreviation": "reg",
        "type": 2
      },
      "week": {
        "number": 13,
        "text": "Week 13"
      },
      "timeValid": true,
      "links": [
        {
          "language": "en-US",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/mlb/game/_/gameId/401814689/red-sox-reds",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=baseball&leagueAbbrev=mlb&gameId=401814689",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        }
      ]
    },
    {
      "id": "401814708",
      "name": "Boston Red Sox at Cincinnati Reds",
      "competitions": [
        {
          "id": "401814708",
          "competitors": [
            {},
            {}
          ],
          "status": {
            "clock": 0.0,
            "displayClock": "0:00",
            "period": 11,
            "type": {},
            "featuredAthletes": [],
            "halfInning": 21,
            "periodPrefix": "Bottom"
          },
          "date": "2026-03-28T20:10Z",
          "attendance": 38298,
          "type": {
            "id": "1",
            "abbreviation": "STD",
            "slug": "standard",
            "text": "Standard",
            "type": "standard"
          },
          "timeValid": true,
          "neutralSite": false,
          "boxscoreAvailable": true,
          "ticketsAvailable": false
        }
      ],
      "season": {
        "displayName": "2026",
        "year": 2026
      },
      "date": "2026-03-28T20:10Z",
      "shortName": "BOS @ CIN",
      "seasonType": {
        "id": "2",
        "name": "Regular Season",
        "abbreviation": "reg",
        "type": 2
      },
      "week": {
        "number": 13,
        "text": "Week 13"
      },
      "timeValid": true,
      "links": [
        {
          "language": "en-US",
          "rel": [
            "summary",
            "desktop"
          ],
          "href": "https://www.espn.com/mlb/game/_/gameId/401814708/red-sox-reds",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [
            "summary",
            "sportscenter"
          ],
          "href": "sportscenter://x-callback-url/showGame?sportName=baseball&leagueAbbrev=mlb&gameId=401814708",
          "text": "Gamecast",
          "shortText": "Summary",
          "isExternal": false,
          "isPremium": false
        }
      ]
    }
  ],
  "team": {
    "id": "17",
    "name": "Reds",
    "displayName": "Cincinnati Reds",
    "abbreviation": "CIN",
    "location": "Cincinnati",
    "clubhouse": "https://www.espn.com/mlb/team/_/name/cin/cincinnati-reds",
    "color": "c6011f",
    "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
    "recordSummary": "20-19",
    "seasonSummary": "2026"
  },
  "status": "success",
  "season": {
    "name": "Regular Season",
    "displayName": "2026",
    "year": 2026,
    "type": 2,
    "half": 1
  },
  "requestedSeason": {
    "name": "Regular Season",
    "displayName": "2026",
    "year": 2026,
    "type": 2
  },
  "timestamp": "2026-05-09T03:57:44Z",
  "allstarsgame": "2026-07-14T04:00Z"
}
```
