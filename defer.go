package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil")
}

func runApplication(val int) {
	defer logging() //tetep akan terpanggil terakhir walau error
	fmt.Println("Run App")
	result := 10 / val
	fmt.Println("res", result)
}
func main() {
	runApplication(10)
}
