# Event Competition

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988` on 2026-05-08.
- This is the richest football core-v2 event object. Most useful subresources hang off this response.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `period` | `int` | Used by some competition child endpoints |
| `sort` | `string` | Used by some competition child endpoints |
| `source` | `string` | Used by some competition child endpoints |
| `showsubplays` | `bool` | Used by some competition child endpoints |

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988?lang=en&region=us",
  "id": "401772988",
  "uid": "s:20~l:28~e:401772988~c:401772988",
  "date": "2026-02-08T23:30Z",
  "attendance": 70823,
  "neutralSite": true,
  "boxscoreAvailable": true,
  "playByPlayAvailable": true,
  "summaryAvailable": true,
  "venue": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/venues/4738?lang=en&region=us",
    "id": "4738",
    "fullName": "Levi's Stadium"
  },
  "competitors": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17?lang=en&region=us",
      "id": "17",
      "homeAway": "home"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/26?lang=en&region=us",
      "id": "26",
      "homeAway": "away"
    }
  ],
  "broadcasts": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/broadcasts?lang=en&region=us"
  },
  "odds": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/odds?lang=en&region=us"
  },
  "officials": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/officials?lang=en&region=us"
  },
  "situation": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/situation?lang=en&region=us"
  }
}
```
