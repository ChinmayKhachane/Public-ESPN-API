# Now Sports News

## https://now.core.api.espn.com/v1/sports/news?sport=hockey

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://now.core.api.espn.com/v1/sports/news?sport=hockey&limit=5`

## Example Response

```json
{
  "resultsCount": 5,
  "resultsLimit": 10,
  "resultsOffset": 0,
  "headlines": [
    {
      "id": 44686270,
      "nowId": "1-44686270",
      "contentKey": "44686270-1-6-1",
      "dataSourceIdentifier": "25ce825f3281f",
      "publishedkey": "44686270",
      "type": "Story",
      "feedDisplayType": "Default",
      "headline": "Stanley Cup odds: Eastern Conference favorite Hurricanes look to sweep Flyers",
      "description": "The Avalanche and Hurricanes are the top two favorites to win the 2026 Stanley Cup.",
      "title": "Stanley Cup odds: Eastern Conference favorite Hurricanes look to sweep Flyers",
      "linkText": "Stanley Cup odds: Eastern Conference favorite Hurricanes look to sweep Flyers",
      "categorized": "2026-05-09T16:18:43Z",
      "originallyPosted": "2026-05-09T15:52:00Z",
      "lastModified": "2026-05-09T16:18:40Z"
    },
    {
      "id": 48714642,
      "nowId": "1-48714642",
      "contentKey": "48714642-1-6-1",
      "dataSourceIdentifier": "ee9ac16994ea7",
      "publishedkey": "48714642",
      "type": "Story",
      "feedDisplayType": "Default",
      "headline": "Stanley Cup playoffs: How a culture of loyalty built the Wild's success",
      "description": "While some teams see veterans come and go, Minnesota has become a place players don't want to leave.",
      "title": "Stanley Cup playoffs: How a culture of loyalty built the Wild's success",
      "linkText": "Where the Wild things are ... and stay: Why players choose to stay in Minnesota",
      "categorized": "2026-05-09T15:55:30Z",
      "originallyPosted": "2026-05-09T12:00:00Z",
      "lastModified": "2026-05-09T10:57:23Z"
    }
  ],
  "breakingNews": [],
  "timestamp": "2026-05-09T17:03:15Z",
  "status": "success"
}
```
