package demo

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_WaitGroup(t *testing.T) {
	t.Parallel()

	var completed bool
	var wg sync.WaitGroup

	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(time.Millisecond * 10)
		fmt.Println("done with waitgroup")

		completed = true

	}(&wg)

	fmt.Println("waiting for waitgroup to unblock")

	wg.Wait()

	fmt.Println("waitgroup is unblocked")

	if !completed {
		t.Fatalf("not completed")
	}
}
