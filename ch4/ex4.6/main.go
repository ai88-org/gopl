package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// and assigns each one a standard number called a Unicode code point or ,
// in Go terminologh.a rune
// byte --> rune -->unicode.IsSpace
func main() {
	b := []byte("我是 你滴  爸爸  爸比 你   懂不懂   。")
	fmt.Println(string(single(b)))
}

// utf8.DecodeRune把utf8编码的字节转换成rune，有自动断言的功能

func single(s []byte) []byte {
	var s2 = s[:0]
	var last rune
	for i := 0; i < len(s); {
		// utf8编码会自动将其断言
		first, size := utf8.DecodeRune(s[i:])
		// 第一次赋值，或者不是空格，或者上一个不是空格，就直接加进去
		if len(s2) == 0 || !unicode.IsSpace(first) || (unicode.IsSpace(first) && !unicode.IsSpace(last)) {
			s2 = append(s2, s[i:size+i]...)
		}
		last = first

		i += size
	}
	return s2
}
