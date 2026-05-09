# MLB Competition Situation

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/situation

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/situ...",
  "lastPlay": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/play..."
  },
  "balls": 0,
  "strikes": 0,
  "outs": 0,
  "situationNotes": [
    {
      "type": "CHANCES_TO_SCORE_1PLUS",
      "text": "Chance of scoring 1+ runs this inning (0 outs, Bases Empty): 30.08%"
    },
    {
      "type": "CHANCES_TO_SCORE_2PLUS",
      "text": "Chance of scoring 2+ runs this inning (0 outs, Bases Empty): 14.70%"
    }
  ]
}
```
