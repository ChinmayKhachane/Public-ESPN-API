package espn

import (
	"net/http"

	"github.com/pseudo-r/Public-ESPN-API/espn-go/endpoints"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/football"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/internal/httpclient"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/models"
)

type (
	BaseURLs             = endpoints.BaseURLs
	JSON                 = models.JSON
	ESPNError            = models.ESPNError
	PlayerHit            = models.PlayerHit
	TeamHit              = models.TeamHit
	ResolvedGame         = models.ResolvedGame
	StatLine             = models.StatLine
	PlayerGameStatsInput = models.PlayerGameStatsInput
	PlayerGameStats      = models.PlayerGameStats
	GameSearch           = models.GameSearch
	IDBundle             = models.IDBundle
	Source               = football.Source
	DataNeed             = football.DataNeed
	SourcePreference     = football.SourcePreference
)

type config struct {
	base     endpoints.BaseURLs
	httpOpts []httpclient.Option
}

// Option configures the root ESPN client.
type Option func(*config)

func WithBaseURLs(base BaseURLs) Option {
	return func(cfg *config) {
		cfg.base = base
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(cfg *config) {
		cfg.httpOpts = append(cfg.httpOpts, httpclient.WithHTTPClient(client))
	}
}

func WithUserAgent(userAgent string) Option {
	return func(cfg *config) {
		cfg.httpOpts = append(cfg.httpOpts, httpclient.WithUserAgent(userAgent))
	}
}

// Client is the root SDK facade. Sport-specific services own the higher-level
// workflows, while domain packages remain importable for raw ESPN payloads.
type Client struct {
	transport *httpclient.Client
	resolver  endpoints.Resolver
}

func NewClient(opts ...Option) *Client {
	cfg := config{}
	for _, opt := range opts {
		opt(&cfg)
	}
	return &Client{
		transport: httpclient.New(cfg.httpOpts...),
		resolver:  endpoints.NewResolver(cfg.base),
	}
}

func (c *Client) Football(league string) *football.Service {
	return football.New(c.transport, c.resolver, league)
}

func ParseAthleteID(uid string) string {
	return football.ParseAthleteID(uid)
}
