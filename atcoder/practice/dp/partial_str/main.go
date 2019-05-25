// https://qiita.com/drken/items/a207e5ae3ea2cf17f4bd
package main

import (
	"fmt"
	"strings"
)

func main() {
	var n string
	fmt.Scan(&n)

	ss := strings.Split(n, "")
	l := len(ss)

	dict := make([][]int64, l+1, l+1)
	dict[l] = make([]int64, 26, 26)
	// res[i][c] := i 文字目以降で最初に文字 c が登場する index (存在しないときは n)
	for index := l - 1; index >= 0; index-- {
		dict[index] = make([]int64, 26, 26)
		for j := 0; j < 26; j++ {
			dict[index][j] =
				dict[index+1][j]
		}
		r := []rune(ss[index])
		dict[index][r[0]-'a'] = int64(index)
	}

	fmt.Printf("%#v\n", dict)

	dp := make([]int64, l+1, l+1)
	dp[0] = 1
}
