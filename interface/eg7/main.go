package main

import "fmt"

func main() {
	var i interface{} = 10
	i = "helloworld"
	i = true

	type T struct {
	}
	i = T{}
	fmt.Println(i)
}
