package fakectx

import (
	"context"
	"errors"
	"fmt"
)

func ExampleCancelled() {
	ctx := Cancelled()
	if !errors.Is(ctx.Err(), context.Canceled) {
		fmt.Println("unexpected error:", ctx.Err())
	}

	fmt.Println("Context is cancelled as expected")
	// Output:
	//Context is cancelled as expected
}
