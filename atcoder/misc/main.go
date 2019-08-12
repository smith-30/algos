package main

import (
	"fmt"
	"strings"
)

func main() {
	l := 30
	ss := make([]int, l+1)
	for i := 1; i <= 30; i++ {
		ss[i] = i
	}

	fmt.Printf("%#v\n", fb(ss))
}

// func fb(vars []int) string {
// 	s := make([]string, 0, len(vars))

// 	for _, item := range vars {
// 		switch {
// 		case item%3 == 0 && item%5 == 0:
// 			s = append(s, "Fizz Buzz")
// 		case item%3 == 0:
// 			s = append(s, "Fizz")
// 		case item%5 == 0:
// 			s = append(s, "Buzz")
// 		}
// 	}

// 	return strings.Join(s, " ")
// }

func fb(vars []int) string {
	s := make([]string, 0, len(vars))

	for _, item := range vars {
		if item%3 == 0 {
			s = append(s, "Fizz")
		}
		if item%5 == 0 {
			s = append(s, "Buzz")
		}
	}

	return strings.Join(s, " ")
}
