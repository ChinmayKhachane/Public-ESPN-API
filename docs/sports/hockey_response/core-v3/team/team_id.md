# Core V3 Team Detail

## https://sports.core.api.espn.com/v3/sports/hockey/{league}/teams/{team}

Notes:
- Verified with `league=nhl`, `season=2026`, `event=401871412`, `competition=401871412`, `competitor=15`, `team=15`, `athlete=4565230`, `play=401871412000000510` on 2026-05-09.
- Live request returned HTTP `200`.
- Content type: `application/json;charset=utf-8`.

Tested URL:
`https://sports.core.api.espn.com/v3/sports/hockey/nhl/teams/15`

## Example Response

```json
{
  "id": "15",
  "uid": "s:70~l:90~t:15",
  "guid": "68aba012-4e93-9371-6861-1bb9a63cfb11",
  "alternateId": "2617",
  "slug": "philadelphia-flyers",
  "location": "Philadelphia",
  "name": "Flyers",
  "nickname": "Philadelphia",
  "abbreviation": "PHI",
  "displayName": "Philadelphia Flyers",
  "shortDisplayName": "Flyers",
  "color": "fe5823",
  "alternateColor": "000000",
  "active": true
}
```
