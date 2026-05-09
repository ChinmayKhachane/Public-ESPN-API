# League Statistics

## https://site.api.espn.com/apis/site/v2/sports/hockey/{league}/statistics

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/statistics`

## Example Response

```json
{
  "timestamp": "2026-05-09T17:01:04Z",
  "status": "success",
  "season": {
    "year": 2026,
    "displayName": "2025-26",
    "type": 3,
    "name": "Postseason"
  },
  "league": {
    "id": "90",
    "name": "National Hockey League",
    "abbreviation": "NHL",
    "shortName": "NHL",
    "slug": "nhl",
    "isTournament": false,
    "links": [
      {
        "language": "en-US",
        "rel": [
          "index",
          "desktop"
        ],
        "href": "https://www.espn.com/nhl/",
        "text": "Index",
        "shortText": "Index",
        "isExternal": false,
        "isPremium": false
      },
      {
        "language": "en-US",
        "rel": [
          "index",
          "sportscenter"
        ],
        "href": "sportscenter://x-callback-url/showClubhouse?uid=s:70~l:90",
        "text": "Index",
        "shortText": "Index",
        "isExternal": false,
        "isPremium": false
      }
    ],
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
    ]
  },
  "stats": {
    "id": "0",
    "name": "Season",
    "abbreviation": "Season",
    "categories": [
      {
        "name": "goals",
        "displayName": "Goals",
        "shortDisplayName": "G",
        "abbreviation": "G",
        "leaders": [
          {},
          {}
        ]
      },
      {
        "name": "assists",
        "displayName": "Assists",
        "shortDisplayName": "A",
        "abbreviation": "A",
        "leaders": [
          {},
          {}
        ]
      }
    ]
  }
}
```
