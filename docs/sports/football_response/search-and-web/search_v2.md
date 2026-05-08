# Search V2

## https://site.web.api.espn.com/apis/search/v2?query={query}

Notes:
- Verified with `query=drake maye`, `query=cowboys`, and `sport=football` on 2026-05-07.
- Top-level `results[]` are result groups such as `player`, `article`, `clips`, or `team`.
- Individual search hits live in each group's `contents[]` array.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `query` | `string` | Search text |
| `sport` | `string` | Sport slug such as `football` |
| `limit` | `int` | Per-group result limit |
| `page` | `int` | Page number |

## Example Response

```json
{
  "totalFound": 3,
  "didYouMean": "drake make",
  "resultTypes": [
    {
      "totalFound": 1,
      "type": "player",
      "displayName": "Players"
    },
    {
      "totalFound": 13104,
      "type": "article",
      "displayName": "Articles"
    },
    {
      "totalFound": 79,
      "type": "clips",
      "displayName": "Clips"
    }
  ],
  "results": [
    {
      "type": "player",
      "displayName": "Players",
      "contents": [
        {
          "uid": "s:20~l:28~a:4431452",
          "type": "player",
          "displayName": "Drake Maye",
          "description": "NFL",
          "subtitle": "New England Patriots",
          "link": {
            "web": "https://www.espn.com/nfl/player/_/id/4431452/drake-maye"
          },
          "defaultLeagueSlug": "nfl",
          "sport": "football"
        }
      ]
    },
    {
      "type": "article",
      "displayName": "Articles",
      "contents": [
        {
          "id": "48697495",
          "type": "headlinenews",
          "displayName": "Patriots QB Maye backs Vrabel, not worried about distraction",
          "link": {
            "web": "https://www.espn.com/nfl/story/_/id/48697495/patriots-qb-maye-backs-vrabel-not-worried-distraction"
          },
          "isPremium": false
        }
      ]
    }
  ]
}
```
