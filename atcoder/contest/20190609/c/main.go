package main

import (
	"fmt"
	"math"
)

const d = 1000000007

func main() {
	var n, p int
	fmt.Scan(&n)
	fmt.Scan(&p)

	var n1 int64
	fmt.Scan(&n1)
	dans := make([]int64, p, p)
	dans[0] = n1
	mdans := map[int64]int64{
		n1: n1,
	}

	for index := 1; index < p; index++ {
		var n1 int64
		fmt.Scan(&n1)
		if dans[index-1] == n1 {
			continue
		}
		dans[index] = n1
		if math.Abs(float64(dans[index-1]-dans[index])) == 1 {
			fmt.Println(0)
			return
		}
		mdans[n1] = n1
	}

	dp := make([]int64, n+1, n+1)
	dp[0] = 1

	for index := 0; index < n; index++ {
		for j := index + 1; j <= Min(n, index+2); j++ {
			if _, ok := mdans[int64(j)]; !ok {
				dp[j] += dp[index]
				dp[j] %= d
			}
		}
	}

	fmt.Printf("%#v\n", dp)
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
