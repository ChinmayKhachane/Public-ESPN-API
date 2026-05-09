# Venue Detail

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/venues/{venue}

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/venues/1845`

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/venues/1845?lang=en&region=us",
  "id": "1845",
  "fullName": "Xfinity Mobile Arena",
  "address": {
    "city": "Philadelphia",
    "state": "PA",
    "country": "USA"
  },
  "grass": false,
  "indoor": true,
  "images": [
    {
      "href": "https://a.espncdn.com/i/venues/nhl/day/1845.jpg",
      "width": 2000,
      "height": 1125,
      "alt": "",
      "rel": [
        "full",
        "day"
      ]
    }
  ]
}
```
