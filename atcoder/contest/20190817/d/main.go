package main

import (
	"bufio"
	"fmt"
	"os"
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

type Node struct {
	number int
	val    int64
	child  []*Node
}

func (a *Node) Flow(num int, v int64) {
	if a.number == num {
		a.val += v
		for _, item := range a.child {
			item.Add(v)
		}
	} else {
		for _, item := range a.child {
			item.Flow(num, v)
		}
	}
}

func (a *Node) Add(v int64) {
	a.val += v
	for _, item := range a.child {
		item.Add(v)
	}
}

func (a *Node) AddNode(key int, number int) {
	if key == a.number {
		a.child = append(a.child, &Node{number: number})
	} else {
		for _, item := range a.child {
			item.AddNode(key, number)
		}
	}
}

func (a *Node) Show() {
	fmt.Printf("%v ", a.val)
	for _, item := range a.child {
		item.Show()
	}
}

func main() {
	n := nextInt()
	q := nextInt()
	N := &Node{number: 1}
	for index := 1; index < n; index++ {
		from, to := nextInt(), nextInt()
		N.AddNode(from, to)
	}

	for index := 0; index < q; index++ {
		target, v := nextInt(), nextInt()
		N.Flow(target, int64(v))
	}

	N.Show()
}
