package main

import (
	"fmt"
)

func sendChan(channel chan<- string) {
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
	receiveChan(channel)

	fmt.Println("About to exit")
}
