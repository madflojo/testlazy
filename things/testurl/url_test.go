package testurl

import (
	"net/url"
	"testing"
)

const (
	testExampleHost   = "example.com"
	testLocalhostHost = "localhost"
)

// TestPredefinedURLFunctions verifies all predefined URL functions return the expected URL string.
func TestPredefinedURLFunctions(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		fn       func() *url.URL
		expected string
	}{
		{"URLHTTPBad", URLHTTPBad, "http://bad-url"},
		{"URLHTTPSBad", URLHTTPSBad, "https://bad-url"},
		{"URLNoScheme", URLNoScheme, "//" + testExampleHost},
		{"URLNoHost", URLNoHost, "http:"},
		{"URLInvalidHost", URLInvalidHost, "http://%E0%A4%A4%E0%A4%A5"},
		{"URLHTTP", URLHTTP, "http://" + testExampleHost + "/"},
		{"URLHTTPS", URLHTTPS, "https://" + testExampleHost + "/"},
		{"URLHTTPWithPort", URLHTTPWithPort, "http://" + testExampleHost + ":8080/"},
		{"URLHTTPSWithPort", URLHTTPSWithPort, "https://" + testExampleHost + ":8443/"},
		{"URLHTTPWithQuery", URLHTTPWithQuery, "http://" + testExampleHost + "/?query=1"},
		{"URLHTTPSWithQuery", URLHTTPSWithQuery, "https://" + testExampleHost + "/?query=1"},
		{"URLHTTPWithPath", URLHTTPWithPath, "http://" + testExampleHost + "/path/to/resource"},
		{"URLHTTPSWithPath", URLHTTPSWithPath, "https://" + testExampleHost + "/path/to/resource"},
		{
			"URLHTTPWithPathAndQuery",
			URLHTTPWithPathAndQuery,
			"http://" + testExampleHost + "/path/to/resource?query=1",
		},
		{
			"URLHTTPSWithPathAndQuery",
			URLHTTPSWithPathAndQuery,
			"https://" + testExampleHost + "/path/to/resource?query=1",
		},
		{"URLHTTPWithFragment", URLHTTPWithFragment, "http://" + testExampleHost + "/#fragment"},
		{"URLHTTPSWithFragment", URLHTTPSWithFragment, "https://" + testExampleHost + "/#fragment"},
		{"URLHTTPLocalhost", URLHTTPLocalhost, "http://" + testLocalhostHost + "/"},
		{"URLHTTPSLocalhost", URLHTTPSLocalhost, "https://" + testLocalhostHost + "/"},
		{
			"URLHTTPLocalhostWithPort",
			URLHTTPLocalhostWithPort,
			"http://" + testLocalhostHost + ":8080/",
		},
		{
			"URLHTTPSLocalhostWithPort",
			URLHTTPSLocalhostWithPort,
			"https://" + testLocalhostHost + ":8443/",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			u := tc.fn()
			if u.String() != tc.expected {
				t.Errorf("Expected %s to be %q, got %q", tc.name, tc.expected, u.String())
			}
		})
	}
}

type mustParseTestCase struct {
	name  string
	url   string
	panic bool
}

func TestMustParse(t *testing.T) {
	t.Parallel()
	tt := []mustParseTestCase{
		{
			name:  "Valid URL",
			url:   "http://example.com",
			panic: false,
		},
		{
			name:  "Malformed URL",
			url:   "http://%E0%A4%A",
			panic: true,
		},
		{
			name:  "Empty URL",
			url:   "",
			panic: false,
		},
		{
			name:  "Valid HTTPS URL",
			url:   "https://example.com",
			panic: false,
		},
		{
			name:  "Valid URL with path",
			url:   "http://example.com/path/to/resource",
			panic: false,
		}}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if tc.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic for URL '%s', but did not panic", tc.url)
					}
				}()
			}
			u := MustParse(tc.url)
			if !tc.panic && u.String() != tc.url {
				t.Errorf("Expected URL '%s', got '%s'", tc.url, u.String())
			}
		})
	}
}
