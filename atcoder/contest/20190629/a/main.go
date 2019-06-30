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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	n := nextLine()
	r := map[string]string{}
	nn := strings.Split(n, "")

	for _, item := range nn {
		r[item] = item
	}

	if len(r) == 2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
