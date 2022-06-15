package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	arr = reverse2(arr)
	fmt.Println(arr)
}
func reverse2(s [5]int) [5]int {
	atr := &s
	for i, j := 0, len(*atr)-1; i < j; i, j = i+1, j-1 {
		atr[i], atr[j] = atr[j], atr[i]
	}
	return *atr
}
