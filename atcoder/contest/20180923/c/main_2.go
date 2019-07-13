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
	s := strings.Split(nextLine(), "")
	t := strings.Split(nextLine(), "")

	n := len(s)
	start := map[string]string{}
	goal := map[string]string{}

	for i := 0; i < n; i++ {
		a := s[i]
		b := t[i]

		// 何かしら変換が行われている
		if start[a] != "" {
			if start[a] != b {
				fmt.Println("No")
				return
			}
		} else {
			start[a] = b
		}

		if goal[b] != "" {
			if goal[b] != a {
				fmt.Println("No")
				return
			}
		} else {
			goal[b] = a
		}
	}

	fmt.Println("Yes")
}
