# MLB Team Depth Charts

## https://site.api.espn.com/apis/site/v2/sports/baseball/mlb/teams/17/depthcharts

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "team": {
    "id": "17",
    "name": "Reds",
    "displayName": "Cincinnati Reds",
    "abbreviation": "CIN",
    "location": "Cincinnati",
    "clubhouse": "https://www.espn.com/mlb/team/_/name/cin/cincinnati-reds",
    "color": "c6011f",
    "logo": "https://a.espncdn.com/i/teamlogos/mlb/500/cin.png",
    "recordSummary": "20-19",
    "seasonSummary": "2026"
  },
  "status": "success",
  "season": {
    "name": "Regular Season",
    "year": 2026,
    "type": 2
  },
  "timestamp": "2026-05-09T03:52:05Z",
  "depthchart": [
    {
      "id": "1",
      "name": "Depth Chart",
      "positions": {
        "rp": {
          "athletes": [
            {},
            {}
          ],
          "position": {
            "id": "0",
            "name": "Relief Pitcher",
            "displayName": "Relief Pitcher",
            "abbreviation": "RP",
            "leaf": true,
            "parent": {}
          }
        },
        "p": {
          "athletes": [
            {},
            {}
          ],
          "position": {
            "id": "1",
            "name": "Pitcher",
            "displayName": "Pitcher",
            "abbreviation": "P",
            "leaf": false
          }
        },
        "c": {
          "athletes": [
            {},
            {}
          ],
          "position": {
            "id": "2",
            "name": "Catcher",
            "displayName": "Catcher",
            "abbreviation": "C",
            "leaf": true,
            "parent": {}
          }
        },
        "1b": {
          "athletes": [
            {},
            {}
          ],
          "position": {
            "id": "3",
            "name": "First Base",
            "displayName": "First Baseman",
            "abbreviation": "1B",
            "leaf": true,
            "parent": {}
          }
        },
        "2b": {
          "athletes": [
            {},
            {}
          ],
          "position": {
            "id": "4",
            "name": "Second Base",
            "displayName": "Second Baseman",
            "abbreviation": "2B",
            "leaf": true,
            "parent": {}
          }
        },
        "3b": {
          "athletes": [
            {},
            {}
          ],
          "position": {
            "id": "5",
            "name": "Third Base",
            "displayName": "Third Baseman",
            "abbreviation": "3B",
            "leaf": true,
            "parent": {}
          }
        },
        "ss": {
          "athletes": [
            {},
            {}
          ],
          "position": {
            "id": "6",
            "name": "Shortstop",
            "displayName": "Shortstop",
            "abbreviation": "SS",
            "leaf": true,
            "parent": {}
          }
        },
        "lf": {
          "athletes": [
            {},
            {}
          ],
          "position": {
            "id": "7",
            "name": "Left Field",
            "displayName": "Left Fielder",
            "abbreviation": "LF",
            "leaf": true,
            "parent": {}
          }
        },
        "cf": {
          "athletes": [
            {},
            {}
          ],
          "position": {
            "id": "8",
            "name": "Center Field",
            "displayName": "Center Fielder",
            "abbreviation": "CF",
            "leaf": true,
            "parent": {}
          }
        },
        "rf": {
          "athletes": [
            {},
            {}
          ],
          "position": {
            "id": "9",
            "name": "Right Field",
            "displayName": "Right Fielder",
            "abbreviation": "RF",
            "leaf": true,
            "parent": {}
          }
        }
      }
    }
  ]
}
```
