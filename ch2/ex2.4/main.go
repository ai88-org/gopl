package main

import "fmt"

func main() {
	fmt.Println(PopCount2_4(13999999))
}

func PopCount2_4(x uint64) int {
	n := 0
	for i := 0; i < 64; i++ {
		if (x>>(1*i))&1 == 1 {
			n++
		}
	}
	return n
}
