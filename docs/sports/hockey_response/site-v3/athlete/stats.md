# Athlete Stats

## https://site.web.api.espn.com/apis/common/v3/sports/hockey/{league}/athletes/{athlete}/stats

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.web.api.espn.com/apis/common/v3/sports/hockey/nhl/athletes/4565230/stats`

## Example Response

```json
{
  "filters": [
    {
      "displayName": "Season",
      "name": "seasontype",
      "value": "2",
      "options": [
        {
          "value": "2",
          "displayValue": "Regular Season"
        },
        {
          "value": "3",
          "displayValue": "Postseason"
        }
      ]
    }
  ],
  "teams": {
    "philadelphia-flyers": {
      "id": "15",
      "uid": "s:70~l:90~t:15",
      "guid": "68aba012-4e93-9371-6861-1bb9a63cfb11",
      "slug": "philadelphia-flyers",
      "location": "Philadelphia",
      "name": "Flyers",
      "nickname": "Flyers",
      "abbreviation": "PHI",
      "displayName": "Philadelphia Flyers",
      "shortDisplayName": "Flyers",
      "color": "fe5823",
      "alternateColor": "000000",
      "isActive": true,
      "isAllStar": false,
      "logos": [
        {
          "href": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
          "width": 500,
          "height": 500,
          "rel": [
            "full",
            "default"
          ]
        },
        {
          "href": "https://a.espncdn.com/i/teamlogos/nhl/500-dark/phi.png",
          "width": 500,
          "height": 500,
          "rel": [
            "full",
            "dark"
          ]
        }
      ],
      "links": [
        {
          "language": "en",
          "rel": [
            "clubhouse",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
          "text": "Clubhouse",
          "shortText": "Clubhouse",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "clubhouse",
            "mobile"
          ],
          "href": "https://m.espn.com/nhl/clubhouse?teamId=15",
          "text": "Clubhouse",
          "shortText": "Clubhouse",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "groups": {},
      "coaches": {}
    },
    "anaheim-ducks": {
      "id": "25",
      "uid": "s:70~l:90~t:25",
      "guid": "183c31f7-2eed-15c0-7fa1-a568d20bf3ad",
      "slug": "anaheim-ducks",
      "location": "Anaheim",
      "name": "Ducks",
      "nickname": "Ducks",
      "abbreviation": "ANA",
      "displayName": "Anaheim Ducks",
      "shortDisplayName": "Ducks",
      "color": "fc4c02",
      "alternateColor": "000000",
      "isActive": true,
      "isAllStar": false,
      "logos": [
        {
          "href": "https://a.espncdn.com/i/teamlogos/nhl/500/ana.png",
          "width": 500,
          "height": 500,
          "rel": [
            "full",
            "default"
          ]
        },
        {
          "href": "https://a.espncdn.com/i/teamlogos/nhl/500-dark/ana.png",
          "width": 500,
          "height": 500,
          "rel": [
            "full",
            "dark"
          ]
        }
      ],
      "links": [
        {
          "language": "en",
          "rel": [
            "clubhouse",
            "desktop"
          ],
          "href": "https://www.espn.com/nhl/team/_/name/ana/anaheim-ducks",
          "text": "Clubhouse",
          "shortText": "Clubhouse",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "clubhouse",
            "mobile"
          ],
          "href": "https://m.espn.com/nhl/clubhouse?teamId=25",
          "text": "Clubhouse",
          "shortText": "Clubhouse",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "groups": {},
      "coaches": {}
    }
  },
  "categories": [
    {
      "name": "center",
      "displayName": "Regular Season ",
      "labels": [
        "GP",
        "G"
      ],
      "names": [
        "games",
        "goals"
      ],
      "displayNames": [
        "Games Played",
        "Goals"
      ],
      "descriptions": [
        "Total games played.",
        "Total goals scored."
      ],
      "statistics": [
        {
          "teamId": "25",
          "teamSlug": "anaheim-ducks",
          "season": {
            "year": 2021,
            "displayName": "20-21"
          },
          "stats": [
            "24",
            "3"
          ],
          "position": "C"
        },
        {
          "teamId": "25",
          "teamSlug": "anaheim-ducks",
          "season": {
            "year": 2022,
            "displayName": "21-22"
          },
          "stats": [
            "75",
            "23"
          ],
          "position": "C"
        }
      ],
      "totals": [
        "349",
        "93"
      ],
      "sortKey": "center"
    }
  ],
  "glossary": [
    {
      "abbreviation": "+/-",
      "displayName": "Plus/Minus Rating"
    },
    {
      "abbreviation": "A",
      "displayName": "Assists"
    }
  ]
}
```
