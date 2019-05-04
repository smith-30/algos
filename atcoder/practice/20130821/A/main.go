package main

import (
	"fmt"
	"math"
)

func main() {
	n := SingleInt()

	chs := make([]int, n)

	for i := 0; i < n; i++ {
		chs[i] = SingleInt()
	}

	// 10002 --> 一つの配点 max - 100点
	//			 問題の数   max - 100
	//
	dp := make([]bool, 10002)
	dp[0] = true

	for i := 0; i < n; i++ {
		// 0から始めると、j+chs[i]で登録した値に対して
		// 同じループ内でindexアクセスができてしまうので逆順から始める
		for j := 10000; j >= 0; j-- {
			if dp[j] {
				dp[j+chs[i]] = true
			}
		}
	}
	sum := 0
	for i := 0; i < 10001; i++ {
		if dp[i] {
			sum++
		}
	}
	fmt.Println(sum)
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
