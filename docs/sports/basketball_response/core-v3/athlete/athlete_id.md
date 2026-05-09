# Core V3 Athlete Detail

## https://sports.core.api.espn.com/v3/sports/basketball/{league}/athletes/{id}

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.
- `enable=team,position,links` produced the richer tested shape.

## Example Response

```json
{
  "id": "3059318",
  "uid": "s:40~l:46~a:3059318",
  "guid": "89410119-0011-5ba0-efed-7b0b6c756c23",
  "displayName": "Joel Embiid",
  "team": {
    "id": "20",
    "uid": "s:40~l:46~t:20",
    "guid": "ca1685ed-b799-53e4-7924-e58ea6eb8f3a",
    "name": "76ers",
    "displayName": "Philadelphia 76ers",
    "shortDisplayName": "76ers",
    "abbreviation": "PHI",
    "slug": "philadelphia-76ers",
    "location": "Philadelphia",
    "nickname": "Philadelphia",
    "color": "1d428a",
    "alternateColor": "e01234"
  },
  "firstName": "Joel",
  "lastName": "Embiid",
  "fullName": "Joel Embiid",
  "shortName": "J. Embiid",
  "weight": 280.0,
  "displayWeight": "280 lbs",
  "height": 84.0
}
```
