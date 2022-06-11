package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Hello,世界"
	fmt.Println(len(s1))
	fmt.Println(utf8.RuneCountInString(s1))

	for i := 0; i < len(s1); {
		r, size := utf8.DecodeRuneInString(s1[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range "Hello,世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	fmt.Println(string("\u6211"))
	fmt.Println(string(25105))

}
