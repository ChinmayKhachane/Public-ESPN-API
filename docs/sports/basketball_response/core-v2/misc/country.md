# Countries

## https://sports.core.api.espn.com/v2/sports/basketball/leagues/{league}/countries

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "count": 258,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 129,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/countries/1?lang=en&region=us",
      "id": "1",
      "name": "USA",
      "abbreviation": "USA",
      "slug": "usa",
      "athletes": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/countries/1/athletes?lang=en&region=us"
      },
      "flag": {
        "href": "https://a.espncdn.com/i/teamlogos/countries/500/usa.png",
        "alt": "USA",
        "rel": []
      }
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/countries/44?lang=en&region=us",
      "id": "44",
      "name": "West Germany",
      "abbreviation": "FRG",
      "slug": "west-germany",
      "athletes": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/basketball/leagues/nba/countries/44/athletes?lang=en&region=us"
      },
      "flag": {
        "href": "https://a.espncdn.com/i/teamlogos/countries/500/frg.png",
        "alt": "West Germany",
        "rel": []
      }
    }
  ]
}
```
