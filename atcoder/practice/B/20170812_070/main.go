package main

import (
	"fmt"
	"math"
)

func main() {
	a, b, c, d := SingleInt(), SingleInt(), SingleInt(), SingleInt()

	max, min := Max(a, c), Min(b, d)

	if max > min {
		fmt.Println(0)
		return
	}

	// (max(A, C) > min(B, D) --- 重複区間なし

	fmt.Println(min - max)
	return
}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

func Max(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton max() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Max(float64(res), float64(nums[i])))
	}
	return res
}
func Min(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton min() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Min(float64(res), float64(nums[i])))
	}
	return res
}
