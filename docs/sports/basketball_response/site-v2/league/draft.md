# Draft

## https://site.api.espn.com/apis/site/v2/sports/basketball/{league}/draft

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "uid": "s:40~l:46~e:draft~y:2026",
  "displayName": "2026 National Basketball Association Draft",
  "shortDisplayName": "2026 NBA Draft",
  "year": 2026,
  "status": {
    "name": "SCHEDULED",
    "round": 1,
    "state": "pre",
    "description": "Scheduled"
  },
  "teams": [
    {
      "id": "1",
      "uid": "s:40~l:46~t:1",
      "name": "Hawks",
      "displayName": "Atlanta Hawks",
      "shortDisplayName": "Hawks",
      "abbreviation": "ATL",
      "record": {
        "season": 2026,
        "summary": "46-36",
        "standing": "6th in Eastern Conference"
      },
      "location": "Atlanta",
      "logo": "https://a.espncdn.com/i/teamlogos/nba/500/scoreboard/atl.png",
      "darkLogo": "https://a.espncdn.com/i/teamlogos/nba/500-dark/scoreboard/atl.png",
      "link": "https://www.espn.com/nba/draft/teams/_/name/atl/atlanta-hawks"
    },
    {
      "id": "17",
      "uid": "s:40~l:46~t:17",
      "name": "Nets",
      "displayName": "Brooklyn Nets",
      "shortDisplayName": "Nets",
      "abbreviation": "BKN",
      "record": {
        "season": 2026,
        "summary": "20-62",
        "standing": "13th in Eastern Conference"
      },
      "location": "Brooklyn",
      "logo": "https://a.espncdn.com/i/teamlogos/nba/500/scoreboard/bkn.png",
      "darkLogo": "https://a.espncdn.com/i/teamlogos/nba/500-dark/scoreboard/bkn.png",
      "link": "https://www.espn.com/nba/draft/teams/_/name/bkn/brooklyn-nets"
    }
  ],
  "rounds": 2,
  "picks": [],
  "broadcasts": [
    {
      "slug": "abc",
      "type": {
        "id": "1",
        "slug": "tv",
        "shortName": "TV",
        "longName": "TV"
      },
      "station": "ABC",
      "priority": 0,
      "market": {
        "id": "1",
        "type": "National"
      },
      "media": {
        "id": "2",
        "name": "ABC",
        "slug": "abc",
        "callLetters": "ABC",
        "shortName": "ABC"
      },
      "lang": "en",
      "region": "us"
    },
    {
      "slug": "nfl-network",
      "type": {
        "id": "1",
        "slug": "tv",
        "shortName": "TV",
        "longName": "TV"
      },
      "station": "NFL Network",
      "priority": 0,
      "market": {
        "id": "1",
        "type": "National"
      },
      "media": {
        "id": "398",
        "name": "NFL Network",
        "slug": "nfl-network",
        "callLetters": "NFL Net",
        "shortName": "NFL Net"
      },
      "lang": "en",
      "region": "us"
    }
  ],
  "positions": [
    {
      "id": "3",
      "displayName": "Guard",
      "abbreviation": "G"
    },
    {
      "id": "7",
      "displayName": "Forward",
      "abbreviation": "F"
    }
  ]
}
```
