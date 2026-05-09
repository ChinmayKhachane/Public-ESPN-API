# Team Roster

## https://site.api.espn.com/apis/site/v2/sports/hockey/{league}/teams/{team}/roster

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/teams/15/roster`

## Example Response

```json
{
  "timestamp": "2026-05-09T17:03:04Z",
  "status": "success",
  "season": {
    "year": 2026,
    "displayName": "2025-26",
    "type": 3,
    "name": "Postseason"
  },
  "athletes": [
    {
      "position": "Centers",
      "items": [
        {
          "id": "4064976",
          "uid": "s:70~l:90~a:4064976",
          "guid": "f0f7c3a7-6e8a-47fa-b290-96724da24d8a",
          "alternateIds": {
            "sdr": "4064976"
          },
          "alternateId": "4064976",
          "firstName": "Rodrigo",
          "lastName": "Abols",
          "fullName": "Rodrigo Abols",
          "displayName": "Rodrigo Abols",
          "shortName": "R. Abols",
          "weight": 206.0,
          "displayWeight": "206 lbs",
          "height": 76.0,
          "displayHeight": "6' 4\"",
          "age": 30,
          "dateOfBirth": "1996-01-05T08:00Z",
          "debutYear": 2016,
          "links": [
            {},
            {}
          ]
        },
        {
          "id": "5173105",
          "uid": "s:70~l:90~a:5173105",
          "guid": "6516db75-a54c-3318-85a3-ecee5c738731",
          "alternateIds": {
            "sdr": "5173105"
          },
          "alternateId": "5173105",
          "firstName": "Denver",
          "lastName": "Barkey",
          "fullName": "Denver Barkey",
          "displayName": "Denver Barkey",
          "shortName": "D. Barkey",
          "weight": 155.0,
          "displayWeight": "155 lbs",
          "height": 69.0,
          "displayHeight": "5' 9\"",
          "age": 21,
          "dateOfBirth": "2005-04-27T07:00Z",
          "links": [
            {},
            {}
          ],
          "birthPlace": {
            "city": "Newmarket",
            "state": "ON",
            "country": "CAN",
            "displayText": "Newmarket, CAN"
          }
        }
      ]
    },
    {
      "position": "Left Wings",
      "items": [
        {
          "id": "5188356",
          "uid": "s:70~l:90~a:5188356",
          "guid": "b672405c-3d77-301d-ad6d-ae5fc1e8d7c3",
          "alternateIds": {
            "sdr": "5188356"
          },
          "alternateId": "5188356",
          "firstName": "Alex",
          "lastName": "Bump",
          "fullName": "Alex Bump",
          "displayName": "Alex Bump",
          "shortName": "A. Bump",
          "weight": 195.0,
          "displayWeight": "195 lbs",
          "height": 72.0,
          "displayHeight": "6' 0\"",
          "age": 22,
          "dateOfBirth": "2003-11-20T08:00Z",
          "links": [
            {},
            {}
          ],
          "birthPlace": {
            "city": "Burnsville",
            "state": "MN",
            "country": "USA",
            "displayText": "Burnsville, MN"
          }
        },
        {
          "id": "4419682",
          "uid": "s:70~l:90~a:4419682",
          "guid": "4353bdca-1851-39a3-8193-675e25d4eda2",
          "alternateIds": {
            "sdr": "4419682"
          },
          "alternateId": "4419682",
          "firstName": "Noah",
          "lastName": "Cates",
          "fullName": "Noah Cates",
          "displayName": "Noah Cates",
          "shortName": "N. Cates",
          "weight": 194.0,
          "displayWeight": "194 lbs",
          "height": 74.0,
          "displayHeight": "6' 2\"",
          "age": 27,
          "dateOfBirth": "1999-02-05T08:00Z",
          "links": [
            {},
            {}
          ],
          "birthPlace": {
            "city": "Stillwater",
            "state": "MN",
            "country": "USA",
            "displayText": "Stillwater, MN"
          }
        }
      ]
    }
  ],
  "coach": [
    {
      "id": "946",
      "firstName": "Rick",
      "lastName": "Tocchet"
    }
  ],
  "team": {
    "id": "15",
    "abbreviation": "PHI",
    "location": "Philadelphia",
    "name": "Flyers",
    "displayName": "Philadelphia Flyers",
    "clubhouse": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
    "color": "fe5823",
    "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
    "recordSummary": "43-27-12",
    "seasonSummary": "2025-26",
    "standingSummary": "2nd in Metropolitan Division"
  }
}
```
