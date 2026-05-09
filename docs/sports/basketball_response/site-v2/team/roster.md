# Team Roster

## https://site.api.espn.com/apis/site/v2/sports/basketball/{league}/teams/{id}/roster

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
  "timestamp": "2026-05-09T03:15:52Z",
  "team": {
    "id": "20",
    "name": "76ers",
    "displayName": "Philadelphia 76ers",
    "abbreviation": "PHI",
    "location": "Philadelphia",
    "clubhouse": "https://www.espn.com/nba/team/_/name/phi/philadelphia-76ers",
    "color": "1d428a",
    "logo": "https://a.espncdn.com/i/teamlogos/nba/500/phi.png",
    "recordSummary": "45-37",
    "seasonSummary": "2025-26",
    "standingSummary": "4th in Atlantic Division"
  },
  "athletes": [
    {
      "id": "4870562",
      "uid": "s:40~l:46~a:4870562",
      "guid": "b6f6857f-9a83-30a6-85a7-2e02c49d5ae5",
      "displayName": "Dominick Barlow",
      "slug": "dominick-barlow",
      "status": {
        "id": "1",
        "name": "Active",
        "abbreviation": "Active",
        "type": "active"
      },
      "teams": [
        {}
      ],
      "injuries": [],
      "alternateIds": {
        "sdr": "4870562"
      },
      "firstName": "Dominick",
      "lastName": "Barlow",
      "fullName": "Dominick Barlow"
    },
    {
      "id": "4432179",
      "uid": "s:40~l:46~a:4432179",
      "guid": "39f93ebb-634b-395c-8eff-c773f4b54a7c",
      "displayName": "MarJon Beauchamp",
      "slug": "marjon-beauchamp",
      "status": {
        "id": "1",
        "name": "Active",
        "abbreviation": "Active",
        "type": "active"
      },
      "teams": [
        {}
      ],
      "injuries": [],
      "alternateIds": {
        "sdr": "4432179"
      },
      "firstName": "MarJon",
      "lastName": "Beauchamp",
      "fullName": "MarJon Beauchamp"
    }
  ],
  "coach": [
    {
      "id": "52085",
      "firstName": "Nick",
      "lastName": "Nurse",
      "experience": 6
    }
  ]
}
```
