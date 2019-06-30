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
	n := nextInt()
	nn := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nn = append(nn, nextInt())
	}

	sort.Sort(sort.IntSlice(nn))
	h := n / 2

	divless := nn[h-1]
	divmore := nn[h]

	fmt.Println(divmore - divless)
}

func IntReverse(vv []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(vv)))
}
