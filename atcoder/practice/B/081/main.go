package main

import (
	"fmt"
)

func main() {
	n := SingleInt()
	nums := ScanNums(n)

	// finish if nums have odd.
	for _, item := range nums {
		if item%2 != 0 {
			fmt.Println(0)
			return
		}
	}

	c := 0
	for {
		b := nums[:0]
		for _, x := range nums {
			if x%2 != 0 {
				fmt.Println(c)
				return
			}
			b = append(b, x/2)
		}
		c++
		nums = b
	}
}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

func ScanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}
