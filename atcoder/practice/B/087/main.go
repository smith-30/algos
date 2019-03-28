package main

import (
	"fmt"
)

func main() {
	a := SingleInt()
	b := SingleInt()
	c := SingleInt()
	yen := SingleInt()

	var re int
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if i*500+j*100+k*50 == yen {
					re++
				}
			}
		}
	}

	fmt.Println(re)
}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
}
