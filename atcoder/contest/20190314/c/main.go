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

	temp := make([]int, n+1, n+1)

	l := len(ss)
	var c int
	for idx := 1; idx < l; idx++ {
		if ss[idx-1] == "A" && ss[idx] == "C" {
			c++
			temp[idx] = c
			continue
		}
		temp[idx] = c
	}
	// fmt.Printf("%#v\n", ss)
	// fmt.Printf("%#v\n", temp)

	for index := 0; index < q; index++ {
		s, e := SingleInt(), SingleInt()
		fmt.Println(temp[e-1] - temp[s-1])
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
