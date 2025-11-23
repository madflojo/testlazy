/*
Package fakectx provides curated context.Context helpers for tests that need
specific cancellation or deadline behavior without the boilerplate. The package
focuses on one-liner constructors so test authors can express failure scenarios
clearly and consistently.
*/
package fakectx

import "context"

// Cancelled returns a context that has already been canceled. The returned
// context should report context.Canceled from Err and have a Done channel that
// is already closed.
func Cancelled() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}
