# Common V3 Team Roster

## https://site.web.api.espn.com/apis/common/v3/sports/hockey/{league}/teams/{team}/roster

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.web.api.espn.com/apis/common/v3/sports/hockey/nhl/teams/15/roster`

## Example Response

```json
{
  "season": {
    "year": 2026,
    "type": 3,
    "name": "Postseason"
  },
  "coach": [
    {
      "id": "946",
      "firstName": "Rick",
      "lastName": "Tocchet"
    }
  ],
  "positionGroups": [
    {
      "type": "centers",
      "displayName": "Centers",
      "athletes": [
        {
          "id": "4064976",
          "alternateIds": [
            {}
          ],
          "guid": "f0f7c3a7-6e8a-47fa-b290-96724da24d8a",
          "firstName": "Rodrigo",
          "lastName": "Abols",
          "displayName": "Rodrigo Abols",
          "fullName": "Rodrigo Abols",
          "shortName": "R. Abols",
          "displayWeight": "206 lbs",
          "weight": 206,
          "displayHeight": "6' 4\"",
          "height": 76,
          "age": 30,
          "displayDOB": "1996-01-05T08:00Z",
          "debutYear": 2016,
          "links": [
            {},
            {}
          ],
          "birthPlace": {
            "city": "Riga",
            "country": "LAT"
          },
          "jersey": "18"
        },
        {
          "id": "5173105",
          "alternateIds": [
            {}
          ],
          "guid": "6516db75-a54c-3318-85a3-ecee5c738731",
          "firstName": "Denver",
          "lastName": "Barkey",
          "displayName": "Denver Barkey",
          "fullName": "Denver Barkey",
          "shortName": "D. Barkey",
          "displayWeight": "155 lbs",
          "weight": 155,
          "displayHeight": "5' 9\"",
          "height": 69,
          "age": 21,
          "displayDOB": "2005-04-27T07:00Z",
          "links": [
            {},
            {}
          ],
          "birthPlace": {
            "city": "Newmarket",
            "state": "ON",
            "country": "CAN"
          },
          "jersey": "52",
          "headshot": {
            "href": "https://a.espncdn.com/i/headshots/nhl/players/full/5173105.png",
            "alt": "Denver Barkey"
          }
        }
      ]
    },
    {
      "type": "defense",
      "displayName": "Defense",
      "athletes": [
        {
          "id": "4697457",
          "alternateIds": [
            {}
          ],
          "guid": "53b8e8bf-3fa0-31f1-ae3a-8e5e442fb810",
          "firstName": "Emil",
          "lastName": "Andrae",
          "displayName": "Emil Andrae",
          "fullName": "Emil Andrae",
          "shortName": "E. Andrae",
          "displayWeight": "189 lbs",
          "weight": 189,
          "displayHeight": "5' 9\"",
          "height": 69,
          "age": 24,
          "displayDOB": "2002-02-23T08:00Z",
          "links": [
            {},
            {}
          ],
          "birthPlace": {
            "city": "Vasteras",
            "country": "SWE"
          },
          "jersey": "36",
          "headshot": {
            "href": "https://a.espncdn.com/i/headshots/nhl/players/full/4697457.png",
            "alt": "Emil Andrae"
          }
        },
        {
          "id": "5149199",
          "alternateIds": [
            {}
          ],
          "guid": "fccacacd-9ecd-3d37-bbb6-ff5132d50e2e",
          "firstName": "Oliver",
          "lastName": "Bonk",
          "displayName": "Oliver Bonk",
          "fullName": "Oliver Bonk",
          "shortName": "O. Bonk",
          "displayWeight": "180 lbs",
          "weight": 180,
          "displayHeight": "6' 2\"",
          "height": 74,
          "age": 21,
          "displayDOB": "2005-01-09T08:00Z",
          "links": [
            {},
            {}
          ],
          "birthPlace": {
            "city": "Ottawa",
            "state": "ON",
            "country": "CAN"
          },
          "jersey": "59",
          "headshot": {
            "href": "https://a.espncdn.com/i/headshots/nhl/players/full/5149199.png",
            "alt": "Oliver Bonk"
          }
        }
      ]
    }
  ],
  "team": {
    "id": "15",
    "abbreviation": "PHI",
    "location": "Philadelphia",
    "name": "Flyers",
    "displayName": "Philadelphia Flyers",
    "venueLink": "http://sports.core.api.espn.pvt/v2/sports/hockey/leagues/nhl/venues/1845?lang=en",
    "clubhouse": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
    "color": "fe5823",
    "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
    "recordSummary": "43-27-12, 98 PTS",
    "seasonSummary": "2026",
    "standingSummary": "3rd in Metropolitan Division"
  }
}
```
