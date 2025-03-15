package main

import (
	"fmt"
	"os"
	"os/signal"
)

func Janis(ch chan string) {
	msg := <-ch
	fmt.Println("Jimi said:", msg)
	ch <- "Hello, Jimi!"
}

func main() {
	phone := make(chan string)
	defer close(phone)

	/* A goroutine is not a system thread or process.
	it is lightweight thread of execution that is managed by the Go runtime scheduler */
	go Janis(phone)

	phone <- "Hello, Janis!"

	msg := <-phone

	fmt.Println("Janis said:", msg)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	s := <-ch
	fmt.Println("Got signal:", s)
}
