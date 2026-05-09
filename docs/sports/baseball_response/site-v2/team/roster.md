# MLB Team Roster

## https://site.api.espn.com/apis/site/v2/sports/baseball/mlb/teams/17/roster

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "team": {
    "id": "17",
    "name": "Reds",
    "displayName": "Cincinnati Reds",
    "abbreviation": "CIN",
    "location": "Cincinnati",
    "clubhouse": "https://www.espn.com/mlb/team/_/name/cin/cincinnati-reds",
    "color": "c6011f",
    "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
    "recordSummary": "20-19",
    "seasonSummary": "2026"
  },
  "athletes": [
    {
      "items": [
        {
          "id": "4414528",
          "uid": "s:1~l:10~a:4414528",
          "guid": "3a81ab59-a326-31e7-9c98-0f9c09e26e18",
          "displayName": "Andrew Abbott",
          "slug": "andrew-abbott",
          "status": {
            "id": "1",
            "name": "Active",
            "abbreviation": "Active",
            "type": "active"
          },
          "alternateIds": {
            "sdr": "4414528"
          },
          "firstName": "Andrew",
          "lastName": "Abbott",
          "fullName": "Andrew Abbott"
        },
        {
          "id": "42355",
          "uid": "s:1~l:10~a:42355",
          "guid": "29463df8-ed05-3797-bb26-1007243007ae",
          "displayName": "Tejay Antone",
          "slug": "tejay-antone",
          "status": {
            "id": "1",
            "name": "Active",
            "abbreviation": "Active",
            "type": "active"
          },
          "alternateIds": {
            "sdr": "4602780"
          },
          "firstName": "Tejay",
          "lastName": "Antone",
          "fullName": "Tejay Antone"
        }
      ],
      "position": "Pitchers"
    },
    {
      "items": [
        {
          "id": "34975",
          "uid": "s:1~l:10~a:34975",
          "guid": "0ac057a9-c508-b8d4-7d33-70bac29964bd",
          "displayName": "Tyler Stephenson",
          "slug": "tyler-stephenson",
          "status": {
            "id": "1",
            "name": "Active",
            "abbreviation": "Active",
            "type": "active"
          },
          "alternateIds": {
            "sdr": "3901080"
          },
          "firstName": "Tyler",
          "lastName": "Stephenson",
          "fullName": "Tyler Stephenson"
        },
        {
          "id": "35268",
          "uid": "s:1~l:10~a:35268",
          "guid": "1726f184-b3fd-552d-09b1-e8efafb52265",
          "displayName": "Jose Trevino",
          "slug": "jose-trevino",
          "status": {
            "id": "1",
            "name": "Active",
            "abbreviation": "Active",
            "type": "active"
          },
          "alternateIds": {
            "sdr": "3983813"
          },
          "firstName": "Jose",
          "lastName": "Trevino",
          "fullName": "Jose Trevino"
        }
      ],
      "position": "Catchers"
    }
  ],
  "status": "success",
  "season": {
    "name": "Regular Season",
    "displayName": "2026",
    "year": 2026,
    "type": 2
  },
  "timestamp": "2026-05-09T03:57:46Z",
  "coach": [
    {
      "id": "2",
      "firstName": "Terry",
      "lastName": "Francona",
      "experience": 28
    }
  ]
}
```
