package main

import (
	"fmt"
	"sort"
)

type pair struct {
	idx  int
	name string
	val  int
}

type Int64Slice []pair

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i].name < p[j].name }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Slice []pair

func (p Slice) Len() int           { return len(p) }
func (p Slice) Less(i, j int) bool { return p[i].val > p[j].val }
func (p Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	var n int
	fmt.Scan(&n)

	all := make(Int64Slice, 0, n)

	for index := 0; index < n; index++ {
		var st string
		fmt.Scan(&st)

		var n int
		fmt.Scan(&n)

		all = append(all, pair{idx: index + 1, name: st, val: n})
	}
	sort.Sort(all)

	p := Slice{}
	now := all[0].name

	for _, item := range all {
		if now == item.name {
			p = append(p, item)
		} else {
			sort.Sort(p)
			for _, it := range p {
				fmt.Println(it.idx)
			}
			p = Slice{item}
			now = item.name
		}
	}
	sort.Sort(p)
	for _, it := range p {
		fmt.Println(it.idx)
	}

	// var re int
	// fmt.Println()
}
