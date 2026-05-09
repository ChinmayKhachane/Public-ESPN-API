# Baseball Scoreboard Header

## https://site.web.api.espn.com/apis/v2/scoreboard/header?sport=baseball&league=mlb

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "sports": [
    {
      "id": "1",
      "uid": "s:1",
      "guid": "e364bfcd-493d-3bfb-ac83-bd27d66fedd0",
      "name": "Baseball",
      "slug": "baseball",
      "leagues": [
        {
          "id": "10",
          "uid": "s:1~l:10",
          "name": "Major League Baseball",
          "abbreviation": "MLB",
          "slug": "mlb",
          "events": [
            {},
            {}
          ],
          "shortName": "MLB",
          "tag": "mlb",
          "isTournament": false,
          "smartdates": [
            "2026-05-07T07:00Z",
            "2026-05-08T07:00Z"
          ]
        }
      ],
      "logos": [
        {
          "href": "https://a.espncdn.com/combiner/i?img=/redesign/assets/img/icons/ESPN-icon-baseball.png",
          "alt": "",
          "rel": [
            "full",
            "default"
          ],
          "width": 500,
          "height": 500
        },
        {
          "href": "https://a.espncdn.com/guid/e364bfcd-493d-3bfb-ac83-bd27d66fedd0/logos/default-dark.png",
          "alt": "",
          "rel": [
            "full",
            "dark"
          ],
          "width": 500,
          "height": 500
        }
      ]
    }
  ]
}
```
