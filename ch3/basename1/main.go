package main

import (
	"fmt"
	"strings"
)

// e.g , a=>a, a.go => a, a/b/c.go => c,a/b.c.go => b.c
func main() {
	fmt.Println(basename1("a"))
	fmt.Println(basename2("a"))
	fmt.Println(basename1("a.go"))
	fmt.Println(basename2("a.go"))
	fmt.Println(basename1("a/b/c.go"))
	fmt.Println(basename2("a/b/c.go"))
	fmt.Println(basename1("a/b.c.go"))
	fmt.Println(basename2("a/b.c.go"))
}

func basename1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	// 找到最后一个斜线的位置
	slash := strings.LastIndex(s, "/")
	// 截取
	s = s[slash+1:]

	// 找不到就是-1 ，找最后一个点的位置，但是取不到
	if dot := strings.LastIndex(s, "."); dot > 0 {
		s = s[:dot]
	}
	return s
}
