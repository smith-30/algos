package main

import (
	"bufio"
	"fmt"
	"math"
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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	n, d := nextInt(), nextInt()
	nn := make([][]int, n, n)

	for i := 0; i < n; i++ {
		nn[i] = make([]int, d, d)
		for index := 0; index < d; index++ {
			nn[i][index] = nextInt()
		}
	}

	var re int
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i == j {
				continue
			}
			var v float64
			for index := 0; index < d; index++ {
				v += math.Pow(math.Abs(float64(nn[i][index]-nn[j][index])), 2)
			}

			vv := fmt.Sprintf("%v", math.Sqrt(v))

			if !strings.Contains(vv, ".") {
				re++
			}
		}
	}

	fmt.Println(re)
}
