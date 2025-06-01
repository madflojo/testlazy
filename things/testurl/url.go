/*
Package testurl provides a set of predefined URL functions for testing HTTP clients and servers.

It includes common URL variants for HTTP and HTTPS:
  - Base URLs (example.com) with and without trailing slash.
  - URLs with ports, query parameters, paths, and fragments.
  - Localhost URLs with optional ports.
  - A malformed URL for testing error handling.

Example usage:

	req := &http.Request{
	    Method: "GET",
	    URL:    testurl.URLHTTPS(),
	}

Use the documented functions directly in your tests.
*/
package testurl

import (
	"fmt"
	"net/url"
)

const (
	exampleHost   = "example.com"
	localhostHost = "localhost"
)

// URLHTTPBad returns a *url.URL for "http://bad-url" which is technically valid but "should" not resolve correctly.
func URLHTTPBad() *url.URL {
	return MalformedURL("http", "bad-url")
}

// URLHTTPSBad returns a *url.URL for "https://bad-url" which is technically valid but "should" not resolve correctly.
func URLHTTPSBad() *url.URL {
	return MalformedURL("https", "bad-url")
}

// URLNoScheme returns a *url.URL for "example.com" (no scheme).
func URLNoScheme() *url.URL {
	return MalformedURL("", exampleHost)
}

// URLNoHost returns a *url.URL for "http://".
func URLNoHost() *url.URL {
	return MalformedURL("http", "")
}

// URLInvalidHost returns a *url.URL for "http://%E0%A4%A4%E0%A4%A5".
func URLInvalidHost() *url.URL {
	return MalformedURL("http", string([]byte{0xE0, 0xA4, 0xA4, 0xE0, 0xA4, 0xA5}))
}

// URLHTTP returns a *url.URL for "http://example.com/".
func URLHTTP() *url.URL {
	return MustParse("http://" + exampleHost + "/")
}

// URLHTTPS returns a *url.URL for "https://example.com/".
func URLHTTPS() *url.URL {
	return MustParse("https://" + exampleHost + "/")
}

// URLHTTPWithPort returns a *url.URL for "http://example.com:8080/".
func URLHTTPWithPort() *url.URL {
	return MustParse("http://" + exampleHost + ":8080/")
}

// URLHTTPSWithPort returns a *url.URL for "https://example.com:8443/".
func URLHTTPSWithPort() *url.URL {
	return MustParse("https://" + exampleHost + ":8443/")
}

// URLHTTPWithQuery returns a *url.URL for "http://example.com/?query=1".
func URLHTTPWithQuery() *url.URL {
	return MustParse("http://" + exampleHost + "/?query=1")
}

// URLHTTPSWithQuery returns a *url.URL for "https://example.com/?query=1".
func URLHTTPSWithQuery() *url.URL {
	return MustParse("https://" + exampleHost + "/?query=1")
}

// URLHTTPWithPath returns a *url.URL for "http://example.com/path/to/resource".
func URLHTTPWithPath() *url.URL {
	return MustParse("http://" + exampleHost + "/path/to/resource")
}

// URLHTTPSWithPath returns a *url.URL for "https://example.com/path/to/resource".
func URLHTTPSWithPath() *url.URL {
	return MustParse("https://" + exampleHost + "/path/to/resource")
}

// URLHTTPWithPathAndQuery returns a *url.URL for "http://example.com/path/to/resource?query=1".
func URLHTTPWithPathAndQuery() *url.URL {
	return MustParse("http://" + exampleHost + "/path/to/resource?query=1")
}

// URLHTTPSWithPathAndQuery returns a *url.URL for "https://example.com/path/to/resource?query=1".
func URLHTTPSWithPathAndQuery() *url.URL {
	return MustParse("https://" + exampleHost + "/path/to/resource?query=1")
}

// URLHTTPWithFragment returns a *url.URL for "http://example.com/#fragment".
func URLHTTPWithFragment() *url.URL {
	return MustParse("http://" + exampleHost + "/#fragment")
}

// URLHTTPSWithFragment returns a *url.URL for "https://example.com/#fragment".
func URLHTTPSWithFragment() *url.URL {
	return MustParse("https://" + exampleHost + "/#fragment")
}

// URLHTTPLocalhost returns a *url.URL for "http://localhost/".
func URLHTTPLocalhost() *url.URL {
	return MustParse("http://" + localhostHost + "/")
}

// URLHTTPSLocalhost returns a *url.URL for "https://localhost/".
func URLHTTPSLocalhost() *url.URL {
	return MustParse("https://" + localhostHost + "/")
}

// URLHTTPLocalhostWithPort returns a *url.URL for "http://localhost:8080/".
func URLHTTPLocalhostWithPort() *url.URL {
	return MustParse("http://" + localhostHost + ":8080/")
}

// URLHTTPSLocalhostWithPort returns a *url.URL for "https://localhost:8443/".
func URLHTTPSLocalhostWithPort() *url.URL {
	return MustParse("https://" + localhostHost + ":8443/")
}

// MustParse is a helper function that parses a URL string and panics if it fails.
// This is useful for defining test URLs, but should be used with caution.
func MustParse(u string) *url.URL {
	v, err := url.Parse(u)
	if err != nil {
		panic(fmt.Sprintf("failed to parse URL %q: %v", u, err))
	}
	return v
}

// MalformedURL returns a malformed URL for testing error handling.
func MalformedURL(scheme, host string) *url.URL {
	return &url.URL{
		Scheme: scheme,
		Host:   host,
	}
}
