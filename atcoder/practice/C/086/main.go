package main

import (
	"fmt"
	"math"
)

func main() {
	n := SingleInt()

	points := [3]int{}
	hit := true
	for i := 0; i < n; i++ {
		nums := ScanNums(3)
		dt := nums[0] - points[0]
		dist := math.Abs(float64(nums[1]-points[1])) + math.Abs(float64(nums[2]-points[2]))
		if dt < int(dist) {
			hit = false
		}
		if dt%2 != int(dist)%2 {
			hit = false
		}
		points[0], points[1], points[2] = nums[0], nums[1], nums[2]
	}

	if hit {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
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
