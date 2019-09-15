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
	a, b := nextInt(), nextInt()
	if b == 1 {
		fmt.Println(0)
		return
	}

	var re, c int

	for {
		if re == 0 {
			re = a
		} else {
			re--
			re += a
		}
		c++

		if re >= b {
			fmt.Println(c)
			return
		}
	}
}
