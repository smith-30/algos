package main

import (
	"fmt"
)

func main() {
	a := SingleInt()
	slice := make([]int, 0, a)

	for i := 0; i < a; i++ {
		slice = append(slice, SingleInt())
	}

	// var accessor int
	// var lastAccessor int
	// c := 1
	// for {
	// 	idx := slice[accessor] - 1 // 2, 0, 2
	// 	lastAccessor = accessor    // 0, 2, 0
	// 	c++
	// 	accessor = idx // 2, 0, 2
	// 	if slice[accessor]-1 == lastAccessor {
	// 		if slice[lastAccessor] == 2 {
	// 			fmt.Println(c)
	// 			return
	// 		}
	// 		fmt.Println(-1)
	// 		return
	// 	}
	// 	if slice[idx] == 2 {
	// 		fmt.Println(c)
	// 		return
	// 	}
	// }

	var now, c int
	for {
		if now == 2 {
			fmt.Println(c)
			return
		}
		if c >= a {
			fmt.Println(-1)
			return
		}
		c++
		if now > 0 {
			now--
		}
		now = slice[now]
	}

	// ボタン押してから現在のボタンの値を取得するようにか
	// N回までに値がでなかったら、重複したボタンを押していたらなんらかのループが起きている
}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

// N 回ボタンを押してもボタン 2 が光らなかったとします。このとき、その間に N + 1 回いずれかのボタン
// が光るので、そのうちのある 2 つは同じボタンです。同じボタンが光っている状態から光っているボタンを
// 押すことを繰り返しても、同じ順にボタンが光るだけなので、この場合ボタン 2 は永遠に光ることはありま
// せん。
