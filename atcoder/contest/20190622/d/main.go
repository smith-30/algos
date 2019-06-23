package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

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

	comb := map[*item]int64{}
	order := make(Int64Slice, 0, n)

	for i := 0; i < n; i++ {
		a, b := nextInt64(), nextInt64()
		it := &item{
			val: b,
		}
		comb[it] = a
		order = append(order, it)
	}

	sort.Sort(order)

	for _, item := range order {
		fmt.Printf("%#v\n", item.val)
	}
	fmt.Printf("%#v\n", "")
	var val int64
	for _, item := range order {
		fmt.Printf("%#v\n", item.val)
		val += comb[item]
		if val > item.val {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
