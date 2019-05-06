package main

import (
	"fmt"
	"math"
)

func main() {
	n := SingleInt()
	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		dp[i] = 999999
		p := 1
		for p <= i {
			dp[i] = Min(dp[i], dp[i-p]+1)
			p *= 6
		}
		p = 1
		for p <= i {
			dp[i] = Min(dp[i], dp[i-p]+1)
			p *= 9
		}
	}

	fmt.Printf("%#v\n", dp[n])
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
