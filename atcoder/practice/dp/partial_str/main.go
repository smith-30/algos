package main

import (
	"fmt"
	"strings"
)

const sep = 1000000007

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

	//
	// debug
	//
	for _, item := range dict {
		fmt.Printf("%#v\n", item)
	}

	fmt.Println()

	dp := make([]int64, l+1, l+1)
	dp[0] = 1
	for i := 0; i < l; i++ {
		for j := 0; j < 26; j++ {
			// 次の文字 j がない場合はスルー
			if dict[i][j] >= int64(l) {
				continue
			}
			fmt.Printf("i: %#v ::: %#v, %#v\n", i, dict[i][j]+1, dp[i])
			dp[dict[i][j]+1] += dp[i]
		}
		fmt.Printf("%#v\n", dp)
	}

	// fmt.Printf("%#v\n", dp)

	var re int64
	for _, item := range dp {
		re += item
	}

	fmt.Println(re % sep)

}
