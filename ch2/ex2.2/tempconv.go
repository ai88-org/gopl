// 人民币与美元的转换
package main

import "fmt"

type Dollar float64
type Renminbi float64

func (d Dollar) String() string {
	return fmt.Sprintf("%g$", d)
}

func (r Renminbi) String() string {
	return fmt.Sprintf("%g¥", r)
}
