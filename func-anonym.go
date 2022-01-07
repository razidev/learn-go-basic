package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}
func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("admin", blacklist)
	registerUser("Razi", blacklist)

	func() {
		fmt.Println("I am anonymous func too")
	}()

	func(x int) {
		fmt.Println("I am anonymous func too with a number of", x)
	}(22)
}
