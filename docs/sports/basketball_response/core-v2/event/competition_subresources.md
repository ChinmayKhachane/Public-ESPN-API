# Competition Broadcasts

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/events/{event}/competitions/{competition}/broadcasts

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.
- See sibling files for competitor, odds, officials, and personnel shapes.

## Example Response

```json
{
  "count": 2,
  "pageIndex": 1,
  "pageSize": 25,
  "pageCount": 1,
  "items": [
    {
      "slug": "prime-video",
      "type": {
        "id": "4",
        "slug": "streaming",
        "shortName": "Streaming",
        "longName": "Streaming"
      },
      "channel": 763,
      "station": "Prime Video",
      "priority": 1,
      "market": {
        "id": "1",
        "type": "National"
      },
      "media": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/media/763?lang=en&region=us",
        "id": "763",
        "name": "Prime Video",
        "slug": "prime-video",
        "callLetters": "Prime Video",
        "shortName": "Prime Video",
        "logos": []
      },
      "lang": "en",
      "region": "us",
      "competition": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161?lang=en&region=us"
      },
      "partnered": false
    },
    {
      "slug": "13715732",
      "type": {
        "id": "5",
        "slug": "radio",
        "shortName": "Radio",
        "longName": "Radio"
      },
      "channel": 133,
      "url": "https://cerebro.espn.com/cerebro/search/airing/a1302900660",
      "station": "13715732",
      "stationKey": "espn",
      "priority": 1,
      "market": {
        "id": "1",
        "type": "National"
      },
      "media": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/media/133?lang=en&region=us",
        "id": "133",
        "name": "ESPN Radio",
        "slug": "espn-radio",
        "callLetters": "ERADM",
        "shortName": "ERADM",
        "logos": [],
        "group": {}
      },
      "lang": "en",
      "region": "us",
      "competition": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/events/401871161/competitions/401871161?lang=en&region=us"
      }
    }
  ]
}
```
