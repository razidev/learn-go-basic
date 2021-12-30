package main

import "fmt"

func main() {
	p1 := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Razi",
		lastName:  "Syahputro",
		age:       100,
	}
	fmt.Println(p1)
}
