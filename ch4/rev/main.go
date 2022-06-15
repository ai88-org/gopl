package main

import "fmt"

func main() {
	var s1 = []int{1, 2, 3, 4, 5, 6}
	var s2 = []int{1, 2, 3, 4, 5, 6, 7}
	reverse(s1)
	reverse(s2)
	fmt.Println(s1)
	fmt.Println(s2)
}

// 翻转切片
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
