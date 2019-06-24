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
	vv := make([]int, 0, n)

	for i := 0; i < n; i++ {
		vv = append(vv, nextInt())
	}

	sort.Sort(sort.Reverse(sort.IntSlice(vv)))

	m := vv[0]

	var sum int
	for _, item := range vv[1:] {
		sum += item
	}

	if m < sum {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
