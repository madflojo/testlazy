/*
Package testlazy is your test helper sidekick—an ever-growing collection of
canonical values, handy fakes, and smart validators so you can write tests
faster and with less boilerplate.

Why waste precious brainpower on repetitive setup when testlazy has your back?
Pick only the packages you need and focus on your test logic, not the plumbing.

Current modules:
  - things/testurl: Pre-built URL helpers for every HTTP scenario.
  - fakes/fakectx: Easy-to-use cancelled or timed-out contexts. (coming soon)
  - validators: Common checkers for HTTP status codes, headers, and more. (coming soon)

# Quick Start

URLs without the boilerplate (things/testurl):

Instead of manual parsing...

	url, err := url.Parse("https://example.com")
	if err != nil {
	    t.Fatalf("failed to parse URL: %v", err)
	}
	req := &http.Request{
	    Method: "GET",
	    URL:    url,
	}

...just grab a ready-made URL:

	req = &http.Request{
	    Method: "GET",
	    URL:    testurl.URLHTTPS(),
	}

And for a malformed URL to test error handling:

	req = &http.Request{
	    Method: "GET",
	    URL:    testurl.URLInvalidHost(),
	}

Contexts without the hassle (fakes/fakectx):

Instead of manually cancelling contexts...

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := doSomething(ctx)
	if err == nil {
	    t.Fatal("expected error for canceled context, got nil")
	}

...use a one-liner fake:

	err := doSomething(fakectx.Cancelled())
	if err == nil {
	    t.Fatal("expected error for canceled context, got nil")
	}

For more details and examples, see the docs in each package.
*/
package testlazy
