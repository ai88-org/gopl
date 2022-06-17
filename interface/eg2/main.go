package main

import "fmt"

func main() {
	var a AInterface
	var m myType
	a = m
	a.SayHello()

	var i integer
	a = i
	a.SayHello()
}

type AInterface interface {
	SayHello()
}

type myType struct {
}

// 自定义类型也能实现接口
type integer int

func (i integer) SayHello() {
	fmt.Println("this is integer say Hello...")
}

func (m myType) SayHello() {
	fmt.Println("Hello...")
}
