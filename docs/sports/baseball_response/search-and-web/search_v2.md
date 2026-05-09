# Baseball Search v2

## https://site.web.api.espn.com/apis/search/v2?query=shohei%20ohtani&sport=baseball

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "totalFound": 4,
  "resultTypes": [
    {
      "displayName": "Players",
      "totalFound": 1,
      "type": "player"
    },
    {
      "displayName": "Films",
      "totalFound": 1,
      "type": "film"
    }
  ],
  "results": [
    {
      "displayName": "Players",
      "type": "player",
      "totalFound": 1,
      "page": 1,
      "limit": 5,
      "contents": [
        {
          "id": "04a76996-4a5c-456d-8022-a9d45e3ff933",
          "uid": "s:1~l:10~a:39832",
          "guid": "04a76996-4a5c-456d-8022-a9d45e3ff933",
          "displayName": "Shohei Ohtani",
          "type": "player",
          "subtitle": "Los Angeles Dodgers",
          "link": {
            "app": "sportscenter://x-callback-url/showClubhouse?uid=s:1~l:10~a:39832",
            "web": "https://www.espn.com/mlb/player/_/id/39832"
          },
          "image": {
            "default": "https://a.espncdn.com/i/headshots/mlb/players/full/39832.png",
            "defaultDark": "https://a.espncdn.com/i/headshots/mlb/players/full/39832.png"
          },
          "defaultLeagueSlug": "mlb",
          "sport": "baseball"
        }
      ]
    },
    {
      "displayName": "Films",
      "type": "film",
      "totalFound": 1,
      "page": 1,
      "limit": 5,
      "contents": [
        {
          "id": "2c4e44a7-7ca6-44d8-b11a-a3bd7b0cde8c",
          "guid": "2c4e44a7-7ca6-44d8-b11a-a3bd7b0cde8c",
          "displayName": "Shohei Ohtani: Beyond the Dream",
          "type": "film",
          "link": {
            "web": "https://www.espn.com/watch/film/2c4e44a7-7ca6-44d8-b11a-a3bd7b0cde8c/shohei-ohtani-beyond-the-dream"
          },
          "image": {
            "default": "http://artwork.espncdn.com/categories/2c4e44a7-7ca6-44d8-b11a-a3bd7b0cde8c/2x3/288x432_20231114151708.jpg"
          }
        }
      ]
    }
  ]
}
```
