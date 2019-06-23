package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Pair struct {
	Key   int64
	Value int64
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

type Int64Slice []*item

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i].val < p[j].val }
func (p Int64Slice) Swap(i, j int)      { p[i].val, p[j].val = p[j].val, p[i].val }

type item struct {
	val int64
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

func nextInt64() int64 {
	sc.Scan()
	i, e := strconv.ParseInt(sc.Text(), 10, 64)
	if e != nil {
		panic(e)
	}
	return i
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	n := nextInt()

	order := make([]Pair, 0, n)

	for i := 0; i < n; i++ {
		a, b := nextInt64(), nextInt64()
		p := Pair{
			Key:   a,
			Value: b,
		}
		order = append(order, p)
	}

	sort.Sort(ByValue(order))

	var val int64
	for _, item := range order {
		val += item.Key
		if val > item.Value {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
