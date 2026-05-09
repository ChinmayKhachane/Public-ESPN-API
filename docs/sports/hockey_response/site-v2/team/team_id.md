# Team Detail

## https://site.api.espn.com/apis/site/v2/sports/hockey/{league}/teams/{team}

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/teams/15`

## Example Response

```json
{
  "team": {
    "id": "15",
    "uid": "s:70~l:90~t:15",
    "slug": "philadelphia-flyers",
    "location": "Philadelphia",
    "name": "Flyers",
    "nickname": "Flyers",
    "abbreviation": "PHI",
    "displayName": "Philadelphia Flyers",
    "shortDisplayName": "Flyers",
    "color": "fe5823",
    "alternateColor": "000000",
    "isActive": true,
    "logos": [
      {
        "href": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
        "width": 500,
        "height": 500,
        "alt": "",
        "rel": [
          "full",
          "default"
        ],
        "lastUpdated": "2024-07-02T15:51Z"
      },
      {
        "href": "https://a.espncdn.com/i/teamlogos/nhl/500-dark/phi.png",
        "width": 500,
        "height": 500,
        "alt": "",
        "rel": [
          "full",
          "dark"
        ],
        "lastUpdated": "2024-07-02T15:46Z"
      }
    ],
    "record": {
      "items": [
        {
          "description": "Overall Record",
          "type": "total",
          "summary": "43-27-12",
          "stats": []
        },
        {
          "description": "Home Record",
          "type": "home",
          "summary": "20-13-8",
          "stats": []
        }
      ]
    }
  }
}
```
