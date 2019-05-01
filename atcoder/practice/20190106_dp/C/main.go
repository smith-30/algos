package main

import (
	"fmt"
	"math"
)

func main() {
	n := SingleInt()

	chs := make([][]int, n)

	for i := 0; i < n; i++ {
		chs[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			chs[i][j] = SingleInt()
		}
	}

	dp := make([][]int, n+1)
	dp[0] = make([]int, 3)
	for i := 0; i < n; i++ {
		dp[i+1] = make([]int, 3)
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if j != k {
					dp[i+1][k] = Max(dp[i][j]+chs[i][k], dp[i+1][k])
				}
			}
		}
	}

	var res int
	for i := 0; i < 3; i++ {
		res = Max(res, dp[n][i])
	}
	fmt.Printf("%#v\n", res)
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
