package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	fib := make([]int, n, n)
	fib[0] = 0
	fib[1] = 1

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	fmt.Println(fib)
}
