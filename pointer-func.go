package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// func ChangeCounteryToIndonesia(address Address) { tidak akan berubah country-nya
func ChangeCounteryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var alamat = Address{
		City:     "Ambarawa",
		Province: "Jawa Tengah",
		Country:  "",
	}

	// ChangeCounteryToIndonesia(alamat) tidak akan berubah country-nya
	ChangeCounteryToIndonesia(&alamat)
	fmt.Print(alamat)
}
