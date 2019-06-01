package main

import (
	"fmt"
	"strings"
)

var RE string
var cache [][]bool

// 通過リスト

func init() {
	RE = "No"
}

// https://atcoder.jp/contests/atc001/tasks/dfs_a

func main() {
	var h, w int
	fmt.Scan(&h)
	fmt.Scan(&w)

	m := make([][]string, h, h)
	cache = make([][]bool, h, h)
	var sl, sr int
	for i := 0; i < h; i++ {
		m[i] = make([]string, w, w)
		cache[i] = make([]bool, w, w)
		var n string
		fmt.Scan(&n)
		ss := strings.Split(n, "")
		for idx, item := range ss {
			if item == "s" {
				sl = i
				sr = idx
			}
			m[i][idx] = item
		}
	}

	fmt.Printf("%#v\n", cache)

	dfs(sl, sr, m)

	fmt.Println(RE)
}

func dfs(l, r int, m [][]string) {
	var v string
	if l >= len(m) || l < 0 || r >= len(m[l]) || r < 0 {
		return
	} else {
		v = m[l][r]
	}
	if v == "#" {
		return
	}
	fmt.Printf("%#v, %#v, %#v\n", v, l, r)
	if v == "g" {
		RE = "Yes"
	}
	if cache[l][r] {
		return
	}

	cache[l][r] = true

	// 右
	dfs(l, r+1, m)

	// 左
	dfs(l, r-1, m)

	// 上
	dfs(l+1, r, m)

	// 下
	dfs(l-1, r, m)
}
