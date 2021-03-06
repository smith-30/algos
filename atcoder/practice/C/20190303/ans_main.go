package main

import (
	"fmt"
	"math"
	"strings"
)

// 文字列 S に 0 も 1 も 1 つ以上残っている場合、どこかで 0 と 1 が隣り合っています。従って、与えら
// れた文字列 S に含まれる 0 の個数と 1 の個数を数え、それぞれ C0, C1 とすると、どのような順番で取り
// 除いても min(C0, C1) 回取り除く操作ができます。よって、答え (取り除けるキューブの個数の最大値) は
// 2 × min(C0, C1) になります。このアルゴリズムは O(|S|) で動作します。

func main() {
	s := SingleStr()
	ss := strings.Split(s, "")

	var oneC, zeroC int
	for _, item := range ss {
		if item == "0" {
			zeroC++
		} else {
			oneC++
		}
	}

	if zeroC == len(ss) || oneC == len(ss) {
		fmt.Println(0)
		return
	}

	fmt.Println(Min(oneC, zeroC) * 2)
}

func SingleStr() string {
	var s string
	fmt.Scan(&s)
	return s
}

func Min(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton min() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Min(float64(res), float64(nums[i])))
	}
	return res
}
