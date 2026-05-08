# Sports News

## https://now.core.api.espn.com/v1/sports/news

Notes:
- Verified on 2026-05-07.
- Returns normalized real-time headline objects.
- Use `sport=football` for football-scoped results.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `sport` | `string` | Sport slug such as `football` |
| `limit` | `int` | Requested headline count; tested `limit=2` returned `resultsLimit=10` |
| `offset` | `int` | Pagination offset |

## Example Response

```json
{
  "resultsCount": 2,
  "resultsLimit": 10,
  "resultsOffset": 0,
  "timestamp": "2026-05-08T15:07:20Z",
  "status": "success",
  "headlines": [
    {
      "headline": "Ex-Ohio coach Brian Smith suing school following December firing",
      "categories": [
        {
          "type": "topic",
          "description": "news"
        },
        {
          "type": "league",
          "description": "NCAA Football"
        },
        {
          "type": "team",
          "description": "Ohio Bobcats"
        }
      ]
    },
    {
      "headline": "Browns' Owusu-Koramoah (neck) to miss second straight season",
      "categories": [
        {
          "type": "league",
          "description": "NFL"
        },
        {
          "type": "team",
          "description": "Cleveland Browns"
        }
      ]
    }
  ],
  "breakingNews": []
}
```
