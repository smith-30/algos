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
	dp := make([][]int, n+1)
	dp[0] = make([]int, w+1)

	for i := 1; i <= n; i++ {
		dp[i] = make([]int, w+1)

		v, ww := nextInt(), nextInt()
		fmt.Printf("%#v\n", ww)
		for weight, val := range dp[i-1] {
			fmt.Printf("%#v\n", weight+ww)
			if weight+ww <= w {
				dp[i][weight+ww] = v + val
				continue
			}

			dp[i][weight] = Max(v, val)
		}
	}

	for _, item := range dp {
		fmt.Printf("%#v\n", item)
	}

	fmt.Printf("%#v\n", dp[n][w])
}
