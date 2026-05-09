# MLB Common v3 Team Roster

## https://site.web.api.espn.com/apis/common/v3/sports/baseball/mlb/teams/17/roster

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
    "venueLink": null,
    "clubhouse": "https://www.espn.com/mlb/team/_/name/cin/cincinnati-reds",
    "color": "c6011f",
    "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
    "recordSummary": "20-19"
  },
  "season": {
    "name": "Regular Season",
    "year": 2026,
    "type": 2
  },
  "coach": [
    {
      "id": "2",
      "firstName": "Terry",
      "lastName": "Francona",
      "experience": 28
    }
  ],
  "positionGroups": [
    {
      "displayName": "Pitchers",
      "athletes": [
        {
          "id": "4414528",
          "guid": "3a81ab59-a326-31e7-9c98-0f9c09e26e18",
          "displayName": "Andrew Abbott",
          "statistics": {
            "splits": {}
          },
          "status": {
            "id": "1",
            "name": "Active",
            "abbreviation": "Active",
            "type": "active"
          },
          "alternateIds": [
            {}
          ],
          "firstName": "Andrew",
          "lastName": "Abbott",
          "fullName": "Andrew Abbott",
          "shortName": "A. Abbott"
        },
        {
          "id": "4842989",
          "guid": "cb37685e-a355-3778-a9c1-bd48609f63b1",
          "displayName": "Kevin Abel",
          "statistics": {
            "splits": {}
          },
          "status": {
            "id": "4",
            "name": "Minors",
            "abbreviation": "Minors",
            "type": "minors"
          },
          "alternateIds": [
            {}
          ],
          "firstName": "Kevin",
          "lastName": "Abel",
          "fullName": "Kevin Abel",
          "shortName": "K. Abel"
        }
      ],
      "type": "pitchers"
    },
    {
      "displayName": "Infielders",
      "athletes": [
        {
          "id": "5016789",
          "guid": "b94c40c8-dd18-3171-a972-884968726201",
          "displayName": "Victor Acosta",
          "statistics": {
            "splits": {}
          },
          "status": {
            "id": "4",
            "name": "Minors",
            "abbreviation": "Minors",
            "type": "minors"
          },
          "alternateIds": [
            {}
          ],
          "firstName": "Victor",
          "lastName": "Acosta",
          "fullName": "Victor Acosta",
          "shortName": "V. Acosta"
        },
        {
          "id": "5274730",
          "guid": "f726bbbb-5c17-310e-a6bf-9e676714c243",
          "displayName": "Alfredo Alcantara",
          "statistics": {
            "splits": {}
          },
          "status": {
            "id": "4",
            "name": "Minors",
            "abbreviation": "Minors",
            "type": "minors"
          },
          "alternateIds": [
            {}
          ],
          "firstName": "Alfredo",
          "lastName": "Alcantara",
          "fullName": "Alfredo Alcantara",
          "shortName": "A. Alcantara"
        }
      ],
      "type": "infielders"
    }
  ]
}
```
