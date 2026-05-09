# Competition Odds

## https://sports.core.api.espn.com/v2/sports/hockey/leagues/{league}/events/{event}/competitions/{competition}/odds

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/odds?limit=2`

## Example Response

```json
{
  "count": 1,
  "pageIndex": 1,
  "pageSize": 2,
  "pageCount": 1,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/events/401871412/competitions/401871412/odds/100?lang=en&region=us",
      "provider": {
        "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/providers/100?lang=en&region=us",
        "id": "100",
        "name": "Draft Kings",
        "priority": 1
      },
      "details": "CAR -166",
      "overUnder": 5.5,
      "spread": 1.5,
      "overOdds": 105.0,
      "underOdds": -125.0,
      "awayTeamOdds": {
        "favorite": true,
        "underdog": false,
        "moneyLine": -166,
        "spreadOdds": 154.0,
        "open": {
          "favorite": true,
          "pointSpread": {},
          "spread": {},
          "moneyLine": {}
        },
        "close": {
          "pointSpread": {},
          "spread": {},
          "moneyLine": {}
        },
        "current": {
          "pointSpread": {},
          "spread": {},
          "moneyLine": {}
        },
        "team": {
          "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/teams/7?lang=en&region=us"
        }
      },
      "homeTeamOdds": {
        "favorite": false,
        "underdog": true,
        "moneyLine": 140,
        "spreadOdds": -185.0,
        "open": {
          "favorite": false,
          "pointSpread": {},
          "spread": {},
          "moneyLine": {}
        },
        "close": {
          "pointSpread": {},
          "spread": {},
          "moneyLine": {}
        },
        "current": {
          "pointSpread": {},
          "spread": {},
          "moneyLine": {}
        },
        "team": {
          "$ref": "http://sports.core.api.espn.com/v2/sports/hockey/leagues/nhl/seasons/2026/teams/15?lang=en&region=us"
        }
      },
      "links": [
        {
          "language": "en-US",
          "rel": [],
          "href": "https://sportsbook.draftkings.com/gateway?s=__s__&wpcid=__wpcid__&wpsrc=413&wpcn=ESPN&wpscn=Widget&wpcrn=BetSlipDeepLink&wpscid=__wpscid__&wpcrid=xx&preurl=https%3A%2F%2Fsportsbook.draftkings.com%2Fevent%2F34104080%3Foutcomes%3D0ML84653447_1",
          "text": "Home Bet",
          "shortText": "Home Bet",
          "isExternal": true,
          "isPremium": false
        },
        {
          "language": "en-US",
          "rel": [],
          "href": "https://sportsbook.draftkings.com/gateway?s=__s__&wpcid=__wpcid__&wpsrc=413&wpcn=ESPN&wpscn=Widget&wpcrn=BetSlipDeepLink&wpscid=__wpscid__&wpcrid=xx&preurl=https%3A%2F%2Fsportsbook.draftkings.com%2Fevent%2F34104080%3Foutcomes%3D0ML84653447_3",
          "text": "Away Bet",
          "shortText": "Away Bet",
          "isExternal": true,
          "isPremium": false
        }
      ],
      "moneylineWinner": false,
      "spreadWinner": false,
      "open": {
        "over": {
          "value": 2.0,
          "displayValue": "1/1",
          "alternateDisplayValue": "+100",
          "decimal": 2.0,
          "fraction": "1/1",
          "american": "+100"
        },
        "under": {
          "value": 1.83,
          "displayValue": "5/6",
          "alternateDisplayValue": "-120",
          "decimal": 1.83,
          "fraction": "5/6",
          "american": "-120"
        },
        "total": {
          "alternateDisplayValue": "5.5",
          "american": "5.5"
        }
      },
      "close": {
        "over": {
          "value": 2.05,
          "displayValue": "21/20",
          "alternateDisplayValue": "+105",
          "decimal": 2.05,
          "fraction": "21/20",
          "american": "+105"
        },
        "under": {
          "value": 1.8,
          "displayValue": "4/5",
          "alternateDisplayValue": "-125",
          "decimal": 1.8,
          "fraction": "4/5",
          "american": "-125"
        },
        "total": {
          "alternateDisplayValue": "5.5",
          "american": "5.5"
        }
      }
    }
  ]
}
```
