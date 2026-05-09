# MLB League Injuries

## https://site.api.espn.com/apis/site/v2/sports/baseball/mlb/injuries

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "status": "success",
  "season": {
    "name": "Regular Season",
    "displayName": "2026",
    "year": 2026,
    "type": 2
  },
  "timestamp": "2026-05-09T03:57:50Z",
  "injuries": [
    {
      "id": "29",
      "displayName": "Arizona Diamondbacks",
      "injuries": [
        {
          "id": "-191657",
          "status": "Day-To-Day",
          "longComment": "It's a rare day off for Marte, who had made 12 consecutive starts. Ildemaro Vargas will cover second base F...",
          "shortComment": "Marte is not in the lineup for Friday's game against the Cubs.",
          "date": "2026-05-09T00:22Z",
          "athlete": {
            "displayName": "Ketel Marte",
            "team": {},
            "status": {},
            "firstName": "Ketel",
            "lastName": "Marte",
            "shortName": "K. Marte",
            "links": [],
            "headshot": {},
            "position": {},
            "notes": {}
          },
          "source": {
            "id": "1",
            "state": "basic"
          },
          "type": {
            "id": "6",
            "name": "INJURY_STATUS_DAYTODAY",
            "abbreviation": "DD"
          },
          "details": {
            "fantasyStatus": {},
            "type": "Illness",
            "returnDate": "2026-05-09"
          }
        },
        {
          "id": "-191495",
          "status": "60-Day-IL",
          "longComment": "Burnes has been throwing occasional bullpen sessions since spring training -- Tuesday's session was his fou...",
          "shortComment": "Burnes (elbow) threw a bullpen session Tuesday, Jose M. Romero of the Arizona Republic reports.",
          "date": "2026-05-08T11:02Z",
          "athlete": {
            "displayName": "Corbin Burnes",
            "team": {},
            "status": {},
            "firstName": "Corbin",
            "lastName": "Burnes",
            "shortName": "C. Burnes",
            "links": [],
            "headshot": {},
            "position": {},
            "notes": {}
          },
          "source": {
            "id": "1",
            "state": "basic"
          },
          "type": {
            "id": "11",
            "name": "INJURY_STATUS_60DAYIL",
            "abbreviation": "IL60"
          },
          "details": {
            "fantasyStatus": {},
            "type": "Elbow",
            "detail": "Surgery",
            "side": "Right",
            "returnDate": "2026-07-17"
          }
        }
      ]
    },
    {
      "id": "11",
      "displayName": "Athletics",
      "injuries": [
        {
          "id": "-191507",
          "status": "10-Day-IL",
          "longComment": "Clarke has been out for the past two weeks with a bone bruise in his right foot but is ready to start rampi...",
          "shortComment": "Clarke (foot) has begun a running and hitting progression, Martin Gallegos of MLB.com reports.",
          "date": "2026-05-08T11:02Z",
          "athlete": {
            "displayName": "Denzel Clarke",
            "team": {},
            "status": {},
            "firstName": "Denzel",
            "lastName": "Clarke",
            "shortName": "D. Clarke",
            "links": [],
            "headshot": {},
            "position": {},
            "notes": {}
          },
          "source": {
            "id": "1",
            "state": "basic"
          },
          "type": {
            "id": "14",
            "name": "INJURY_STATUS_10DAYIL",
            "abbreviation": "IL10"
          },
          "details": {
            "fantasyStatus": {},
            "type": "Foot",
            "detail": "Bruise",
            "side": "Right",
            "returnDate": "2026-05-18"
          }
        },
        {
          "id": "-191280",
          "status": "60-Day-IL",
          "longComment": "Hoglund has been out all season with knee and back injuries, so the injection for his hip could be related ...",
          "shortComment": "Hoglund (knee/back) was given a cortisone injection in his left hip May 1, Martin Gallegos of MLB.com reports.",
          "date": "2026-05-08T11:02Z",
          "athlete": {
            "displayName": "Gunnar Hoglund",
            "team": {},
            "status": {},
            "firstName": "Gunnar",
            "lastName": "Hoglund",
            "shortName": "G. Hoglund",
            "links": [],
            "headshot": {},
            "position": {},
            "notes": {}
          },
          "source": {
            "id": "1",
            "state": "basic"
          },
          "type": {
            "id": "11",
            "name": "INJURY_STATUS_60DAYIL",
            "abbreviation": "IL60"
          },
          "details": {
            "fantasyStatus": {},
            "type": "Knee",
            "detail": "Strain",
            "side": "Not Specified",
            "returnDate": "2026-06-02"
          }
        }
      ]
    }
  ]
}
```
