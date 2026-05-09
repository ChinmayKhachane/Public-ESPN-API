# Athlete Bio

## https://site.web.api.espn.com/apis/common/v3/sports/hockey/{league}/athletes/{athlete}/bio

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.web.api.espn.com/apis/common/v3/sports/hockey/nhl/athletes/4565230/bio`

## Example Response

```json
{
  "teamHistory": [
    {
      "id": "15",
      "uid": "s:70~l:90~t:15",
      "slug": "philadelphia-flyers",
      "displayName": "Philadelphia Flyers",
      "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
      "seasons": "2025-CURRENT",
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
      "seasonCount": "1",
      "isActive": true
    },
    {
      "id": "25",
      "uid": "s:70~l:90~t:25",
      "slug": "anaheim-ducks",
      "displayName": "Anaheim Ducks",
      "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/ana.png",
      "seasons": "2020-2025",
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
      "seasonCount": "5",
      "isActive": true
    }
  ]
}
```
