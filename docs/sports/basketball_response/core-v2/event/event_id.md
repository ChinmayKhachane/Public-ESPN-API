# Event Detail

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/events/{event}

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161?lang=en&region=us",
  "id": "401871161",
  "uid": "s:40~l:46~e:401871161",
  "name": "New York Knicks at Philadelphia 76ers",
  "season": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026?lang=en&region=us"
  },
  "competitions": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161?lang=en&region=us",
      "id": "401871161",
      "uid": "s:40~l:46~e:401871161~c:401871161",
      "guid": "55df9197-9744-326e-b9f2-ef8e591d40d5",
      "type": {
        "id": "15",
        "abbreviation": "QTR",
        "slug": "quarterfinal",
        "type": "quarterfinal",
        "text": "Quarterfinal"
      },
      "status": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/status?lang=en&region=us"
      },
      "competitors": [
        {},
        {}
      ],
      "venue": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/venues/1845?lang=en&region=us",
        "id": "1845",
        "guid": "5c91d4ea-46af-3b85-a87d-54422ed1e8c3",
        "fullName": "Xfinity Mobile Arena",
        "shortName": "Wells Fargo Center",
        "address": {},
        "grass": false,
        "indoor": true,
        "images": []
      },
      "date": "2026-05-08T23:00Z",
      "timeValid": true,
      "broadcasts": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/broadcasts?lang=en&region=us"
      },
      "officials": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161/officials?lang=en&region=us"
      }
    }
  ],
  "league": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba?lang=en&region=us"
  },
  "date": "2026-05-08T23:00Z",
  "timeValid": true,
  "shortName": "NY @ PHI",
  "seasonType": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/seasons/2026/types/3?lang=en&region=us"
  },
  "links": [
    {
      "text": "Gamecast",
      "shortText": "Summary",
      "language": "en-US",
      "rel": [
        "summary",
        "desktop"
      ],
      "href": "https://www.espn.com/nba/game/_/gameId/401871161/knicks-76ers",
      "isExternal": false,
      "isPremium": false
    },
    {
      "text": "Gamecast",
      "shortText": "Summary",
      "language": "en-US",
      "rel": [
        "summary",
        "sportscenter"
      ],
      "href": "sportscenter://x-callback-url/showGame?sportName=basketball&leagueAbbrev=nba&gameId=401871161",
      "isExternal": false,
      "isPremium": false
    }
  ]
}
```
