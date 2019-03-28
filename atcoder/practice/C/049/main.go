package main

import (
	"fmt"
)

func main() {
	s := SingleStr()

	a := []string{"dreamer", "eraser"}
	b := []string{"dream", "erase"}

	slen := len(s)
	var startIdx, endIdx int
	var hit bool

	fmt.Printf("slen %#v\n", slen)

	for {
		endIdx = startIdx + 6

		if slen >= endIdx {
			if startIdx != 0 {
				startIdx -= 1
			}

			if StrSearch(a, s[startIdx:endIdx]) {
				fmt.Printf("1 %#v\n", s[startIdx:endIdx])
				startIdx = endIdx
				hit = true
			}
		}

		if !hit {
			endIdx = startIdx + 5
			if slen >= endIdx {
				if StrSearch(b, s[startIdx:endIdx]) {
					startIdx = endIdx + 1
					hit = true
				}
			}
		}

		if hit {
			fmt.Printf("endIdx %#v %#v\n", endIdx, s[startIdx:endIdx])
			if slen == endIdx {
				break
			}
			hit = false
			continue
		}

		break
	}

	if hit {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}

func SingleStr() string {
	var s string
	fmt.Scan(&s)
	return s
}

func StrSearch(a []string, b string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == b {
			return true
		}
	}
	return false
}
