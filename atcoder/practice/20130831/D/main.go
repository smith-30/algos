package main

import (
	"fmt"
	"math"
)

func main() {
	n, d := SingleInt(), SingleInt()

	a := math.Pow(6, float64(n))
	var p int
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		
	}

	for i := n; i > 0; i-- {
		for j := 1; j <= 6; j++ {

		}
	}
	fmt.Println(float64(p) / a)
}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

func Max(nums ...int) int {
	if len(nums) == 0 {
		panic("function max() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Max(float64(res), float64(nums[i])))
	}
	return res
}
