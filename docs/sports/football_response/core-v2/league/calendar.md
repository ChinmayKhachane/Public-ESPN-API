# League Calendar

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/calendar

Notes:
- Verified with `league=nfl` on 2026-05-08.
- This returns a small collection of calendar buckets, not a single expanded calendar object.

## Optional Parameters

| Param | Type | Description |
| --- | --- | --- |
| `dates` | `string` | Filter by date range when supported by downstream calendar refs |
| `weeks` | `string` | Filter by week when supported by downstream calendar refs |
| `seasontype` | `int` | Filter by season type when supported by downstream calendar refs |

## Example Response

```json
{
  "count": 4,
  "pageIndex": 1,
  "pageSize": 25,
  "pageCount": 1,
  "items": [
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/calendar/ondays?lang=en&region=us",
      "type": "list",
      "startDate": "2025-07-31T07:00Z",
      "endDate": "2026-02-12T07:59Z",
      "eventDate": {
        "type": "ondays",
        "dates": [
          "2025-07-31T07:00Z",
          "2025-08-07T07:00Z",
          "2025-08-08T07:00Z",
          "2025-08-09T07:00Z",
          "2025-08-10T07:00Z",
          "2025-08-15T07:00Z",
          "2025-08-16T07:00Z",
          "2025-08-17T07:00Z",
          "2025-08-18T07:00Z",
          "2025-08-21T07:00Z",
          "2025-08-22T07:00Z",
          "2025-08-23T07:00Z",
          "2025-09-04T07:00Z",
          "2025-09-05T07:00Z",
          "2025-09-07T07:00Z",
          "2025-09-08T07:00Z",
          "2025-09-11T07:00Z",
          "2025-09-14T07:00Z",
          "2025-09-15T07:00Z",
          "2025-09-18T07:00Z",
          "2025-09-21T07:00Z",
          "2025-09-22T07:00Z",
          "2025-09-25T07:00Z",
          "2025-09-28T07:00Z",
          "2025-09-29T07:00Z",
          "2025-10-02T07:00Z",
          "2025-10-05T07:00Z",
          "2025-10-06T07:00Z",
          "2025-10-09T07:00Z",
          "2025-10-12T07:00Z",
          "2025-10-13T07:00Z",
          "2025-10-16T07:00Z",
          "2025-10-19T07:00Z",
          "2025-10-20T07:00Z",
          "2025-10-23T07:00Z",
          "2025-10-26T07:00Z",
          "2025-10-27T07:00Z",
          "2025-10-30T07:00Z",
          "2025-11-02T07:00Z",
          "2025-11-03T08:00Z",
          "2025-11-06T08:00Z",
          "2025-11-09T08:00Z",
          "2025-11-10T08:00Z",
          "2025-11-13T08:00Z",
          "2025-11-16T08:00Z",
          "2025-11-17T08:00Z",
          "2025-11-20T08:00Z",
          "2025-11-23T08:00Z",
          "2025-11-24T08:00Z",
          "2025-11-27T08:00Z",
          "2025-11-28T08:00Z",
          "2025-11-30T08:00Z",
          "2025-12-01T08:00Z",
          "2025-12-04T08:00Z",
          "2025-12-07T08:00Z",
          "2025-12-08T08:00Z",
          "2025-12-11T08:00Z",
          "2025-12-14T08:00Z",
          "2025-12-15T08:00Z",
          "2025-12-18T08:00Z",
          "2025-12-20T08:00Z",
          "2025-12-21T08:00Z",
          "2025-12-22T08:00Z",
          "2025-12-25T08:00Z",
          "2025-12-27T08:00Z",
          "2025-12-28T08:00Z",
          "2025-12-29T08:00Z",
          "2026-01-03T08:00Z",
          "2026-01-04T08:00Z",
          "2026-01-10T08:00Z",
          "2026-01-11T08:00Z",
          "2026-01-12T08:00Z",
          "2026-01-17T08:00Z",
          "2026-01-18T08:00Z",
          "2026-01-25T08:00Z",
          "2026-02-03T08:00Z",
          "2026-02-08T08:00Z"
        ]
      },
      "sections": [
        {
          "label": "Preseason",
          "value": "1",
          "startDate": "2025-07-31T07:00Z",
          "endDate": "2025-09-04T06:59Z",
          "entries": [
            {
              "label": "Hall of Fame Weekend",
              "alternateLabel": "HOF",
              "detail": "Jul 31-Aug 6",
              "value": "1",
              "startDate": "2025-07-31T07:00Z",
              "endDate": "2025-08-07T06:59Z"
            },
            {
              "label": "Preseason Week 1",
              "alternateLabel": "Pre Wk 1",
              "detail": "Aug 7-13",
              "value": "2",
              "startDate": "2025-08-07T07:00Z",
              "endDate": "2025-08-14T06:59Z"
            },
            {
              "label": "Preseason Week 2",
              "alternateLabel": "Pre Wk 2",
              "detail": "Aug 14-20",
              "value": "3",
              "startDate": "2025-08-14T07:00Z",
              "endDate": "2025-08-21T06:59Z"
            },
            {
              "label": "Preseason Week 3",
              "alternateLabel": "Pre Wk 3",
              "detail": "Aug 21-Sep 3",
              "value": "4",
              "startDate": "2025-08-21T07:00Z",
              "endDate": "2025-09-04T06:59Z"
            }
          ],
          "seasonType": {}
        },
        {
          "label": "Regular Season",
          "value": "2",
          "startDate": "2025-09-04T07:00Z",
          "endDate": "2026-01-07T07:59Z",
          "entries": [
            {
              "label": "Week 1",
              "alternateLabel": "Week 1",
              "detail": "Sep 4-9",
              "value": "1",
              "startDate": "2025-09-04T07:00Z",
              "endDate": "2025-09-10T06:59Z"
            },
            {
              "label": "Week 2",
              "alternateLabel": "Week 2",
              "detail": "Sep 10-16",
              "value": "2",
              "startDate": "2025-09-10T07:00Z",
              "endDate": "2025-09-17T06:59Z"
            },
            {
              "label": "Week 3",
              "alternateLabel": "Week 3",
              "detail": "Sep 17-23",
              "value": "3",
              "startDate": "2025-09-17T07:00Z",
              "endDate": "2025-09-24T06:59Z"
            },
            {
              "label": "Week 4",
              "alternateLabel": "Week 4",
              "detail": "Sep 24-30",
              "value": "4",
              "startDate": "2025-09-24T07:00Z",
              "endDate": "2025-10-01T06:59Z"
            },
            {
              "label": "Week 5",
              "alternateLabel": "Week 5",
              "detail": "Oct 1-7",
              "value": "5",
              "startDate": "2025-10-01T07:00Z",
              "endDate": "2025-10-08T06:59Z"
            },
            {
              "label": "Week 6",
              "alternateLabel": "Week 6",
              "detail": "Oct 8-14",
              "value": "6",
              "startDate": "2025-10-08T07:00Z",
              "endDate": "2025-10-15T06:59Z"
            },
            {
              "label": "Week 7",
              "alternateLabel": "Week 7",
              "detail": "Oct 15-21",
              "value": "7",
              "startDate": "2025-10-15T07:00Z",
              "endDate": "2025-10-22T06:59Z"
            },
            {
              "label": "Week 8",
              "alternateLabel": "Week 8",
              "detail": "Oct 22-28",
              "value": "8",
              "startDate": "2025-10-22T07:00Z",
              "endDate": "2025-10-29T06:59Z"
            },
            {
              "label": "Week 9",
              "alternateLabel": "Week 9",
              "detail": "Oct 29-Nov 4",
              "value": "9",
              "startDate": "2025-10-29T07:00Z",
              "endDate": "2025-11-05T07:59Z"
            },
            {
              "label": "Week 10",
              "alternateLabel": "Week 10",
              "detail": "Nov 5-11",
              "value": "10",
              "startDate": "2025-11-05T08:00Z",
              "endDate": "2025-11-12T07:59Z"
            },
            {
              "label": "Week 11",
              "alternateLabel": "Week 11",
              "detail": "Nov 12-18",
              "value": "11",
              "startDate": "2025-11-12T08:00Z",
              "endDate": "2025-11-19T07:59Z"
            },
            {
              "label": "Week 12",
              "alternateLabel": "Week 12",
              "detail": "Nov 19-25",
              "value": "12",
              "startDate": "2025-11-19T08:00Z",
              "endDate": "2025-11-26T07:59Z"
            },
            {
              "label": "Week 13",
              "alternateLabel": "Week 13",
              "detail": "Nov 26-Dec 2",
              "value": "13",
              "startDate": "2025-11-26T08:00Z",
              "endDate": "2025-12-03T07:59Z"
            },
            {
              "label": "Week 14",
              "alternateLabel": "Week 14",
              "detail": "Dec 3-9",
              "value": "14",
              "startDate": "2025-12-03T08:00Z",
              "endDate": "2025-12-10T07:59Z"
            },
            {
              "label": "Week 15",
              "alternateLabel": "Week 15",
              "detail": "Dec 10-16",
              "value": "15",
              "startDate": "2025-12-10T08:00Z",
              "endDate": "2025-12-17T07:59Z"
            },
            {
              "label": "Week 16",
              "alternateLabel": "Week 16",
              "detail": "Dec 17-23",
              "value": "16",
              "startDate": "2025-12-17T08:00Z",
              "endDate": "2025-12-24T07:59Z"
            },
            {
              "label": "Week 17",
              "alternateLabel": "Week 17",
              "detail": "Dec 24-30",
              "value": "17",
              "startDate": "2025-12-24T08:00Z",
              "endDate": "2025-12-31T07:59Z"
            },
            {
              "label": "Week 18",
              "alternateLabel": "Week 18",
              "detail": "Dec 31-Jan 6",
              "value": "18",
              "startDate": "2025-12-31T08:00Z",
              "endDate": "2026-01-07T07:59Z"
            }
          ],
          "seasonType": {}
        },
        {
          "label": "Postseason",
          "value": "3",
          "startDate": "2026-01-07T08:00Z",
          "endDate": "2026-02-12T07:59Z",
          "entries": [
            {
              "label": "Wild Card",
              "alternateLabel": "Wild Card",
              "detail": "Jan 7-13",
              "value": "1",
              "startDate": "2026-01-07T08:00Z",
              "endDate": "2026-01-14T07:59Z"
            },
            {
              "label": "Divisional Round",
              "alternateLabel": "Div Rd",
              "detail": "Jan 14-20",
              "value": "2",
              "startDate": "2026-01-14T08:00Z",
              "endDate": "2026-01-21T07:59Z"
            },
            {
              "label": "Conference Championship",
              "alternateLabel": "Conf Champ",
              "detail": "Jan 21-27",
              "value": "3",
              "startDate": "2026-01-21T08:00Z",
              "endDate": "2026-01-28T07:59Z"
            },
            {
              "label": "Pro Bowl",
              "alternateLabel": "Pro Bowl",
              "detail": "Jan 28-Feb 3",
              "value": "4",
              "startDate": "2026-01-28T08:00Z",
              "endDate": "2026-02-04T07:59Z"
            },
            {
              "label": "Super Bowl",
              "alternateLabel": "Super Bowl",
              "detail": "Feb 4-11",
              "value": "5",
              "startDate": "2026-02-04T08:00Z",
              "endDate": "2026-02-12T07:59Z"
            }
          ],
          "seasonType": {}
        },
        {
          "label": "Off Season",
          "value": "4",
          "startDate": "2026-02-12T08:00Z",
          "endDate": "2026-08-06T06:59Z",
          "entries": [
            {
              "label": "Week 1",
              "alternateLabel": "Week 1",
              "detail": "Feb 12-Jul 31",
              "value": "1",
              "startDate": "2026-02-12T08:00Z",
              "endDate": "2026-08-01T06:59Z"
            }
          ],
          "seasonType": {}
        }
      ],
      "season": {}
    },
    {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/calendar/offdays?lang=en&region=us",
      "type": "list",
      "startDate": "2025-07-31T07:00Z",
      "endDate": "2026-02-12T07:59Z",
      "eventDate": {
        "type": "offdays",
        "dates": [
          "2025-08-01T07:00Z",
          "2025-08-02T07:00Z",
          "2025-08-03T07:00Z",
          "2025-08-04T07:00Z",
          "2025-08-05T07:00Z",
          "2025-08-06T07:00Z",
          "2025-08-11T07:00Z",
          "2025-08-12T07:00Z",
          "2025-08-13T07:00Z",
          "2025-08-14T07:00Z",
          "2025-08-19T07:00Z",
          "2025-08-20T07:00Z",
          "2025-08-24T07:00Z",
          "2025-08-25T07:00Z",
          "2025-08-26T07:00Z",
          "2025-08-27T07:00Z",
          "2025-08-28T07:00Z",
          "2025-08-29T07:00Z",
          "2025-08-30T07:00Z",
          "2025-08-31T07:00Z",
          "2025-09-01T07:00Z",
          "2025-09-02T07:00Z",
          "2025-09-03T07:00Z",
          "2025-09-06T07:00Z",
          "2025-09-09T07:00Z",
          "2025-09-10T07:00Z",
          "2025-09-12T07:00Z",
          "2025-09-13T07:00Z",
          "2025-09-16T07:00Z",
          "2025-09-17T07:00Z",
          "2025-09-19T07:00Z",
          "2025-09-20T07:00Z",
          "2025-09-23T07:00Z",
          "2025-09-24T07:00Z",
          "2025-09-26T07:00Z",
          "2025-09-27T07:00Z",
          "2025-09-30T07:00Z",
          "2025-10-01T07:00Z",
          "2025-10-03T07:00Z",
          "2025-10-04T07:00Z",
          "2025-10-07T07:00Z",
          "2025-10-08T07:00Z",
          "2025-10-10T07:00Z",
          "2025-10-11T07:00Z",
          "2025-10-14T07:00Z",
          "2025-10-15T07:00Z",
          "2025-10-17T07:00Z",
          "2025-10-18T07:00Z",
          "2025-10-21T07:00Z",
          "2025-10-22T07:00Z",
          "2025-10-24T07:00Z",
          "2025-10-25T07:00Z",
          "2025-10-28T07:00Z",
          "2025-10-29T07:00Z",
          "2025-10-31T07:00Z",
          "2025-11-01T07:00Z",
          "2025-11-04T08:00Z",
          "2025-11-05T08:00Z",
          "2025-11-07T08:00Z",
          "2025-11-08T08:00Z",
          "2025-11-11T08:00Z",
          "2025-11-12T08:00Z",
          "2025-11-14T08:00Z",
          "2025-11-15T08:00Z",
          "2025-11-18T08:00Z",
          "2025-11-19T08:00Z",
          "2025-11-21T08:00Z",
          "2025-11-22T08:00Z",
          "2025-11-25T08:00Z",
          "2025-11-26T08:00Z",
          "2025-11-29T08:00Z",
          "2025-12-02T08:00Z",
          "2025-12-03T08:00Z",
          "2025-12-05T08:00Z",
          "2025-12-06T08:00Z",
          "2025-12-09T08:00Z",
          "2025-12-10T08:00Z",
          "2025-12-12T08:00Z",
          "2025-12-13T08:00Z",
          "2025-12-16T08:00Z",
          "2025-12-17T08:00Z",
          "2025-12-19T08:00Z",
          "2025-12-23T08:00Z",
          "2025-12-24T08:00Z",
          "2025-12-26T08:00Z",
          "2025-12-30T08:00Z",
          "2025-12-31T08:00Z",
          "2026-01-01T08:00Z",
          "2026-01-02T08:00Z",
          "2026-01-05T08:00Z",
          "2026-01-06T08:00Z",
          "2026-01-07T08:00Z",
          "2026-01-08T08:00Z",
          "2026-01-09T08:00Z",
          "2026-01-13T08:00Z",
          "2026-01-14T08:00Z",
          "2026-01-15T08:00Z",
          "2026-01-16T08:00Z",
          "2026-01-19T08:00Z",
          "2026-01-20T08:00Z",
          "2026-01-21T08:00Z",
          "2026-01-22T08:00Z",
          "2026-01-23T08:00Z",
          "2026-01-24T08:00Z",
          "2026-01-26T08:00Z",
          "2026-01-27T08:00Z",
          "2026-01-28T08:00Z",
          "2026-01-29T08:00Z",
          "2026-01-30T08:00Z",
          "2026-01-31T08:00Z",
          "2026-02-01T08:00Z",
          "2026-02-02T08:00Z",
          "2026-02-04T08:00Z",
          "2026-02-05T08:00Z",
          "2026-02-06T08:00Z",
          "2026-02-07T08:00Z",
          "2026-02-09T08:00Z",
          "2026-02-10T08:00Z",
          "2026-02-11T08:00Z"
        ]
      },
      "sections": [
        {
          "label": "Preseason",
          "value": "1",
          "startDate": "2025-07-31T07:00Z",
          "endDate": "2025-09-04T06:59Z",
          "entries": [
            {
              "label": "Hall of Fame Weekend",
              "alternateLabel": "HOF",
              "detail": "Jul 31-Aug 6",
              "value": "1",
              "startDate": "2025-07-31T07:00Z",
              "endDate": "2025-08-07T06:59Z"
            },
            {
              "label": "Preseason Week 1",
              "alternateLabel": "Pre Wk 1",
              "detail": "Aug 7-13",
              "value": "2",
              "startDate": "2025-08-07T07:00Z",
              "endDate": "2025-08-14T06:59Z"
            },
            {
              "label": "Preseason Week 2",
              "alternateLabel": "Pre Wk 2",
              "detail": "Aug 14-20",
              "value": "3",
              "startDate": "2025-08-14T07:00Z",
              "endDate": "2025-08-21T06:59Z"
            },
            {
              "label": "Preseason Week 3",
              "alternateLabel": "Pre Wk 3",
              "detail": "Aug 21-Sep 3",
              "value": "4",
              "startDate": "2025-08-21T07:00Z",
              "endDate": "2025-09-04T06:59Z"
            }
          ],
          "seasonType": {}
        },
        {
          "label": "Regular Season",
          "value": "2",
          "startDate": "2025-09-04T07:00Z",
          "endDate": "2026-01-07T07:59Z",
          "entries": [
            {
              "label": "Week 1",
              "alternateLabel": "Week 1",
              "detail": "Sep 4-9",
              "value": "1",
              "startDate": "2025-09-04T07:00Z",
              "endDate": "2025-09-10T06:59Z"
            },
            {
              "label": "Week 2",
              "alternateLabel": "Week 2",
              "detail": "Sep 10-16",
              "value": "2",
              "startDate": "2025-09-10T07:00Z",
              "endDate": "2025-09-17T06:59Z"
            },
            {
              "label": "Week 3",
              "alternateLabel": "Week 3",
              "detail": "Sep 17-23",
              "value": "3",
              "startDate": "2025-09-17T07:00Z",
              "endDate": "2025-09-24T06:59Z"
            },
            {
              "label": "Week 4",
              "alternateLabel": "Week 4",
              "detail": "Sep 24-30",
              "value": "4",
              "startDate": "2025-09-24T07:00Z",
              "endDate": "2025-10-01T06:59Z"
            },
            {
              "label": "Week 5",
              "alternateLabel": "Week 5",
              "detail": "Oct 1-7",
              "value": "5",
              "startDate": "2025-10-01T07:00Z",
              "endDate": "2025-10-08T06:59Z"
            },
            {
              "label": "Week 6",
              "alternateLabel": "Week 6",
              "detail": "Oct 8-14",
              "value": "6",
              "startDate": "2025-10-08T07:00Z",
              "endDate": "2025-10-15T06:59Z"
            },
            {
              "label": "Week 7",
              "alternateLabel": "Week 7",
              "detail": "Oct 15-21",
              "value": "7",
              "startDate": "2025-10-15T07:00Z",
              "endDate": "2025-10-22T06:59Z"
            },
            {
              "label": "Week 8",
              "alternateLabel": "Week 8",
              "detail": "Oct 22-28",
              "value": "8",
              "startDate": "2025-10-22T07:00Z",
              "endDate": "2025-10-29T06:59Z"
            },
            {
              "label": "Week 9",
              "alternateLabel": "Week 9",
              "detail": "Oct 29-Nov 4",
              "value": "9",
              "startDate": "2025-10-29T07:00Z",
              "endDate": "2025-11-05T07:59Z"
            },
            {
              "label": "Week 10",
              "alternateLabel": "Week 10",
              "detail": "Nov 5-11",
              "value": "10",
              "startDate": "2025-11-05T08:00Z",
              "endDate": "2025-11-12T07:59Z"
            },
            {
              "label": "Week 11",
              "alternateLabel": "Week 11",
              "detail": "Nov 12-18",
              "value": "11",
              "startDate": "2025-11-12T08:00Z",
              "endDate": "2025-11-19T07:59Z"
            },
            {
              "label": "Week 12",
              "alternateLabel": "Week 12",
              "detail": "Nov 19-25",
              "value": "12",
              "startDate": "2025-11-19T08:00Z",
              "endDate": "2025-11-26T07:59Z"
            },
            {
              "label": "Week 13",
              "alternateLabel": "Week 13",
              "detail": "Nov 26-Dec 2",
              "value": "13",
              "startDate": "2025-11-26T08:00Z",
              "endDate": "2025-12-03T07:59Z"
            },
            {
              "label": "Week 14",
              "alternateLabel": "Week 14",
              "detail": "Dec 3-9",
              "value": "14",
              "startDate": "2025-12-03T08:00Z",
              "endDate": "2025-12-10T07:59Z"
            },
            {
              "label": "Week 15",
              "alternateLabel": "Week 15",
              "detail": "Dec 10-16",
              "value": "15",
              "startDate": "2025-12-10T08:00Z",
              "endDate": "2025-12-17T07:59Z"
            },
            {
              "label": "Week 16",
              "alternateLabel": "Week 16",
              "detail": "Dec 17-23",
              "value": "16",
              "startDate": "2025-12-17T08:00Z",
              "endDate": "2025-12-24T07:59Z"
            },
            {
              "label": "Week 17",
              "alternateLabel": "Week 17",
              "detail": "Dec 24-30",
              "value": "17",
              "startDate": "2025-12-24T08:00Z",
              "endDate": "2025-12-31T07:59Z"
            },
            {
              "label": "Week 18",
              "alternateLabel": "Week 18",
              "detail": "Dec 31-Jan 6",
              "value": "18",
              "startDate": "2025-12-31T08:00Z",
              "endDate": "2026-01-07T07:59Z"
            }
          ],
          "seasonType": {}
        },
        {
          "label": "Postseason",
          "value": "3",
          "startDate": "2026-01-07T08:00Z",
          "endDate": "2026-02-12T07:59Z",
          "entries": [
            {
              "label": "Wild Card",
              "alternateLabel": "Wild Card",
              "detail": "Jan 7-13",
              "value": "1",
              "startDate": "2026-01-07T08:00Z",
              "endDate": "2026-01-14T07:59Z"
            },
            {
              "label": "Divisional Round",
              "alternateLabel": "Div Rd",
              "detail": "Jan 14-20",
              "value": "2",
              "startDate": "2026-01-14T08:00Z",
              "endDate": "2026-01-21T07:59Z"
            },
            {
              "label": "Conference Championship",
              "alternateLabel": "Conf Champ",
              "detail": "Jan 21-27",
              "value": "3",
              "startDate": "2026-01-21T08:00Z",
              "endDate": "2026-01-28T07:59Z"
            },
            {
              "label": "Pro Bowl",
              "alternateLabel": "Pro Bowl",
              "detail": "Jan 28-Feb 3",
              "value": "4",
              "startDate": "2026-01-28T08:00Z",
              "endDate": "2026-02-04T07:59Z"
            },
            {
              "label": "Super Bowl",
              "alternateLabel": "Super Bowl",
              "detail": "Feb 4-11",
              "value": "5",
              "startDate": "2026-02-04T08:00Z",
              "endDate": "2026-02-12T07:59Z"
            }
          ],
          "seasonType": {}
        },
        {
          "label": "Off Season",
          "value": "4",
          "startDate": "2026-02-12T08:00Z",
          "endDate": "2026-08-06T06:59Z",
          "entries": [
            {
              "label": "Week 1",
              "alternateLabel": "Week 1",
              "detail": "Feb 12-Jul 31",
              "value": "1",
              "startDate": "2026-02-12T08:00Z",
              "endDate": "2026-08-01T06:59Z"
            }
          ],
          "seasonType": {}
        }
      ],
      "season": {}
    },
    {
      "error": {
        "message": "Calendar type whens is not supported.",
        "code": 400
      }
    },
    {
      "error": {
        "message": "Calendar type events is not supported.",
        "code": 400
      }
    }
  ]
}
```
