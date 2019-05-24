package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n string
	fmt.Scan(&n)

	nn := strings.Split(n, "")
	a := nn[0] + nn[1]
	b := nn[2] + nn[3]

	c, _ := strconv.Atoi(a)
	d, _ := strconv.Atoi(b)

	switch {
	case 0 < c && c <= 12 && 0 < d && d <= 12:
		fmt.Println("AMBIGUOUS")
	case (0 == c) || (0 == d || d > 12):
		fmt.Println("NA")
	case c > 12 && 0 < d && d <= 12:
		fmt.Println("YYMM")
	case d > 12 && 0 < c && c <= 12:
		fmt.Println("MMYY")
	}
}
