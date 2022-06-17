// 自定义类型可以实现多个接口
package main

import "fmt"

func main() {
	var a A
	var b B
	var j JustX

	a = j
	b = j
	fmt.Println(a.Say)
	fmt.Println(b.Hello)
}

type A interface {
	Say()
}

type B interface {
	Hello()
}

type JustX struct {
}

func (s JustX) Say() {
	fmt.Println("s say ...")
}

func (s JustX) Hello() {
	fmt.Println("s hello ...")
}
