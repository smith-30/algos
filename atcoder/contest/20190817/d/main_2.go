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

var (
	ans  []int
	tree [][]int
)

func main() {
	n := nextInt()
	q := nextInt()

	tree = make([][]int, n)

	for index := 0; index < n-1; index++ {
		from, to := nextInt()-1, nextInt()-1

		tree[from] = append(tree[from], to)
		tree[to] = append(tree[to], from)
	}

	ans = make([]int, n)
	for index := 0; index < q; index++ {
		target, v := nextInt()-1, nextInt()
		ans[target] += v
	}
	dfs(0, -1)

	for _, item := range ans {
		fmt.Printf("%v ", item)
	}
	fmt.Println()
}

func dfs(n, from int) {
	for _, item := range tree[n] {
		if from == item {
			continue
		}
		ans[item] += ans[n]
		dfs(item, n)
	}
}
