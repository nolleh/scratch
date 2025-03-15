package main

import "fmt"
import "time"

type Monitor struct{
  done chan struct{}
}

// return read only version of channel
func (m Monitor) Done() <-chan struct{} {
  return m.done
}

func (m Monitor) Start(quit chan struct{}) {
	tick := time.NewTicker(10 * time.Millisecond)
	defer tick.Stop()

	for {
		select {
		case <-quit:
			fmt.Println("shutting down monitor")
			return
		case <-tick.C:
			fmt.Println("monitor check")
		}
	}
}
