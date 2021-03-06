package lib

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func SInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

func DInt() (int, int) {
	var n int
	fmt.Scan(&n)

	var m int
	fmt.Scan(&m)
	return n, m
}

func SStr() string {
	var s string
	fmt.Scan(&s)
	return s
}

func StrSlice() []string {
	var s string
	fmt.Scan(&s)
	return strings.Split(s, "")
}

type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// An IntHeap is a min-heap of ints.
type Int64Heap []int64

func (h Int64Heap) Len() int           { return len(h) }
func (h Int64Heap) Less(i, j int) bool { return h[i] > h[j] }
func (h Int64Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Int64Heap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int64))
}

func (h *Int64Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

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

func modpow(a, p int64) int64 {
	if p == 0 {
		return 1
	}

	if p%2 == 0 {
		//pが偶数の時
		halfP := p / 2
		half := modpow(a, halfP)
		//a^(p/2) をhalfとして、half*halfを計算
		return half * half % MOD
	} else {
		//pが奇数の時は、
		//pを偶数にするために1減らす
		return a * modpow(a, p-1) % MOD
	}
}

//(10*9*8)/(3*2*1);
//10*9*8 -> ansMul
//3*2*1 -> ansDiv
func calcComb(a, b int64) int64 {
	if b > a-b {
		return calcComb(a, a-b)
	}

	var ansMul int64 = 1
	var ansDiv int64 = 1
	var i int64
	for i = 0; i < b; i++ {
		ansMul *= (a - i)
		ansDiv *= (i + 1)
		ansMul %= MOD
		ansDiv %= MOD
	}
	//ansMul / ansDivをやりたい
	//ansDivの逆元を使って求めよう！

	ans := ansMul * modpow(ansDiv, MOD-2) % MOD
	return ans
}

func Divisors(n int) []int {
	divisors := make([]int, 0, 1000)
	end := int(math.Pow(float64(n), 0.5) + 1)
	for i := 1; i < end; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if v := n / i; i != n/i {
				divisors = append(divisors, v)
			}
		}

	}
	return divisors
}

type PriorityQueue []int

func (h PriorityQueue) Len() int            { return len(h) }
func (h PriorityQueue) Less(i, j int) bool  { return h[i] > h[j] }
func (h PriorityQueue) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *PriorityQueue) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *PriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

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
