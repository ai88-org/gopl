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
	f3 := PopCount2_4
	f4 := PopCount2_5
	TimeConsuming(f1, 13999999)
	TimeConsuming(f2, 13999999)
	TimeConsuming(f3, 13999999)
	TimeConsuming(f4, 13999999)
	TimeConsuming(f1, 98765)
	TimeConsuming(f2, 98765)
	TimeConsuming(f3, 98765)
	TimeConsuming(f4, 98765)
}

// 计算函数运行时间
func TimeConsuming(f func(uint64) int, x uint64) {
	start := time.Now()
	f(x)
	fmt.Printf("elapsed: %d ns count:= %d\n", time.Since(start).Nanoseconds(), f(x))
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

// 习题2.4的方法
func PopCount2_4(x uint64) int {
	n := 0
	for i := 0; i < 64; i++ {
		if (x>>(1*i))&1 == 1 {
			n++
		}
	}
	return n
}

// 习题2.5的方法
func PopCount2_5(x uint64) int {
	var n int
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	return n
}
