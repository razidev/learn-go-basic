package main

import (
	"Udemy-PZN/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
