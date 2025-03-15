package main

import (
	"os"
	"os/signal"
	"time"
  "fmt"
)

func main() {
	// quit := make(chan struct{})
	//
	// mon := Monitor{}
	// go mon.Start(quit)
	//
	// time.Sleep(10 * time.Second)
	//
	// close(quit)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	quit := make(chan struct{})
	// mon := Monitor{}
	mon := Monitor{
		done: make(chan struct{}),
	}

	go mon.Start(quit)

	<-sig

	close(quit)

	select {

	case <-mon.Done():
		os.Exit(0)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("timed out while trying to shut down the monitor")
		os.Exit(1)
	}

}
