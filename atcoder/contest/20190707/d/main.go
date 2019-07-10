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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	n := nextInt()
	lis := make([]int, 0, n)

	for i := 0; i < n; i++ {
		lis = append(lis, nextInt())
	}

	var x2 int
	for i, item := range lis {
		if i%2 == 0 {
			x2 += item
		} else {
			x2 -= item
		}
	}

	ans := make([]int, len(lis), len(lis))

	ans[0] = x2 / 2
	for i := 1; i < n; i++ {
		ans[i] = lis[i-1] - ans[i-1]
	}

	for _, item := range ans {
		fmt.Printf("%#v ", item*2)
	}
	fmt.Println()
}
