package itunes

// WithClient allows for override of http.DefaultClient
func WithClient(inputClient Doer) Option {
	return func(client *Client) {
		client.client = inputClient
	}
}

// SetDomain allows for the overwrite of default domain
func SetDomain(domain string) Option {
	return func(client *Client) {
		client.domain = domain
	}
}

// Option allows overriding of default properties and configuration of client
type Option func(client *Client)
