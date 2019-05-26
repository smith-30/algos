package main

import (
	"fmt"
	"strings"
)

// https://qiita.com/drken/items/a207e5ae3ea2cf17f4bd
func main() {
	var n string
	fmt.Scan(&n)

	ss := strings.Split(n, "")
	l := len(ss)

	dict := make([][]int64, l+1, l+1)
	dict[l] = make([]int64, 26, 26)
	for j := 0; j < 26; j++ {
		dict[l][j] = int64(l)
	}
	// res[i][c] := i 文字目以降で最初に文字 c が登場する index (存在しないときは n)
	for index := l - 1; index >= 0; index-- {
		dict[index] = make([]int64, 26, 26)
		for j := 0; j < 26; j++ {
			dict[index][j] = int64(l)
			dict[index][j] = dict[index+1][j]
		}
		r := []rune(ss[index])
		dict[index][r[0]-'a'] = int64(index)
	}

	for _, item := range dict {
		fmt.Printf("%#v\n", item)
	}

	fmt.Println()

	dp := make([]int64, n+1, n+1)
	dp[0] = 1
	for index := 0; index < n; index++ {
		for j := 0; j < 26; j++ {
			if dict[i][j]
		}
	}

	// for index := l - 1; index >= 0; index-- {
	// 	dict[index] = make([]int64, 26, 26)
	// 	for j := 0; j < 26; j++ {
	// 		dict[index][j] = dict[index+1][j]
	// 	}
	// 	r := []rune(ss[index])
	// 	fmt.Printf("%#v\n", r[0]-'a')
	// 	dict[index][r[0]-'a'] = int64(index)
	// }

	dp := make([]int64, l+1, l+1)
	dp[0] = 1

}
