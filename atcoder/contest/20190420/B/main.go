package main

import (
	"fmt"
	"strings"
)

func main() {
	_, b, c := SingleInt(), SingleStr(), SingleInt()
	strs := strings.Split(b, "")
	k := strs[c-1]

	for _, v := range strs {
		if k == v {
			fmt.Printf("%v", v)
		} else {
			fmt.Printf("%v", "*")
		}

	}
	fmt.Println()
}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

func SingleStr() string {
	var s string
	fmt.Scan(&s)
	return s
}
