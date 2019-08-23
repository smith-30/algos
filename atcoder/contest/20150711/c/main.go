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
	Idx   int
	Val   int64
	Child []*Node
}

func main() {
	n := nextInt()

	ns := make([]*Node, n)
	for i := 0; i < n; i++ {
		ns[i] = &Node{
			Child: make([]*Node, 0, 20),
		}
	}
	fmt.Printf("%#v\n", ns)
	for i := 1; i < n; i++ {
		num := nextInt() - 1
		ns[num].Child = append(ns[num].Child, &Node{Idx: i})
	}

	for idx, item := range ns {
		if len(item.Child) == 0 {
			item.Val = 1
			item.Idx = idx
		}
	}

	for _, item := range ns {
		fmt.Printf("%#v\n", item)
		// for _, i := range item.Child {
		// 	fmt.Printf("%#v\n", i)
		// }
	}

	fmt.Println(ns)
}
