package main

import (
	"fmt"
	"math"
)

type sw struct {
	vals []int
	cond int
}

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)

	sws := make([]sw, m)

	for i := 0; i < m; i++ {
		var n int
		fmt.Scan(&n)

		for j := 0; j < n; j++ {
			var v int
			fmt.Scan(&v)
			sws[i].vals = append(sws[i].vals, v-1)
		}
	}

	for i := 0; i < m; i++ {
		var n int
		fmt.Scan(&n)

		sws[i].cond = n
	}

	pattern := int(math.Pow(2, float64(n)))
	on_off := map[int]int{}
	re := 0

	for index := 0; index < pattern; index++ {
		for j := 0; j < n; j++ {
			on_off[j] = int(uint16(index) >> uint16(j) & 1)
		}

		ok := true
		for _, item := range sws {
			sum := 0
			for _, v := range item.vals {
				sum += on_off[v]
			}
			if item.cond != sum%2 {
				ok = false
				break
			}
		}
		if ok {
			re++
		}
	}

	fmt.Printf("%#v\n", re)
}
