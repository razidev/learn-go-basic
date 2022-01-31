package main

import (
	"fmt"
	"time"
)

func sendChan(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Michelangelo"
}

func receiveChan(channel <-chan string) {
	data := <-channel
	fmt.Println("receiver:", data)
}

func main() {
	channel := make(chan string)
	defer close(channel)

	go sendChan(channel)
	go receiveChan(channel)

	time.Sleep(5 * time.Second)
}
