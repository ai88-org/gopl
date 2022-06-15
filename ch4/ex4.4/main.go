package main

import "fmt"

func main() {
	var s = []string{"a", "b", "c", "d", "e"}
	fmt.Println(rotate(s, 2))
}

// 翻转函数
func rotate(s []string, n int) []string {
	var tmp []string
	// 先把从0到n位置的元素放到最后
	tmp = append(s, s[:n]...)
	// 再把前n个截取掉
	return tmp[n:]
}
