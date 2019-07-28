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

func main() {
	n := nextInt()

	ss := make([]int, 0, n)
	for i := 0; i < n; i++ {
		ss = append(ss, nextInt())
	}

	var c int
	dd := make([]int, n)
	copy(dd, ss)
	sort.Sort(sort.IntSlice(dd))

	for i := 0; i < n; i++ {
		if dd[i] != ss[i] {
			c++
		}
	}

	if c == 0 || c == 2 {
		fmt.Println("YES")
		return
	}

	fmt.Println("NO")
}
