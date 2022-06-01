package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func main() {
	c := FToC(212)
	fmt.Println(c.String()) // 100°C
	fmt.Printf("%v\n", c)   // 100°C
	fmt.Printf("%s\n", c)   // 100°C ; 没有显示调用String()
	fmt.Println(c)          // 100°C
	fmt.Printf("%g\n", c)   // 100
	fmt.Println(float64(c)) // 100
}

func FToC(f Fahrenheit) Celsius {
	return Celsius(f-32) * 5 / 9
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
