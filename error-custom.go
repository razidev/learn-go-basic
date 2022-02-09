package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "this is the custom error",
	}

	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
	fmt.Println(e.(customErr).info)
}
