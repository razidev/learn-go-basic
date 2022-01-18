package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type userSlice []User

func (value userSlice) Len() int {
	return len(value)
}

func (value userSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value userSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"Leonardo da Vinci", 25},
		{"Michelangelo", 30},
		{"Rembrandt", 40},
		{"Vermeer", 23},
	}

	fmt.Println(users)
	sort.Sort(userSlice(users))
	fmt.Println(users)
}
