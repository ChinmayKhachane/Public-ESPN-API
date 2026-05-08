# Athlete Detail

## https://sports.core.api.espn.com/v3/sports/football/{league}/athletes/{id}

Notes:
- Verified with `league=nfl`, `id=4431452` on 2026-05-07.
- The default response is sparse.
- Add `enable=team,position,links` to include team, position, and ESPN web links.

## Example Response

```json
{
  "id": "4431452",
  "uid": "s:20~l:28~a:4431452",
  "displayName": "Drake Maye",
  "displayHeight": "6' 4\"",
  "displayWeight": "225 lbs",
  "age": 23,
  "jersey": "10",
  "active": true,
  "position": {
    "id": "8",
    "name": "Quarterback",
    "displayName": "Quarterback",
    "abbreviation": "QB",
    "leaf": true
  },
  "team": {
    "id": "17",
    "slug": "new-england-patriots",
    "abbreviation": "NE",
    "displayName": "New England Patriots"
  },
  "links": [
    {
      "rel": ["playercard", "desktop", "athlete"],
      "href": "https://www.espn.com/nfl/player/_/id/4431452/drake-maye"
    },
    {
      "rel": ["stats", "desktop", "athlete"],
      "href": "https://www.espn.com/nfl/player/stats/_/id/4431452/drake-maye"
    }
  ]
}
```
