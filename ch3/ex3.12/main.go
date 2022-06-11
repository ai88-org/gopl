package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(anagrams("bagbgc", "agbcdg"))
}

func anagrams(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	// 防止有重复的字符
	for i := 0; i < len(s1); i++ {
		if !strings.Contains(s2, string(s1[i])) {
			return false
		}
	}

	for i := 0; i < len(s2); i++ {
		if !strings.Contains(s1, string(s2[i])) {
			return false
		}
	}
	return true
}
