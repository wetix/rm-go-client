package rm

import "golang.org/x/oauth2"

// Option :
type Option func(*Client)

// WithTokenSource :
func WithTokenSource(cb oauth2.TokenSource) Option {
	return func(c *Client) {
		c.oauth2 = cb
	}
}
