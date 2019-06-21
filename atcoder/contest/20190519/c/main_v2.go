package main

import "fmt"

func main() {
	var n, k float64
	fmt.Scan(&n)
	fmt.Scan(&k)

	var re float64

	var nocoin int
	for i := 1; i <= int(n); i++ {
		if float64(i) >= k {
			nocoin++
			continue
		}

		// f := i << 1

		c := float64(1)
		f := float64(i)
		for {
			if f >= k {
				break
			}

			f *= 2
			c *= 2
		}

		fmt.Printf("%#v\n", c)

		re += (1 / c)
	}

	re += (float64(nocoin) / float64(n))
	fmt.Printf("%.9f\n", re)
}
