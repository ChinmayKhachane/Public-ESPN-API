# Team Detail

## https://sports.core.api.espn.com/v3/sports/football/{league}/teams/{id}

Notes:
- Verified with `league=nfl`, `id=6` on 2026-05-07.
- The default response contains identity and color fields.
- Add `enable=logos,links,groups,venue,record` for a richer team response.

## Example Response

```json
{
  "id": "6",
  "uid": "s:20~l:28~t:6",
  "slug": "dallas-cowboys",
  "displayName": "Dallas Cowboys",
  "abbreviation": "DAL",
  "color": "002a5c",
  "alternateColor": "b0b7bc",
  "active": true,
  "logos": [
    {
      "href": "https://a.espncdn.com/i/teamlogos/nfl/500/dal.png",
      "rel": ["full", "default"]
    },
    {
      "href": "https://a.espncdn.com/i/teamlogos/nfl/500-dark/dal.png",
      "rel": ["full", "dark"]
    }
  ],
  "venue": {
    "id": "3687",
    "fullName": "AT&T Stadium",
    "address": {
      "city": "Arlington",
      "state": "TX",
      "zipCode": "76011"
    },
    "grass": false,
    "indoor": true
  },
  "record": {
    "total": {
      "stats": {
        "wins": {
          "value": 7.0
        },
        "losses": {
          "value": 9.0
        },
        "ties": {
          "value": 1.0
        }
      }
    }
  }
}
```
