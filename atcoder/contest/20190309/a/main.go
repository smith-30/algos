package main

import "fmt"

func main() {
	var H, W int
	fmt.Scan(&H)
	fmt.Scan(&W)

	var h, w int
	fmt.Scan(&h)
	fmt.Scan(&w)

	fmt.Println((H - h) * (W - w))
}
