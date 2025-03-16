package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for i := range 5 {
		go listener(ctx, i)
	}

	time.Sleep(time.Millisecond * 500)

	fmt.Println("canceling the context")

	cancel()

	time.Sleep(time.Millisecond * 500)

}
