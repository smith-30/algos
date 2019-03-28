package lib

import (
	"fmt"
	"math"
	"sort"
)

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

func ScanNums(len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}

func ScanStrings(len int) (strings []string) {
	var str string
	for i := 0; i < len; i++ {
		fmt.Scanf("%s", &str)
		strings = append(strings, str)
	}
	return
}

func Abs(a int) int {
	return int(math.Abs(float64(a)))
}

func Pow(p, q int) int {
	return int(math.Pow(float64(p), float64(q)))
}

func StrSearch(a []string, b string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == b {
			return true
		}
	}
	return false
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

func Max(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton max() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Max(float64(res), float64(nums[i])))
	}
	return res
}

func Sum(fn func(key, value int) int) func(m []int) int {
	return func(m []int) int {
		var result int
		for key, value := range m {
			result += fn(key, value)
		}
		return result
	}
}

func SumDigits(number int) int {
	var remainder, sumResult int
	for number != 0 {
		remainder = number % 10
		sumResult += remainder
		number = number / 10
	}
	return sumResult
}

func NoopSum(key, value int) int {
	return value
}

func MakeRangeInts(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func ReverseIntSlice(s []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
}
