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

	if n >= m {
		fmt.Println(0)
		return
	}

	vv := make([]int, 0, m)
	for i := 0; i < m; i++ {
		vv = append(vv, nextInt())
	}

	sort.Sort(sort.Reverse(sort.IntSlice(vv)))

	nn := make([]int, 0, m)
	for i, item := range vv {
		if i+1 == len(vv) {
			break
		}
		nn = append(nn, item-vv[i+1])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nn)))

	var sum int
	for _, item := range nn[n-1:] {
		sum += item
	}

	fmt.Printf("%#v\n", sum)
}
