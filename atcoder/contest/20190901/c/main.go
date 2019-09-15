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

func nextInt64() int64 {
	sc.Scan()
	i, e := strconv.ParseInt(sc.Text(), 10, 64)
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	n := nextInt64()
	mm := make([]int64, 0, n)

	start := nextInt64()
	var result int64
	for index := int64(1); index < n; index++ {
		now := nextInt64()
		if start >= now {
			result++
		} else {
			mm = append(mm, result)
			result = 0
		}
		start = now
	}

	if result > 0 {
		mm = append(mm, result)
	}

	if len(mm) == 0 {
		fmt.Println(0)
		return
	}

	fmt.Println(Max(mm...))
}

func Max(nums ...int64) int64 {
	if len(nums) == 0 {
		panic("function max() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int64(math.Max(float64(res), float64(nums[i])))
	}
	return res
}
