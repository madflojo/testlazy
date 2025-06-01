package testurl

import (
	"testing"
)

func TestVars(t *testing.T) {
	t.Parallel()
	t.Run("URLHTTP", func(t *testing.T) {
		u := URLHTTP
		if u.Scheme != "http" {
			t.Errorf("Expected scheme 'http', got '%s'", u.Scheme)
		}
		if u.Host != "example.com" {
			t.Errorf("Expected host 'example.com', got '%s'", u.Host)
		}
		if u.Path != "/" {
			t.Errorf("Expected path '/', got '%s'", u.Path)
		}
	})

	t.Run("URLHTTPS", func(t *testing.T) {
		u := URLHTTPS
		if u.Scheme != "https" {
			t.Errorf("Expected scheme 'https', got '%s'", u.Scheme)
		}
		if u.Host != "example.com" {
			t.Errorf("Expected host 'example.com', got '%s'", u.Host)
		}
		if u.Path != "/" {
			t.Errorf("Expected path '/', got '%s'", u.Path)
		}
	})

	t.Run("URLMalformed", func(t *testing.T) {
		u := URLMalformed
		if u.Scheme != "http" {
			t.Errorf("Expected scheme 'http', got '%s'", u.Scheme)
		}
		if u.Host != "%E0%A4%A" {
			t.Errorf("Expected host '%s', got '%s'", "%E0%A4%A", u.Host)
		}
		if u.Path != "" {
			t.Errorf("Expected path '', got '%s'", u.Path)
		}
		if u.RawQuery != "" {
			t.Errorf("Expected query '', got '%s'", u.RawQuery)
		}
		if u.Fragment != "" {
			t.Errorf("Expected fragment '', got '%s'", u.Fragment)
		}
	})
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
