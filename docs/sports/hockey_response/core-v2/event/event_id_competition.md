# Event Competition

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/events/{event}/competitions/{competition}

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412`

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412?lang=en&region=us",
  "id": "401871412",
  "guid": "85f1df1b-499b-35d6-9a30-7adc9d7c5140",
  "uid": "s:70~l:90~e:401871412~c:401871412",
  "date": "2026-05-08T00:00Z",
  "attendance": 19970,
  "type": {
    "id": "15",
    "text": "Quarterfinal",
    "abbreviation": "QTR",
    "slug": "quarterfinal",
    "type": "quarterfinal"
  },
  "necessary": true,
  "timeValid": true,
  "neutralSite": false,
  "previewAvailable": true,
  "recapAvailable": true,
  "boxscoreAvailable": true,
  "lineupAvailable": false,
  "gamecastAvailable": true,
  "playByPlayAvailable": true,
  "conversationAvailable": true,
  "commentaryAvailable": false
}
```
