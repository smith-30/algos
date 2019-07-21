package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
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
	return s[i].Value > s[j].Value
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.ParseInt(sc.Text(), 10, 64)
	if e != nil {
		panic(e)
	}
	return int(i)
}

func main() {
	n := nextInt()

	kvs := make([]Pair, 0, n)

	for i := 1; i <= n; i++ {
		v := nextInt()
		kvs = append(kvs, Pair{Key: i, Value: v})
	}

	sort.Sort(ByValue(kvs))

	for i := 1; i <= n; i++ {
		for _, item := range kvs {
			if item.Key != i {
				fmt.Println(item.Value)
				break
			}
		}
	}
}
