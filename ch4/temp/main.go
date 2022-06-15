package main

import "fmt"

func main() {

	var ages2 map[string]int
	//ages2["carol"] = 21
	fmt.Println(ages2)

	var ages = make(map[string]int)
	fmt.Println(ages)
	ages["carol"] = 21
	fmt.Println(ages)

}
