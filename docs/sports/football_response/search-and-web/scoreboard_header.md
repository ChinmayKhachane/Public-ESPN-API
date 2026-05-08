# Scoreboard Header

## https://site.web.api.espn.com/apis/v2/scoreboard/header

Notes:
- Verified on 2026-05-07.
- The unfiltered endpoint returns multiple sports.
- `sport=football&league=nfl` returned a football/NFL-specific header.
- `sport=football` without `league=nfl` returned HTTP `400` in the tested request.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `sport` | `string` | Sport slug. Tested with `football`. |
| `league` | `string` | League slug. Tested with `nfl`. |

## Example Response

```json
{
  "sports": [
    {
      "name": "Football",
      "slug": "football",
      "leagues": [
        {
          "name": "National Football League",
          "abbreviation": "NFL",
          "events": [
            {
              "id": "401772988",
              "name": "Seattle Seahawks at New England Patriots",
              "date": "2026-02-08T23:30:00Z"
            }
          ]
        }
      ]
    }
  ]
}
```
