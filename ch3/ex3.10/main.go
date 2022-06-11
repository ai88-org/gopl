package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma1("12345678"))
}

func comma1(s string) string {
	// 定义一个buffer
	var buf bytes.Buffer
	// 原始数字的长度，字节长度
	n := len(s)
	// 小于3位直接返回
	if n <= 3 {
		return s
	}

	for i := 0; i < len(s); i++ {
		buf.WriteByte(s[i])
		if i%3 == 1 && i < len(s)-1 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}
