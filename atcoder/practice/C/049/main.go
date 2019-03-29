package main

import (
	"fmt"
)

func main() {
	s := SingleStr()
	strs := []string{"dreamer", "eraser", "dream", "erase"}
	var hit bool

	s = StrReverse(s)
	for i, item := range strs {
		strs[i] = StrReverse(item)
	}

	for i := 0; i < len(s); {
		var strHit bool
		for _, item := range strs {
			l := len(item)
			if len(s) >= i+l {
				if item == s[i:i+l] {
					strHit = true
					i += l
					break
				}
			}
		}
		if len(s) == i && strHit {
			hit = true
			break
		}
		if !strHit {
			break
		}
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

func StrReverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
