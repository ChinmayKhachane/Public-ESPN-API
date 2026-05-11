package common

import (
	"context"
	"net/url"
	"strings"

	"github.com/pseudo-r/Public-ESPN-API/espn-go/endpoints"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/internal/httpclient"
	"github.com/pseudo-r/Public-ESPN-API/espn-go/models"
)

// Client wraps site.web.api.espn.com common-v3 sports resources.
type Client struct {
	transport *httpclient.Client
	resolver  endpoints.Resolver
	sport     string
	league    string
}

func New(transport *httpclient.Client, resolver endpoints.Resolver, sport, league string) *Client {
	return &Client{transport: transport, resolver: resolver, sport: sport, league: league}
}

func (c *Client) get(ctx context.Context, suffix string, q url.Values) (models.JSON, error) {
	return c.transport.GetJSON(ctx, c.resolver.CommonV3(c.sport, c.league, strings.TrimLeft(suffix, "/"), cloneValues(q)))
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

func (c *Client) Athlete(ctx context.Context, athleteID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "athletes/"+athleteID, q)
}

func (c *Client) AthleteOverview(ctx context.Context, athleteID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "athletes/"+athleteID+"/overview", q)
}

func (c *Client) AthleteStats(ctx context.Context, athleteID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "athletes/"+athleteID+"/stats", q)
}

func (c *Client) AthleteGameLog(ctx context.Context, athleteID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "athletes/"+athleteID+"/gamelog", q)
}

func (c *Client) AthleteSplits(ctx context.Context, athleteID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "athletes/"+athleteID+"/splits", q)
}

func (c *Client) AthleteBio(ctx context.Context, athleteID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "athletes/"+athleteID+"/bio", q)
}

func (c *Client) TeamRoster(ctx context.Context, teamID string, q url.Values) (models.JSON, error) {
	return c.get(ctx, "teams/"+teamID+"/roster", q)
}

func (c *Client) Athletes(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "athletes", q)
}

func (c *Client) StatisticsByAthlete(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "statistics/byathlete", q)
}

func (c *Client) StatisticsByTeam(ctx context.Context, q url.Values) (models.JSON, error) {
	return c.get(ctx, "statistics/byteam", q)
}

func (c *Client) URL(suffix string, q url.Values) string {
	return c.resolver.CommonV3(c.sport, c.league, strings.TrimLeft(suffix, "/"), cloneValues(q))
}
