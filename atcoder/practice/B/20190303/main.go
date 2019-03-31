package main

import (
	"fmt"
	"math"
)

func main() {
	a, b, k := SingleInt(), SingleInt(), SingleInt()

	re := []int{}
	ra := Min(a, b)

	for i := ra; 1 <= i; i-- {
		if b%i == 0 && a%i == 0 {
			re = append(re, i)
		}
	}
	fmt.Println(re[k-1])
}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
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
