package main

import "fmt"

func main() {
	var a A
	var x JustX
	a = x
	a.yeah()
	a.Say()
	a.Hello()
}

type B interface {
	Say()
}

type C interface {
	Hello()
}

// A接口继承了B、C接口
type A interface {
	B
	C
	yeah()
}

type JustX struct {
}

func (j JustX) Say() {
	fmt.Println("say...")
}

func (j JustX) Hello() {
	fmt.Println("hello...")
}

func (j JustX) yeah() {
	fmt.Println("yeah...")
}
