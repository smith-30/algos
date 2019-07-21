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

func main() {
	n, d := nextInt(), nextInt()

	var ans int
	i := d + 1
	for {
		ans++
		end := i + d
		if end >= n {
			fmt.Println(ans)
			return
		}
		i = end + 1 + d
		// if i >= n {
		// 	fmt.Println(ans)
		// 	return
		// }
	}
}
