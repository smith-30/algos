package main

import "fmt"

func main() {
	n, q := SingleInt(), SingleInt()
	s := SingleStr()
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
