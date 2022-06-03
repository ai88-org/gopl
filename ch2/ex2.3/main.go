package main

import (
	"fmt"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	f1 := PopCount1
	f2 := PopCount2
	TimeConsuming(f1, 13999999) // 236 ns
	TimeConsuming(f2, 13999999) // 157 ns
	TimeConsuming(f2, 98765)    // 75 ns
	TimeConsuming(f1, 98765)    // 92 ns
}

// 计算函数运行时间
func TimeConsuming(f func(uint64) int, x uint64) {
	start := time.Now()
	f(x)
	fmt.Printf("elapsed: %d ns\n", time.Since(start).Nanoseconds())
}

func PopCount1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	var n int
	for i := 0; i < 8; i++ {
		n += int(pc[byte(x>>(i*8))])
	}
	return n
}
