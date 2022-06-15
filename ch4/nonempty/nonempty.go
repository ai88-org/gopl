package main

import "fmt"

func main() {
	var s1 = []int{1, 2, 3, 4, 5, 6}
	s1 = remove(s1, 4)
	fmt.Printf("%v\n", s1)

}

func nonempty(strings []string) []string {
	var i = 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	var output = strings[:0]
	for _, s := range strings {
		if s != "" {
			output = append(output, s)
		}
	}
	return output
}

// 删除切片中间的元素
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
