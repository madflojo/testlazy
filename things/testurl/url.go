/*
Package testurl makes URL generation in Go tests a breeze—no more boilerplate
url.Parse calls or manual query-string concat. Simply import testurl and pick
the helper you need to focus on your test logic, not the URL plumbing.

	github.com/madflojo/testlazy/things/testurl

Example usage

	req := &http.Request{
	    Method: "GET",
	    URL:    testurl.URLHTTPS(),
	}

Your tests deserve clarity and speed—let testurl handle the URL creation work.
*/
package testurl

import (
	"fmt"
	"net/url"
)

const (
	ExampleHost   = "example.com"
	LocalhostHost = "localhost"
	HTTPPort      = "8080"
	HTTPSPort     = "8443"
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
	return MalformedURL("", ExampleHost)
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
	return MustParse("http://" + ExampleHost + "/")
}

// URLHTTPS returns a *url.URL for "https://example.com/".
func URLHTTPS() *url.URL {
	return MustParse("https://" + ExampleHost + "/")
}

// URLHTTPWithPort returns a *url.URL for "http://example.com:8080/".
func URLHTTPWithPort() *url.URL {
	return MustParse("http://" + ExampleHost + ":" + HTTPPort + "/")
}

// URLHTTPSWithPort returns a *url.URL for "https://example.com:8443/".
func URLHTTPSWithPort() *url.URL {
	return MustParse("https://" + ExampleHost + ":" + HTTPSPort + "/")
}

// URLHTTPWithQuery returns a *url.URL for "http://example.com/?query=1".
func URLHTTPWithQuery() *url.URL {
	return MustParse("http://" + ExampleHost + "/?query=1")
}

// URLHTTPSWithQuery returns a *url.URL for "https://example.com/?query=1".
func URLHTTPSWithQuery() *url.URL {
	return MustParse("https://" + ExampleHost + "/?query=1")
}

// URLHTTPWithPath returns a *url.URL for "http://example.com/path/to/resource".
func URLHTTPWithPath() *url.URL {
	return MustParse("http://" + ExampleHost + "/path/to/resource")
}

// URLHTTPSWithPath returns a *url.URL for "https://example.com/path/to/resource".
func URLHTTPSWithPath() *url.URL {
	return MustParse("https://" + ExampleHost + "/path/to/resource")
}

// URLHTTPWithPathAndQuery returns a *url.URL for "http://example.com/path/to/resource?query=1".
func URLHTTPWithPathAndQuery() *url.URL {
	return MustParse("http://" + ExampleHost + "/path/to/resource?query=1")
}

// URLHTTPSWithPathAndQuery returns a *url.URL for "https://example.com/path/to/resource?query=1".
func URLHTTPSWithPathAndQuery() *url.URL {
	return MustParse("https://" + ExampleHost + "/path/to/resource?query=1")
}

// URLHTTPWithFragment returns a *url.URL for "http://example.com/#fragment".
func URLHTTPWithFragment() *url.URL {
	return MustParse("http://" + ExampleHost + "/#fragment")
}

// URLHTTPSWithFragment returns a *url.URL for "https://example.com/#fragment".
func URLHTTPSWithFragment() *url.URL {
	return MustParse("https://" + ExampleHost + "/#fragment")
}

// URLHTTPLocalhost returns a *url.URL for "http://localhost/".
func URLHTTPLocalhost() *url.URL {
	return MustParse("http://" + LocalhostHost + "/")
}

// URLHTTPSLocalhost returns a *url.URL for "https://localhost/".
func URLHTTPSLocalhost() *url.URL {
	return MustParse("https://" + LocalhostHost + "/")
}

// URLHTTPLocalhostWithPort returns a *url.URL for "http://localhost:8080/".
func URLHTTPLocalhostWithPort() *url.URL {
	return MustParse("http://" + LocalhostHost + ":" + HTTPPort + "/")
}

// URLHTTPSLocalhostWithPort returns a *url.URL for "https://localhost:8443/".
func URLHTTPSLocalhostWithPort() *url.URL {
	return MustParse("https://" + LocalhostHost + ":" + HTTPSPort + "/")
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
