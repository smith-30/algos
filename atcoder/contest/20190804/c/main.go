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

func nextInt64() int64 {
	sc.Scan()
	i, e := strconv.ParseInt(sc.Text(), 10, 64)
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	n := nextInt()
	yes := "Yes"
	no := "No"

	if n == 1 {
		fmt.Println(yes)
		return
	}

	nn := make([]int64, 0, n)
	for i := 0; i < n; i++ {
		nn = append(nn, nextInt64())
	}

	nn = append(nn, 1000000000)

	for i := 0; i < n; i++ {
		// if i+1 == n {
		// 	break
		// }
		if nn[i+1] == nn[i] {
			continue
		}
		if nn[i+1]-1 < nn[i] {
			// fmt.Printf("%#v %#v %#v\n", i, nn[i+1], nn[i])
			fmt.Println(no)
			return
		}

		if nn[i+1] > nn[i] {
			nn[i+1]--
		}
	}

	fmt.Println(yes)
}
