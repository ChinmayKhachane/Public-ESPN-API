# Core V3 Athletes

## https://sports.core.api.espn.com/v3/sports/basketball/{league}/athletes

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "count": 844,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 422,
  "items": [
    {
      "id": "4432932",
      "uid": "s:40~l:46~a:4432932",
      "guid": "8dc1be8c-8e64-3b3a-bc7c-075d4154efd0",
      "displayName": "Max Abmas",
      "firstName": "Max",
      "lastName": "Abmas",
      "fullName": "Max Abmas",
      "shortName": "M. Abmas",
      "weight": 175.0,
      "displayWeight": "175 lbs",
      "height": 72.0,
      "displayHeight": "6' 0\""
    },
    {
      "id": "4431679",
      "uid": "s:40~l:46~a:4431679",
      "guid": "79540e31-b541-3156-9ed2-fea920e18b17",
      "displayName": "Precious Achiuwa",
      "firstName": "Precious",
      "lastName": "Achiuwa",
      "fullName": "Precious Achiuwa",
      "shortName": "P. Achiuwa",
      "weight": 243.0,
      "displayWeight": "243 lbs",
      "height": 80.0,
      "displayHeight": "6' 8\""
    }
  ]
}
```
