package site

import (
	"context"
	"net/url"
	"strings"

	"github.com/pseudo-r/Public-ESPN-API/espn-go/endpoints"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/internal/httpclient"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/models"
)

// Client wraps site.api.espn.com site-v2 and adjacent site API endpoints.
type Client struct {
	transport *httpclient.Client
	resolver  endpoints.Resolver
	sport     string
	league    string
}

func New(transport *httpclient.Client, resolver endpoints.Resolver, sport, league string) *Client {
	return &Client{
		transport: transport,
		resolver:  resolver,
		sport:     sport,
		league:    league,
	}
}

func (c *Client) get(ctx context.Context, suffix string, q url.Values) (models.JSON, error) {
	return c.transport.GetJSON(ctx, c.resolver.SiteV2(c.sport, c.league, suffix, cloneValues(q)))
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

func (c *Client) Scoreboard(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "scoreboard", q)
}

func (c *Client) Summary(ctx context.Context, eventID string) (models.JSON, error) {
	q := url.Values{"event": {eventID}}
	return c.get(ctx, "summary", q)
}

func (c *Client) Teams(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "teams", q)
}

func (c *Client) Team(ctx context.Context, teamID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "teams/"+teamID, q)
}

func (c *Client) Roster(ctx context.Context, teamID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "teams/"+teamID+"/roster", q)
}

func (c *Client) Schedule(ctx context.Context, teamID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "teams/"+teamID+"/schedule", q)
}

func (c *Client) DepthCharts(ctx context.Context, teamID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "teams/"+teamID+"/depthcharts", q)
}

func (c *Client) AthleteNews(ctx context.Context, athleteID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "athletes/"+athleteID+"/news", q)
}

func (c *Client) Draft(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "draft", q)
}

func (c *Client) Calendar(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "calendar", q)
}

func (c *Client) Groups(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "groups", q)
}

func (c *Client) Injuries(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "injuries", q)
}

func (c *Client) TeamInjuries(ctx context.Context, teamID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "teams/"+teamID+"/injuries", q)
}

func (c *Client) TeamRecord(ctx context.Context, teamID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "teams/"+teamID+"/record", q)
}

func (c *Client) TeamNews(ctx context.Context, teamID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "teams/"+teamID+"/news", q)
}

func (c *Client) TeamLeaders(ctx context.Context, teamID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "teams/"+teamID+"/leaders", q)
}

func (c *Client) TeamHistory(ctx context.Context, teamID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "teams/"+teamID+"/history", q)
}

func (c *Client) TeamTransactions(ctx context.Context, teamID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "teams/"+teamID+"/transactions", q)
}

func (c *Client) News(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "news", q)
}

func (c *Client) Standings(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.transport.GetJSON(ctx, c.resolver.SiteStandings(c.sport, c.league, cloneValues(q)))
}

// StandingsStub calls the site-v2 standings route that football docs observed
// as a link-only response. Use Standings for the populated standings tree.
func (c *Client) StandingsStub(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "standings", q)
}

func (c *Client) Statistics(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "statistics", q)
}

func (c *Client) Rankings(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "rankings", q)
}

func (c *Client) Transactions(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "transactions", q)
}

func (c *Client) URL(suffix string, q url.Values) string {
	return c.resolver.SiteV2(c.sport, c.league, strings.TrimLeft(suffix, "/"), cloneValues(q))
}
