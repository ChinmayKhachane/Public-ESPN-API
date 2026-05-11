package endpoints

import (
	"net/url"
	"testing"
)

func TestResolverBuildsDomainSpecificURLs(t *testing.T) {
	resolver := NewResolver(BaseURLs{
		Site: "https://site.test/",
		Core: "https://core.test/",
		Web:  "https://web.test/",
		CDN:  "https://cdn.test/",
		Now:  "https://now.test/",
	})

	tests := map[string]string{
		resolver.SiteV2("football", "nfl", "scoreboard", url.Values{"week": {"1"}}):        "https://site.test/apis/site/v2/sports/football/nfl/scoreboard?week=1",
		resolver.SiteStandings("football", "nfl", nil):                                     "https://site.test/apis/v2/sports/football/nfl/standings",
		resolver.CoreV2("football", "nfl", "events/401/competitions/401/plays", nil):       "https://core.test/v2/sports/football/leagues/nfl/events/401/competitions/401/plays",
		resolver.CommonV3("football", "nfl", "athletes/4431452/stats", nil):                "https://web.test/apis/common/v3/sports/football/nfl/athletes/4431452/stats",
		resolver.SearchV2(url.Values{"query": {"drake maye"}, "sport": {"football"}}):      "https://web.test/apis/search/v2?query=drake+maye&sport=football",
		resolver.CDN("nfl", "boxscore", url.Values{"xhr": {"1"}, "gameId": {"401772988"}}): "https://cdn.test/core/nfl/boxscore?gameId=401772988&xhr=1",
		resolver.NowNews(url.Values{"sport": {"football"}}):                                "https://now.test/v1/sports/news?sport=football",
	}

	for got, want := range tests {
		if got != want {
			t.Fatalf("url mismatch\nwant: %s\n got: %s", want, got)
		}
	}
}
