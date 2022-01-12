package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var myKtp NoKTP = "123455677854"
	var marriage Married = false

	fmt.Println(myKtp)
	fmt.Println(marriage)
}

// type declaration uses for initialize with the name you want
