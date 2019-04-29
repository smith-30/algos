package main

import (
	"fmt"
	"math"
)

// 与えられた数列を R, G, B いずれかで塗って
// それぞれの総和で三角形を作る
// そのとき三角形の作り方は何通りあるか

// 三角形の条件
// | a - b | < c < a + b
// a < b + c

// a + b + c = S
// max(a, b, c) < S/2
func main() {
	n := SingleInt()

	for i := 0; i < n; i++ {
		fmt.Printf("%#v\n", SingleInt())
	}

	a := math.Pow(3, float64(n))
	fmt.Printf("%#v\n", a)
}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

func SingleStr() string {
	var s string
	fmt.Scan(&s)
	return s
}
