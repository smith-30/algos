package main

import (
	"fmt"
	"strings"
)

func main() {
	b := SingleStr()
	bs := strings.Split(b, "")
	var max int

	for idx, _ := range bs {
		t := bs[idx:]
		var c int
		for _, v := range t {
			var end bool
			switch v {
			case "A", "C", "G", "T":
				c++
			default:
				end = true
			}
			if end {
				break
			}
		}
		if max < c {
			max = c
		}
	}
	fmt.Println(max)
}

func SingleStr() string {
	var s string
	fmt.Scan(&s)
	return s
}
