# Team Schedule

## https://site.api.espn.com/apis/site/v2/sports/hockey/{league}/teams/{team}/schedule

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/teams/15/schedule?season=2026`

## Example Response

```json
{
  "timestamp": "2026-05-09T17:01:02Z",
  "status": "success",
  "season": {
    "year": 2026,
    "type": 3,
    "name": "Postseason",
    "displayName": "2025-26",
    "half": 1
  },
  "team": {
    "id": "15",
    "abbreviation": "PHI",
    "location": "Philadelphia",
    "name": "Flyers",
    "displayName": "Philadelphia Flyers",
    "clubhouse": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
    "color": "fe5823",
    "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
    "recordSummary": "43-27-12",
    "seasonSummary": "2025-26",
    "standingSummary": "2nd in Metropolitan Division",
    "groups": {
      "id": "33",
      "parent": {
        "id": "7"
      },
      "isConference": false
    }
  },
  "events": [
    {
      "id": "401869717",
      "date": "2026-04-19T00:00Z",
      "name": "Philadelphia Flyers at Pittsburgh Penguins",
      "shortName": "PHI @ PIT",
      "season": {
        "year": 2026,
        "displayName": "2025-26"
      },
      "seasonType": {
        "id": "3",
        "type": 3,
        "name": "Postseason",
        "abbreviation": "post"
      },
      "timeValid": true,
      "competitions": [
        {
          "id": "401869717",
          "date": "2026-04-19T00:00Z",
          "attendance": 18346,
          "type": {},
          "timeValid": true,
          "neutralSite": false,
          "boxscoreAvailable": true,
          "ticketsAvailable": false,
          "venue": {},
          "competitors": [],
          "notes": [],
          "broadcasts": [],
          "status": {}
        }
      ],
      "links": [
        {
          "language": "en-US",
          "rel": [],
          "href": "https://www.espn.com/nhl/game/_/gameId/401869717/flyers-penguins",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [],
          "href": "sportscenter://x-callback-url/showGame?sportName=hockey&leagueAbbrev=nhl&gameId=401869717",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        }
      ]
    },
    {
      "id": "401869794",
      "date": "2026-04-20T23:00Z",
      "name": "Philadelphia Flyers at Pittsburgh Penguins",
      "shortName": "PHI @ PIT",
      "season": {
        "year": 2026,
        "displayName": "2025-26"
      },
      "seasonType": {
        "id": "3",
        "type": 3,
        "name": "Postseason",
        "abbreviation": "post"
      },
      "timeValid": true,
      "competitions": [
        {
          "id": "401869794",
          "date": "2026-04-20T23:00Z",
          "attendance": 18308,
          "type": {},
          "timeValid": true,
          "neutralSite": false,
          "boxscoreAvailable": true,
          "ticketsAvailable": false,
          "venue": {},
          "competitors": [],
          "notes": [],
          "broadcasts": [],
          "status": {}
        }
      ],
      "links": [
        {
          "language": "en-US",
          "rel": [],
          "href": "https://www.espn.com/nhl/game/_/gameId/401869794/flyers-penguins",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [],
          "href": "sportscenter://x-callback-url/showGame?sportName=hockey&leagueAbbrev=nhl&gameId=401869794",
          "text": "Gamecast",
          "shortText": "Gamecast",
          "isExternal": false,
          "isPremium": false
        }
      ]
    }
  ],
  "requestedSeason": {
    "year": 2026,
    "type": 3,
    "name": "Postseason",
    "displayName": "2025-26"
  }
}
```
