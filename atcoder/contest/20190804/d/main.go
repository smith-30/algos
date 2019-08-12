package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	var s string
	fmt.Scan(&s)

	res := make([]int, len(s))
	for i, _ := range res {
		res[i] = 1
	}

	resCache := make([]int, len(s))

	lll := strings.Split(s, "R")
	rrr := strings.Split(s, "L")

	ss := strings.Split(s, "")
	var rc, lc int
	for _, item := range lll {
		if l := len(item); lc < l {
			lc = l
		}
	}
	for _, item := range rrr {
		if l := len(item); rc < l {
			rc = l
		}
	}

	var l int
	if rc >= lc {
		l = rc
	} else {
		l = lc
	}

	if l%2 != 0 {
		l++
	}

	for index := 0; index < l; index++ {
		copy(resCache, res)
		for idx, item := range ss {
			switch item {
			case "R":
				res[idx+1] += resCache[idx]
				res[idx] -= resCache[idx]
			case "L":
				res[idx-1] += resCache[idx]
				res[idx] -= resCache[idx]
			}
		}
	}

	for _, item := range res {
		fmt.Printf("%v ", item)
	}
	fmt.Println()
}
