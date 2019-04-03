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

	line := 1
	row := 1
	c := points[line][row]

	for {
		switch line {
		case 1:
			if row == n {
				line++
			}
			if points[line][row+1] >= points[line+1][row] {
				if line != 2 {
					row++
				}

			} else {
				var rowSome, lineSome int
				for i := row + 1; i <= n; i++ {
					rowSome += points[line][row]
					lineSome += points[line+1][row]
				}
				lineSome += points[line+1][row]
				if lineSome > rowSome {
					line++
				}
			}

			c += points[line][row]
		case 2:
			row++
			fmt.Printf("%#v\n", points[line][row])
			c += points[line][row]
		}

		if line == 2 && row == n {
			fmt.Println(c)
			return
		}
	}

	// (2, N)がゴール
	// 下か右にしか移動できない。
	// 上には戻れない(下に行くのは1回のみ 2*Nのmapなので)

	// 下を選択したら、もう右にしか行けない
	// どのタイミングで下に行くかが重要か。
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
