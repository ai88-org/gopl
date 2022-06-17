package main

import "fmt"

func main() {
}

type B interface {
	// 返回值类型应该一样
	Test01() int
	Test02()
}

type C interface {
	Test01() int
	Test03()
}

type A interface {
	B
	C
}

type JustX struct {
}

func (x JustX) Test01() {
	fmt.Println("this is test01")
}

func (x JustX) Test02() {
	fmt.Println("this is test02")
}

func (x JustX) Test03() {
	fmt.Println("this is test03")
}
