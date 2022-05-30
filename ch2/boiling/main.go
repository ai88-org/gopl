package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g F or %g C", f, c) // 摄氏度
	//var d = 212.5
	//var e = 45
	//var h = d - e 常量或者字面量都可以又编译器推导，但是变量推导不了，得做显示的类型转换
	//fmt.Println(h)
}
