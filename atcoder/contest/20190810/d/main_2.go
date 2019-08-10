package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type Pair struct {
	Index int
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

func main() {
	n, w := nextInt(), nextInt()+1
	ps := make([]Pair, 0, n)

	for i := 0; i < n; i++ {
		vv, w := nextInt(), nextInt()
		ps = append(ps, Pair{
			Index: i,
			Key:   vv,
			Value: w,
		})
	}

	used := map[int]struct{}{}
	sort.Sort(sort.Reverse(ByValue(ps)))

	var c int64
	for i := 0; i < n; i++ {
		for _, item := range ps {
			if item.Key <= w {
				if _, ok := used[item.Index]; !ok {
					// fmt.Printf("%#v\n", item.Value)
					c += int64(item.Value)
					used[item.Index] = struct{}{}
					w -= item.Key
					break
				}
			}
		}
	}

	fmt.Println(c)
}
