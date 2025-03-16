package main

import (
	"context"
	"fmt"
)

func listener(ctx context.Context, i int) {
	fmt.Printf("listener %d is waiting\n", i)
	<-ctx.Done()

	fmt.Printf("listener %d is exiting\n", i)
}
