package counter

import (
	"errors"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	c := New()
	if c == nil {
		t.Fatal("NewCounter returned nil")
	}

	t.Run("Initial Value", func(t *testing.T) {
		if got := c.Value(); got != 0 {
			t.Fatalf("initial value: want 0, got %d", got)
		}
	})

	t.Run("WaitAbove", func(t *testing.T) {
		// Start a goroutine to increment the counter
		go func() {
			for i := 0; i < 10; i++ {
				c.Increment()
				time.Sleep(10 * time.Millisecond)
			}
		}()

		// Wait for the counter to go above 5
		if err := <-c.WaitAbove(5, 2*time.Second); err != nil {
			t.Fatal("WaitAbove returned error:", err)
		}
	})

	t.Run("WaitBelow", func(t *testing.T) {
		// Start a goroutine to decrement the counter
		go func() {
			for i := 0; i < 10; i++ {
				c.Decrement()
				time.Sleep(10 * time.Millisecond)
			}
		}()

		// Wait for the counter to go below 5
		if err := <-c.WaitBelow(5, 2*time.Second); err != nil {
			t.Fatal("WaitBelow returned error:", err)
		}
	})

	t.Run("Set", func(t *testing.T) {
		c.Set(42)
		if got := c.Value(); got != 42 {
			t.Fatalf("after set: want 42, got %d", got)
		}
	})

	t.Run("Reset", func(t *testing.T) {
		c.Reset()
		if got := c.Value(); got != 0 {
			t.Fatalf("after reset: want 0, got %d", got)
		}
	})

	t.Run("Add", func(t *testing.T) {
		c.Add(10)
		if got := c.Value(); got != 10 {
			t.Fatalf("after add: want 10, got %d", got)
		}
	})

	t.Run("Subtract", func(t *testing.T) {
		c.Subtract(5)
		if got := c.Value(); got != 5 {
			t.Fatalf("after subtract: want 5, got %d", got)
		}
	})
}

func TestWaitAboveTimeout(t *testing.T) {
	c := New()
	if c == nil {
		t.Fatal("NewCounter returned nil")
	}

	err := <-c.WaitAbove(1, 1*time.Second)
	if err == nil || !errors.Is(err, ErrTimeout) {
		t.Fatal("WaitAbove should have timed out but didn't")
	}
}

func TestWaitBelowTimeout(t *testing.T) {
	c := New()
	if c == nil {
		t.Fatal("NewCounter returned nil")
	}

	c.Set(10)

	err := <-c.WaitBelow(5, 1*time.Second)
	if err == nil || !errors.Is(err, ErrTimeout) {
		t.Fatal("WaitBelow should have timed out but didn't")
	}
}
