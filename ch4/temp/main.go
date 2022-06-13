package main

import (
	"fmt"
)

func main() {

	x := 3
	fmt.Println(&x)
	test(x)

}

func test(x int) {
	// 复制了参数，一个新的int类型在内存生成了
	fmt.Println(&x)
}
