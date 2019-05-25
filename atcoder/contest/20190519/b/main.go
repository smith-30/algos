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

	if 0 < c && c <= 12 {
		if 0 < d && d <= 12 {
			fmt.Println("AMBIGUOUS")
		} else {
			fmt.Println("MMYY")
		}
	} else {
		if 0 < d && d <= 12 {
			fmt.Println("YYMM")
		} else {
			fmt.Println("NA")
		}
	}

}
