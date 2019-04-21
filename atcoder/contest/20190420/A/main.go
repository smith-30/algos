package main

import (
	"fmt"
)

func main() {
	a, b, c := SingleInt(), SingleInt(), SingleInt()
	if a < c && b > c {
		fmt.Println("Yes")
		return
	}
	if a > c && b < c {
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
