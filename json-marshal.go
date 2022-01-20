package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	p1 := struct {
		Name    string
		Hobbies []string
		Books   map[int]string
	}{
		Name:    "Razi",
		Hobbies: []string{"Badminton", "Games"},
		Books: map[int]string{
			1: "Atomic Habits",
			2: "The Intelligent Investor",
			3: "Supernova",
		},
	}

	bytes, _ := json.Marshal(p1)
	fmt.Println(string(bytes))
}
