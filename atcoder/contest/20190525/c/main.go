package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)

	var lma, rmin int

	rmin = 9999999999999

	for index := 0; index < m; index++ {
		var l, r int
		fmt.Scan(&l)
		fmt.Scan(&r)

		if lma < l {
			lma = l
		}
		if rmin > r {
			rmin = r
		}
	}

	var re int

	for index := 1; index <= n; index++ {
		if lma <= index && rmin >= index {
			re++
		}
	}

	fmt.Println(re)
}
