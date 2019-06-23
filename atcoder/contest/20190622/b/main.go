package main

import (
	"bufio"
	"fmt"
	"math"
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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	n, l := nextInt(), nextInt()

	apps := make([]int, 0, n)

	c := 99999

	var re int
	for index := 1; index <= n; index++ {
		v := l + index - 1
		k := int(math.Abs(float64(v)))
		if k < int(math.Abs(float64(c))) {
			c = v
		}

		apps = append(apps, l+index-1)
	}

	for _, item := range apps {
		if item == c {
			continue
		}
		re += item
	}

	fmt.Printf("%#v\n", re)
}
