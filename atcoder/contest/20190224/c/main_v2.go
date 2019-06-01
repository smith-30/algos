package main

import (
	"fmt"
	"math"
)

var N, A, B, C int
var RE = 1000000000
var arr []int

func main() {
	fmt.Scan(&N, &A, &B, &C)
	arr = make([]int, N)
	for i := 0; i < N; i++ {
		var tmp int
		fmt.Scan(&tmp)
		arr[i] = tmp
	}
	pattern := make([]int, N, N)
	dfs(0, pattern)
	fmt.Println(RE)
}

// 選択肢 = 4 (A, B, C, 何も選ばない) と、何個の竹があるか (N)
// 4^N バターンについて考えればいい
// 下記でNまでの深さの木を作っている
func dfs(depth int, arr []int) {
	if depth == N {
		calc(arr)
		return
	}
	for i := 0; i < 4; i++ {
		arr[depth] = i
		dfs(depth+1, arr)
	}
}

func calc(v []int) {
	var cost, a, b, c, asum, bsum, csum, tmp int
	for idx, item := range v {
		switch item {
		// A使う
		case 0:
			a++
			asum += arr[idx]
		// B使う
		case 1:
			b++
			bsum += arr[idx]
		// C使う
		case 2:
			c++
			csum += arr[idx]
		// 使わない
		case 3:
			tmp += arr[idx]
		}
	}
	// a, b, c に当てはまるものがない場合は、正当な結果が得られないのでskip
	if a == 0 || b == 0 || c == 0 {
		return
	}
	addSum := (a-1)*10 + (b-1)*10 + (c-1)*10
	cost = addSum + abs(A-asum) + abs(B-bsum) + abs(C-csum)
	if RE > cost {
		RE = cost
	}
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}
