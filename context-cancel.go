package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return destination
}

func main() {
	fmt.Println("total goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)

	fmt.Println("total goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	cancel()

	time.Sleep(2 * time.Second)
	fmt.Println("total goroutine", runtime.NumGoroutine())
}
