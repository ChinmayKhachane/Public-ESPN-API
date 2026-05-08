# Competition Subresources

Verified with `league=nfl`, `event=401772988`, and `competition=401772988` on 2026-05-08.

---

## Broadcasts

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/broadcasts`

```json
{
  "count": 2,
  "pageIndex": 1,
  "pageSize": 25,
  "items": [
    {
      "type": {
        "id": "1",
        "shortName": "TV",
        "longName": "Television",
        "slug": "tv"
      },
      "station": "NBC",
      "slug": "nbc",
      "market": {
        "id": "1",
        "type": "National"
      },
      "media": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/media/379?lang=en&region=us",
        "id": "379",
        "name": "NBC"
      }
    }
  ]
}
```

---

## Competitor

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/competitors/{competitor}`

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17?lang=en&region=us",
  "id": "17",
  "uid": "s:20~l:28~t:17",
  "type": "team",
  "order": 0,
  "homeAway": "home",
  "winner": false,
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2025/teams/17?lang=en&region=us"
  },
  "score": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17/score?lang=en&region=us"
  },
  "statistics": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/competitors/17/statistics/0?lang=en&region=us"
  }
}
```

---

## Odds

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/odds`

```json
{
  "count": 1,
  "pageIndex": 1,
  "pageSize": 25,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/odds/100?lang=en&region=us",
      "provider": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/providers/100?lang=en&region=us",
        "id": "100",
        "name": "Draft Kings",
        "priority": 1
      },
      "details": "SEA -4.5",
      "overUnder": 45.5,
      "spread": 4.5
    }
  ]
}
```

---

## Officials

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/officials`

```json
{
  "count": 7,
  "pageIndex": 1,
  "pageSize": 25,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/events/401772988/competitions/401772988/officials/17655?lang=en&region=us",
      "id": "17655",
      "firstName": "Dana",
      "lastName": "McKenzie",
      "fullName": "Dana McKenzie",
      "position": {
        "name": "Down Judge",
        "id": "112"
      }
    }
  ]
}
```

---

## Play Personnel

`https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/events/{event}/competitions/{competition}/plays/{play}/personnel`

Notes:
- Using the last play ref from the competition situation (`play=4017729884651`), NFL currently returned HTTP 500.

```json
{
  "error": {
    "message": "application error",
    "code": 500
  }
}
```
