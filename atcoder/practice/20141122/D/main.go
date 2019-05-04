package main

import (
	"fmt"
	"math"
)

func main() {
	wid := SingleInt()
	n, k := SingleInt(), SingleInt()

	chs := make([][]int, n)
	width := 0
	value := 1

	for i := 0; i < n; i++ {
		chs[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			chs[i][j] = SingleInt()
		}
	}

	fmt.Printf("%#v\n", chs)
	fmt.Printf("%#v\n", k)
	fmt.Printf("%#v\n", wid)

	dp := make([][][]int, n+1)
	dp[0] = make([][]int, wid+1)
	for i := 0; i <= k; i++ {
		dp[0][i] = make([]int, k+1)
	}

	// 幅 -> chs[i][width]
	// 次 -> chs[i+1][width]

	for i := 0; i < n; i++ {
		dp[i+1] = make([][]int, wid+1)
		dp[i+1][wid] = make([]int, k+1)
		fmt.Printf("1 %#v\n", 1)
		for j := 0; j <= k; j++ {

			if wid >= chs[i][width] {
				fmt.Printf("aa %#v\n", dp[i+1][wid][j])
				fmt.Printf("aaa %#v\n", wid-chs[i][width])
				fmt.Printf("%#v\n", dp[i][wid])
				fmt.Printf("aaaa %#v\n", dp[i][wid][j])
				dp[i+1][wid][j] = Max(dp[i][wid-chs[i][width]][j]+chs[i][value], dp[i][wid][j])
			} else {
				dp[i+1][wid][j] = dp[i][wid][j]
			}
		}
	}
	fmt.Printf("%#v\n", dp)

	fmt.Printf("%#v\n", dp[n][wid][k])
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
