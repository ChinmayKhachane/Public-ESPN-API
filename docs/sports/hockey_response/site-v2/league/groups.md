# Groups

## https://site.api.espn.com/apis/site/v2/sports/hockey/{league}/groups

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/groups`

## Example Response

```json
{
  "status": "success",
  "groups": [
    {
      "name": "Eastern Conference",
      "abbreviation": "East",
      "children": [
        {
          "name": "Atlantic Division",
          "abbreviation": "ATL",
          "teams": []
        },
        {
          "name": "Metropolitan Division",
          "abbreviation": "MET",
          "teams": []
        }
      ]
    },
    {
      "name": "Western Conference",
      "abbreviation": "West",
      "children": [
        {
          "name": "Central Division",
          "abbreviation": "CEN",
          "teams": []
        },
        {
          "name": "Pacific Division",
          "abbreviation": "PAC",
          "teams": []
        }
      ]
    }
  ]
}
```
