package fakectx

import (
	"context"
	"testing"
	"time"
)

func TestCancelled(t *testing.T) {
	t.Parallel()

	ctx := Cancelled()
	if ctx == nil {
		t.Fatal("Cancelled returned nil context")
	}

	t.Run("ErrIsCanceled", func(t *testing.T) {
		t.Parallel()

		if err := ctx.Err(); err != context.Canceled {
			t.Fatalf("expected context.Canceled, got %v", err)
		}
	})

	t.Run("DoneChannelClosed", func(t *testing.T) {
		t.Parallel()

		select {
		case <-ctx.Done():
			// expected immediate close
		case <-time.After(1 * time.Millisecond):
			t.Fatal("Done channel was not closed immediately")
		}
	})
}

func TestDeadlineExceeded(t *testing.T) {
	t.Parallel()

	ctx := DeadlineExceeded()
	if ctx == nil {
		t.Fatal("DeadlineExceeded returned nil context")
	}

	t.Run("ErrIsDeadlineExceeded", func(t *testing.T) {
		t.Parallel()

		if err := ctx.Err(); err != context.DeadlineExceeded {
			t.Fatalf("expected context.DeadlineExceeded, got %v", err)
		}
	})

	t.Run("DoneChannelClosed", func(t *testing.T) {
		t.Parallel()

		select {
		case <-ctx.Done():
			// expected immediate close
		case <-time.After(1 * time.Millisecond):
			t.Fatal("Done channel was not closed immediately")
		}
	})

	t.Run("DeadlineReportedInPast", func(t *testing.T) {
		t.Parallel()

		deadline, ok := ctx.Deadline()
		if !ok {
			t.Fatal("expected Deadline() to be set")
		}

		if deadline.After(time.Now()) {
			t.Fatalf("expected deadline in the past, got %v", deadline)
		}
	})
}

func TestTimedOut(t *testing.T) {
	t.Parallel()

	ctx := TimedOut()
	if ctx == nil {
		t.Fatal("TimedOut returned nil context")
	}

	t.Run("ErrIsDeadlineExceeded", func(t *testing.T) {
		t.Parallel()

		if err := ctx.Err(); err != context.DeadlineExceeded {
			t.Fatalf("expected context.DeadlineExceeded, got %v", err)
		}
	})

	t.Run("DoneChannelClosed", func(t *testing.T) {
		t.Parallel()

		select {
		case <-ctx.Done():
		case <-time.After(1 * time.Millisecond):
			t.Fatal("Done channel was not closed immediately")
		}
	})

	t.Run("DeadlineReportedInPast", func(t *testing.T) {
		t.Parallel()

		deadline, ok := ctx.Deadline()
		if !ok {
			t.Fatal("expected Deadline() to be set")
		}

		if deadline.After(time.Now()) {
			t.Fatalf("expected deadline in the past, got %v", deadline)
		}
	})
}

func TestTimesOutAfter(t *testing.T) {
	t.Parallel()

	t.Run("NotExpiredImmediately", func(t *testing.T) {
		t.Parallel()

		timeout := 10 * time.Millisecond
		ctx := TimesOutAfter(timeout)
		if ctx == nil {
			t.Fatal("TimesOutAfter returned nil context")
		}

		select {
		case <-ctx.Done():
			t.Fatal("context timed out immediately")
		case <-time.After(timeout / 2):
		}

		if err := ctx.Err(); err != nil {
			t.Fatalf("expected nil before timeout, got %v", err)
		}
	})

	t.Run("ExpiresAfterTimeout", func(t *testing.T) {
		t.Parallel()

		timeout := 5 * time.Millisecond
		ctx := TimesOutAfter(timeout)
		if ctx == nil {
			t.Fatal("TimesOutAfter returned nil context")
		}

		select {
		case <-ctx.Done():
		case <-time.After(timeout * 4):
			t.Fatal("context did not time out")
		}

		if err := ctx.Err(); err != context.DeadlineExceeded {
			t.Fatalf("expected context.DeadlineExceeded, got %v", err)
		}
	})

	t.Run("DeadlineInFuture", func(t *testing.T) {
		t.Parallel()

		timeout := 20 * time.Millisecond
		ctx := TimesOutAfter(timeout)
		deadline, ok := ctx.Deadline()
		if !ok {
			t.Fatal("expected Deadline() to be set")
		}

		if !deadline.After(time.Now()) {
			t.Fatalf("expected future deadline, got %v", deadline)
		}
	})
}

func TestCancelledWithCallback(t *testing.T) {
	t.Skip("CancelledWithCallback helper not implemented yet")
}
