package main

import "fmt"

// 摄氏温度
type Celsius float64

// 华氏温度
type Fahrenheit float64

func main() {
	// 相同类型的比较
	var c1 Celsius = 100
	fmt.Println(c1.String())
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
