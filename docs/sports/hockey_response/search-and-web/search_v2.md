# Search V2

## https://site.web.api.espn.com/apis/search/v2?query={query}&sport=hockey

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.web.api.espn.com/apis/search/v2?query=flyers&sport=hockey`

## Example Response

```json
{
  "totalFound": 5,
  "resultTypes": [
    {
      "totalFound": 38,
      "type": "team",
      "displayName": "Teams"
    },
    {
      "totalFound": 18189,
      "type": "article",
      "displayName": "Articles"
    }
  ],
  "results": [
    {
      "type": "team",
      "totalFound": 5,
      "page": 1,
      "limit": 5,
      "displayName": "Teams",
      "contents": [
        {
          "id": "68aba012-4e93-9371-6861-1bb9a63cfb11",
          "uid": "s:70~l:90~t:15",
          "guid": "68aba012-4e93-9371-6861-1bb9a63cfb11",
          "type": "team",
          "displayName": "Philadelphia Flyers",
          "subtitle": "NHL",
          "link": {},
          "image": {},
          "defaultLeagueSlug": "nhl",
          "sport": "hockey"
        },
        {
          "id": "a2827104-45ed-950b-3444-0192dd5a1723",
          "uid": "s:40~l:41~t:2168",
          "guid": "a2827104-45ed-950b-3444-0192dd5a1723",
          "type": "team",
          "displayName": "Dayton Flyers",
          "subtitle": "NCAAM",
          "link": {},
          "image": {},
          "defaultLeagueSlug": "mens-college-basketball",
          "sport": "basketball"
        }
      ]
    },
    {
      "type": "article",
      "totalFound": 5,
      "page": 1,
      "limit": 5,
      "displayName": "Articles",
      "contents": [
        {
          "id": "48709047",
          "nowId": "1-48709047",
          "type": "headlinenews",
          "displayName": "Flyers in 3-0 hole vs. Canes as penalties, power play costly",
          "link": {},
          "images": [],
          "byline": "Greg Wyshynski",
          "date": "2026-05-08T05:21:08.000+00:00",
          "isPremium": false,
          "categories": []
        },
        {
          "id": "48708465",
          "nowId": "1-48708465",
          "type": "recap",
          "displayName": "Hurricanes beat the Flyers 4-1 in Game 3, take a 3-0 series lead",
          "link": {},
          "images": [],
          "byline": "AP",
          "date": "2026-05-08T04:13:01.000+00:00",
          "isPremium": false,
          "categories": []
        }
      ]
    }
  ]
}
```
