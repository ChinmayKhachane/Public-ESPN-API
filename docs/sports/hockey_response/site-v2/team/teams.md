# Teams

## https://site.api.espn.com/apis/site/v2/sports/hockey/{league}/teams

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json`.

Tested URL:
`https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/teams`

## Example Response

```json
{
  "sports": [
    {
      "id": "70",
      "leagues": [
        {
          "abbreviation": "NHL",
          "id": "90",
          "name": "National Hockey League",
          "season": {},
          "shortName": "NHL",
          "slug": "nhl",
          "teams": [],
          "uid": "s:70~l:90",
          "year": 2026
        }
      ],
      "name": "Ice Hockey",
      "slug": "hockey",
      "uid": "s:70"
    }
  ]
}
```
