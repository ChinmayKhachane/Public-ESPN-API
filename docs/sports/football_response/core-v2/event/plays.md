# Competition Plays

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/plays

Notes:
- Verified with `league=nfl`, `event=401772988`, `competition=401772988` on 2026-05-08.
- This is the main play-by-play collection.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `page` | `int` | Page number |
| `limit` | `int` | Page size |
| `period` | `int` | Filter by period |
| `sort` | `string` | Sort order |
| `source` | `string` | Feed/source selector |
| `showsubplays` | `bool` | Include subplays where available |

## Example Response

```json
{
  "count": 206,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 103,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/plays/4017729881?lang=en&region=us",
      "id": "4017729881",
      "sequenceNumber": "100",
      "type": {
        "id": "70",
        "text": "Coin Toss"
      },
      "text": "GAME",
      "awayScore": 0,
      "homeScore": 0
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/plays/40177298840?lang=en&region=us",
      "id": "40177298840",
      "sequenceNumber": "4000",
      "type": {
        "id": "53",
        "text": "Kickoff",
        "abbreviation": "K"
      },
      "text": "A.Borregales kicks 65 yards from NE 35 to end zone, Touchback to the SEA 35.",
      "probability": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/probabilities/40177298840?lang=en&region=us"
      }
    }
  ]
}
```
