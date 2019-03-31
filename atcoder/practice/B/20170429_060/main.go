package main

import (
	"fmt"
	"math"
)

func main() {
	a, b, c := SingleInt(), SingleInt(), SingleInt()

	// sum := x*a + y*a + ...
	// c = sum % b

}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
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
