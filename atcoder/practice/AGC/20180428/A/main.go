package main

import (
	"fmt"
	"sort"
)

func main() {
	n := SingleInt()
	orig := ScanNums(n)
	ss := make([]int, n+1, n+1)

	for i, v := range orig {
		ss[i+1] = ss[i] + v
	}
	sort.Sort(sort.IntSlice(ss))

	base := 0
	cnt := 0
	ress := make([]int, 0, n+1)
	for _, v := range ss {
		if v == base {
			cnt++
		} else {
			if cnt > 1 {
				ress = append(ress, cnt)
			}
			cnt = 1
			base = v
		}
	}
	if cnt > 1 {
		ress = append(ress, cnt)
	}
	var ans int
	for _, v := range ress {
		for 
	}
	fmt.Printf("%#v\n", ress)
	fmt.Printf("%#v\n", cnt-1)
	fmt.Printf("%#v\n", ss)
}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

func ScanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}
