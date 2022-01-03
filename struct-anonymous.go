package main

import "fmt"

func main() {
	p1 := struct {
		name    string
		hobbies []string
		books   map[int]string
	}{
		name:    "Razi",
		hobbies: []string{"Bandminto", "Games"},
		books: map[int]string{
			1: "Atomic Habits",
			2: "the intelligent investor",
			3: "Supernova",
		},
	}
	fmt.Println(p1.name)
	fmt.Println(p1.hobbies)
	fmt.Println(p1.books)
}
