# Team Detail

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/seasons/{season}/teams/{team}

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/teams/15`

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/teams/15?lang=en&region=us",
  "id": "15",
  "guid": "68aba012-4e93-9371-6861-1bb9a63cfb11",
  "uid": "s:70~l:90~t:15",
  "alternateIds": {
    "sdr": "2617"
  },
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
  "isAllStar": false,
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
    "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/types/2/teams/15/record?lang=en&region=us"
  }
}
```
