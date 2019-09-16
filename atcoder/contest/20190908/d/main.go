package main

import (
	"fmt"
	"math"
	"strings"
)

func SInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

func DInt() (int, int) {
	var n int
	fmt.Scan(&n)

	var m int
	fmt.Scan(&m)
	return n, m
}

func SIntSlice(n int) []int {
	re := make([]int, 0, n)
	for i := 0; i < n; i++ {
		re = append(re, SInt())
	}
	return re
}

func SStr() string {
	var s string
	fmt.Scan(&s)
	return s
}

func StrSlice() []string {
	var s string
	fmt.Scan(&s)
	return strings.Split(s, "")
}

func main() {
	n, k := DInt()
	ans := n - 1
	s := SStr()
	ss := strings.Split(s, "")
	if len(ss) == 1 {
		ans = 0
	}
	var re int
	for index := 0; index < len(ss)-1; index++ {
		if ss[index] == ss[index+1] {
			re++
		}
	}
	ansre := re + 2*k

	fmt.Println(Min(ans, ansre))
}

func Min(nums ...int) int {
	if len(nums) == 0 {
		panic("function min() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Min(float64(res), float64(nums[i])))
	}
	return res
}
