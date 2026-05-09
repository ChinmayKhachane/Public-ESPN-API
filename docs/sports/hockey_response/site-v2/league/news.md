# League News

## https://site.api.espn.com/apis/site/v2/sports/hockey/{league}/news

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=UTF-8`.

Tested URL:
`https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/news?limit=2`

## Example Response

```json
{
  "header": "NHL News",
  "link": {
    "language": "en",
    "rel": [
      "index",
      "desktop"
    ],
    "href": "https://www.espn.com/nhl/",
    "text": "All NHL News",
    "shortText": "All News",
    "isExternal": false,
    "isPremium": false
  },
  "articles": [
    {
      "id": 44686270,
      "nowId": "1-44686270",
      "contentKey": "44686270-1-6-1",
      "dataSourceIdentifier": "25ce825f3281f",
      "type": "Story",
      "headline": "Stanley Cup odds: Eastern Conference favorite Hurricanes look to sweep Flyers",
      "description": "The Avalanche and Hurricanes are the top two favorites to win the 2026 Stanley Cup.",
      "lastModified": "2026-05-09T16:18:40Z",
      "published": "2026-05-09T16:18:40Z",
      "images": [
        {
          "dataSourceIdentifier": "36f5d7465189e",
          "id": 48666250,
          "type": "header",
          "name": "Carolina Hurricanes [1296x729]",
          "caption": "The Carolina Hurricanes are the current favorites to win the Eastern Conference.",
          "alt": "Mark Jankowski #77 of the Carolina Hurricanes faces off against Luke Glendening #41 of the Philadelphia Flyers in Game One of the Second Round of the 2026 Stanley Cup Playoffs at Lenovo Center on May 02, 2026 in Raleigh, North Carolina.",
          "credit": "Josh Lavallee/NHLI via Getty Images",
          "height": 729,
          "width": 1296,
          "url": "https://a.espncdn.com/photo/2026/0503/r1652863_1296x729_16-9.jpg"
        },
        {
          "dataSourceIdentifier": "e09e98e94bb75",
          "id": 48704137,
          "type": "header",
          "name": "Josh Doan [1296x729]",
          "caption": "The Sabres' Game 1 victory has their odds on the move.",
          "alt": "Buffalo Sabres right wing Josh Doan puts the puck past Montreal Canadiens goaltender Jakub Dobes (75) during the first period in Game 1 of a second-round NHL hockey Stanley Cup playoff series, Wednesday, May 6, 2026, in Buffalo, N.Y.",
          "credit": "AP Photo/Jeffrey T. Barnes",
          "height": 729,
          "width": 1296,
          "url": "https://a.espncdn.com/photo/2026/0507/r1654913_1296x729_16-9.jpg"
        }
      ],
      "categories": [
        {
          "id": 611584,
          "type": "league",
          "guid": "550a1e9e-377e-3019-bf5e-19c027d50a49",
          "description": "Sports Betting",
          "sportId": 22000,
          "leagueId": 22000,
          "league": {}
        },
        {
          "id": 137817,
          "type": "topic",
          "guid": "14f579db-79ec-a898-9405-8fec63e16063",
          "description": "news \u2013 sports betting",
          "sportId": 0,
          "topicId": 179
        }
      ],
      "premium": false,
      "links": {
        "web": {
          "href": "https://www.espn.com/espn/betting/story/_/id/44686270/2026-nhl-stanley-cup-playoffs-championship-odds"
        },
        "mobile": {
          "href": "http://m.espn.go.com/wireless/story?storyId=44686270"
        },
        "api": {
          "self": {}
        },
        "app": {
          "sportscenter": {}
        }
      },
      "byline": "Doug Greenberg"
    },
    {
      "id": 48714642,
      "nowId": "1-48714642",
      "contentKey": "48714642-1-6-1",
      "dataSourceIdentifier": "ee9ac16994ea7",
      "type": "Story",
      "headline": "Stanley Cup playoffs: How a culture of loyalty built the Wild's success",
      "description": "While some teams see veterans come and go, Minnesota has become a place players don't want to leave.",
      "lastModified": "2026-05-09T10:57:23Z",
      "published": "2026-05-09T10:57:23Z",
      "images": [
        {
          "dataSourceIdentifier": "35b4d2b893d72",
          "id": 48714797,
          "type": "header",
          "name": "Wild celebrate [1296x729]",
          "credit": "Bruce Kluckhohn/NHLI via Getty Images",
          "height": 729,
          "width": 1296,
          "url": "https://a.espncdn.com/photo/2026/0508/r1655517_1296x729_16-9.jpg"
        },
        {
          "dataSourceIdentifier": "424f1ee9a257e",
          "id": 48714776,
          "type": "header",
          "name": "Marcus Foligno [1296x729]",
          "caption": "Marcus Foligno is one of several veterans who chose to stay in Minnesota instead of potentially landing a bigger contract elsewhere.",
          "credit": "Luke Schmidt/NHLI via Getty Images",
          "height": 729,
          "width": 1296,
          "url": "https://a.espncdn.com/photo/2026/0508/r1655515_1296x729_16-9.jpg"
        }
      ],
      "categories": [
        {
          "id": 9580,
          "type": "league",
          "uid": "s:70~l:90",
          "guid": "1a5f0227-a13e-396c-8cea-8961bc288666",
          "description": "NHL",
          "sportId": 90,
          "leagueId": 90,
          "league": {}
        },
        {
          "id": 6601,
          "type": "team",
          "uid": "s:70~l:90~t:30",
          "guid": "b024c635-e192-6545-f8c1-b5da0df5ff6c",
          "description": "Minnesota Wild",
          "sportId": 90,
          "teamId": 30,
          "team": {}
        }
      ],
      "premium": false,
      "links": {
        "web": {
          "href": "https://www.espn.com/nhl/story/_/id/48714642/2026-nhl-playoffs-stanley-cup-minnesota-wild-franchise-contracts-free-agents-trades"
        },
        "mobile": {
          "href": "http://m.espn.go.com/nhl/story?storyId=48714642"
        },
        "api": {
          "self": {}
        },
        "app": {
          "sportscenter": {}
        }
      },
      "byline": "Ryan Clark"
    }
  ]
}
```
