# Athlete ID

## https://sports.core.api.espn.com/v2/sports/football/leagues/{league}/athletes/{id}


## Optional Parameters

| Param | Type | Description|
| --- | --- | --- |
|  |  |  |

## Example Response

```json
{
  "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/athletes/4429202?lang=en&region=us",
  "id": "4429202",
  "uid": "s:20~l:28~a:4429202",
  "guid": "2fcbb6af-36f8-3b49-9438-3fdef938d6da",
  "type": "football",
  "alternateIds": {
    "sdr": "4429202"
  },
  "firstName": "Israel",
  "lastName": "Abanikanda",
  "fullName": "Israel Abanikanda",
  "displayName": "Israel Abanikanda",
  "shortName": "I. Abanikanda",
  "weight": 215,
  "displayWeight": "215 lbs",
  "height": 71,
  "displayHeight": "5' 11\"",
  "age": 23,
  "dateOfBirth": "2002-10-05T07:00Z",
  "links": [
    {
      "language": "en-US",
      "rel": [
        "playercard",
        "desktop",
        "athlete"
      ],
      "href": "https://www.espn.com/nfl/player/_/id/4429202/israel-abanikanda",
      "text": "Player Card",
      "shortText": "Player Card",
      "isExternal": false,
      "isPremium": false
    },
    {
      "language": "en-US",
      "rel": [
        "stats",
        "desktop",
        "athlete"
      ],
      "href": "https://www.espn.com/nfl/player/stats/_/id/4429202/israel-abanikanda",
      "text": "Stats",
      "shortText": "Stats",
      "isExternal": false,
      "isPremium": false
    },
    {
      "language": "en-US",
      "rel": [
        "splits",
        "desktop",
        "athlete"
      ],
      "href": "https://www.espn.com/nfl/player/splits/_/id/4429202/israel-abanikanda",
      "text": "Splits",
      "shortText": "Splits",
      "isExternal": false,
      "isPremium": false
    },
    {
      "language": "en-US",
      "rel": [
        "gamelog",
        "desktop",
        "athlete"
      ],
      "href": "https://www.espn.com/nfl/player/gamelog/_/id/4429202/israel-abanikanda",
      "text": "Game Log",
      "shortText": "Game Log",
      "isExternal": false,
      "isPremium": false
    },
    {
      "language": "en-US",
      "rel": [
        "news",
        "desktop",
        "athlete"
      ],
      "href": "https://www.espn.com/nfl/player/news/_/id/4429202/israel-abanikanda",
      "text": "News",
      "shortText": "News",
      "isExternal": false,
      "isPremium": false
    },
    {
      "language": "en-US",
      "rel": [
        "bio",
        "desktop",
        "athlete"
      ],
      "href": "https://www.espn.com/nfl/player/bio/_/id/4429202/israel-abanikanda",
      "text": "Bio",
      "shortText": "Bio",
      "isExternal": false,
      "isPremium": false
    },
    {
      "language": "en-US",
      "rel": [
        "overview",
        "desktop",
        "athlete"
      ],
      "href": "https://www.espn.com/nfl/player/_/id/4429202/israel-abanikanda",
      "text": "Overview",
      "shortText": "Overview",
      "isExternal": false,
      "isPremium": false
    }
  ],
  "birthPlace": {
    "city": "Brooklyn",
    "state": "NY",
    "country": "USA"
  },
  "college": {
    "$ref": "http://sports.core.api.espn.com/v2/colleges/221?lang=en&region=us"
  },
  "slug": "israel-abanikanda",
  "headshot": {
    "href": "https://a.espncdn.com/i/headshots/nfl/players/full/4429202.png",
    "alt": "Israel Abanikanda"
  },
  "jersey": "30",
  "position": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/positions/9?lang=en&region=us",
    "id": "9",
    "name": "Running Back",
    "displayName": "Running Back",
    "abbreviation": "RB",
    "leaf": false,
    "parent": {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/positions/70?lang=en&region=us"
    }
  },
  "linked": true,
  "team": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2026/teams/6?lang=en&region=us"
  },
  "statistics": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/athletes/4429202/statistics?lang=en&region=us"
  },
  "contracts": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/athletes/4429202/contracts?lang=en&region=us"
  },
  "experience": {
    "years": 3
  },
  "collegeAthlete": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/college-football/athletes/4429202?lang=en&region=us"
  },
  "active": true,
  "draft": {
    "displayText": "Year: 2023 Round: 5 Pick: 143",
    "round": 5,
    "year": 2023,
    "selection": 143,
    "team": {
      "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2023/teams/20?lang=en&region=us"
    }
  },
  "status": {
    "id": "1",
    "name": "Active",
    "type": "active",
    "abbreviation": "Active"
  },
  "statisticslog": {
    "$ref": "http://sports.core.api.espn.com/v2/sports/football/leagues/nfl/athletes/4429202/statisticslog?lang=en&region=us"
  }
}
```