package fakectx

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func ExampleCancelled() {
	ctx := Cancelled()

	<-ctx.Done()
	fmt.Println("Context done")
	// Output:
	// Context done
}

func ExampleDeadlineExceeded() {
	ctx := DeadlineExceeded()
	deadline, ok := ctx.Deadline()

	if ok && deadline.Before(time.Now()) && errors.Is(ctx.Err(), context.DeadlineExceeded) {
		fmt.Println("Deadline exceeded")
	}
	// Output:
	// Deadline exceeded
}

func ExampleTimedOut() {
	ctx := TimedOut()

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		fmt.Println("Timed out")
	}
	// Output:
	// Timed out
}

func ExampleTimesOutAfter() {
	ctx := TimesOutAfter(5 * time.Millisecond)

	if ctx.Err() != nil {
		fmt.Println("Timeout too soon")
	}

	time.Sleep(10 * time.Millisecond)

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		fmt.Println("Timed out")
	}
	// Output:
	// Timed out
}

func ExampleCancelledWithCallback() {
	_, cancel := CancelledWithCallback(func() {
		fmt.Println("Cancelled callback called")
	})

	cancel()
	// Output:
	// Cancelled callback called
}
