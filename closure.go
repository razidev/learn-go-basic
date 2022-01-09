package main

import "fmt"

var num int

func main() {
	name := "Razi"
	counter := 0

	fmt.Println("num", num)

	increment := func() {
		name := "aziz"
		counter++
		num++
		fmt.Println("name", name)
	}

	increment()
	fmt.Println("num2", num)
	fmt.Println("counter", counter)
	fmt.Println("name1", name)
	increase()
	fmt.Println("num3", num)

	{
		x := 30
		fmt.Println("x", x)
	}

	a := incrementor()
	b := incrementor()
	fmt.Println("a1", a())
	fmt.Println("a2", a())
	fmt.Println("b1", b())
	fmt.Println("b2", b())
}

func increase() {
	num++
}
func incrementor() func() int {
	var y int
	return func() int {
		y++
		return y
	}
}

//hati-hati menggunakan closure bisa merubah value yg sudah ada
