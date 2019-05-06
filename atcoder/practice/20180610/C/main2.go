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
			if i == 36 {
				fmt.Printf("tt6 - %#v \n", t)
			}
			t /= 6
			if i == 36 {
				fmt.Printf("c - %#v \n", c)
			}
		}
		if i == 36 {
			fmt.Printf("idx: %v c - %#v ", i, c)
		}

		temp := c
		// 127 - 110 = 17
		t = n - i
		if i == 36 {
			fmt.Printf("t(n-i) %#v ", t)
		}
		for t > 0 {
			c += t % 9
			if i == 36 {
				fmt.Printf("tt6 - %#v \n", t)
			}
			t /= 9
			if i == 36 {
				fmt.Printf("c - %#v \n", c)
			}
		}
		if i == 36 {
			fmt.Printf("inc %v c - %#v\n", c-temp, c)
		}

		if c < m {
			m = c
		}
	}
	Println(m)
}
