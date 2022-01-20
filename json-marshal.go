package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string         `json:"name"`
	Hobbies []string       `json:"hobbies"`
	Books   map[int]string `json:"books"`
}

func main() {
	p1 := Person{
		Name:    "Razi",
		Hobbies: []string{"Badminton", "Games"},
		Books: map[int]string{
			1: "Atomic Habits",
			2: "The Intelligent Investor",
			3: "Supernova",
		},
	}
	p2 := Person{
		Name:    "Syahputro",
		Hobbies: []string{"Football", "Boxing"},
		Books: map[int]string{
			1: "The Lord of The Ring",
			2: "Harry Potter",
			3: "The Great Gatsby",
		},
	}

	bytes1, _ := json.Marshal(p1)
	bytes2, _ := json.Marshal(p2)
	fmt.Println(string(bytes1))

	slice := []string{string(bytes1), string(bytes2)}
	fmt.Println(slice)
}
