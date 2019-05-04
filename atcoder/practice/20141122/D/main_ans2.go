package main

import (
	"fmt"
)

const INF = 1000000000

type screanshot struct {
	width      int
	importance int
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var w, n, k int
	fmt.Scan(&w, &n, &k)
	ss := make([]screanshot, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ss[i].width, &ss[i].importance)
	}

	ansBoard := make([][]int, k+1)
	for i := range ansBoard {
		line := make([]int, w+1)
		for j := range line {
			switch {
			case i == 0 && j == 0:
				line[j] = 0
			default:
				line[j] = -INF
			}
		}
		ansBoard[i] = line
	}

	ans := 0
	for _, s := range ss {
		for i := k; i >= 1; i-- {
			for j := s.width; j <= w; j++ {
				ansBoard[i][j] = maxInt(ansBoard[i][j], ansBoard[i-1][j-s.width]+s.importance)
				ans = maxInt(ans, ansBoard[i][j])
			}
		}
		fmt.Printf("%#v\n", ansBoard)
	}

	fmt.Println(ans)
}
