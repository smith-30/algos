package main

import (
	"fmt"
	"sort"
)

func main() {
	n := SingleInt()
	nums := ScanNums(n)

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	var a, b int
	for idx, item := range nums {
		if idx%2 == 0 {
			a += item
		} else {
			b += item
		}
	}
	fmt.Println(a - b)
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
