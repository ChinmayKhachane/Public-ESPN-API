# CDN Game

## https://cdn.espn.com/core/nba/game?xhr=1&gameId={event}

Notes:
- Verified with `league=nba`, `season=2026`, `event=401871161`, `competition=401871161`, `competitor=20`, `team=20`, `athlete=3059318`, `play=4018711617` on 2026-05-09.
- Live request returned HTTP `200`.
- Useful game data is under `gamepackageJSON`.

## Example Response

```json
{
  "type": "game",
  "content": {
    "league": "nba",
    "title": "Knicks vs. 76ers - Game Summary - May 8, 2026 - ESPN",
    "description": "Get a summary of the New York Knicks vs. Philadelphia 76ers basketball game",
    "og_type": "website",
    "sport": "basketball",
    "tab": {
      "layout": "cbc",
      "pageType": "Summary",
      "metaDescription": "Get a summary of the {aDisplayName} vs. {hDisplayName} basketball game",
      "columnsModuleTypes": {
        "default": [],
        "tablet": [],
        "mobile": []
      },
      "metaTitle": "{aName} vs. {hName} - Game Summary - {date}"
    },
    "statusState": "post",
    "canonical": "http://www.espn.com/nba/game/_/gameId/401871161",
    "tabType": "game"
  },
  "gamepackageJSON": {
    "leaders": [
      {
        "team": {},
        "leaders": []
      },
      {
        "team": {},
        "leaders": []
      }
    ],
    "boxscore": {
      "teams": [
        {},
        {}
      ],
      "players": [
        {},
        {}
      ]
    },
    "header": {
      "id": "401871161",
      "uid": "s:40~l:46~e:401871161",
      "season": {
        "year": 2026,
        "type": 3,
        "current": true
      },
      "competitions": [
        {}
      ],
      "league": {
        "id": "46",
        "uid": "s:40~l:46",
        "name": "National Basketball Association",
        "abbreviation": "NBA",
        "slug": "nba",
        "links": [],
        "logos": [],
        "isTournament": false
      },
      "timeValid": true,
      "links": [
        {},
        {}
      ],
      "gameNote": "East Semifinals - Game 3"
    },
    "plays": [
      {
        "id": "4018711614",
        "type": {},
        "team": {},
        "text": "Karl-Anthony Towns vs. Joel Embiid (Kelly Oubre Jr. gains possession)",
        "period": {},
        "clock": {},
        "homeScore": 0,
        "awayScore": 0,
        "scoringPlay": false,
        "scoreValue": 0,
        "shootingPlay": false,
        "sequenceNumber": "4"
      },
      {
        "id": "4018711617",
        "type": {},
        "team": {},
        "text": "Kelly Oubre Jr. makes driving layup",
        "period": {},
        "clock": {},
        "homeScore": 2,
        "awayScore": 0,
        "scoringPlay": true,
        "scoreValue": 2,
        "shootingPlay": true,
        "sequenceNumber": "7"
      }
    ],
    "news": {
      "articles": [
        {},
        {}
      ],
      "header": "NBA News",
      "link": {
        "text": "All NBA News",
        "shortText": "All News",
        "isExternal": false,
        "rel": [],
        "language": "en-US",
        "href": "https://www.espn.com/nba/",
        "isPremium": false
      }
    },
    "broadcasts": [
      {
        "type": {},
        "market": {},
        "station": "Prime Video",
        "media": {},
        "lang": "en",
        "region": "us",
        "isNational": true
      },
      {
        "type": {},
        "market": {},
        "station": "13715732",
        "stationKey": "espn",
        "media": {},
        "lang": "en",
        "region": "us",
        "isNational": true
      }
    ],
    "pickcenter": [
      {
        "header": {},
        "underOdds": -110,
        "awayTeamOdds": {},
        "footer": {},
        "link": {},
        "pointSpread": {},
        "spread": -3.5,
        "overUnder": 214.5,
        "total": {},
        "provider": {},
        "details": "PHI -3.5",
        "links": []
      }
    ],
    "winprobability": [
      {
        "play": {},
        "homeWinPercentage": 0.36,
        "playId": "4018711614",
        "tiePercentage": 0
      },
      {
        "play": {},
        "homeWinPercentage": 0.388,
        "playId": "4018711617",
        "tiePercentage": 0
      }
    ],
    "gameInfo": {
      "venue": {
        "id": "1845",
        "guid": "5c91d4ea-46af-3b85-a87d-54422ed1e8c3",
        "images": [],
        "address": {},
        "grass": false,
        "fullName": "Xfinity Mobile Arena",
        "shortName": "Wells Fargo Center"
      },
      "officials": [
        {},
        {}
      ],
      "attendance": 19746
    },
    "videos": [
      {
        "id": 48716904,
        "cerebroId": "69fe9255c35ca554fbcc4c84",
        "thumbnail": "https://a.espncdn.com/media/motion/wsc/2026/0509/cf187f6b-a84f-43a8-9c12-d620f9cee16c/cf187f6b-a84f-43a8-9c12-d620f9cee16c.jpg",
        "ad": {},
        "timeRestrictions": {},
        "description": "New York Knicks vs. Philadelphia 76ers: Game Highlights",
        "geoRestrictions": {},
        "source": "espn",
        "tracking": {},
        "duration": 78,
        "deviceRestrictions": {},
        "originalPublishDate": "2026-05-09T01:48:06Z"
      },
      {
        "id": 48715632,
        "cerebroId": "69fe6f11c4ffca195770685e",
        "thumbnail": "https://a.espncdn.com/media/motion/wsc/2026/0508/fb64a52e-2d75-4347-b31e-77dd34eecb6b/fb64a52e-2d75-4347-b31e-77dd34eecb6b.jpg",
        "ad": {},
        "timeRestrictions": {},
        "description": "Embiid finds George for early Sixers triple",
        "geoRestrictions": {},
        "source": "espn",
        "tracking": {},
        "duration": 24,
        "deviceRestrictions": {},
        "originalPublishDate": "2026-05-08T23:17:37Z"
      }
    ],
    "standings": {
      "header": "2025-26 Standings",
      "groups": [
        {}
      ],
      "fullViewLink": {
        "text": "Full Standings",
        "href": "https://www.espn.com/nba/standings"
      },
      "isSameConference": true
    },
    "article": {
      "id": 48717006,
      "type": "Recap",
      "categories": [
        {},
        {}
      ],
      "dataSourceIdentifier": "1180d2a624701",
      "keywords": [
        "NBA",
        "Philadelphia 76ers"
      ],
      "publishedkey": "nba401871161",
      "description": "— Jalen Brunson scored 33 points and sealed the game with big buckets late to the delight of roaring Knicks fans, leading New York...",
      "categorized": "2026-05-09T02:29:32Z",
      "section": "NBA",
      "source": "AP",
      "video": [
        {}
      ],
      "nowId": "1-48717006"
    }
  },
  "gameId": 401871161,
  "customStyleSheet": "game-package-basketball",
  "__gamepackage__": {
    "playerHash": {
      "4251": {
        "homeAway": "home",
        "json": {},
        "null": {},
        "teamColor": "1d428a"
      },
      "4870562": {
        "homeAway": "home",
        "json": {},
        "null": {},
        "teamColor": "1d428a"
      },
      "3062679": {
        "homeAway": "away",
        "json": {},
        "null": {},
        "teamColor": "1d428a"
      },
      "2528426": {
        "homeAway": "away",
        "json": {},
        "null": {},
        "teamColor": "1d428a"
      },
      "3136195": {
        "homeAway": "away",
        "json": {},
        "null": {},
        "teamColor": "1d428a"
      },
      "5211983": {
        "homeAway": "away",
        "json": {},
        "null": {},
        "teamColor": "1d428a"
      },
      "4432446": {
        "homeAway": "home",
        "json": {},
        "null": {},
        "teamColor": "1d428a"
      },
      "4431675": {
        "homeAway": "home",
        "json": {},
        "null": {},
        "teamColor": "1d428a"
      },
      "4433159": {
        "homeAway": "home",
        "json": {},
        "null": {},
        "teamColor": "1d428a"
      },
      "4351852": {
        "homeAway": "away",
        "json": {},
        "null": {},
        "teamColor": "1d428a"
      },
      "5124612": {
        "homeAway": "home",
        "json": {},
        "null": {},
        "teamColor": "1d428a"
      },
      "3133603": {
        "homeAway": "home",
        "json": {},
        "null": {},
        "teamColor": "1d428a"
      }
    },
    "awayTeam": {
      "id": "18",
      "uid": "s:40~l:46~t:18",
      "team": {
        "id": "18",
        "uid": "s:40~l:46~t:18",
        "guid": "61719eb2-11c3-4e3d-90c3-0a1319fd850b",
        "name": "Knicks",
        "displayName": "New York Knicks",
        "abbreviation": "NY",
        "alternateColor": "f58426",
        "color": "1d428a",
        "location": "New York",
        "links": [],
        "logos": []
      },
      "score": "108",
      "linescores": [
        {},
        {}
      ],
      "record": [
        {},
        {}
      ],
      "winner": true,
      "homeAway": "away",
      "possession": false,
      "order": 1
    },
    "homeTeam": {
      "id": "20",
      "uid": "s:40~l:46~t:20",
      "team": {
        "id": "20",
        "uid": "s:40~l:46~t:20",
        "guid": "ca1685ed-b799-53e4-7924-e58ea6eb8f3a",
        "name": "76ers",
        "displayName": "Philadelphia 76ers",
        "abbreviation": "PHI",
        "alternateColor": "e01234",
        "color": "1d428a",
        "location": "Philadelphia",
        "links": [],
        "logos": []
      },
      "score": "94",
      "linescores": [
        {},
        {}
      ],
      "record": [
        {},
        {}
      ],
      "winner": false,
      "homeAway": "home",
      "possession": false,
      "order": 0
    },
    "awayTeamLogo": "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nba/500/ny.png&h=100&w=100",
    "homeTeamLogo": "https://a.espncdn.com/combiner/i?img=/i/teamlogos/nba/500/phi.png&h=100&w=100",
    "highlightPlayers": false,
    "airingsHash": {
      "onWatch": false,
      "onDTC": false,
      "gameOnEPlus": false,
      "isOOM": false,
      "airingsAll": [],
      "airingsTVE": [],
      "airingsDTC": [],
      "networkHashTVE": {},
      "networkHashDTC": {},
      "userIsEntitledTVE": false,
      "userIsEntitledDTC": false
    }
  },
  "analytics": {
    "metrics": {
      "league": "nba",
      "page_url": "/nba/game/_/gameId/401871161",
      "site": "espn",
      "game_state": "post",
      "content_type": "gamecast",
      "page_infrastructure": "sCore",
      "page_type": "game",
      "game_detail": "401871161 New York Knicks vs Philadelphia 76ers",
      "page_name": "espn:nba:game:gamecast",
      "section": "basketball",
      "sport": "basketball"
    },
    "omniture": {
      "league": "nba",
      "gameInfo": "401871161 New York Knicks vs Philadelphia 76ers",
      "countryRegion": "en-us",
      "hier1": "nba:game:gamecast",
      "section": "basketball",
      "pageName": "nba:game:gamecast",
      "sections": "basketball:gamecast",
      "site": "espn",
      "premium": "premium-no",
      "appearance": "light",
      "convrSport": "basketball",
      "pageURL": "www.espn.com/nba/game/_/gameId/401871161"
    },
    "chartbeat": {
      "domain": "www.espn.com",
      "sections": "basketball",
      "authors": "gamecast",
      "path": "/nba/game/_/gameId/401871161",
      "title": "Knicks vs. 76ers - Game Summary - May 8, 2026 - ESPN",
      "zone": "www.espn.com.us.basketball",
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
  "ads": {
    "id": 12129264,
    "page_url": "https://www.espn.com/nba/game/_/gameId/401871161",
    "prebidAdConfig": {
      "usePrebidBids": true,
      "timeout": 1000
    },
    "level": "espn.com/nba/game",
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
  "targeting": {},
  "meta": {
    "type": "gamepackage",
    "imageWidth": 576,
    "image": "https://s.espncdn.com/stitcher/sports/basketball/nba/events/401871161.png?templateId=espn.com.share.1",
    "twitter_card": "summary",
    "og_site_name": "ESPN.com",
    "twitter_app_id_iphone": "317469184",
    "description": "Get a summary of the New York Knicks vs. Philadelphia 76ers basketball game",
    "og_type": "website",
    "twitter_app_name_googleplay": "ESPN",
    "label": "NBA",
    "canonical": "https://www.espn.com/nba/game/_/gameId/401871161",
    "title": "Knicks vs. 76ers - Game Summary - May 8, 2026 - ESPN"
  },
  "nowFeedSupported": true,
  "customNav": "<div id=\"gamepackage-header-wrap\" class=\" details-header\"><div id=\"gamepackage-matchup-wrap\"><header class=\"game-strip game-package nba post away-winner\"><div class=\"game-detail..."
}
```
