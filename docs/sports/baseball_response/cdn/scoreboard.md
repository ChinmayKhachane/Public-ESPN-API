# MLB CDN Scoreboard

## https://cdn.espn.com/core/mlb/scoreboard?xhr=1

Notes:
- Verified with `league=mlb`, `season=2026`, `event=401815256`, `competition=401815256`, `competitor=17` (Cincinnati Reds), `team=17`, `athlete=4414528` (Andrew Abbott), `play=4018152560000000059` on 2026-05-08.
- Live request returned HTTP `200`.
- HTTP `200` with JSON payload.

## Example Response

```json
{
  "news": {
    "header": "MLB News",
    "link": {
      "isExternal": false,
      "shortText": "All News",
      "rel": [
        "index",
        "desktop"
      ],
      "language": "en-US",
      "href": "https://www.espn.com/mlb/",
      "text": "All MLB News",
      "isPremium": false
    },
    "articles": [
      {
        "id": 48717767,
        "categories": [
          {
            "id": 80311,
            "uid": "s:1~l:10",
            "guid": "b38f959b-7865-31ac-8841-b88355519e10",
            "sportId": 10,
            "leagueId": 10,
            "league": {},
            "type": "league"
          },
          {
            "id": 724,
            "uid": "s:1~l:10~t:25",
            "guid": "4dec648c-3eb9-055c-aebc-2711f30975a0",
            "team": {},
            "sportId": 10,
            "teamId": 25,
            "type": "team"
          }
        ],
        "contentKey": "48717767-1-5-1",
        "images": [
          {
            "id": 13368714,
            "name": "Baseball seams 150802 [600x400]",
            "dataSourceIdentifier": "886e0be6d23ad",
            "alt": "Baseball seams",
            "width": 600,
            "type": "header",
            "credit": "Pouya Dianat/Getty Images",
            "url": "https://a.espncdn.com/photo/2015/0802/mlb_baseball_b1_600x400.jpg",
            "height": 400
          }
        ],
        "dataSourceIdentifier": "f60e80c46aee3",
        "published": "2026-05-09T03:43:19Z",
        "type": "HeadlineNews",
        "nowId": "1-48717767",
        "premium": false,
        "links": {
          "app": {
            "sportscenter": {}
          },
          "web": {
            "href": "https://www.espn.com/mlb/story/_/id/48717767/padres-prospect-pleads-guilty-charge-transporting-noncitizen-i..."
          },
          "mobile": {
            "href": "http://m.espn.go.com/mlb/story?storyId=48717767"
          },
          "api": {
            "self": {}
          }
        }
      },
      {
        "id": 48717747,
        "categories": [
          {
            "id": 525402,
            "uid": "s:1~l:10~a:4142424",
            "guid": "c8f56866-3f5a-304f-9961-10d6b8970628",
            "sportId": 10,
            "athleteId": 4142424,
            "athlete": {},
            "type": "athlete"
          },
          {
            "id": 468339,
            "uid": "s:1~l:10~a:4717833",
            "guid": "dcbeba5a-3fe9-3bf0-9f19-c077ef2a6f78",
            "sportId": 10,
            "athleteId": 4717833,
            "athlete": {},
            "type": "athlete"
          }
        ],
        "contentKey": "48717747-1-293-1",
        "images": [
          {
            "name": "Chicago Cubs vs. Texas Rangers: Game Highlights",
            "alt": "",
            "width": 576,
            "url": "https://a.espncdn.com/media/motion/wsc/2026/0509/e4a1a620-c370-4340-9b36-3c6550c5edc9/e4a1a620-c370-4340-9b...",
            "height": 324
          }
        ],
        "dataSourceIdentifier": "c9a50ebd56d4c",
        "published": "2026-05-09T03:30:14Z",
        "type": "Media",
        "nowId": "1-48717747",
        "premium": false,
        "links": {
          "sportscenter": {
            "href": "sportscenter://x-callback-url/showVideo?videoID=48717747&videoDSI=c9a50ebd56d4c"
          },
          "web": {
            "self": {},
            "href": "https://www.espn.com/video/clip/_/id/48717747/game-highlights"
          },
          "api": {
            "self": {},
            "artwork": {}
          }
        }
      }
    ]
  },
  "pinnedCount": 0,
  "nowFeedMD5Hash": "864d43377c023715f628d50f4a8b6a32",
  "type": "scoreboard",
  "content": {
    "league": "mlb",
    "sbGroup": {
      "pageTitle": "MLB",
      "altTitle": "Major League Baseball",
      "scheduleStartDate": "2002-04-01",
      "isCollege": false,
      "league": "mlb",
      "sport": "baseball"
    },
    "sbData": {
      "events": [
        {
          "id": "401815256",
          "uid": "s:1~l:10~e:401815256",
          "name": "Houston Astros at Cincinnati Reds",
          "competitions": [
            {}
          ],
          "status": {
            "period": 9,
            "displayClock": "0:00",
            "clock": 0,
            "type": {}
          },
          "season": {
            "slug": "regular-season",
            "year": 2026,
            "type": 2
          },
          "date": "2026-05-08T22:10Z",
          "links": [
            {},
            {}
          ],
          "shortName": "HOU @ CIN"
        },
        {
          "id": "401815255",
          "uid": "s:1~l:10~e:401815255",
          "name": "Colorado Rockies at Philadelphia Phillies",
          "competitions": [
            {}
          ],
          "status": {
            "period": 11,
            "displayClock": "0:00",
            "clock": 0,
            "type": {}
          },
          "season": {
            "slug": "regular-season",
            "year": 2026,
            "type": 2
          },
          "date": "2026-05-08T22:40Z",
          "links": [
            {},
            {}
          ],
          "shortName": "COL @ PHI"
        }
      ],
      "season": {
        "year": 2026,
        "type": 2
      },
      "leagues": [
        {
          "id": "10",
          "uid": "s:1~l:10",
          "name": "Major League Baseball",
          "abbreviation": "MLB",
          "slug": "mlb",
          "season": {
            "displayName": "2026",
            "year": 2026,
            "endDate": "2026-11-12T07:59Z",
            "type": {},
            "startDate": "2026-02-19T08:00Z"
          },
          "calendarIsWhitelist": false,
          "calendar": [
            "2026-02-19T08:00Z",
            "2026-07-13T07:00Z"
          ],
          "calendarType": "day",
          "calendarEndDate": "2026-11-12T07:59Z"
        }
      ],
      "provider": {
        "id": "100",
        "name": "Draft Kings",
        "displayName": "Draft Kings",
        "priority": 1,
        "logos": [
          {
            "rel": [],
            "href": "https://a.espncdn.com/i/betting/Draftkings_Light.svg"
          },
          {
            "rel": [],
            "href": "https://a.espncdn.com/i/betting/Draftkings_Dark.svg"
          }
        ]
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
      "2026-02-19T08:00Z",
      "2026-07-13T07:00Z"
    ],
    "defaults": {
      "scoDate": "20260508"
    },
    "title": "MLB Baseball Scores - MLB Scoreboard - ESPN",
    "og_type": "website",
    "canonical": "/mlb/scoreboard"
  },
  "analytics": {
    "metrics": {
      "page_url": "/mlb/scoreboard",
      "site": "espn",
      "content_type": "scoreboard",
      "page_infrastructure": "sCore",
      "page_type": "scoreboard",
      "league": "mlb",
      "page_name": "espn:mlb:scoreboard",
      "section": "mlb",
      "sport": "baseball"
    },
    "omniture": {
      "espn3ContentType": "scoreboard:null:mlb",
      "league": "mlb",
      "countryRegion": "en-us",
      "hier1": "mlb:scoreboard",
      "section": "mlb",
      "pageName": "mlb:scoreboard",
      "sections": "mlb:scoreboard",
      "site": "espn",
      "premium": "premium-no",
      "appearance": "light"
    },
    "chartbeat": {
      "domain": "www.espn.com",
      "sections": "mlb",
      "authors": "scoreboard",
      "path": "/mlb/scoreboard",
      "title": "MLB Baseball Scores - MLB Scoreboard - ESPN",
      "zone": "www.espn.com.us.mlb",
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
      }
    },
    "device": "desktop",
    "cto": true,
    "qualtrics": false
  },
  "nowFeed": [
    {
      "id": 41126698,
      "categories": [
        {
          "id": 581,
          "uid": "s:1~l:10~t:17",
          "guid": "04b65a0b-3cca-d795-0e21-23606470418a",
          "team": {
            "id": 17,
            "links": {},
            "logos": {}
          },
          "sportId": 10,
          "teamId": 17,
          "type": "team"
        },
        {
          "id": 474,
          "uid": "s:1~l:10~t:22",
          "guid": "ff1e263a-f6a6-93c3-1373-418623652ff0",
          "team": {
            "id": 22,
            "links": {},
            "logos": {}
          },
          "sportId": 10,
          "teamId": 22,
          "type": "team"
        }
      ],
      "contentKey": "41126698-21-4-1",
      "images": [],
      "dataSourceIdentifier": "e4568fc3464e8",
      "externalId": "41126698",
      "linkText": "Marcus Giamatti, whose late father Bart banned Pete Rose in 1989, told ESPN he opposes the Hall of Fame can...",
      "categorized": "2025-05-15T03:02:45Z",
      "published": "2025-05-13T20:00:07Z",
      "type": "Shortstop"
    },
    {
      "id": 41126697,
      "categories": [
        {
          "id": 581,
          "uid": "s:1~l:10~t:17",
          "guid": "04b65a0b-3cca-d795-0e21-23606470418a",
          "team": {
            "id": 17,
            "links": {},
            "logos": {}
          },
          "sportId": 10,
          "teamId": 17,
          "type": "team"
        },
        {
          "id": 474,
          "uid": "s:1~l:10~t:22",
          "guid": "ff1e263a-f6a6-93c3-1373-418623652ff0",
          "team": {
            "id": 22,
            "links": {},
            "logos": {}
          },
          "sportId": 10,
          "teamId": 22,
          "type": "team"
        }
      ],
      "contentKey": "41126697-21-4-1",
      "images": [],
      "dataSourceIdentifier": "3ea6d44c9051d",
      "externalId": "41126697",
      "linkText": "\"Reinstating him, even after he died, puts in peril the future & integrity of the game … (it) lets the dama...",
      "categorized": "2025-05-15T03:02:44Z",
      "published": "2025-05-13T19:57:26Z",
      "type": "Shortstop"
    }
  ],
  "ads": {
    "id": 12129264,
    "page_url": "https://www.espn.com/mlb/scoreboard",
    "prebidAdConfig": {
      "usePrebidBids": true,
      "timeout": 1000
    },
    "level": "espn.com/mlb/scoreboard",
    "sizesEspnPlus": {
      "banner-index": {
        "excludedSize": [
          "728,90"
        ],
        "mappings": [
          {
            "viewport": [],
            "slot": []
          },
          {
            "viewport": [],
            "slot": []
          }
        ],
        "defaultSize": [
          970,
          66
        ],
        "excludedProfile": [
          "xl"
        ],
        "includedCountries": [
          "us"
        ],
        "pbjs": {
          "s": [
            []
          ],
          "xl": [
            []
          ],
          "l": [
            []
          ],
          "m": [
            []
          ]
        }
      },
      "gamecast": {
        "mappings": [
          {
            "viewport": [],
            "slot": []
          }
        ],
        "defaultSize": [
          320,
          50
        ]
      },
      "banner-scoreboard": {
        "excludedSize": [
          "970,250"
        ],
        "mappings": [
          {
            "viewport": [],
            "slot": []
          },
          {
            "viewport": [],
            "slot": []
          }
        ],
        "defaultSize": [
          970,
          66
        ],
        "includedCountries": [
          "us"
        ],
        "pbjs": {
          "s": [
            []
          ],
          "xl": [
            []
          ],
          "l": [
            []
          ],
          "m": [
            []
          ]
        }
      },
      "banner": {
        "mappings": [
          {
            "viewport": [],
            "slot": []
          },
          {
            "viewport": [],
            "slot": []
          }
        ],
        "defaultSize": [
          970,
          66
        ],
        "pbjs": {
          "s": [
            []
          ],
          "xl": [
            [],
            []
          ],
          "l": [
            [],
            []
          ],
          "m": [
            []
          ]
        }
      },
      "incontent-betting": {
        "mappings": [
          {
            "viewport": [],
            "slot": []
          },
          {
            "viewport": [],
            "slot": []
          }
        ],
        "defaultSize": [
          300,
          251
        ]
      },
      "native-betting": {
        "mappings": [
          {
            "viewport": [],
            "slot": []
          }
        ],
        "defaultSize": "fluid"
      },
      "instream": {
        "mappings": [
          {
            "viewport": [],
            "slot": []
          }
        ],
        "defaultSize": [
          1,
          3
        ]
      },
      "incontent": {
        "mappings": [
          {
            "viewport": [],
            "slot": []
          }
        ],
        "defaultSize": [
          300,
          250
        ]
      }
    },
    "delayInPageAdSlots": true,
    "incontentPositions": {
      "defaults": {
        "favorites": -1,
        "news": 4,
        "now": 4
      },
      "index": {
        "top": {
          "favorites": -1
        },
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
    "network": "21783347309"
  },
  "nowFeedCount": 25,
  "meta": {
    "imageWidth": 1200,
    "image": "https://a.espncdn.com/i/espn/espn_logos/espn_red.png",
    "twitter_card": "summary",
    "og_site_name": "ESPN.com",
    "twitter_app_id_iphone": "317469184",
    "og_type": "website",
    "twitter_app_name_googleplay": "ESPN",
    "label": "MLB",
    "canonical": "https://www.espn.com/mlb/scoreboard",
    "type": "scoreboard"
  },
  "nowFeedSupported": true,
  "sport": [
    "mlb"
  ],
  "tier2Nav": {
    "subNavMenu": {
      "navigation": {
        "$ref": "/v2/navigation/12001873",
        "id": 12001873,
        "items": [
          {
            "$ref": "/v2/navigation/12001915",
            "id": 12001915,
            "links": [],
            "title": "MLB Home"
          },
          {
            "$ref": "/v2/navigation/11586778",
            "id": 11586778,
            "links": [],
            "title": "MLB Scores"
          }
        ],
        "links": [
          {
            "isExternal": false,
            "shortText": "MLB",
            "rel": [],
            "text": "MLB",
            "href": "/mlb/",
            "isPremium": false
          }
        ],
        "attributes": {
          "sport_id": "10",
          "root": "mlb",
          "league": true
        },
        "text": "MLB",
        "title": "MLB Menu - LIVE"
      },
      "navId": 12001873,
      "fallback": false
    }
  }
}
```
