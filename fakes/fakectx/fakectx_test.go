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
