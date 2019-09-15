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

func nextInt64() int64 {
	sc.Scan()
	i, e := strconv.ParseInt(sc.Text(), 10, 64)
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	n := nextInt()
	k := nextInt64()
	q := nextInt()

	ps := make([]int64, n)

	// for index := 0; index < n; index++ {
	// 	ps[index] = k
	// }

	for index := 0; index < q; index++ {
		a := nextInt() - 1
		ps[a]++
	}

	for _, item := range ps {
		if item > int64(q)-k {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
