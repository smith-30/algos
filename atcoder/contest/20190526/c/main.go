package main

import (
	"fmt"
	"math"
)

type sw struct {
	status bool
}

type denkyu struct {
	condition int
	relation  []int
}

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)

	sws := make([]*sw, 0, n)
	for index := 0; index < n; index++ {
		sws = append(sws, &sw{status: true})
	}
	ds := make([]denkyu, 0, m)

	for index := 0; index < m; index++ {
		var n int
		fmt.Scan(&n)

		d := denkyu{}
		for j := 0; j < n; j++ {
			var sw int
			fmt.Scan(&sw)
			d.relation = append(d.relation, sw-1)
		}
		ds = append(ds, d)
	}

	for index := 0; index < m; index++ {
		var n int
		fmt.Scan(&n)

		ds[index].condition = n
	}

	var re int
	pattern := int(math.Pow(2, float64(n)))
	for index := 0; index < pattern; index++ {
		pa := map[int]int{}
		for j := 0; j < n; j++ {
			pa[j] = 0
			if uint16(index)>>uint16(j)&1 == 1 {
				pa[j] = 1
			}
		}
		ok := true
		for _, item := range ds {
			var c int
			for _, r := range item.relation {
				if pa[r] == 1 {
					c++
				}
			}
			if c%2 != item.condition {
				ok = false
			}
		}
		if ok {
			re++
		}
	}

	fmt.Println(re)
}
