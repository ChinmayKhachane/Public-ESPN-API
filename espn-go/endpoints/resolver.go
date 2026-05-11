package endpoints

import (
	"net/url"
	"strings"
)

const (
	DefaultSiteURL = "https://site.api.espn.com"
	DefaultCoreURL = "https://sports.core.api.espn.com"
	DefaultWebURL  = "https://site.web.api.espn.com"
	DefaultCDNURL  = "https://cdn.espn.com"
	DefaultNowURL  = "https://now.core.api.espn.com"
)

// BaseURLs contains ESPN domain roots.
type BaseURLs struct {
	Site string
	Core string
	Web  string
	CDN  string
	Now  string
}

// Resolver owns ESPN URL construction. Domain clients ask it for URLs instead
// of assembling hostnames themselves.
type Resolver struct {
	base BaseURLs
}

// NewResolver creates a URL resolver with optional base URL overrides.
func NewResolver(base BaseURLs) Resolver {
	defaults := BaseURLs{
		Site: DefaultSiteURL,
		Core: DefaultCoreURL,
		Web:  DefaultWebURL,
		CDN:  DefaultCDNURL,
		Now:  DefaultNowURL,
	}
	if base.Site != "" {
		defaults.Site = base.Site
	}
	if base.Core != "" {
		defaults.Core = base.Core
	}
	if base.Web != "" {
		defaults.Web = base.Web
	}
	if base.CDN != "" {
		defaults.CDN = base.CDN
	}
	if base.Now != "" {
		defaults.Now = base.Now
	}
	defaults.Site = strings.TrimRight(defaults.Site, "/")
	defaults.Core = strings.TrimRight(defaults.Core, "/")
	defaults.Web = strings.TrimRight(defaults.Web, "/")
	defaults.CDN = strings.TrimRight(defaults.CDN, "/")
	defaults.Now = strings.TrimRight(defaults.Now, "/")
	return Resolver{base: defaults}
}

func (r Resolver) build(base, path string, q url.Values) string {
	u := strings.TrimRight(base, "/") + "/" + strings.TrimLeft(path, "/")
	if len(q) == 0 {
		return u
	}
	return u + "?" + q.Encode()
}

func (r Resolver) SiteV2(sport, league, suffix string, q url.Values) string {
	path := "/apis/site/v2/sports/" + sport + "/" + league
	if suffix != "" {
		path += "/" + strings.TrimLeft(suffix, "/")
	}
	return r.build(r.base.Site, path, q)
}

// SiteStandings returns the non-site /apis/v2 standings route. The football
// docs record /apis/site/v2 standings as a web-link stub.
func (r Resolver) SiteStandings(sport, league string, q url.Values) string {
	return r.build(r.base.Site, "/apis/v2/sports/"+sport+"/"+league+"/standings", q)
}

func (r Resolver) CoreV2(sport, league, suffix string, q url.Values) string {
	path := "/v2/sports/" + sport + "/leagues/" + league
	if suffix != "" {
		path += "/" + strings.TrimLeft(suffix, "/")
	}
	return r.build(r.base.Core, path, q)
}

func (r Resolver) CoreV3(sport, league, suffix string, q url.Values) string {
	path := "/v3/sports/" + sport + "/" + league
	if suffix != "" {
		path += "/" + strings.TrimLeft(suffix, "/")
	}
	return r.build(r.base.Core, path, q)
}

func (r Resolver) CommonV3(sport, league, suffix string, q url.Values) string {
	path := "/apis/common/v3/sports/" + sport + "/" + league
	if suffix != "" {
		path += "/" + strings.TrimLeft(suffix, "/")
	}
	return r.build(r.base.Web, path, q)
}

func (r Resolver) SearchV2(q url.Values) string {
	return r.build(r.base.Web, "/apis/search/v2", q)
}

func (r Resolver) ScoreboardHeader(q url.Values) string {
	return r.build(r.base.Web, "/apis/v2/scoreboard/header", q)
}

func (r Resolver) CDN(cdnSport, view string, q url.Values) string {
	return r.build(r.base.CDN, "/core/"+cdnSport+"/"+view, q)
}

func (r Resolver) NowNews(q url.Values) string {
	return r.build(r.base.Now, "/v1/sports/news", q)
}
