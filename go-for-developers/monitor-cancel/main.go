package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {

	ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)
	// defer cancel()

	// deadline := time.Date(2029, 3, 14, 15, 45, 0, 0, time.UTC)
	//  fmt.Println("deadline:", deadline.Format(time.RFC3339))

	// ctx, cancel := context.WithDeadline(ctx, deadline)

	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)

	mon := Monitor{}

	ctx = mon.Start(ctx)

	// go func() {
	// 	time.Sleep(time.Millisecond * 50)
	// 	cancel()
	// }()
	defer cancel()

	select {
	case <-ctx.Done():
		os.Exit(0)
	case <-time.After(time.Second * 2):
		fmt.Println("timed out while trying to shut down the monitor")

		if err := ctx.Err(); err != nil {
			fmt.Printf("error: %s\n", err)
		}
		os.Exit(1)
	}

}
