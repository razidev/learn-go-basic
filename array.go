package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "razi"
	names[1] = "aziz"
	names[2] = "syahputro"
	fmt.Println(names)

	var values = [3]int{
		22,
		04,
		1996,
	}

	fmt.Println(values[1])

	var checkLength [10]int
	fmt.Println(len(checkLength))
	fmt.Println(checkLength)
}
