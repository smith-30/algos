package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	ls := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var nn int
		fmt.Scan(&nn)
		ls = append(ls, nn)
	}

	m := Max(ls...)

	var exs bool
	var mm int
	for _, item := range ls {
		if item != m {
			mm += item
		} else {
			if exs {
				mm += item
			}
			exs = true
		}
	}

	if mm > m {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
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
