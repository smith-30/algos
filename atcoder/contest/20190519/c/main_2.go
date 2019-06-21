package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
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
	sc.Split(bufio.ScanWords)
	n, k := nextInt(), nextInt()

	var re float64
	nn := float64(n)
	for i := 1; i <= n; i++ {
		if i >= k {
			re += 1 / nn
			continue
		}

		x := math.Ceil(math.Log2(float64(k) / float64(i)))
		re += (1 / nn) * (1 / math.Pow(2, x))
	}

	fmt.Printf("%.15f\n", re)
}
