package main

import (
	"fmt"
)

func main() {
	a, b, x := SingleInt(), SingleInt(), SingleInt()

	fn := divide(x)

	// 5 36 5 --- 36/5...7 + 1(0の分) -- `aの値は含まれていてよい` のでa-1で範囲外にする
	// 範囲外の中に割り切れる数がある場合は、その数を引く。

	// 重複範囲を削るため -fn(a-1) を行う
	fmt.Println(fn(b) - fn(a-1))
}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

// fn は0以上を対象としているため、商に+1している
func divide(x int) func(int) int {
	return func(a int) int {
		if a < 0 {
			return 0
		}
		return a/x + 1
	}
}
