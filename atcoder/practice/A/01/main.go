package main

import "fmt"

func main() {
	a := SingleInt()
	bc := ScanNums(2)
	b, c := bc[0], bc[1]

	s := SingleStr()

	fmt.Printf("%d %s\n", (a + b + c), s)
}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

func SingleStr() string {
	var s string
	fmt.Scan(&s)
	return s
}

func ScanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}
