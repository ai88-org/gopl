package main

import "fmt"

func main() {

	var runes []rune
	for _, s := range "Hello, 世界" {
		runes = append(runes, s)
	}

	fmt.Printf("%c", runes)

}
