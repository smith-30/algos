package main

import (
	"fmt"
	"strings"
)

func main() {
	s := SingleStr()
	ss := strings.Split(s, "")
	l := len(ss)

	if l == 0 {
		fmt.Printf("%#v\n", 0)
		return
	}

	var re int
	base := ss[0]
	for i := 1; i < l; i++ {
		if base == ss[i] {
			re++
			base = reverse(base)
		} else {
			base = ss[i]
		}
	}
	fmt.Printf("%#v\n", re)
}

func reverse(v string) string {
	if v == "0" {
		return "1"
	}
	return "0"
}

func SingleStr() string {
	var s string
	fmt.Scan(&s)
	return s
}
