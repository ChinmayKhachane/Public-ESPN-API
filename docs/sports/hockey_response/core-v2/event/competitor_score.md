# Competitor Score

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}/score

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15/score`

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/competitors/15/scores/2?lang=en&region=us",
  "value": 1.0,
  "displayValue": "1",
  "winner": false,
  "source": {
    "id": "2",
    "description": "feed",
    "state": "full"
  }
}
```
