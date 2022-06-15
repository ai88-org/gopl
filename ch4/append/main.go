package main

import "fmt"

func main() {
	var x = make([]int, 3, 5)

	//for i := 0; i < 10; i++ {
	//	//x = appendInt(x, i)
	//	x = appendInt(x, i)
	//	fmt.Printf("x: %v\tlen: %d\tcap: %d\n", x, len(x), cap(x))
	//}
	x = appendIntPlus(x, 1, 2, 3)
	fmt.Printf("x: %v\tlen: %d\tcap: %d\n", x, len(x), cap(x))

}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1

	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		// 已然超过了容量，先赋值，再去修正
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func appendIntPlus(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)

	// 还在容量之内，不用变数组
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		// 超过容量，需要依赖新的数组
		zcap := zlen
		// 应该是两倍原来长度
		if zcap < 2*cap(x) {
			zcap = 2 * cap(x)
		}
		z = make([]int, zlen, zcap)
	}
	copy(z[len(x):], y)
	return z

}
