# MLB Event By ID

## https://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256?lang=en&region=us",
  "id": "401815256",
  "uid": "s:1~l:10~e:401815256",
  "name": "Houston Astros at Cincinnati Reds",
  "competitions": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256?lang...",
      "id": "401815256",
      "uid": "s:1~l:10~e:401815256~c:401815256",
      "guid": "af84a724-273f-3ff8-a416-075e02c117d5",
      "competitors": [
        {
          "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp...",
          "id": "17",
          "uid": "s:1~l:10~t:17",
          "team": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/17?lang=en&region=us"
          },
          "statistics": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp..."
          },
          "type": "team",
          "order": 0,
          "homeAway": "home",
          "winner": false,
          "score": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp..."
          }
        },
        {
          "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp...",
          "id": "18",
          "uid": "s:1~l:10~t:18",
          "team": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/teams/18?lang=en&region=us"
          },
          "statistics": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp..."
          },
          "type": "team",
          "order": 1,
          "homeAway": "away",
          "winner": true,
          "score": {
            "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/comp..."
          }
        }
      ],
      "status": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/events/401815256/competitions/401815256/stat..."
      },
      "date": "2026-05-08T22:10Z",
      "timeOfDay": "DAY",
      "attendance": 24347,
      "type": {
        "id": "1",
        "abbreviation": "STD",
        "slug": "standard",
        "text": "Standard",
        "type": "standard"
      }
    }
  ],
  "season": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026?lang=en&region=us"
  },
  "date": "2026-05-08T22:10Z",
  "shortName": "HOU @ CIN",
  "seasonType": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/types/2?lang=en&region=us"
  },
  "week": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/seasons/2026/types/2/weeks/19?lang=en&region=us"
  },
  "timeValid": true,
  "links": [
    {
      "language": "en-US",
      "rel": [
        "summary",
        "desktop"
      ],
      "href": "https://www.espn.com/mlb/game/_/gameId/401815256/astros-reds",
      "text": "Gamecast",
      "shortText": "Summary",
      "isExternal": false,
      "isPremium": false
    },
    {
      "language": "en-US",
      "rel": [
        "summary",
        "sportscenter"
      ],
      "href": "sportscenter://x-callback-url/showGame?sportName=baseball&leagueAbbrev=mlb&gameId=401815256",
      "text": "Gamecast",
      "shortText": "Summary",
      "isExternal": false,
      "isPremium": false
    }
  ],
  "venues": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb/venues/83?lang=en&region=us"
    }
  ],
  "league": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/baseball/leagues/mlb?lang=en&region=us"
  }
}
```
