package main

import (
	"fmt"
	"math"
)

// 与えられた数列を R, G, B いずれかで塗って
// それぞれの総和で三角形を作る
// そのとき三角形の作り方は何通りあるか

// 三角形の条件
// | a - b | < c < a + b
// a < b + c

// a + b + c = S
// max(a, b, c) < S/2
func main() {
	n := SingleInt()
	chs := make([]int64, 0, n)
	var sum int64
	// mod := int64(998244353)

	for i := 0; i < n; i++ {
		item := int64(SingleInt())
		sum += item
		chs = append(chs, item)
	}

	halfSum := CeilInt64(sum, 2)

	fmt.Printf("halfSum %#v\n", halfSum)

	dp1 := make([][]int64, n)
	// dp2 := make([][]int64, n)
	dp1[0] = make([]int64, sum+1)
	dp1[0][0] = 0
	for i := 1; i < n; i++ {
		dp1[i] = make([]int64, sum+1)
		for j := 1; j <= sum; j++ {
			dp1[i][j] = Ma
		}
	}
	fmt.Printf("%#v\n", dp1[n][sum])

	// a := int64(math.Pow(3, float64(n)))
	// fmt.Printf("%#v\n", (a-(3*dp1[n-1])+(3*dp2[n-1]))%mod)
}

// 0, []int{2, 1, 0, 0, 0, 0}
// 1, []int{4, 4, 1, 0, 0, 0}
// 2, []int{8, 12, 6, 1, 0, 0}
// 3, []int{16, 24, 20, 14, 6, 1}

func CeilInt64(a, b int64) int64 {
	d := float64(a) / float64(b)
	return int64(math.Ceil(d))
}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
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
