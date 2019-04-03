package main

import (
	"fmt"
)

func main() {
	n := SingleInt()

	points := map[int]map[int]int{}
	for i := 1; i <= 2; i++ {
		points[i] = map[int]int{}
		ps := ScanNums(n)
		for idx, item := range ps {
			points[i][idx+1] = item
		}
	}

	
	max := 0

	for i := 0; i <= n; i++ {
		line := 1
	row := 1
	c := points[line][row]
		var re int
		for row := 1; row < n; row++ {
			if i == 0 {
				l
			}
		}
	}

	// (2, N)がゴール
	// 下か右にしか移動できない。
	// 上には戻れない(下に行くのは1回のみ 2*Nのmapなので)

	// 下を選択したら、もう右にしか行けない
	// どのタイミングで下に行くかが重要か。
	fmt.Printf("%#v\n", c)
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
