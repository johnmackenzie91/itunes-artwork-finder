package itunes

import (
	"net/url"

	"github.com/johnmackenzie91/commonlogger"
)

// WithClient allows for override of http.DefaultClient
func WithClient(inputClient Doer) Option {
	return func(c *Client) error {
		c.client = inputClient
		return nil
	}
}

func WithLogger(logger commonlogger.ErrorInfoDebugger) Option {
	return func(c *Client) error {
		c.logger = logger
		return nil
	}
}

// SetDomain allows for the overwrite of default domain
func SetDomain(domain string) Option {
	return func(c *Client) error {
		if domain == "" {
			c.domain = defaultDomain
			return nil
		}
		if _, err := url.ParseRequestURI(domain); err != nil {
			return errBadURL(domain)
		}
		c.domain = domain
		return nil
	}
}

// Option allows overriding of default properties and configuration of client
type Option func(client *Client) error
