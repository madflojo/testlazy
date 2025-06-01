/*
Package testlazy provides an encyclopedia of test values, fakes, and validators for Go.

It helps reduce boilerplate in your tests by offering predefined test values and helpers.

Current components include:
  - things/testurl: Predefined URLs for common use cases.
  - fakes/fakectx: Fake contexts for testing cancellation and timeouts.
  - validators/validhttp: Validators for HTTP status codes and responses.

Quick Start:

URLs without the boilerplate (things/testurl):

	// Instead of parsing URLs manually:
	url, err := url.Parse("https://example.com")
	if err != nil {
	    t.Fatalf("failed to parse URL: %v", err)
	}
	req := &http.Request{
	    Method: "GET",
	    URL:    url,
	}

	// Use testlazy to get predefined URLs:
	req = &http.Request{
	    Method: "GET",
	    URL:    testurl.URLHTTPS(),
	}
	// Or a malformed URL for error handling:
	req = &http.Request{
	    Method: "GET",
	    URL:    testurl.URLInvalidHost(),
	}

Contexts without the hassle (fakes/fakectx):

	// Simulate a canceled context without boilerplate:
	err := doSomething(fakectx.Cancelled())
	if err == nil {
	    t.Fatal("expected error for canceled context, got nil")
	}

For more details and usage examples, see the individual packages under this module.
*/
package testlazy
