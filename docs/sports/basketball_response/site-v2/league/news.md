# News

## https://site.api.espn.com/apis/site/v2/sports/basketball/{league}/news

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.

## Example Response

```json
{
  "articles": [
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
      "headline": "Jalen Brunson shuts door on 76ers as Knicks go up 3-0",
      "description": "Behind 33 points from Jalen Brunson, including big buckets late, the Knicks held off the 76ers to take a commanding 3-0 lead in the semifinal series.",
      "lastModified": "2026-05-09T02:49:38Z",
      "published": "2026-05-09T02:49:38Z",
      "images": [
        {},
        {}
      ],
      "premium": false
    },
    {
      "id": 48676289,
      "type": "Story",
      "categories": [
        {},
        {}
      ],
      "nowId": "1-48676289",
      "contentKey": "48676289-1-6-1",
      "dataSourceIdentifier": "67dd576b8b9d4",
      "headline": "2026 NBA playoffs: Conference semifinals takeaways",
      "description": "Here's what we've learned -- and what's next -- for the 76ers-Knicks, Cavs-Pistons, Wolves-Spurs and Lakers-Thunder series.",
      "lastModified": "2026-05-09T02:30:20Z",
      "published": "2026-05-09T02:30:20Z",
      "images": [
        {},
        {}
      ],
      "premium": false
    }
  ],
  "header": "NBA News",
  "link": {
    "text": "All NBA News",
    "shortText": "All News",
    "language": "en",
    "rel": [
      "index",
      "desktop"
    ],
    "href": "https://www.espn.com/nba/",
    "isExternal": false,
    "isPremium": false
  }
}
```
