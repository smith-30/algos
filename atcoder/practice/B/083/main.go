package main

import (
	"fmt"
)

func main() {
	nums := ScanNums(3)
	n, a, b := nums[0], nums[1], nums[2]

	var sum, re int
	for i := 1; i <= n; i++ {
		sum = SumDigits(i)
		if sum >= a && sum <= b {
			re += i
		}
	}
	fmt.Println(re)
}

func ScanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}

func SumDigits(number int) int {
	var remainder, sumResult int
	for number != 0 {
		remainder = number % 10
		sumResult += remainder
		number = number / 10
	}
	return sumResult
}
