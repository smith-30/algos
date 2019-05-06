package main

import (
	"fmt"
	"math"
)

func main() {
	n := SingleInt()
	res := n

	// 1 ≤ N ≤ 100000
	// 6で利用できる累乗は6まで
	// 9で利用できる累乗は5まで
	// for i := 0; i < n; i++ {

	// }
	var cc, t int
	t = n
	for t > 0 {
		cc += t % 6
		fmt.Printf("--- %#v %#v\n", t, t%6)
		t /= 6
	}

	fmt.Printf("%#v\n", cc)

	var cc2 int
	t = n
	for t > 0 {
		cc2 += t % 9
		t /= 9
	}

	fmt.Printf("%#v\n", Min(res, cc, cc2))
}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

func Min(nums ...int) int {
	if len(nums) == 0 {
		panic("function min() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Min(float64(res), float64(nums[i])))
	}
	return res
}
