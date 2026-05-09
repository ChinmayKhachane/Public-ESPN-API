# Scoreboard Header

## https://site.web.api.espn.com/apis/v2/scoreboard/header?sport=basketball&league=nba

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "sports": [
    {
      "id": "40",
      "uid": "s:40",
      "guid": "cd70a58e-a830-330c-93ed-52360b51b632",
      "name": "Basketball",
      "slug": "basketball",
      "logos": [
        {},
        {}
      ],
      "leagues": [
        {}
      ]
    }
  ]
}
```
