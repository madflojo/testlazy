package fakectx

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func ExampleCancelled() {
	ctx := fakectx.Cancelled()

	fmt.Println(errors.Is(ctx.Err(), context.Canceled))
	// Output:
	// true
}

func ExampleDeadlineExceeded() {
	ctx := DeadlineExceeded()
	deadline, ok := ctx.Deadline()

	fmt.Println(ok)
	fmt.Println(deadline.Before(time.Now()))
	fmt.Println(errors.Is(ctx.Err(), context.DeadlineExceeded))
	// Output:
	// true
	// true
	// true
}

func ExampleTimedOut() {
	ctx := TimedOut()

	fmt.Println(errors.Is(ctx.Err(), context.DeadlineExceeded))
	// Output:
	// true
}

func ExampleTimesOutAfter() {
	ctx := TimesOutAfter(5 * time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println(errors.Is(ctx.Err(), context.DeadlineExceeded))
	case <-time.After(50 * time.Millisecond):
		fmt.Println("timeout waiting for context")
	}
	// Output:
	// true
}
