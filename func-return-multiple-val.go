package main

import "fmt"

func getFullName() (string, string, string) {
	return "razi", "aziz", "syahputro"
}

func getFullName2() (firstName, middleName, lastName string) {
	firstName = "nurul"
	middleName = "iza"
	lastName = "hiyah"

	return
}

func main() {
	firstName, _, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println((lastName))

	a, _, c := getFullName2()
	fmt.Println(a)
	fmt.Println(c)
}
