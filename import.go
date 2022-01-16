package main

import (
	"Udemy-PZN/helper"
	"fmt"
)

func main() {
	helper.SayHello("Razi")
	// helper.sayGoodBye("Razi") //not public
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) //not public
}
