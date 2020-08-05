package client

import (
	"net/http"
	"net/url"
)

type (
	// Option applies some option to clientOptions.
	Option func(opts *clientOptions)
)

// WithHttpClient sets custom http.Client.
func WithHttpClient(c *http.Client) Option {
	return func(opts *clientOptions) {
		opts.httpClient = c
	}
}

// WithHttpClient sets custom http.Client.
func WithEndpointURL(endpointURL string) Option {
	endpointURLParsed, err := url.Parse(endpointURL)
	if err != nil { // TODO(Smoreg): BC 2 no panic
		panic(err)
	}
	return func(opts *clientOptions) {
		opts.endpointURL = endpointURLParsed
	}
}

// WithCredentialProvider sets credential provider.
func WithCredentialProvider(c CredentialProvider) Option {
	return func(opts *clientOptions) {
		opts.credentialProvider = c
	}
}

func applyOptions(options *clientOptions, opts ...Option) {
	for _, o := range opts {
		o(options)
	}
}
