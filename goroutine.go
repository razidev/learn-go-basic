package main

import (
	"fmt"
	"time"
)

func showHelloWolrd() {
	fmt.Println("Hello World")
}

func main() {
	go showHelloWolrd()
	fmt.Println("shows here")
	time.Sleep(1 * time.Second)
}
