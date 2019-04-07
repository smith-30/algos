package main

import (
	"fmt"
)

type P struct {
	complete bool
	count    int64
}

func main() {
	n := SingleInt()

	vehicles := ScanNums(5)
	for i := len(vehicles)/2 - 1; i >= 0; i-- {
		opp := len(vehicles) - 1 - i
		vehicles[i], vehicles[opp] = vehicles[opp], vehicles[i]
	}
	var t int64

	status := []P{
		P{complete: false, count: 0},
		P{complete: false, count: 0},
		P{complete: false, count: 0},
		P{complete: false, count: 0},
		P{complete: false, count: 0},
		P{complete: true, count: n},
	}

	lastIdx := len(status) - 1
	for {
		t++
		for k, _ := range status {
			if k != lastIdx {
				if 0 < status[k+1].count {
					c := plus(vehicles[k], status[k+1].count)
					status[k].count += c
					status[k+1].count = minus(status[k+1].count, c)
					if status[k].count == n {
						status[k].complete = true
					}
				}
			}
		}

		if status[0].count == n {
			fmt.Println(t)
			return
		}
	}
}

func moveCount(n, t int64) int64 {
	if n < t {
		return n
	}
	return t
}

func plus(n, v int64) int64 {
	if n < v {
		return n
	}
	return v
}

func minus(n, v int64) int64 {
	if n < v {
		return 0
	}
	return n - v
}

func SingleInt() int64 {
	var n int64
	fmt.Scan(&n)
	return n
}

func ScanNums(len int) (nums []int64) {
	var num int64
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}
