package main

import "fmt"

func main() {
	var n, p int
	fmt.Scan(&n)
	fmt.Scan(&p)

	ss := make([]int, n+1)

	for index := 0; index < n; index++ {
		var n int
		fmt.Scan(&n)

		ss[index] = n
	}

	var re, sum, r int

	for l := 0; l < n; l++ {
		for {
			if sum >= p || r == n {
				break
			}

			sum += ss[r]
			r++
		}

		if sum < p {
			break
		}

		re += n - r + 1
		sum -= ss[l]
	}

	fmt.Printf("%#v\n", re)
}
