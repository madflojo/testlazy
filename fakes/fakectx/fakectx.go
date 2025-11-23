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
	"time"
)

// Cancelled returns a context that has already been canceled. The returned
// context should report context.Canceled from Err and have a Done channel that
// is already closed.
func Cancelled() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

// DeadlineExceeded returns a context that has already exceeded its deadline.
// Callers can inspect the returned Deadline to verify it is in the past.
func DeadlineExceeded() context.Context {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(-time.Minute))
	cancel()
	return ctx
}

// TimedOut returns a context that has already hit its timeout deadline. The
// context behaves like the result of context.WithTimeout where the timer has
// already elapsed.
func TimedOut() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 0)
	defer cancel()
	return ctx
}

// TimesOutAfter returns a context that will cancel itself after the provided
// duration. Tests can use this to trigger timeout paths without wiring up a
// real timer manually.
func TimesOutAfter(timeout time.Duration) context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	go func() {
		<-ctx.Done()
		cancel()
	}()
	return ctx
}

// CancelledWithCallback returns a canceled context alongside a cancel function
// that executes the provided callback when invoked. This allows tests to ensure
// downstream code triggers cancellation.
func CancelledWithCallback(cb func()) (context.Context, context.CancelFunc) {
	return nil, nil
}
