package main

import (
	"fmt"
	"math"
)

func main() {
	wid := SingleInt()
	n, k := SingleInt(), SingleInt()

	chs := make([][]int, n)
	width := 0
	value := 1

	for i := 0; i < n; i++ {
		chs[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			chs[i][j] = SingleInt()
		}
	}

	dp := make([][]int, k+1)
	for idx, _ := range dp {
		dp[idx] = make([]int, wid+1)
	}

	fmt.Printf("%#v\n", dp)

	var ans int

	// 幅 -> vals[width]
	// 次 -> vals[width]
	for _, vals := range chs {
		// なぜindexを後ろから進めているかというと
		// 20130831-Aと同じように、ある選択(index)のときにキャッシュさせたい値が
		// 同一ループ内で再アクセスされるのを防ぎたいから
		// 基本的にdpの一次元目のindexを持つ値には再アクセスさせないように工夫するっぽい
		// index足りなくてout of range error 出てきたらループの回し方がおかしいと思ったほうがいい
		for i := k; i >= 1; i-- {
			fmt.Printf("%#v, %#v\n", i, vals[value])
			for j := vals[width]; j <= wid; j++ {
				dp[i][j] = Max(dp[i][j-vals[width]]+vals[value], dp[i][j])
				ans = Max(ans, dp[i][j])
			}
		}
		fmt.Printf("%#v\n", dp)
	}

	fmt.Printf("%#v\n", ans)
}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

func Max(nums ...int) int {
	if len(nums) == 0 {
		panic("function max() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Max(float64(res), float64(nums[i])))
	}
	return res
}
