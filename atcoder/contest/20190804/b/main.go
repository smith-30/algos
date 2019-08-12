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

func main() {
	n := nextInt()

	switch {
	case n < 10:
		fmt.Println(n)
		return
	case n >= 10 && n < 100:
		fmt.Println(9)
		return
	case 100 <= n && n < 1000:
		fmt.Println(9 + (n - 100 + 1))
		return
	case 1000 <= n && n < 10000:
		fmt.Println(9 + 900)
		return
	case 10000 <= n && n < 100000:
		fmt.Println(9 + 900 + (n - 10000 + 1))
		return
	case n == 100000:
		fmt.Println(90909)
		return
	}
}
