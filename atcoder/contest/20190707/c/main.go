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

	if r-l >= 2019 {
		fmt.Println(0)
		return
	}

	re := int64(9999999999)
	for i := l; i <= r; i++ {
		for j := i; j <= r; j++ {
			if i == j {
				continue
			}
			v := i * j % div
			if re > v {
				re = v
			}
			if re == 0 {
				fmt.Println(re)
				return
			}
		}
	}

	fmt.Println(re)
}
