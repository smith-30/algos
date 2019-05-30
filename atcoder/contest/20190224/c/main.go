package main

import "fmt"

func main() {
	var n, a, b, c int
	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	list := map[int]int{}
	for index := 0; index < n; index++ {
		var l int
		fmt.Scan(&l)

		list[l] = l
	}

	fn := func() int {
		return 0
	}

	fmt.Println(fn())
}
