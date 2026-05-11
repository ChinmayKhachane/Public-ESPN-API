package cdn

import (
	"context"
	"net/url"

	"github.com/pseudo-r/Public-ESPN-API/espn-go/endpoints"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/internal/httpclient"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/models"
)

// Client wraps cdn.espn.com page-package endpoints.
type Client struct {
	transport *httpclient.Client
	resolver  endpoints.Resolver
	cdnSport  string
}

func New(transport *httpclient.Client, resolver endpoints.Resolver, cdnSport string) *Client {
	return &Client{transport: transport, resolver: resolver, cdnSport: cdnSport}
}

func (c *Client) get(ctx context.Context, view string, q url.Values) (models.JSON, error) {
	return c.transport.GetJSON(ctx, c.resolver.CDN(c.cdnSport, view, cloneValues(q)))
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

func eventQuery(eventID string) url.Values {
	return url.Values{"xhr": {"1"}, "gameId": {eventID}}
}

func (c *Client) Scoreboard(ctx context.Context, q url.Values) (models.JSON, error) {
	values := cloneValues(q)
	if values == nil {
		values = url.Values{}
	}
	if values.Get("xhr") == "" {
		values.Set("xhr", "1")
	}
	return c.get(ctx, "scoreboard", values)
}

func (c *Client) Game(ctx context.Context, eventID string) (models.JSON, error) {
	return c.get(ctx, "game", eventQuery(eventID))
}

func (c *Client) Boxscore(ctx context.Context, eventID string) (models.JSON, error) {
	return c.get(ctx, "boxscore", eventQuery(eventID))
}

func (c *Client) Matchup(ctx context.Context, eventID string) (models.JSON, error) {
	return c.get(ctx, "matchup", eventQuery(eventID))
}

func (c *Client) PlayByPlay(ctx context.Context, eventID string) (models.JSON, error) {
	return c.get(ctx, "playbyplay", eventQuery(eventID))
}

func (c *Client) URL(view string, q url.Values) string {
	return c.resolver.CDN(c.cdnSport, view, cloneValues(q))
}
