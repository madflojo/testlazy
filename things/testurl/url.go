/*
Package testurl provides a set of predefined URLs for testing purposes.
*/
package testurl

import (
	"net/url"
  "fmt"
)

var (
	URLHTTP                   = MustParse("http://example.com/")
	URLHTTPS                  = MustParse("https://example.com/")
	URLHTTPWithPort           = MustParse("http://example.com:8080/")
	URLHTTPSWithPort          = MustParse("https://example.com:8443/")
	URLHTTPWithQuery          = MustParse("http://example.com/?query=1")
	URLHTTPSWithQuery         = MustParse("https://example.com/?query=1")
	URLHTTPWithPath           = MustParse("http://example.com/path/to/resource")
	URLHTTPSWithPath          = MustParse("https://example.com/path/to/resource")
	URLHTTPWithPathAndQuery   = MustParse("http://example.com/path/to/resource?query=1")
	URLHTTPSWithPathAndQuery  = MustParse("https://example.com/path/to/resource?query=1")
	URLHTTPWithFragment       = MustParse("http://example.com/#fragment")
	URLHTTPSWithFragment      = MustParse("https://example.com/#fragment")
	URLHTTPLocalhost          = MustParse("http://localhost/")
	URLHTTPSLocalhost         = MustParse("https://localhost/")
	URLHTTPLocalhostWithPort  = MustParse("http://localhost:8080/")
	URLHTTPSLocalhostWithPort = MustParse("https://localhost:8443/")
	URLMalformed              = urlMalformed()
)

func MustParse(u string) *url.URL {
	v, err := url.Parse(u)
	if err != nil {
		panic(fmt.Sprintf("failed to parse URL %q: %v", u, err))
	}
	return v
}


func urlMalformed() *url.URL {
  return &url.URL{
    Scheme:   "http",
    Host:     "%E0%A4%A",
  }
}
