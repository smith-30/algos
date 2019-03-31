package main

import (
	"fmt"
	"math"
)

func main() {
	h := SingleInt()
	p := 1
	m := math.Pow(10, 9) + 7

	for i := 1; i <= h; i++ {
		p = i * p % int(m)
	}

	fmt.Printf("%v\n", p)
}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
}
