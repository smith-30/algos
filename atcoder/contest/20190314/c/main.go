package main

import (
	"fmt"
	"strings"
)

func main() {
	n, q := SingleInt(), SingleInt()

	s := SingleStr()
	ss := make([]string, 0, n)
	ss = strings.Split(s, "")

	temp := make([]string, n+1, n+1)

	l := len(ss)

	for i, item := range ss {
		if i+1 < l {
			if item == "A" && ss[i+1] == "C" {
				temp[i] = "*"
				ss[i+1] = ""
				continue
			}
		}
		temp[i] = item
	}

	for index := 0; index < q; index++ {
		s, e := SingleInt(), SingleInt()
		var c int
		for j := s - 1; j < e-1; j++ {
			if temp[j] == "*" {
				c++
			}
		}
		fmt.Println(c)
	}

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
