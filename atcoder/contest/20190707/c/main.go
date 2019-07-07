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

const div = 2019

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
	l, r := nextInt(), nextInt()

	ll := l % div
	rr := r % div

	var s, e int64
	if ll > rr {
		s = rr
		e = ll
	} else {
		s = ll
		e = rr
	}

	if s == e {
		fmt.Println(s % div)
		return
	}

	if s == 0 {
		fmt.Println(0)
		return
	}

	re := int64(9999999999)
	for i := s; i <= e; i++ {
		for j := i; j <= e; j++ {
			if i == j {
				continue
			}
			v := i * j % div
			if re > v {
				re = v
			}
		}
	}

	fmt.Println(re)
}
