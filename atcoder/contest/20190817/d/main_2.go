package main

import (
	"bufio"
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

func main() {
	n := nextInt()
	q := nextInt()

	line := make([]int64, n)
	for index := 1; index <= n; index++ {
		from, to := nextInt(), nextInt()
	}

	for index := 0; index < q; index++ {
		target, v := nextInt(), nextInt()
	}
}
