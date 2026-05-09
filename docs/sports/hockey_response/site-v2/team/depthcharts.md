# Team Depth Charts

## https://site.api.espn.com/apis/site/v2/sports/hockey/{league}/teams/{team}/depthcharts

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/teams/15/depthcharts`

## Example Response

```json
{
  "timestamp": "2026-05-09T17:01:03Z",
  "status": "success",
  "season": {
    "year": 2026,
    "type": 3,
    "name": "Postseason"
  },
  "team": {
    "id": "15",
    "abbreviation": "PHI",
    "location": "Philadelphia",
    "name": "Flyers",
    "displayName": "Philadelphia Flyers",
    "clubhouse": "https://www.espn.com/nhl/team/_/name/phi/philadelphia-flyers",
    "color": "fe5823",
    "logo": "https://a.espncdn.com/i/teamlogos/nhl/500/phi.png",
    "recordSummary": "43-27-12",
    "seasonSummary": "2025-26",
    "standingSummary": "2nd in Metropolitan Division",
    "groups": {
      "$ref": "http://sports.core.api.espn.pvt/v2/sports/hockey/leagues/nhl/seasons/2026/types/2/groups/33?lang=en&region=us"
    }
  }
}
```
