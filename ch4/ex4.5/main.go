package main

import "fmt"

func main() {
	s := []string{"a", "b", "b", "b", "c", "b"}
	s = remove(s)
	fmt.Println(s)
}

func remove(s []string) []string {
	var tmp = s[:0]
	for _, v := range s {
		if len(tmp) == 0 || tmp[len(tmp)-1] != v {
			tmp = append(tmp, v)
		}
	}
	return tmp
}
