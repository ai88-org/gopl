package main

import "fmt"

func main() {
	// new 创建了匿名变量，返回的是指针类型，初始化并给变量赋0值
	fmt.Println(6 % 12)
}

func incr(p *int) int {
	*p++
	return *p
}
