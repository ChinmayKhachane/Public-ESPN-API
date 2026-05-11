package now

import (
	"context"
	"net/url"

	"github.com/pseudo-r/Public-ESPN-API/espn-go/endpoints"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/internal/httpclient"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/models"
)

// Client wraps now.core.api.espn.com news endpoints.
type Client struct {
	transport *httpclient.Client
	resolver  endpoints.Resolver
	sport     string
}

func New(transport *httpclient.Client, resolver endpoints.Resolver, sport string) *Client {
	return &Client{transport: transport, resolver: resolver, sport: sport}
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

func (c *Client) Headlines(ctx context.Context, q url.Values) (models.JSON, error) {
	values := cloneValues(q)
	if values == nil {
		values = url.Values{}
	}
	if values.Get("sport") == "" {
		values.Set("sport", c.sport)
	}
	return c.transport.GetJSON(ctx, c.resolver.NowNews(values))
}
