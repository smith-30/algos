package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.ParseInt(sc.Text(), 10, 64)
	if e != nil {
		panic(e)
	}
	return int(i)
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	n := nextInt()

	nn := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nn = append(nn, nextInt())
	}

	var re int

	r := 2
	for i := 0; r < n; i++ {
		nums := make([]int, 0, 3)
		for index := i; index <= r; index++ {
			nums = append(nums, nn[index])
		}

		tmp := nums[1]
		if Min(nums...) != tmp && Max(nums...) != tmp {
			re++
		}
		r++
	}

	fmt.Println(re)
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
