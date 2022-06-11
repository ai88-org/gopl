package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("123456789"))
	fmt.Println(commaF("123456789.65"))
}

// 这种非常有规律的截取，可以考虑递归的思想
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:(n-3)]) + "," + s[n-3:]
}

func commaF(s string) string {
	n := len(s)
	if dot := strings.LastIndex(s, "."); dot > 0 {
		n = dot
	}
	if n <= 3 {
		return s
	}
	return comma(s[:(n-3)]) + "," + s[n-3:]
}
