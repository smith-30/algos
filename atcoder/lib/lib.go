package lib

import (
	"fmt"
	"math"
	"sort"
)

type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

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

func ReverseInt(target []int) {
	for i := len(target)/2 - 1; i >= 0; i-- {
		opp := len(target) - 1 - i
		target[i], target[opp] = target[opp], target[i]
	}
}

func ReverseInt64(target []int64) {
	for i := len(target)/2 - 1; i >= 0; i-- {
		opp := len(target) - 1 - i
		target[i], target[opp] = target[opp], target[i]
	}
}

func CeilInt(a, b int) int {
	d := float64(a) / float64(b)
	return int(math.Ceil(d))
}

func CeilInt64(a, b int64) int64 {
	d := float64(a) / float64(b)
	return int64(math.Ceil(d))
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

func StrReverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func IntPow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

type Pair struct {
	Key   int
	Value int
}

// sort.Sort(ByKey(pairs))
type ByKey []Pair

func (s ByKey) Len() int {
	return len(s)
}

func (s ByKey) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByKey) Less(i, j int) bool {
	return s[i].Key < s[j].Key
}

// sort.Sort(ByValue(pairs))
type ByValue []Pair

func (s ByValue) Len() int {
	return len(s)
}

func (s ByValue) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByValue) Less(i, j int) bool {
	return s[i].Value < s[j].Value
}

// greatest common divisor (GCD) via Euclidean algorithm
func _GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func _LCM(a, b int, integers ...int) int {
	result := a * b / _GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

/* Function without return declaration*/
func LCM(temp1 int, temp2 int) int {
	var lcmnum int = 1
	if temp1 > temp2 {
		lcmnum = temp1
	} else {
		lcmnum = temp2
	}
	/* Use of For Loop as a While Loop*/
	for {
		if lcmnum%temp1 == 0 && lcmnum%temp2 == 0 { // And operator
			/*  Print Statement with multiple variables   */
			return lcmnum
		}
		lcmnum++
	}
	return lcmnum // Return without any value
}

func __gcd(temp1 int, temp2 int) {
	var gcdnum int
	/* Use of And operator in For Loop */
	for i := 1; i <= temp1 && i <= temp2; i++ {
		if temp1%i == 0 && temp2%i == 0 {
			gcdnum = i
		}
	}
	fmt.Printf("GCD of %d and %d is %d", temp1, temp2, gcdnum)
	return
}

// greatest common divisor (GCD) via Euclidean algorithm
func Gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func Lcm(a, b int, integers ...int) int {
	result := a * b / Gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = Lcm(result, integers[i])
	}

	return result
}

func IntReverse(vv []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(vv)))
}

func nCr(n, r int64) int64 {
	if r == 0 {
		return 0
	}

	var a int64 = 1
	for i := 0; i < int(r); i++ {
		a *= (n - int64(i))
	}

	var b int64 = 1
	for i := 1; i <= int(r); i++ {
		b *= int64(i)
	}

	return a / b
}

const COMNUM = 2000

const MOD = 1000000007

var Com [][]int64

func CalcCom(comnum int) {
	Com = make([][]int64, comnum, comnum)
	Com[0] = make([]int64, comnum, comnum)
	Com[0][0] = 1
	for i := 1; i < comnum; i++ {
		Com[i] = make([]int64, comnum, comnum)
	}
	for i := 1; i < comnum; i++ {
		Com[i][0] = 1
		for j := 1; j < comnum; j++ {
			Com[i][j] = (Com[i-1][j-1] + Com[i-1][j]) % MOD
		}
	}
}

const MAX = 2001
const _MOD = 1000000007

var fac, finv, inv [MAX]int

func cominit() {
	fac[0], fac[1] = 1, 1
	finv[0], finv[1] = 1, 1
	inv[1] = 1
	for i := 2; i < MAX; i++ {
		fac[i] = fac[i-1] * i % MOD
		inv[i] = MOD - inv[MOD%i]*(MOD/i)%MOD
		finv[i] = finv[i-1] * inv[i] % MOD
	}
}

func combination(n, k int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	return fac[n] * (finv[k] * finv[n-k] % MOD) % MOD
}
