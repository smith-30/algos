package main

import (
	"fmt"
	"math"
)

// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=DPL_1_B&lang=jp
func main() {
	n, ww := SingleInt(), SingleInt()
	vals, weis := make([]int, 0, n), make([]int, 0, n)

	for i := 0; i < n; i++ {
		vals = append(vals, SingleInt())
		weis = append(weis, SingleInt())
	}

	// dp[i+1][w] := i 番目までの品物の中から重さが w を超えないように選んだときの、価値の総和の最大値
	// dp[n][ww] を index　進めて最大値で更新していく

	// index + 1 に対して、

	dp := make([][]int, n+1)
	dp[0] = make([]int, ww+1)
	for i := 0; i < n; i++ {
		dp[i+1] = make([]int, ww+1)
		for w := 0; w <= ww; w++ {
			if w >= weis[i] {
				// dp[i][w-weis[i]]+vals[i]
				// 一つ前の重さの値に現在地を足し込む

				// Max(dp[i][w-weis[i]]+vals[i], dp[i][w])
				// 一つ前の最大値(dp[i][w])と、新規に値を追加したときの大きい方をとる

				fmt.Printf("%v, %#v, %v\n", w-weis[i], dp[i][w-weis[i]]+vals[i], dp[i][w])
				dp[i+1][w] = Max(dp[i][w-weis[i]]+vals[i], dp[i][w])
			} else {
				dp[i+1][w] = dp[i][w]
			}
		}
		fmt.Printf("%#v\n", dp[i+1])
	}

	fmt.Printf("%#v\n", dp[n][ww]/10000)

	// a := int64(math.Pow(3, float64(n)))
	// fmt.Printf("%#v\n", (a-(3*dp1[n-1])+(3*dp2[n-1]))%mod)
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

func Min64(nums ...int64) int64 {
	if len(nums) == 0 {
		panic("function min() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int64(math.Min(float64(res), float64(nums[i])))
	}
	return res
}
