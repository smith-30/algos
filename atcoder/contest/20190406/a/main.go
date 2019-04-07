package main

import (
	"fmt"
	"math"
)

func main() {
	ss := ScanNums(5)
	k := SingleInt()

	for idx, _ := range ss {
		for _, item := range ss {
			if math.Abs(float64(ss[idx]-item)) > float64(k) {
				fmt.Println(":(")
				return
			}
		}
	}

	fmt.Println("Yay!")
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
