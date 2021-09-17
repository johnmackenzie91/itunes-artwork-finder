package itunes

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

func generateSearchURL(domain, term, country, entity string) (*url.URL, error) {
	u, err := url.Parse(domain)

	if err != nil {
		return nil, err
	}

	v := url.Values{}
	v.Add("term", term)
	v.Add("country", country)
	v.Add("entity", entity)

	encoded := v.Encode()
	u.RawQuery = encoded

	return u, nil
}

func (c Client) buildSearchRequest(ctx context.Context, term, country, entity string) (*http.Request, error) {
	u, err := generateSearchURL(c.domain, term, country, entity)

	if err != nil {
		return nil, fmt.Errorf("failed to generate url: %w", err)
	}
	return http.NewRequestWithContext(ctx, "GET", u.String(), nil)
}

func (c Client) attemptRequest(r *http.Request) (*http.Response, error) {
	c.logger.Debug("attempting  itunes search request", r)

	res, err := c.client.Do(r)

	if err != nil {
		return nil, fmt.Errorf("search itunes failed with network error: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, errNon2XX(res.StatusCode)
	}
	return res, nil
}
