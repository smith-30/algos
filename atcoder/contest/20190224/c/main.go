package main

import (
	"fmt"
	"math"
)

var N, A, B, C int

var RE int

func main() {
	fmt.Scan(&N, &A, &B, &C)
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		var tmp int
		fmt.Scan(&tmp)
		arr[i] = tmp
	}
	fmt.Println(DFS(0, 0, 0, 0, arr))
	fmt.Printf("%#v\n", RE)
}

func DFS(depth, a, b, c int, arr []int) int {
	if depth >= N {
		RE++
		if a > 0 && b > 0 && c > 0 {
			return abs(A-a) + abs(B-b) + abs(C-c) - 30
		} else {
			return math.MaxInt32
		}
	}
	w := DFS(depth+1, a, b, c, arr)
	x := DFS(depth+1, a+arr[depth], b, c, arr) + 10
	y := DFS(depth+1, a, b+arr[depth], c, arr) + 10
	z := DFS(depth+1, a, b, c+arr[depth], arr) + 10

	fmt.Printf("%#v\n", []int{a, b, c})
	return min1([]int{w, x, y, z})
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func min1(arr []int) int {
	a := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < a {
			a = arr[i]
		}
	}
	return a
}
