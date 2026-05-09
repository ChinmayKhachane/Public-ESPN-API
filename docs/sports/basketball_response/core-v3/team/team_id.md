# Core V3 Team Detail

## https://sports.core.api.espn.com/v3/sports/basketball/{league}/teams/{id}

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.
- `enable=logos,links,groups,venue,record` produced the richer tested shape.

## Example Response

```json
{
  "id": "20",
  "uid": "s:40~l:46~t:20",
  "guid": "ca1685ed-b799-53e4-7924-e58ea6eb8f3a",
  "name": "76ers",
  "displayName": "Philadelphia 76ers",
  "shortDisplayName": "76ers",
  "abbreviation": "PHI",
  "slug": "philadelphia-76ers",
  "record": {
    "total": {
      "stats": {
        "wins": {},
        "losses": {},
        "ties": {},
        "OTWins": {},
        "OTLosses": {},
        "points": {},
        "pointsFor": {},
        "pointsAgainst": {},
        "avgPointsFor": {},
        "avgPointsAgainst": {},
        "gamesPlayed": {},
        "winPercent": {}
      }
    },
    "lastTenGames": {
      "stats": {
        "wins": {},
        "losses": {},
        "ties": {},
        "OTWins": {},
        "OTLosses": {},
        "points": {},
        "pointsFor": {},
        "pointsAgainst": {},
        "avgPointsFor": {},
        "avgPointsAgainst": {},
        "gamesPlayed": {},
        "winPercent": {}
      }
    }
  },
  "venue": {
    "id": "1845",
    "fullName": "Xfinity Mobile Arena",
    "shortName": "Wells Fargo Center",
    "address": {
      "city": "Philadelphia",
      "state": "PA"
    },
    "grass": false,
    "indoor": true
  },
  "groups": [
    {
      "id": "5",
      "uid": "s:40~l:46~g:5",
      "name": "Eastern Conference",
      "abbreviation": "East",
      "slug": "eastern-conference",
      "isConference": true
    },
    {
      "id": "1",
      "uid": "s:40~l:46~g:1",
      "name": "Atlantic",
      "abbreviation": "AT",
      "slug": "atlantic",
      "isConference": false
    }
  ],
  "location": "Philadelphia"
}
```
