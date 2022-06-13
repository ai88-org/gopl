package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println(compareSHA256("x", "X"))
}

// 比较SHA256 HASHES
func compareSHA256(x1 string, x2 string) int {
	c1 := sha256.Sum256([]byte(x1))
	c2 := sha256.Sum256([]byte(x2))

	// 最终返回的结果
	var bit1 byte = 0
	var bit2 byte = 0
	var res = 0

	// 32个字节长度
	for i := 0; i < 32; i++ {
		for j := 0; j < 8; j++ {
			// 拿到最右边的1bit，然后进行比较
			bit1 = (c1[i] >> j) & 1
			bit2 = (c2[i] >> j) & 1
			if bit1 != bit2 {
				res++
			}
		}
	}
	return res
}
