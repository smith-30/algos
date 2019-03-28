package main

import (
	"fmt"
)

func main() {
	n := SingleInt()
	total := SingleInt()

	re := [3]int{}

	a, b, c, hit := getResult(total, n)
	if hit {
		re[0] = a
		re[1] = b
		re[2] = c
	}

	if hit {
		fmt.Printf("%d %d %d\n", re[0], re[1], re[2])
		return
	}
	fmt.Printf("%d %d %d\n", -1, -1, -1)
}

func getResult(total, n int) (int, int, int, bool) {
	for i := 0; i <= n; i++ {
		for j := 0; i+j <= n; j++ { // ここ注意
			k := n - i - j
			if i*10000+j*5000+k*1000 == total {
				return i, j, k, true
			}
		}
	}
	return 0, 0, 0, false
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
