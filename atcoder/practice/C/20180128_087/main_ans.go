package main

import (
	"fmt"
)

func main() {
	n := SingleInt()

	points := map[int]map[int]int{}
	for i := 1; i <= 2; i++ {
		points[i] = map[int]int{}
		ps := ScanNums(n)
		for idx, item := range ps {
			points[i][idx+1] = item
		}
	}

	line := 1
	row := 1
	max := points[line][row]
	line++
	max += points[line][row]
	for i := 1; i < n; i++ {
		max += points[line][row+i]
	}

	for i := 2; i < n; i++ {
		line := 1
		c := points[line][row]
		for row := 2; row <= n; row++ {
			if i == row {
				c += points[line][row]
				line++
				c += points[line][row]
				continue
			}
			c += points[line][row]
		}
		if c > max {
			max = c
		}
	}

	fmt.Printf("%v\n", max)
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
