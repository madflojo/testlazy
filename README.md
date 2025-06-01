# TestLazy 😴

![TestLazy Logo](/docs/img/testlazy.png)

**An encyclopedia of test values, fakes, and validators for Go - so you can type less and test more.**

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/madflojo/testlazy)
[![codecov](https://codecov.io/gh/madflojo/testlazy/branch/main/graph/badge.svg?token=0TTTEWHLVN)](https://codecov.io/gh/madflojo/testlazy)
[![Go Reference](https://pkg.go.dev/badge/github.com/madflojo/testlazy.svg)](https://pkg.go.dev/github.com/madflojo/testlazy)
[![Go Report Card](https://goreportcard.com/badge/github.com/madflojo/testlazy)](https://goreportcard.com/report/github.com/madflojo/testlazy)

---

Are you tired of writing the same three lines of code to create a `url.URL` for example.com?

How about repeatedly copying the same "canceled context" setup to check timeout logic?

Or maybe you want to generate a random set of bytes quickly?

Same here, and that's why I got lazy and created **TestLazy**.

---

## 🧪 What is TestLazy?

TestLazy is a growing collection of pre-built test values, fakes, and validators for Go - made to reduce boilerplate in your tests.
Instead of writing the same four lines every time you need a malformed URL or a canceled context, you can use TestLazy to get those values with a single function call.

- Use **canonical values** (`url.URL`, `net.IP`, etc.) without the setup.
- Simulate **broken things** (like a canceled context) with a single function call.
- Validate common responses (like HTTP status codes) with built-in validators.

Whether you're testing APIs or database interactions or don't want to think about test values, TestLazy has you covered.

---

## 🚀 Getting Started

### URLs without the boilerplate | `github.com/madflojo/testlazy/things/testurl`

Instead of this:

```go
// Create a URL for example.com
url, err := url.Parse("https://example.com")
if err != nil {
    t.Fatalf("Failed to parse URL: %v", err)
}

// Create HTTP Request
req := &http.Request{
    Method: "GET",
    URL:    url,
}
```

You can now just do this:

```go
req := &http.Request{
    Method: "GET",
    URL:    testurl.URLHTTPS(),
}
```

Or, maybe you want a malformed URL:

```go
req := &http.Request{
    Method: "GET",
    URL:    testurl.URLInvalidHost(),
}
```

### Contexts without the hassle | `github.com/madflojo/testlazy/fakes/fakectx`

Instead of this:

```go
// Create a context that is already canceled
ctx, cancel := context.WithCancel(context.Background())
cancel() // Cancel the context immediately

err := doSomething(ctx)
if err == nil {
    t.Fatal("Expected error for canceled context, got nil")
}
```

You can now just do this:

```go
err := doSomething(fakectx.Cancelled())
if err == nil {
    t.Fatal("Expected error for canceled context, got nil")
}
```

---

## 🧱 Structure

TestLazy is a set of individual packages; each focused on a specific category of test value or functionality.
It allows you to take on only the dependencies you need.

| Package | Description | Go Package Reference |
|---------|-------------|----------------------|
| `github.com/madflojo/testlazy/things/testurl` | Pre-built URLs for common use cases | [![Go Reference](https://pkg.go.dev/badge/github.com/madflojo/testlazy/things/testurl.svg)](https://pkg.go.dev/github.com/madflojo/testlazy/things/testurl) |
| `github.com/madflojo/testlazy/fakes/fakectx` | Fake contexts for testing | [![Go Reference](https://pkg.go.dev/badge/github.com/madflojo/testlazy/fakes/fakectx.svg)](https://pkg.go.dev/github.com/madflojo/testlazy/fakes/fakectx) |
| `github.com/madflojo/testlazy/validators/validhttp` | Validators for HTTP responses | [![Go Reference](https://pkg.go.dev/badge/github.com/madflojo/testlazy/validators/validhttp.svg)](https://pkg.go.dev/github.com/madflojo/testlazy/validators/validhttp) |

---

## 🦥 Stay Lazy, Test More!

More is on the way - TestLazy is just getting started.
But if you cannot bear to write the same test setup code repeatedly, contributions (pull requests) and suggestions (issues) are welcome!
Let's make testing less of a chore and more of a breeze. 🏝️
