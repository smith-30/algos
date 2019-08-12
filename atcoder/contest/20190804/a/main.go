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
	lim := nextInt()
	now := nextInt()
	c := nextInt()

	if lim == now {
		fmt.Println(c)
		return
	}

	if now+c <= lim {
		fmt.Println(0)
		return
	} else {
		fmt.Println(c - (lim - now))
	}
}
