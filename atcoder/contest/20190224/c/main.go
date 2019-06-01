package main

import (
	"fmt"
	"math"
)

const (
	inf = 1000000000
)

var N int
var L []int

func main() {
	var n, a, b, c int
	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	N = n

	list := []int{}
	for index := 0; index < n; index++ {
		var l int
		fmt.Scan(&l)

		list[index] = l
	}

	L = list

	fmt.Println(sel())
}

func sel(cur, a,b,c int) int {
	if N == cur {
		return 
	}
	return 0
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
