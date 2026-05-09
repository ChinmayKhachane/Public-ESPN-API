# Plays

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/events/{event}/competitions/{competition}/plays

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/plays?limit=2`

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/plays?lang=en&region=us",
  "count": 323,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 162,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/plays/401871412000000520?lang=en&region=us",
      "id": "401871412000000520",
      "sequenceNumber": "2",
      "type": {
        "id": "518",
        "text": "Period Start",
        "abbreviation": "period-start"
      },
      "text": "Start of 1st Period",
      "shortText": "Start of 1st Period",
      "alternativeText": "Start of 1st Period",
      "shortAlternativeText": "Start of 1st Period",
      "awayScore": 0,
      "homeScore": 0,
      "period": {
        "number": 1,
        "displayValue": "1st"
      },
      "clock": {
        "value": 0.0,
        "displayValue": "0:00"
      },
      "scoringPlay": false,
      "priority": false
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/plays/401871412000000510?lang=en&region=us",
      "id": "401871412000000510",
      "sequenceNumber": "3",
      "type": {
        "id": "502",
        "text": "Face Off",
        "abbreviation": "faceoff"
      },
      "text": "Luke Glendening faceoff won against Jordan Staal",
      "shortText": "Luke Glendening faceoff won against Jordan Staal",
      "alternativeText": "Luke Glendening faceoff won against Jordan Staal",
      "shortAlternativeText": "Luke Glendening faceoff won against Jordan Staal",
      "awayScore": 0,
      "homeScore": 0,
      "period": {
        "number": 1,
        "displayValue": "1st"
      },
      "clock": {
        "value": 0.0,
        "displayValue": "0:00"
      },
      "scoringPlay": false,
      "priority": false
    }
  ]
}
```
