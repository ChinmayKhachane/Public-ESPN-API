# Common V3 Team Roster

## https://site.web.api.espn.com/apis/common/v3/sports/basketball/{league}/teams/{id}/roster

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.
- Response is organized by `positionGroups[]`.

## Example Response

```json
{
  "season": {
    "name": "Postseason",
    "year": 2026,
    "type": 3
  },
  "team": {
    "id": "20",
    "name": "76ers",
    "displayName": "Philadelphia 76ers",
    "abbreviation": "PHI",
    "location": "Philadelphia",
    "venueLink": "http://sports.core.api.espn.pvt/v2/sports/basketball/leagues/nba/venues/1845?lang=en",
    "clubhouse": "https://www.espn.com/nba/team/_/name/phi/philadelphia-76ers",
    "color": "1d428a",
    "logo": "https://a.espncdn.com/i/teamlogos/nba/500/phi.png",
    "recordSummary": "45-37",
    "seasonSummary": "2026",
    "standingSummary": "4th in Atlantic"
  },
  "coach": [
    {
      "id": "52085",
      "firstName": "Nick",
      "lastName": "Nurse",
      "experience": 6
    }
  ],
  "positionGroups": [
    {
      "displayName": "All",
      "type": "all",
      "athletes": [
        {},
        {}
      ]
    }
  ]
}
```
