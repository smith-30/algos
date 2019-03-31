package main

import "fmt"

func main() {
	a, b, c := SingleInt(), SingleInt(), SingleInt()

	re := b / a
	if c <= re {
		fmt.Println(c)
	}
	if re < c {
		fmt.Println(re)
	}
}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
}
