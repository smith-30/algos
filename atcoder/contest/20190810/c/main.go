package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
	mm := map[string]int64{}
	for i := 0; i < n; i++ {
		s := nextLine()
		ss := strings.Split(s, "")
		sort.Sort(sort.StringSlice(ss))
		v := strings.Join(ss, "")
		mm[v]++
	}

	var re int64
	for _, v := range mm {
		if v > 1 {
			re += nCr(v, 2)
		}
	}

	fmt.Println(re)
}

func nCr(n, r int64) int64 {
	if r == 0 {
		return 0
	}

	var a int64 = 1
	for i := 0; i < int(r); i++ {
		a *= (n - int64(i))
	}

	var b int64 = 1
	for i := 1; i <= int(r); i++ {
		b *= int64(i)
	}

	return a / b
}
