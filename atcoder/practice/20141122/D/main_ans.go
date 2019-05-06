package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var W, N, K int
	fmt.Scan(&W)
	fmt.Scan(&N, &K)

	dp := make([][]int, 51)
	for i := 0; i < 51; i++ {
		dp[i] = make([]int, 10001)
	}

	ans := 0
	for i := 0; i < N; i++ {
		var A, B int
		fmt.Scan(&A, &B)
		for j := K; j > 0; j-- {
			for k := A; k <= W; k++ {
				dp[j][k] = max(dp[j][k], dp[j-1][k-A]+B)
				ans = max(ans, dp[j][k])
			}
		}
	}

	fmt.Println(ans)
}
