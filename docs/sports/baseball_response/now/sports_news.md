# Baseball Now News

## https://now.core.api.espn.com/v1/sports/news?sport=baseball

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "status": "success",
  "resultsCount": 10,
  "resultsLimit": 10,
  "resultsOffset": 0,
  "headlines": [
    {
      "id": 48589245,
      "categories": [
        {
          "id": 611584,
          "guid": "550a1e9e-377e-3019-bf5e-19c027d50a49",
          "type": "league",
          "sportId": 22000,
          "leagueId": 22000,
          "league": {
            "id": 22000
          }
        },
        {
          "id": 137817,
          "guid": "14f579db-79ec-a898-9405-8fec63e16063",
          "type": "topic",
          "sportId": 0,
          "topicId": 179
        }
      ],
      "nowId": "1-48589245",
      "contentKey": "48589245-1-6-1",
      "dataSourceIdentifier": "37c41f0d075d7",
      "publishedkey": "48589245",
      "type": "Story",
      "feedDisplayType": "Default",
      "headline": "MLB betting tips for Saturday: Nick Kurtz set up for success",
      "title": "MLB betting tips for Saturday: Nick Kurtz set up for success"
    },
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
      "publishedkey": "48717767",
      "type": "HeadlineNews",
      "feedDisplayType": "Default",
      "headline": "Padres prospect pleads guilty to charge of transporting noncitizen immigrants",
      "title": "Padres prospect pleads guilty to charge of transporting noncitizen immigrants"
    }
  ],
  "breakingNews": [],
  "timestamp": "2026-05-09T03:58:01Z"
}
```
