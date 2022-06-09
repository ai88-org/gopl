package main

import "fmt"

func main() {
	fmt.Println(PopCount2_5(13999999))
}

func PopCount2_5(x uint64) int {
	var n int
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	return n
}
