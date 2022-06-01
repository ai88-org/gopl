package main

// 求两个数的最大公约数 这个算法是真的巧妙啊
// 最大公约数就是其中一个，通过循环的办法，让其中一个数等于最大公约数

func main() {
	println(gcd(100, 13))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
