# Season Draft

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/seasons/{season}/draft

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `available` | `bool` | Availability filter |
| `position` | `string` | Position filter |
| `team` | `string` | Team filter |

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/draft?lang=en&region=us",
  "uid": "s:40~l:46~e:draft~y:2026",
  "displayName": "2026 National Basketball Association Draft",
  "shortDisplayName": "2026 NBA Draft",
  "year": 2026,
  "status": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/draft/status?lang=en&region=us"
  },
  "athletes": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/draft/athletes?lang=en&region=us"
  },
  "rounds": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/draft/rounds?lang=en&region=us"
  },
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
  "numberOfRounds": 2,
  "positions": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/positions/3?lang=en&region=us",
      "id": "3",
      "name": "Guard",
      "displayName": "Guard",
      "abbreviation": "G",
      "leaf": false
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/positions/7?lang=en&region=us",
      "id": "7",
      "name": "Forward",
      "displayName": "Forward",
      "abbreviation": "F",
      "leaf": false
    }
  ],
  "links": [
    {
      "text": "NBA Draft",
      "shortText": "NBA Draft",
      "language": "en-US",
      "rel": [
        "index",
        "desktop"
      ],
      "href": "https://www.espn.com/nba/draft",
      "isExternal": false,
      "isPremium": false
    }
  ]
}
```
