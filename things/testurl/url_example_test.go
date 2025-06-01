package testurl

import (
	"fmt"
)

func ExampleURLHTTPBad() {
	fmt.Println(URLHTTPBad().String())
	// Output: http://bad-url
}

func ExampleURLHTTPSBad() {
	fmt.Println(URLHTTPSBad().String())
	// Output: https://bad-url
}

func ExampleURLNoScheme() {
	fmt.Println(URLNoScheme().String())
	// Output: //example.com
}

func ExampleURLNoHost() {
	fmt.Println(URLNoHost().String())
	// Output: http:
}

func ExampleURLInvalidHost() {
	fmt.Println(URLInvalidHost().String())
	// Output: http://%E0%A4%A4%E0%A4%A5
}

func ExampleURLHTTP() {
	fmt.Println(URLHTTP().String())
	// Output: http://example.com/
}

func ExampleURLHTTPS() {
	fmt.Println(URLHTTPS().String())
	// Output: https://example.com/
}

func ExampleURLHTTPWithPort() {
	fmt.Println(URLHTTPWithPort().String())
	// Output: http://example.com:8080/
}

func ExampleURLHTTPSWithPort() {
	fmt.Println(URLHTTPSWithPort().String())
	// Output: https://example.com:8443/
}

func ExampleURLHTTPWithQuery() {
	fmt.Println(URLHTTPWithQuery().String())
	// Output: http://example.com/?query=1
}

func ExampleURLHTTPSWithQuery() {
	fmt.Println(URLHTTPSWithQuery().String())
	// Output: https://example.com/?query=1
}

func ExampleURLHTTPWithPath() {
	fmt.Println(URLHTTPWithPath().String())
	// Output: http://example.com/path/to/resource
}

func ExampleURLHTTPSWithPath() {
	fmt.Println(URLHTTPSWithPath().String())
	// Output: https://example.com/path/to/resource
}

func ExampleURLHTTPWithPathAndQuery() {
	fmt.Println(URLHTTPWithPathAndQuery().String())
	// Output: http://example.com/path/to/resource?query=1
}

func ExampleURLHTTPSWithPathAndQuery() {
	fmt.Println(URLHTTPSWithPathAndQuery().String())
	// Output: https://example.com/path/to/resource?query=1
}

func ExampleURLHTTPWithFragment() {
	fmt.Println(URLHTTPWithFragment().String())
	// Output: http://example.com/#fragment
}

func ExampleURLHTTPSWithFragment() {
	fmt.Println(URLHTTPSWithFragment().String())
	// Output: https://example.com/#fragment
}

func ExampleURLHTTPLocalhost() {
	fmt.Println(URLHTTPLocalhost().String())
	// Output: http://localhost/
}

func ExampleURLHTTPSLocalhost() {
	fmt.Println(URLHTTPSLocalhost().String())
	// Output: https://localhost/
}

func ExampleURLHTTPLocalhostWithPort() {
	fmt.Println(URLHTTPLocalhostWithPort().String())
	// Output: http://localhost:8080/
}

func ExampleURLHTTPSLocalhostWithPort() {
	fmt.Println(URLHTTPSLocalhostWithPort().String())
	// Output: https://localhost:8443/
}

// ExampleMustParse demonstrates parsing a URL string into a *url.URL.
func ExampleMustParse() {
	fmt.Println(MustParse("http://example.com/resource").String())
	// Output: http://example.com/resource
}

// ExampleMalformedURL demonstrates creating a malformed *url.URL without parsing.
func ExampleMalformedURL() {
	fmt.Println(MalformedURL("http", "bad-host").String())
	// Output: http://bad-host
}
