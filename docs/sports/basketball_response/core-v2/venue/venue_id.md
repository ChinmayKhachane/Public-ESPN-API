# Venue Detail

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/venues/{id}

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/venues/1845?lang=en&region=us",
  "id": "1845",
  "guid": "5c91d4ea-46af-3b85-a87d-54422ed1e8c3",
  "fullName": "Xfinity Mobile Arena",
  "shortName": "Wells Fargo Center",
  "address": {
    "city": "Philadelphia",
    "state": "PA"
  },
  "grass": false,
  "indoor": true,
  "images": [
    {
      "href": "https://a.espncdn.com/i/venues/nba/day/1845.jpg",
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
