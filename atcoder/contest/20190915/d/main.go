package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] > p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	n, m := nextInt(), nextInt()
	nn := make(Int64Slice, 0, n)

	for index := 0; index < n; index++ {
		a := nextInt64()
		nn = append(nn, a)
	}

	if len(nn) == 1 {
		ans := nn[0] / int64(math.Pow(2, float64(m)))
		fmt.Println(ans)
		return
	}

	sort.Sort(nn)

	var index int

	for {
		if m == 0 {
			break
		}
		var peek, now int64

		peek = nn[index+1]
		now = nn[index]

		// fmt.Printf("%#v, %#v\n", now, peek)

		for {
			m--
			if now = now / 2; now < peek {
				nn[index+1] = now
				nn[index] = peek
				break
			}
			if m == 0 {
				nn[index] = now
				break
			}
		}
	}

	var re int64
	for _, item := range nn {
		re += item
	}
	fmt.Println(re)
}
