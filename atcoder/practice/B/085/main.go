package main

import (
	"fmt"
)

func main() {
	n := SingleInt()
	nums := ScanNums(n)

	re := map[int]struct{}{}
	for _, item := range nums {
		re[item] = struct{}{}
	}
	fmt.Println(len(re))
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
