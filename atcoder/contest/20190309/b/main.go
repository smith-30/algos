package main

import "fmt"

func main() {
	var n, m, c int
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&c)

	b := ScanNums(m)
	var re int
	for index := 0; index < n; index++ {
		a := ScanNums(m)
		var rs int
		for j := 0; j < m; j++ {
			rs += a[j] * b[j]
		}
		if rs+c > 0 {
			re++
		}
	}
	fmt.Printf("%#v\n", re)
}

func ScanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}
