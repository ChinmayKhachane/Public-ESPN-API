# League Statistics

## https://site.api.espn.com/apis/site/v2/sports/basketball/{league}/statistics

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
    "type": 3
  },
  "status": "success",
  "timestamp": "2026-05-09T03:15:31Z",
  "league": {
    "id": "46",
    "name": "National Basketball Association",
    "abbreviation": "NBA",
    "slug": "nba",
    "shortName": "NBA",
    "isTournament": false,
    "links": [
      {
        "text": "Index",
        "shortText": "Index",
        "language": "en-US",
        "rel": [],
        "href": "https://www.espn.com/nba/",
        "isExternal": false,
        "isPremium": false
      },
      {
        "text": "Index",
        "shortText": "Index",
        "language": "en-US",
        "rel": [],
        "href": "sportscenter://x-callback-url/showClubhouse?uid=s:40~l:46",
        "isExternal": false,
        "isPremium": false
      }
    ],
    "logos": [
      {
        "href": "https://a.espncdn.com/i/teamlogos/leagues/500/nba.png",
        "width": 500,
        "height": 500,
        "alt": "",
        "rel": [],
        "lastUpdated": "2018-06-05T12:07Z"
      },
      {
        "href": "https://a.espncdn.com/combiner/i?img=/i/teamlogos/leagues/500-dark/nba.png&w=500&h=500&transparent=true",
        "width": 500,
        "height": 500,
        "alt": "",
        "rel": [],
        "lastUpdated": "2026-05-08T05:17Z"
      }
    ]
  }
}
```
