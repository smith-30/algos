package main

import "fmt"

func main() {
	bc := ScanNums(2)
	b, c := bc[0], bc[1]

	re := b * c
	if re%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func ScanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}
