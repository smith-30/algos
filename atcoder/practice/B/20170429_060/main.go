package main

import (
	"fmt"
	"math"
)

func main() {
	a, b, c := SingleInt(), SingleInt(), SingleInt()

	// sum := x*a + y*a + ...
	// sum := (x + y + ...)*a --- x * a  10 * 5
	// c = sum % b
	// sum / b + c =

	// 100 % 33 = 1 --- 34 % 33 = 1
	// bで割ったときのあまりを作る最小の手段
	// b + c = (x + y + ...)*a / b
	// 終了条件は?
	// bで割ったときのあまりとして考えられるのは、b - n - 1 = 0 -- n = b -1

	if a == 1 {
		fmt.Println("YES")
		return
	}

	if a%2 == 0 && c%2 != 0 {
		fmt.Println("NO")
		return
	}

	for i := 1; i <= 100; i++ {
		for j := 1; j <= 100; j++ {
			if b*i+c == j*a {
				fmt.Println("YES")
				return
			}
		}
	}
	fmt.Println("NO")
}

func SingleInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

func Min(nums ...int) int {
	if len(nums) == 0 {
		panic("function min() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Min(float64(res), float64(nums[i])))
	}
	return res
}
