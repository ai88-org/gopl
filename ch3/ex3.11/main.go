package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("1234567"))
	fmt.Println(comma("1234567.45"))
	fmt.Println(comma("-1234567.45"))
	fmt.Println(comma("-12345"))
}

// 可以处理浮点数以及正负数
func comma(s string) string {
	n := len(s)
	// 处理浮点数的情况
	var dots string
	if dot := strings.LastIndex(s, "."); dot > 0 {
		n = dot
		dots = s[dot:]
	}

	// 处理负数情况
	var sign byte
	if strings.HasPrefix(s, "-") {
		n--
		s = strings.Replace(s, "-", "", 1)
		sign = '-'
	}
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	buf.WriteByte(sign)
	for i := 0; i < n; i++ {
		buf.WriteByte(s[i])
		if i%3 == 1 && i != n-1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteString(dots)
	return buf.String()
}
