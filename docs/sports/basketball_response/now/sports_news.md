# Now Sports News

## https://now.core.api.espn.com/v1/sports/news?sport=basketball

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "status": "success",
  "timestamp": "2026-05-09T03:16:14Z",
  "headlines": [
    {
      "id": 48717144,
      "type": "Recap",
      "categories": [
        {},
        {}
      ],
      "nowId": "1-48717144",
      "contentKey": "48717144-1-21-1",
      "dataSourceIdentifier": "c317de6004053",
      "publishedkey": "wnba401856891",
      "gameId": "401856891",
      "headline": "Shakira Austin hits winning free throws as Mystics spoil Tempo opener with 68-65 win",
      "description": "— Shakira Austin made four free throws down the stretch and the Washington Mystics held on to beat the Toronto Tempo 68-65 on Friday night, spoiling the Tempo's first game.",
      "linkText": "Shakira Austin hits winning free throws as Mystics spoil Tempo opener with 68-65 win",
      "categorized": "2026-05-09T02:28:38Z"
    },
    {
      "id": 48717062,
      "type": "HeadlineNews",
      "categories": [
        {},
        {}
      ],
      "nowId": "1-48717062",
      "contentKey": "48717062-1-5-1",
      "dataSourceIdentifier": "e02cdd9cd0568",
      "publishedkey": "48717062",
      "feedDisplayType": "Default",
      "headline": "Jalen Brunson shuts door on 76ers as Knicks go up 3-0",
      "description": "Behind 33 points from Jalen Brunson, including big buckets late, the Knicks held off the 76ers to take a commanding 3-0 lead in the semifinal series.",
      "title": "Jalen Brunson shuts door on 76ers as Knicks go up 3-0",
      "linkText": "Brunson shuts door on 76ers as Knicks go up 3-0"
    }
  ],
  "resultsCount": 2,
  "resultsLimit": 10,
  "resultsOffset": 0,
  "breakingNews": []
}
```
