# MLB Competition Plays

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/plays?limit=5

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.
- MLB play objects include baseball fields such as `atBatId`, `pitchCount`, `resultCount`, `outs`, and pitch/hit coordinates when available.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/play...",
  "count": 553,
  "pageIndex": 1,
  "pageSize": 5,
  "pageCount": 111,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/play...",
      "id": "4018152560000000059",
      "team": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/18?lang=en&region=us"
      },
      "sequenceNumber": "1",
      "type": {
        "id": "59",
        "text": "Start Inning",
        "type": "start-inning"
      },
      "text": "Top of the 1st inning",
      "shortText": "Top of the 1st inning",
      "alternativeText": "Top of the 1st inning",
      "shortAlternativeText": "Top of the 1st inning",
      "awayScore": 0
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/play...",
      "id": "4018152560001010001",
      "team": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/18?lang=en&region=us"
      },
      "sequenceNumber": "1",
      "type": {
        "id": "1",
        "text": "Start Batter/Pitcher",
        "alternativeText": "Now at bat",
        "type": "start-batterpitcher"
      },
      "text": "Nick Lodolo pitches to Jose Altuve",
      "shortText": "Nick Lodolo pitches to Jose Altuve",
      "alternativeText": "Nick Lodolo pitches to Jose Altuve",
      "shortAlternativeText": "Nick Lodolo pitches to Jose Altuve",
      "awayScore": 0
    }
  ]
}
```
