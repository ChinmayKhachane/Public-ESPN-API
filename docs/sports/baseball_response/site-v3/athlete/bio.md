# MLB Athlete Bio

## https://site.web.api.espn.com/apis/common/v3/sports/baseball/mlb/athletes/4414528/bio

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "teamHistory": [
    {
      "id": "17",
      "uid": "s:1~l:10~t:17",
      "displayName": "Cincinnati Reds",
      "slug": "cincinnati-reds",
      "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
      "seasons": "2023-CURRENT",
      "links": [
        {
          "language": "en",
          "rel": [
            "clubhouse",
            "desktop"
          ],
          "href": "https://www.espn.com/mlb/team/_/name/cin/cincinnati-reds",
          "text": "Clubhouse",
          "shortText": "Clubhouse",
          "isExternal": false,
          "isPremium": false
        },
        {
          "language": "en",
          "rel": [
            "roster",
            "desktop"
          ],
          "href": "https://www.espn.com/mlb/team/roster/_/name/cin/cincinnati-reds",
          "text": "Roster",
          "shortText": "Roster",
          "isExternal": false,
          "isPremium": false
        }
      ],
      "seasonCount": "4",
      "isActive": true
    }
  ]
}
```
