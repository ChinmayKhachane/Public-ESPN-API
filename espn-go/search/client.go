package search

import (
	"context"
	"net/url"

	"github.com/pseudo-r/Public-ESPN-API/espn-go/endpoints"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/internal/httpclient"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/models"
)

// Client wraps site.web.api.espn.com search and scoreboard-header endpoints.
type Client struct {
	transport *httpclient.Client
	resolver  endpoints.Resolver
	sport     string
	league    string
}

func New(transport *httpclient.Client, resolver endpoints.Resolver, sport, league string) *Client {
	return &Client{transport: transport, resolver: resolver, sport: sport, league: league}
}

func cloneValues(q url.Values) url.Values {
	if q == nil {
		return nil
	}
	out := make(url.Values, len(q))
	for key, values := range q {
		out[key] = append([]string(nil), values...)
	}
	return out
}

func (c *Client) Query(ctx context.Context, query string, q url.Values) (models.JSON, error) {
	values := cloneValues(q)
	if values == nil {
		values = url.Values{}
	}
	values.Set("query", query)
	if values.Get("sport") == "" {
		values.Set("sport", c.sport)
	}
	if values.Get("league") == "" && c.league != "" {
		values.Set("league", c.league)
	}
	return c.transport.GetJSON(ctx, c.resolver.SearchV2(values))
}

func (c *Client) ScoreboardHeader(ctx context.Context, q url.Values) (models.JSON, error) {
	values := cloneValues(q)
	if values == nil {
		values = url.Values{}
	}
	if values.Get("sport") == "" {
		values.Set("sport", c.sport)
	}
	if values.Get("league") == "" && c.league != "" {
		values.Set("league", c.league)
	}
	return c.transport.GetJSON(ctx, c.resolver.ScoreboardHeader(values))
}
