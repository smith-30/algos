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

func nextInt64() int64 {
	sc.Scan()
	i, e := strconv.ParseInt(sc.Text(), 10, 64)
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	n := nextInt64()
	m := nextInt64()

	if n == 0 || m == 0 {
		fmt.Println("IMPOSSIBLE")
		return
	}

	if int64(math.Abs(float64(n-m)))%2 == 1 {
		fmt.Println("IMPOSSIBLE")
		return
	}

	k := (math.Pow(float64(m), 2) - math.Pow(float64(n), 2)) / float64((2 * (n - m)))
	fmt.Println(int64(math.Abs(k)))
}
