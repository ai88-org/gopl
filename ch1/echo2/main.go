package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// 确实strings.Join效率更高
func main() {
	countTime(p1)
	countTime(p2)
}

func countTime(f func()) {
	startTime := time.Now().UnixNano()
	f()
	endTime := time.Now().UnixNano()
	fmt.Println(endTime)
	fmt.Println(startTime)
	fmt.Printf("一共花费%d\n", endTime-startTime)
}

func p1() {
	s, sep := "", " "
	for _, arg := range os.Args[1:] {
		s += arg + sep
	}
	fmt.Println(s)
}

func p2() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
