package main

import (
	"fmt"
	. "fmt"
)

func main() {
	var n, m int
	Scan(&n)
	m = n + 1
	// 127
	for i := 0; i < n+1; i++ {
		var c, t int
		c = 0
		// i = 110
		t = i
		for t > 0 {
			c += t % 6
			t /= 6
		}

		temp := c
		// 127 - 110 = 17
		t = n - i
		for t > 0 {
			c += t % 9
			t /= 9
		}
		fmt.Printf("***** %#v\n", c)

		if c < m {
			m = c
		}
	}
	Println(m)
}
