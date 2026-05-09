# MLB News

## https://site.api.espn.com/apis/site/v2/sports/baseball/mlb/news

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "header": "MLB News",
  "link": {
    "language": "en",
    "rel": [
      "index",
      "desktop"
    ],
    "href": "https://www.espn.com/mlb/",
    "text": "All MLB News",
    "shortText": "All News",
    "isExternal": false,
    "isPremium": false
  },
  "articles": [
    {
      "id": 48717767,
      "categories": [
        {
          "id": 80311,
          "uid": "s:1~l:10",
          "guid": "b38f959b-7865-31ac-8841-b88355519e10",
          "type": "league",
          "sportId": 10,
          "leagueId": 10,
          "league": {
            "id": 10,
            "abbreviation": "MLB",
            "links": {}
          }
        },
        {
          "id": 724,
          "uid": "s:1~l:10~t:25",
          "guid": "4dec648c-3eb9-055c-aebc-2711f30975a0",
          "team": {
            "id": 25,
            "links": {}
          },
          "type": "team",
          "sportId": 10,
          "teamId": 25
        }
      ],
      "nowId": "1-48717767",
      "contentKey": "48717767-1-5-1",
      "dataSourceIdentifier": "f60e80c46aee3",
      "type": "HeadlineNews",
      "headline": "Padres prospect pleads guilty to charge of transporting noncitizen immigrants",
      "lastModified": "2026-05-09T03:43:19Z",
      "published": "2026-05-09T03:43:19Z",
      "images": [
        {
          "id": 13368714,
          "name": "Baseball seams 150802 [600x400]",
          "dataSourceIdentifier": "886e0be6d23ad",
          "type": "header",
          "alt": "Baseball seams",
          "credit": "Pouya Dianat/Getty Images",
          "height": 400,
          "width": 600,
          "url": "https://a.espncdn.com/photo/2015/0802/mlb_baseball_b1_600x400.jpg"
        }
      ]
    },
    {
      "id": 48717747,
      "categories": [
        {
          "id": 525402,
          "uid": "s:1~l:10~a:4142424",
          "guid": "c8f56866-3f5a-304f-9961-10d6b8970628",
          "type": "athlete",
          "sportId": 10,
          "athleteId": 4142424,
          "athlete": {
            "id": 4142424,
            "links": {}
          }
        },
        {
          "id": 468339,
          "uid": "s:1~l:10~a:4717833",
          "guid": "dcbeba5a-3fe9-3bf0-9f19-c077ef2a6f78",
          "type": "athlete",
          "sportId": 10,
          "athleteId": 4717833,
          "athlete": {
            "id": 4717833,
            "links": {}
          }
        }
      ],
      "nowId": "1-48717747",
      "contentKey": "48717747-1-293-1",
      "dataSourceIdentifier": "c9a50ebd56d4c",
      "type": "Media",
      "headline": "Chicago Cubs vs. Texas Rangers: Game Highlights",
      "lastModified": "2026-05-09T03:30:14Z",
      "published": "2026-05-09T03:30:14Z",
      "images": [
        {
          "name": "Chicago Cubs vs. Texas Rangers: Game Highlights",
          "alt": "",
          "height": 324,
          "width": 576,
          "url": "https://a.espncdn.com/media/motion/wsc/2026/0509/e4a1a620-c370-4340-9b36-3c6550c5edc9/e4a1a620-c370-4340-9b..."
        }
      ]
    }
  ]
}
```
