package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)

	var re float64
	for index := 1; index <= n; index++ {
		b := 1 / float64(n)
		if k-1 < index {
			re += b
			continue
		}
		logb := float64(k) / float64(index)
		re += b * math.Pow(0.5, math.Ceil(math.Log2(logb)))
	}

	fmt.Printf("%v\n", re)
}
