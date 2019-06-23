package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
}

func nextInt() int {
	sc.Scan()

	i, e := strconv.ParseInt(sc.Text(), 10, 64)
	if e != nil {
		panic(e)
	}
	return int(i)
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
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

func main() {
	n, w := nextInt(), nextInt()
	vals := make([]int, 0, n)
	wws := make([]int, 0, n)

	for i := 0; i < n; i++ {
		vv, w := nextInt(), nextInt()
		vals = append(vals, vv)
		wws = append(wws, w)
	}

	dp := make([][]int, n+1)
	dp[0] = make([]int, w+1)

	for i := 0; i < n; i++ {
		dp[i+1] = make([]int, w+1)

		for j := 0; j <= w; j++ {
			// j = 6
			// wws[i] = 2
			// dp[i+1][6] = Max(dp[i][4] + vals[i], dp[i][6])
			if j >= wws[i] {
				dp[i+1][j] = Max(
					dp[i][j-wws[i]]+vals[i], // 一つ前のdp結果の値にアクセスして足し算行って、大きい方の値を採用する
					dp[i][j],
				)
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}

	for _, item := range dp {
		fmt.Printf("%#v\n", item)
		fmt.Println()
	}

	fmt.Printf("%#v\n", dp[n][w])
}
