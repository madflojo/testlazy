/*
Package counter provides a tiny, thread-safe counter focused on
testing scenarios. It supports atomic increments/decrements, direct
reads/sets, and simple wait helpers that block until the value crosses a
threshold.

	github.com/madflojo/testlazy/helpers/counter

# Why use it

- Minimal API designed for tests.
- Safe for concurrent use by multiple goroutines.
- Non-blocking "wait" helpers that return a channel for easy select/timeout.

Quick example

	c := counter.New()
	c.Increment()
	c.Add(9)
	<-c.WaitAbove(10, time.Second) // wait until value >= 10
*/
package counter

import (
	"errors"
	"sync/atomic"
	"time"
)

// defaultPollInterval controls how often waiters check the counter value.
const defaultPollInterval = 10 * time.Millisecond

// ErrTimeout indicates a WaitAbove or WaitBelow call timed out before the
// condition was met.
var ErrTimeout = errors.New("timeout waiting for counter condition")

// Counter is a thread-safe counter.
// All methods are safe to call from multiple goroutines.
type Counter struct {
	value int64
}

// New creates a new Counter with initial value 0.
func New() *Counter {
	return &Counter{}
}

// Increment increases the counter by 1.
func (c *Counter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

// Decrement decreases the counter by 1.
func (c *Counter) Decrement() {
	atomic.AddInt64(&c.value, -1)
}

// Value returns the current value of the counter.
func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

// Reset sets the counter back to 0.
func (c *Counter) Reset() {
	atomic.StoreInt64(&c.value, 0)
}

// Add increases the counter by the given delta.
func (c *Counter) Add(delta int64) {
	atomic.AddInt64(&c.value, delta)
}

// Subtract decreases the counter by the given delta.
func (c *Counter) Subtract(delta int64) {
	atomic.AddInt64(&c.value, -delta)
}

// Set sets the counter to the given value.
func (c *Counter) Set(value int64) {
	atomic.StoreInt64(&c.value, value)
}

// WaitAbove returns a channel that will receive a single error when the counter
// value is >= target (inclusive) or when the timeout elapses.
//
// On success the error is nil. On timeout the error is ErrTimeout.
func (c *Counter) WaitAbove(target int64, timeout time.Duration) <-chan error {
	result := make(chan error, 1)
	go func() {
		ticker := time.NewTicker(defaultPollInterval)
		defer ticker.Stop()
		timeoutCh := time.After(timeout)
		for {
			select {
			case <-ticker.C:
				if c.Value() >= target {
					result <- nil
					return
				}
			case <-timeoutCh:
				result <- ErrTimeout
				return
			}
		}
	}()
	return result
}

// WaitBelow returns a channel that will receive a single error when the counter
// value is <= target (inclusive) or when the timeout elapses.
//
// On success the error is nil. On timeout the error is ErrTimeout.
func (c *Counter) WaitBelow(target int64, timeout time.Duration) <-chan error {
	result := make(chan error, 1)
	go func() {
		ticker := time.NewTicker(defaultPollInterval)
		defer ticker.Stop()
		timeoutCh := time.After(timeout)
		for {
			select {
			case <-ticker.C:
				if c.Value() <= target {
					result <- nil
					return
				}
			case <-timeoutCh:
				result <- ErrTimeout
				return
			}
		}
	}()
	return result
}
