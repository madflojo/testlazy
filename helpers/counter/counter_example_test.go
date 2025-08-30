package counter_test

import (
	"fmt"
	"time"

	"github.com/madflojo/testlazy/helpers/counter"
)

func ExampleNew() {
	// Create a new counter.
	c := counter.New()

	// Initial value is 0.
	fmt.Println(c.Value())

	// Increment the counter by 1.
	c.Increment()

	fmt.Println(c.Value())
	// Output:
	// 0
	// 1
}

func ExampleCounter_Add() {
	// Create a new counter.
	c := counter.New()

	// Initial value is 0.
	fmt.Println(c.Value())

	// Add 5 to the counter.
	c.Add(5)

	fmt.Println(c.Value())
	// Output:
	// 0
	// 5
}

func ExampleCounter_Subtract() {
	// Create a new counter.
	c := counter.New()

	// Initial value is 0.
	fmt.Println(c.Value())

	// Add 5 to the counter.
	c.Add(5)

	fmt.Println(c.Value())

	// Subtract 3 from the counter.
	c.Subtract(3)

	fmt.Println(c.Value())
	// Output:
	// 0
	// 5
	// 2
}

func ExampleCounter_WaitAbove() {
	// Create a new counter.
	c := counter.New()

	// Bump the counter in the background.
	go func() {
		for i := 0; i < 5; i++ {
			c.Increment()
			time.Sleep(5 * time.Millisecond)
		}
	}()

	// Wait until value >= 5 (or time out).
	if err := <-c.WaitAbove(5, time.Second); err != nil {
		fmt.Println("timeout")
		return
	}
	fmt.Println("done")
	// Output:
	// done
}

func ExampleCounter_WaitBelow() {
	// Create a new counter.
	c := counter.New()

	// Set the counter to 5.
	c.Add(5)

	// Drop the counter in the background.
	go func() {
		for i := 0; i < 5; i++ {
			c.Decrement()
			time.Sleep(5 * time.Millisecond)
		}
	}()

	// Wait until value <= 2 (or time out).
	if err := <-c.WaitBelow(2, time.Second); err != nil {
		fmt.Println("timeout")
		return
	}
	fmt.Println("done")
	// Output:
	// done
}
