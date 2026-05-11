package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/pseudo-r/Public-ESPN-API/espn-go/models"
)

// Client performs ESPN HTTP requests.
type Client struct {
	httpClient *http.Client
	userAgent  string
}

// Option configures Client.
type Option func(*Client)

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		if httpClient != nil {
			c.httpClient = httpClient
		}
	}
}

func WithUserAgent(userAgent string) Option {
	return func(c *Client) {
		if userAgent != "" {
			c.userAgent = userAgent
		}
	}
}

func New(opts ...Option) *Client {
	c := &Client{
		httpClient: &http.Client{Timeout: 30 * time.Second},
		userAgent:  "espn-go/0.1",
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Client) GetJSON(ctx context.Context, requestURL string) (models.JSON, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, &models.ESPNError{
			StatusCode: resp.StatusCode,
			URL:        requestURL,
			Body:       string(body),
		}
	}

	var out models.JSON
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, fmt.Errorf("decode espn response %s: %w", requestURL, err)
	}
	return out, nil
}
