package main

import (
	"Udemy-PZN/helper"
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	defer close(channel)

	go helper.GiveMeResponse(channel)

	data := <-channel
	fmt.Println("data:", data)

	time.Sleep(5 * time.Second)
}
