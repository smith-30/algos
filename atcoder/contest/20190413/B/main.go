package main

import (
	"fmt"
)

func main() {
	n := SingleInt()
	ss := ScanNums(n)

	re := 1
	lim := ss[0]
	for i := 1; i < n; i++ {
		if lim <= ss[i] {
			re++
			lim = ss[i]
		}
	}
	fmt.Printf("%#v\n", re)
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
