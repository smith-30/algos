package main

import (
	"fmt"
	"math"
)

func main() {
	var n, p, q int
	fmt.Scan(&n)
	fmt.Scan(&p)
	fmt.Scan(&q)

	a1 := n + p
	a2 := n + q
	a3 := p + q

	min := Min(a1, a2, a3)

	fmt.Println(min)
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
