package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	if b >= a*c {
		fmt.Println(c)
		return
	}

	if a > b {
		fmt.Println(0)
		return
	}

	fmt.Println(b / a)
}
