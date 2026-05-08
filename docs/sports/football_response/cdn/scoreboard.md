# CDN Scoreboard

## https://cdn.espn.com/core/nfl/scoreboard?xhr=1

Notes:
- Verified on 2026-05-07.
- Returns a CDN page shell, not just the raw site scoreboard object.
- Scoreboard data is nested under `content.sbData`.
- Tested response contained one NFL event, `401772988`.

## Example Response

```json
{
  "sport": ["nfl"],
  "topKeys": [
    "news",
    "pinnedCount",
    "nowFeedMD5Hash",
    "type",
    "content",
    "analytics",
    "nowFeed",
    "ads",
    "nowFeedCount",
    "meta",
    "nowFeedSupported",
    "sport",
    "tier2Nav"
  ],
  "content": {
    "league": {},
    "sbGroup": {},
    "sbData": {
      "events": [
        {
          "id": "401772988",
          "name": "Seattle Seahawks at New England Patriots",
          "date": "2026-02-08T23:30Z",
          "status": {
            "type": {
              "name": "STATUS_FINAL"
            }
          },
          "competitions": [{}]
        }
      ]
    }
  }
}
```
