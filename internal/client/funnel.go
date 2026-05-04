package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// Provider defines the contract for any external data source.
// The Funnel handles all network operations; the provider only
// describes how to build a request and how to parse the response.
type Provider interface {
	Name() string
	BuildRequest(ctx context.Context, target string) (*http.Request, error)
	ParseResponse(body []byte) (any, error)
}

// Funnel is the common execution path for all network calls.
// It centralises rate limiting, retries, timeouts, and error handling.
type Funnel struct {
	HTTPClient *http.Client
	Limiter    *rate.Limiter
	MaxRetries int
}

// New creates a Funnel with sensible defaults:
//   - 30 second HTTP timeout
//   - 5 requests/second rate limit (burstable)
//   - 3 retries with exponential backoff
func New() *Funnel {
	return &Funnel{
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		Limiter:    rate.NewLimiter(rate.Every(time.Second/5), 5),
		MaxRetries: 3,
	}
}

// Default is the package-level shared Funnel used by all sources.
var Default = New()

// Fetch executes a Provider through the shared funnel.
// It applies rate limiting, retries on 429/5xx, and delegates
// request construction and response parsing to the Provider.
func (f *Funnel) Fetch(ctx context.Context, p Provider, target string) (any, error) {
	var lastErr error

	for attempt := 0; attempt < f.MaxRetries; attempt++ {
		if attempt > 0 {
			backoff := time.Duration(attempt*attempt) * time.Second
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(backoff):
			}
		}

		req, err := p.BuildRequest(ctx, target)
		if err != nil {
			return nil, err
		}

		if err := f.Limiter.Wait(ctx); err != nil {
			return nil, err
		}

		resp, err := f.HTTPClient.Do(req)
		if err != nil {
			lastErr = err
			continue
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			lastErr = err
			continue
		}

		switch {
		case resp.StatusCode == http.StatusTooManyRequests:
			lastErr = fmt.Errorf("%s: rate limited (429)", p.Name())
			continue
		case resp.StatusCode >= 500:
			lastErr = fmt.Errorf("%s: server error %d", p.Name(), resp.StatusCode)
			continue
		case resp.StatusCode == 404:
			return nil, fmt.Errorf("%s: not-found error %d", p.Name(), resp.StatusCode)
		case resp.StatusCode >= 400:
			return nil, fmt.Errorf("%s: HTTP %d: %s", p.Name(), resp.StatusCode, string(body))
		}

		return p.ParseResponse(body)
	}

	return nil, fmt.Errorf("after %d attempts: %w", f.MaxRetries, lastErr)
}
