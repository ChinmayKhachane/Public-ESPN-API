# Athlete Overview

## https://site.web.api.espn.com/apis/common/v3/sports/football/{league}/athletes/{id}/overview

Notes:
- Verified with `league=nfl`, `id=4429202` on 2026-05-08.
- This endpoint combines season snapshot stats, news, recent game log metadata, Rotowire blurbs, and fantasy metadata.

## Example Response

```json
{
  "statistics": {
    "displayName": "2025 Offense",
    "categories": [
      {
        "name": "rushing",
        "displayName": "Rushing",
        "count": 5
      },
      {
        "name": "receiving",
        "displayName": "Receiving",
        "count": 5
      }
    ],
    "splits": [
      {
        "displayName": "Regular Season"
      },
      {
        "displayName": "Career"
      }
    ]
  },
  "news": [
    {
      "headline": "Jets confident about youngest running back room in the NFL"
    }
  ],
  "nextGame": {},
  "gameLog": {
    "displayName": "Recent Games"
  },
  "rotowire": {
    "headline": "Dallas signed Abanikanda to a reserve/future contract Tuesday."
  },
  "fantasy": {
    "draftRank": "1440"
  }
}
```
