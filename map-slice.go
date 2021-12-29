package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`razi`:      []string{`racing`, `boxing`, `kicking`},
		`aziz`:      []string{`study`, `Literature`, `Computer Science`},
		`syahputro`: []string{`Badminton`, `Ice cream`, `Sunsets`},
	}

	// fmt.Println(m)

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
