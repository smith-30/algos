package main

import (
	"fmt"
	"strings"
)

var coos [][]string
var history [][]bool
var h, w int
var exists bool

func main() {
	h, w = SingleInt(), SingleInt()
	coos = make([][]string, h)
	history = make([][]bool, h)
	var sk, sr int

	for i := 0; i < h; i++ {
		coos[i] = make([]string, w)
		history[i] = make([]bool, w)
		vs := strings.Split(SingleStr(), "")
		for j, item := range vs {
			if item == "s" {
				sk = i
				sr = j
			}
			coos[i][j] = item
		}
	}

	search(sk, sr)
	if exists {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

// 再帰の勘所としては、探索したい方向の数だけ関数内で呼び出す
// 自身がアクセスした場所を覚えておくようにglobal領域にキャッシュを置く
// どういうときに来た道を戻るべきか、終了条件や成約をはっきりさせる
func search(x, y int) {
	if exists {
		return
	}
	if !(x >= 0 && x < h && y >= 0 && y < w) {
		return
	}
	v := coos[x][y]
	if v == "g" {
		exists = true
		return
	}

	if v == "#" {
		return
	}

	if history[x][y] {
		return
	}

	history[x][y] = true

	search(x+1, y)
	search(x-1, y)
	search(x, y+1)
	search(x, y-1)
}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

func SingleStr() string {
	var s string
	fmt.Scan(&s)
	return s
}
