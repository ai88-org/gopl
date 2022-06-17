package main

import "fmt"

func main() {
	var c = Cat{}
	var d Dog
	fmt.Println(d)
	var interface2 i = c
	if v, ok := interface2.(Cat); ok {
		fmt.Printf("%v\t%T\n", v, v)
	}

	if v, ok := interface2.(Dog); !ok {
		fmt.Printf("%v\t%T\n", v, v)
	}
}

type i interface {
	Say()
}

type Cat struct {
}

type Dog int

func (s Cat) Say() {
	fmt.Println("cat hello....")
}

func (s Dog) Say() {
	fmt.Println("dog hello....")
}
