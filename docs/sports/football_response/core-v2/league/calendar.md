# League Calendar

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/calendar

Notes:
- Verified with `league=nfl` on 2026-05-08.
- This returns a small collection of calendar buckets, not a single expanded calendar object.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `dates` | `string` | Filter by date range when supported by downstream calendar refs |
| `weeks` | `string` | Filter by week when supported by downstream calendar refs |
| `seasontype` | `int` | Filter by season type when supported by downstream calendar refs |

## Example Response

```json
{
  "count": 4,
  "pageIndex": 1,
  "pageSize": 25,
  "pageCount": 1,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/calendar/ondays?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/calendar/offdays?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/calendar/whens?lang=en&region=us"
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/calendar/events?lang=en&region=us"
    }
  ]
}
```
