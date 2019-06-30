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
	n, m := nextInt(), nextInt()
	pp := map[int][]Pair{}
	vv := make([]Pair, 0, n)

	for i := 0; i < m; i++ {
		k, v := nextInt(), nextInt()
		pp[k] = append(pp[k], Pair{
			Value: v,
		})
		vv = append(vv, Pair{
			Key:   k,
			Value: v,
		})
	}

	sort.Sort(ByValue(pp[1]))

	res := map[int]map[int]string{}

	for c, item := range pp {
		sort.Sort(ByValue(item))
		res[c] = map[int]string{}

		for idx, i := range item {
			res[c][i.Value] = fmt.Sprintf("%06d", idx+1)
		}
	}

	for _, item := range vv {
		fmt.Printf("%06d%s\n", item.Key, res[item.Key][item.Value])
	}
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
