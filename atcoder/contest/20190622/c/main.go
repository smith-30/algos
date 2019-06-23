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

func nextInt() int64 {
	sc.Scan()
	i, e := strconv.ParseInt(sc.Text(), 10, 64)
	if e != nil {
		panic(e)
	}
	return i
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	a, b, c, d := nextInt(), nextInt(), nextInt(), nextInt()
	var re int64

	if c == 1 || d == 1 {
		fmt.Println(0)
		return
	}

	if c > a {
		re += c - a
	}

	cd := c * d

	for i := c; i <= d; i++ {

	}

	for i := a; i <= b; i++ {
		if i%c == 0 || i%d == 0 {
			continue
		}
		re++
	}

	fmt.Printf("%#v\n", 2%4)

	fmt.Println(re)
}
