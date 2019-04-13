package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	n := SingleInt()
	str := SingleStr()
	strs := strings.Split(str, "")

	r := make([]int, 0, n)
	l := make([]int, 0, n)
	var nw, ne, wc, ec int

	for _, v := range strs {
		nw = wc
		ne = ec

		if v == "W" {
			wc++
		} else {
			ec++
		}

		if nw != wc {
			l = append(l, wc-1)
		} else {
			l = append(l, wc)
		}

		if ne != ec {
			r = append(r, ec-1)
		} else {
			r = append(r, ec)
		}
	}

	fmt.Printf("%#v\n", l)
	fmt.Printf("%#v\n", r)

	var min int
	var once sync.Once
	for i := 0; i < n; i++ {
		re := r[i] + l[n-i-1]
		once.Do(func() {
			min = re
		})
		if min > re {
			min = re
		}
	}

	fmt.Printf("%#v\n", min)
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
