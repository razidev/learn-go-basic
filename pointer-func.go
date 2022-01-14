package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// func ChangeCounteryToIndonesia(address Address) { it won't change the country
func ChangeCounteryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var alamat = Address{
		City:     "Ambarawa",
		Province: "Jawa Tengah",
		Country:  "",
	}

	// ChangeCounteryToIndonesia(alamat) it won't change the country
	ChangeCounteryToIndonesia(&alamat)
	fmt.Print(alamat)
}
