package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)

	vals := make([]int, 0, n)
	for index := 0; index < n; index++ {
		var n int
		fmt.Scan(&n)

		vals = append(vals, n)
	}

	var res []int
	var re int
	for _, item := range res {
		re += item
	}
	fmt.Println(re)
}
