# Scoreboard Header

## https://site.web.api.espn.com/apis/v2/scoreboard/header?sport=hockey&league={league}

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.web.api.espn.com/apis/v2/scoreboard/header?sport=hockey&league=nhl`

## Example Response

```json
{
  "sports": [
    {
      "id": "70",
      "uid": "s:70",
      "guid": "2512ac76-a335-39cb-af51-b9afffc6571d",
      "name": "Ice Hockey",
      "slug": "hockey",
      "logos": [
        {
          "href": "https://a.espncdn.com/redesign/assets/img/icons/ESPN-icon-hockey.png",
          "alt": "",
          "rel": [],
          "width": 500,
          "height": 500
        },
        {
          "href": "https://a.espncdn.com/guid/2512ac76-a335-39cb-af51-b9afffc6571d/logos/default-dark.png",
          "alt": "",
          "rel": [],
          "width": 500,
          "height": 500
        }
      ],
      "leagues": [
        {
          "id": "90",
          "uid": "s:70~l:90",
          "name": "National Hockey League",
          "abbreviation": "NHL",
          "shortName": "NHL",
          "slug": "nhl",
          "tag": "nhl",
          "isTournament": false,
          "smartdates": [],
          "events": []
        }
      ]
    }
  ]
}
```
