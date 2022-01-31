package main

import (
	"fmt"
	"time"
)

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Michelangelo"
}

func main() {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println("data:", data)

	time.Sleep(5 * time.Second)
}
