package main

import (
	"fmt"
	"strconv"
	"time"
)

const noDelay time.Duration = 0
const timeout = 5 * time.Minute

type weekday int

const (
	Sunday weekday = iota
	Monday
	Tuesday
	wednesday
	Thursday
	Friday
	Saturday
)

const (
	a = 1
	b
	c = true
	d
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Printf("%q\n", y)

	println(strconv.FormatInt(int64(x), 2))

	fmt.Println(a, b, c, d)

	fmt.Println(Sunday, Monday, Tuesday, wednesday, Thursday, Friday, Saturday)

}
