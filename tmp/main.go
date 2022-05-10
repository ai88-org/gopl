package main

import "fmt"

func main() {
	count1 := make(map[string]map[string]int)
	count2 := make(map[string]int)

	count2["a"] = 1
	count1["b"] = count2

	fmt.Println(count1)
	fmt.Println(count2)
}
