# CDN Scoreboard

## https://cdn.espn.com/core/nba/scoreboard?xhr=1

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.
- Returns a CDN page shell; scoreboard data is nested under `content.sbData`.

## Example Response

```json
{
  "type": "scoreboard",
  "content": {
    "league": "nba",
    "sbGroup": {
      "league": "nba",
      "pageTitle": "NBA",
      "altTitle": "National Basketball Association",
      "scheduleStartDate": "2002-10-10",
      "isCollege": false,
      "sport": "basketball"
    },
    "sbData": {
      "season": {
        "year": 2026,
        "type": 3
      },
      "events": [
        {},
        {}
      ],
      "leagues": [
        {}
      ],
      "provider": {
        "id": "100",
        "name": "Draft Kings",
        "displayName": "Draft Kings",
        "priority": 1,
        "logos": []
      },
      "day": {
        "date": "2026-05-08"
      }
    },
    "isWeekOriented": false,
    "dateParams": {
      "date": "20260508"
    },
    "calendar": [
      "2025-10-01T07:00Z",
      "2025-10-18T07:00Z"
    ],
    "defaults": {
      "scoDate": "20260508"
    },
    "title": "NBA Basketball Scores - NBA Scoreboard - ESPN",
    "description": "Real-time NBA Basketball scores on ESPN",
    "og_type": "website",
    "canonical": "/nba/scoreboard"
  },
  "news": {
    "articles": [
      {
        "id": 48717062,
        "type": "HeadlineNews",
        "categories": [],
        "contentKey": "48717062-1-5-1",
        "images": [],
        "dataSourceIdentifier": "e02cdd9cd0568",
        "description": "Behind 33 points from Jalen Brunson, including big buckets late, the Knicks held off the 76ers to take a commanding 3-0 lead in the semifinal series.",
        "published": "2026-05-09T02:49:38Z",
        "nowId": "1-48717062",
        "premium": false,
        "links": {},
        "lastModified": "2026-05-09T02:49:38Z"
      },
      {
        "id": 48676289,
        "type": "Story",
        "categories": [],
        "contentKey": "48676289-1-6-1",
        "images": [],
        "dataSourceIdentifier": "67dd576b8b9d4",
        "description": "Here's what we've learned -- and what's next -- for the 76ers-Knicks, Cavs-Pistons, Wolves-Spurs and Lakers-Thunder series.",
        "published": "2026-05-09T02:30:20Z",
        "nowId": "1-48676289",
        "premium": false,
        "links": {},
        "lastModified": "2026-05-09T02:30:20Z"
      }
    ],
    "header": "NBA News",
    "link": {
      "text": "All NBA News",
      "shortText": "All News",
      "isExternal": false,
      "rel": [
        "index",
        "desktop"
      ],
      "language": "en-US",
      "href": "https://www.espn.com/nba/",
      "isPremium": false
    }
  },
  "pinnedCount": 0,
  "nowFeedMD5Hash": "5e1fdd3a4404947b10084434bd1e72f0",
  "analytics": {
    "metrics": {
      "league": "nba",
      "page_url": "/nba/scoreboard",
      "site": "espn",
      "content_type": "scoreboard",
      "page_infrastructure": "sCore",
      "page_type": "scoreboard",
      "page_name": "espn:nba:scoreboard",
      "section": "nba",
      "sport": "basketball"
    },
    "omniture": {
      "league": "nba",
      "espn3ContentType": "scoreboard:null:nba",
      "countryRegion": "en-us",
      "hier1": "nba:scoreboard",
      "section": "nba",
      "pageName": "nba:scoreboard",
      "sections": "nba:scoreboard",
      "site": "espn",
      "premium": "premium-no",
      "appearance": "light",
      "convrSport": "basketball",
      "pageURL": "www.espn.com/nba/scoreboard"
    },
    "chartbeat": {
      "domain": "www.espn.com",
      "sections": "nba",
      "authors": "scoreboard",
      "path": "/nba/scoreboard",
      "title": "NBA Basketball Scores - NBA Scoreboard - ESPN",
      "zone": "www.espn.com.us.nba",
      "loadPubJS": false,
      "loadVidJS": true
    },
    "nielsen": {
      "espnuk": {
        "apid": "P07264C85-15CD-4A80-8E56-B5BFA6D93296",
        "vc": "b01"
      },
      "espnau": {
        "apid": "P07264C85-15CD-4A80-8E56-B5BFA6D93296",
        "vc": "b01"
      },
      "espn": {
        "apid": "P07264C85-15CD-4A80-8E56-B5BFA6D93296",
        "vc": "b01"
      },
      "fantasy": {
        "apid": "P302B69D5-F1DD-4E7A-BF8D-3E60F0EB5E5A",
        "vc": "c07"
      },
      "espndeportes": {
        "apid": "P890E2723-EDBC-4CCE-96BA-F35EA3E50650",
        "vc": "c02"
      },
      "espnfc": {
        "apid": "PE6995AAE-0C49-4372-B5E7-54C61BFE2AA5",
        "vc": "c03"
      },
      "espnww": {
        "apid": "P07264C85-15CD-4A80-8E56-B5BFA6D93296",
        "vc": "b01"
      },
      "general": {
        "ci": "us-600140",
        "assetid": "N/A",
        "segB": "N/A",
        "sfcode": "dcr",
        "segA": "N/A",
        "section": "N/A",
        "segC": "N/A",
        "apn": "espnCOM"
      },
      "espnza": {
        "apid": "P07264C85-15CD-4A80-8E56-B5BFA6D93296",
        "vc": "b01"
      },
      "espnin": {
        "apid": "P07264C85-15CD-4A80-8E56-B5BFA6D93296",
        "vc": "b01"
      },
      "watchespn": {
        "apid": "P07264C85-15CD-4A80-8E56-B5BFA6D93296",
        "vc": "b01"
      },
      "cricinfo": {
        "apid": "PED8CDAC2-F114-41BE-8B98-AFA06FAEA06E",
        "vc": "c04"
      }
    },
    "device": "desktop",
    "cto": true,
    "qualtrics": false
  },
  "nowFeed": [
    {
      "id": 41131522,
      "type": "Shortstop",
      "categories": [
        {},
        {}
      ],
      "contentKey": "41131522-21-4-1",
      "images": [
        {}
      ],
      "dataSourceIdentifier": "8485ee852e0ff",
      "linkText": "New York Knicks forward OG Anunoby has been diagnosed with right hamstring strain and will be day to day, sources tell ESPN. Anunoby is the Knicks second-leading scorer this pos...",
      "categorized": "2026-05-07T17:31:23Z",
      "published": "2026-05-07T17:29:06Z",
      "title": "New York Knicks forward OG Anunoby has been diagnosed with right hamstring strain and will be day to day, sources tell ESPN. Anunoby is the Knicks second-leading scorer this pos...",
      "allowContentReactions": true,
      "nowId": "21-41131522"
    },
    {
      "id": 41131518,
      "type": "Shortstop",
      "categories": [
        {},
        {}
      ],
      "contentKey": "41131518-21-4-1",
      "images": [
        {}
      ],
      "dataSourceIdentifier": "e345d33dd41d6",
      "linkText": "Just in: Philadelphia 76ers star Joel Embiid has been ruled out for Game 2 tonight against the New York Knicks due to ankle and hip injuries, sources tell ESPN.",
      "categorized": "2026-05-06T17:01:51Z",
      "published": "2026-05-06T16:58:16Z",
      "title": "Just in: Philadelphia 76ers star Joel Embiid has been ruled out for Game 2 tonight against the New York Knicks due to ankle and hip injuries, sources tell ESPN.",
      "allowContentReactions": true,
      "nowId": "21-41131518"
    }
  ],
  "ads": {
    "id": 12129264,
    "page_url": "https://www.espn.com/nba/scoreboard",
    "prebidAdConfig": {
      "usePrebidBids": true,
      "timeout": 1000
    },
    "level": "espn.com/nba/scoreboard",
    "sizesEspnPlus": {
      "banner-index": {
        "excludedSize": [],
        "mappings": [],
        "defaultSize": [],
        "excludedProfile": [],
        "includedCountries": [],
        "pbjs": {}
      },
      "gamecast": {
        "mappings": [],
        "defaultSize": []
      },
      "banner-scoreboard": {
        "excludedSize": [],
        "mappings": [],
        "defaultSize": [],
        "includedCountries": [],
        "pbjs": {}
      },
      "banner": {
        "mappings": [],
        "defaultSize": [],
        "pbjs": {}
      },
      "incontent-betting": {
        "mappings": [],
        "defaultSize": []
      },
      "native-betting": {
        "mappings": [],
        "defaultSize": "fluid"
      },
      "instream": {
        "mappings": [],
        "defaultSize": []
      },
      "incontent": {
        "mappings": [],
        "defaultSize": []
      }
    },
    "delayInPageAdSlots": true,
    "incontentPositions": {
      "defaults": {
        "news": 4,
        "favorites": -1,
        "now": 4
      },
      "index": {
        "top": {},
        "nfl": {}
      }
    },
    "showEspnPlusAds": false,
    "kvpsEspnPlus": [
      {
        "name": "ed",
        "value": "us"
      },
      {
        "name": "eplus",
        "value": "true"
      }
    ],
    "network": "21783347309",
    "refreshOnBreakpointChange": true,
    "webviewOverride": {
      "banner": {
        "roster": "banner-webview",
        "mlb/stats": "banner-webview",
        "cfb/rankings": "banner-webview",
        "team/stats": "banner-webview",
        "nba/stats": "banner-webview",
        "ncaaw/rankings": "banner-webview",
        "nfl/stats": "banner-webview",
        "standings": "banner-webview",
        "cfb/stats": "banner-webview",
        "ncb/rankings": "banner-webview"
      }
    }
  },
  "nowFeedCount": 25,
  "meta": {
    "type": "scoreboard",
    "imageWidth": 1200,
    "image": "https://a.espncdn.com/i/espn/espn_logos/espn_red.png",
    "twitter_card": "summary",
    "og_site_name": "ESPN.com",
    "twitter_app_id_iphone": "317469184",
    "description": "Real-time NBA Basketball scores on ESPN",
    "og_type": "website",
    "twitter_app_name_googleplay": "ESPN",
    "label": "NBA",
    "canonical": "https://www.espn.com/nba/scoreboard",
    "title": "NBA Basketball Scores - NBA Scoreboard - ESPN"
  },
  "nowFeedSupported": true,
  "sport": [
    "nba"
  ]
}
```
