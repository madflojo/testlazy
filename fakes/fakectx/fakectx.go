/*
Package fakectx provides curated context.Context helpers for tests that need
specific cancellation or deadline behavior without the boilerplate. The package
focuses on one-liner constructors so test authors can express failure scenarios
clearly and consistently. Start with Cancelled for immediate
context.Canceled states and layer in additional helpers as your tests grow more
complex.
*/
package fakectx

import (
	"context"
	"sync"
	"time"
)

// Cancelled returns a context that has already been canceled. The returned
// context reports context.Canceled, closes Done immediately, and carries no
// deadline or values—perfect for forcing cancel-only branches.
func Cancelled() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	return ctx
}

// DeadlineExceeded returns a context that has already exceeded its deadline.
// The Deadline reported by the context is always in the past and Err returns
// context.DeadlineExceeded immediately.
func DeadlineExceeded() context.Context {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(-time.Minute))
	cancel()

	return ctx
}

// TimedOut returns a context that has already hit its timeout deadline. It is
// equivalent to a context.WithTimeout call whose timer has fired.
func TimedOut() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 0)
	defer cancel()

	return ctx
}

// TimesOutAfter returns a context that will cancel itself after the provided
// duration. The Deadline is set in the future so callers can assert how much
// time remains before it expires.
func TimesOutAfter(timeout time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), timeout)

	return ctx
}

// CancelledWithCallback returns a canceled context alongside a cancel function
// that executes the provided callback when invoked. This allows tests to ensure
// downstream code triggers cancellation.
func CancelledWithCallback(cb func()) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())

	var once sync.Once
	wrapped := func() {
		once.Do(func() {
			cancel()
			if cb != nil {
				cb()
			}
		})
	}

	return ctx, wrapped
}
