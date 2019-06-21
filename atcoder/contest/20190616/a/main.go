package main

import "fmt"

func main() {
	var n, p int
	fmt.Scan(&n)
	fmt.Scan(&p)

	if n < p {
		fmt.Println(0)
	} else {
		fmt.Println(10)
	}
}
