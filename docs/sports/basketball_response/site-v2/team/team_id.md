# Team Detail

## https://site.api.espn.com/apis/site/v2/sports/basketball/{league}/teams/{id}

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "team": {
    "id": "20",
    "uid": "s:40~l:46~t:20",
    "name": "76ers",
    "displayName": "Philadelphia 76ers",
    "shortDisplayName": "76ers",
    "abbreviation": "PHI",
    "slug": "philadelphia-76ers",
    "record": {
      "items": [
        {},
        {}
      ]
    },
    "groups": {
      "id": "1",
      "parent": {
        "id": "5"
      },
      "isConference": false
    },
    "location": "Philadelphia",
    "color": "1d428a",
    "alternateColor": "e01234"
  }
}
```
