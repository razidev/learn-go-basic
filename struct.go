package main

import "fmt"

type MySelf struct {
	Name, Address string
	Age           int
}

//struct function
func (me MySelf) sayHi(name string) {
	fmt.Println("Hello", name, "my name is", me.Name)
}

func main() {
	var Me MySelf
	Me.Name = "Razi"
	Me.Address = "Kp. Mangga"
	Me.Age = 25
	fmt.Println(Me)
	Me.sayHi("Joko")

	aziz := MySelf{
		Name:    "Aziz",
		Address: "Gg. Teladan 5",
		Age:     24,
	}
	fmt.Println(aziz)

	//tidak disarankan krn ketergantungan posisi
	syahputro := MySelf{"Syahputro", "Jakarta", 23}
	fmt.Println(syahputro)
}
