package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "razi",
		"age":  "23",
	}

	person["job"] = "programmer"
	fmt.Println(person)
	fmt.Println(person["name"])

	var book map[string]string = make(map[string]string)
	book["title"] = "learn go"
	book["author"] = "eko"
	book["delete"] = "please"

	fmt.Println(book)
	delete(book, "delete")
	fmt.Println(book)

}
