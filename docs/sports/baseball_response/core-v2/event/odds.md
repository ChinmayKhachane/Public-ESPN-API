# MLB Competition Odds

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/odds?limit=5

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "count": 1,
  "pageIndex": 1,
  "pageSize": 5,
  "pageCount": 1,
  "items": [
    {
      "provider": {
        "id": "100",
        "name": "DraftKings",
        "priority": 1
      },
      "details": "CIN -132",
      "overUnder": 9.0,
      "spread": -1.5,
      "initialSpread": 0.0,
      "initialOverUnder": 0.0,
      "price": 0.0,
      "overOdds": -108.0,
      "underOdds": -112.0,
      "awayTeamOdds": {
        "team": {
          "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/18?lang=en&region=us"
        },
        "favorite": false,
        "underdog": true,
        "moneyLine": 109,
        "open": {
          "favorite": false,
          "pointSpread": {
            "alternateDisplayValue": "+1.5",
            "american": "+1.5"
          },
          "spread": {
            "value": 1.53,
            "displayValue": "50/93",
            "alternateDisplayValue": "-186",
            "decimal": 1.53,
            "fraction": "50/93",
            "american": "-186"
          },
          "moneyLine": {
            "value": 2.09,
            "displayValue": "109/100",
            "alternateDisplayValue": "+109",
            "decimal": 2.09,
            "fraction": "109/100",
            "american": "+109"
          }
        },
        "close": {
          "pointSpread": {
            "alternateDisplayValue": "+1.5",
            "american": "+1.5"
          },
          "spread": {
            "value": 1.53,
            "displayValue": "25/47",
            "alternateDisplayValue": "-188",
            "decimal": 1.53,
            "fraction": "25/47",
            "american": "-188"
          },
          "moneyLine": {
            "value": 2.09,
            "displayValue": "109/100",
            "alternateDisplayValue": "+109",
            "decimal": 2.09,
            "fraction": "109/100",
            "american": "+109"
          }
        },
        "current": {
          "pointSpread": {
            "alternateDisplayValue": "+1.5",
            "american": "+1.5"
          },
          "spread": {
            "value": 1.53,
            "displayValue": "25/47",
            "alternateDisplayValue": "-188",
            "decimal": 1.53,
            "fraction": "25/47",
            "american": "-188"
          },
          "moneyLine": {
            "value": 2.09,
            "displayValue": "109/100",
            "alternateDisplayValue": "+109",
            "decimal": 2.09,
            "fraction": "109/100",
            "american": "+109"
          }
        }
      }
    }
  ]
}
```
